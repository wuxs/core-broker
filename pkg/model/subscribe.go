package model

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/tkeel-io/core-broker/pkg/util"
	"github.com/tkeel-io/kit/log"
	"gorm.io/gorm"
)

var ErrUndeleteable = errors.New("undeleteable")

type Subscribe struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	UserID      string `gorm:"index"`
	Endpoint    string `gorm:"index"`
	IsDefault   bool   `gorm:"default:false"`
}

func (s *Subscribe) BeforeCreate(tx *gorm.DB) error {
	if s.Endpoint == "" {
		s.Endpoint = util.GenerateSubscribeEndpoint()
	}
	return nil
}

func (s *Subscribe) BeforeDelete(tx *gorm.DB) error {
	if s.IsDefault {
		return NewUndeleteable("this is default subscribe")
	}
	destroyEndpoint(tx, s.Endpoint)
	destroyRelevant(tx, s.ID)
	return nil
}

func destroyEndpoint(tx *gorm.DB, endpoint string) {
	// TODO: destroy endpoint
	log.Debug("destroy endpoint: %s", endpoint)
}

func destroyRelevant(tx *gorm.DB, id uint) {
	tx.Delete(&SubscribeEntities{}, "subscribe_id = ?", id)
	log.Debugf("destroy Relevant subscribe id: %d", id)
}

type SubscribeEntities struct {
	EntityID    string `gorm:"index;not null"`
	UniqueKey   string `gorm:"index;unique;size:255"`
	SubscribeID uint   `gorm:"index;not null"`

	Subscribe Subscribe
}

func (e *SubscribeEntities) AfterCreate(tx *gorm.DB) error {
	tx.Model(&e.Subscribe).Where("id = ?", e.SubscribeID).First(&e.Subscribe)
	log.Debug("creation of SubscribeEntities:", *e)
	if err := createCoreSubscription(e.EntityID, e.Subscribe.Endpoint); err != nil {
		err = errors.Wrap(err, "create core subscription err")
		log.Error(err)
		return err
	}
	if err := updateEntitySubscribeEndpoint(e.EntityID,
		strings.Join([]string{e.Subscribe.Title, strconv.FormatUint(uint64(e.SubscribeID), 10),
			makeAMQPAddress(e.Subscribe.Endpoint)}, "@"),
		add); err != nil {
		err = errors.Wrap(err, "update entity subscribe endpoint err")
		log.Error(err)
		return err
	}
	return nil
}

func (e *SubscribeEntities) BeforeDelete(tx *gorm.DB) error {
	// this condition will skip by destroyRelevant() function
	if e.EntityID == "" && e.Subscribe.Endpoint == "" {
		return nil
	}
	log.Debug("deleted of SubscribeEntities:", *e)
	if err := deleteCoreSubscription(e.EntityID, e.Subscribe.Endpoint); err != nil {
		log.Error(err)
		return err
	}
	if err := updateEntitySubscribeEndpoint(e.EntityID,
		strings.Join([]string{e.Subscribe.Title, strconv.FormatUint(uint64(e.SubscribeID), 10),
			makeAMQPAddress(e.Subscribe.Endpoint)}, "@"),
		reduce); err != nil {
		return err
	}
	return nil
}

func createCoreSubscription(entityID string, topic string) error {
	return coreClient.Subscribe(entityID, topic)
}

func deleteCoreSubscription(entityID string, topic string) error {
	return coreClient.Unsubscribe(entityID, topic)
}

type choice uint8

const (
	add choice = iota + 1
	reduce
)

func updateEntitySubscribeEndpoint(entityID, endpoint string, c choice) error {
	separator := ","
	patchData := make([]map[string]interface{}, 0)

	device, err := coreClient.GetDeviceEntity(entityID)
	if err != nil {
		log.Error("get entity err:", err)
		return err
	}
	subscribeAddr := endpoint
	switch c {
	case add:
		if strings.Contains(device.Properties.SysField.SubscribeAddr, endpoint) {
			return nil
		}
		if device.Properties.SysField.SubscribeAddr != "" {
			subscribeAddr = strings.Join([]string{device.Properties.SysField.SubscribeAddr, endpoint}, separator)
		}
	case reduce:
		addrs := strings.Split(device.Properties.SysField.SubscribeAddr, separator)
		validAddresses := make([]string, 0, len(addrs))
		for i := range addrs {
			if addrs[i] != endpoint {
				validAddresses = append(validAddresses, addrs[i])
			}
		}
		if len(validAddresses) != 0 {
			subscribeAddr = strings.Join(validAddresses, separator)
		} else {
			subscribeAddr = ""
		}
	}

	patchData = append(patchData, map[string]interface{}{
		"operator": "replace",
		"path":     "sysField._subscribeAddr",
		"value":    subscribeAddr,
	})

	log.Debug("patchData:", patchData)
	log.Debug("call patch on choice (add 1, reduce 2):", c)

	if err = coreClient.PatchEntity(entityID, patchData); err != nil {
		err = errors.Wrap(err, "patch entity err")
		return err
	}

	return nil
}

func NewUndeleteable(content string) error {
	return errors.Wrap(ErrUndeleteable, content)
}
