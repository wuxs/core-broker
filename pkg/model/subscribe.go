package model

import (
	"github.com/pkg/errors"
	"github.com/tkeel-io/core-broker/pkg/util"
	"github.com/tkeel-io/kit/log"
	"gorm.io/gorm"
	"strings"
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

func (s Subscribe) AfterDelete(tx *gorm.DB) error {
	if s.IsDefault {
		return NewUndeleteable("this is default subscribe")
	}
	destroyEndpoint(s.Endpoint)
	destroyRelevant(s.ID)
	return nil
}

func destroyEndpoint(endpoint string) {
	// TODO: destroy endpoint
	log.Debug("endpoint: %s", endpoint)
}

func destroyRelevant(id uint) {
	DB().Delete(SubscribeEntities{}, "subscribe_id = ?", id)
	log.Debug("destroyRelevant")
}

type SubscribeEntities struct {
	EntityID    string `gorm:"index;not null"`
	UniqueKey   string `gorm:"index;unique;size:255"`
	SubscribeID uint   `gorm:"index;not null"`

	Subscribe Subscribe
}

func (e *SubscribeEntities) AfterCreate(tx *gorm.DB) error {
	if err := createCoreSubscription(e.EntityID, e.UniqueKey); err != nil {
		log.Error(err)
		return err
	}
	if err := updateEntitySubscribeEndpoint(e.EntityID, e.Subscribe.Endpoint, add); err != nil {
		return err
	}
	return nil
}

func (e *SubscribeEntities) AfterDelete(tx *gorm.DB) error {
	if err := deleteCoreSubscription(e.EntityID); err != nil {
		log.Error(err)
		return err
	}
	if err := updateEntitySubscribeEndpoint(e.EntityID, e.Subscribe.Endpoint, reduce); err != nil {
		return err
	}
	return nil
}

func createCoreSubscription(entityID string, topic string) error {
	return coreClient.Subscribe(entityID, topic)
}

func deleteCoreSubscription(entityID string) error {
	return coreClient.UnSubscribe(entityID)
}

type choice uint8

const (
	add choice = iota + 1
	reduce
)

func updateEntitySubscribeEndpoint(entityID, endpoint string, c choice) error {
	separator := ","
	patchData := make([]map[string]interface{}, 0)

	device, err := coreClient.GetEntity(entityID)
	if err != nil {
		return err
	}
	subscribeAddr := endpoint
	switch c {
	case add:
		subscribeAddr = strings.Join([]string{device.Properties.SysField.SubscribeAddr, endpoint}, separator)
	case reduce:
		addrs := strings.Split(device.Properties.SysField.SubscribeAddr, separator)
		validAddresses := make([]string, 0, len(addrs))
		for i := range addrs {
			if addrs[i] != endpoint {
				validAddresses = append(validAddresses, addrs[i])
			}
		}
		subscribeAddr = strings.Join(validAddresses, separator)
	}

	patchData = append(patchData, map[string]interface{}{
		"operator": "replace",
		"path":     "sysField._subscribeAddr",
		"value":    subscribeAddr,
	})
	return coreClient.PatchEntity(entityID, patchData)
}

func NewUndeleteable(content string) error {
	return errors.Wrap(ErrUndeleteable, content)
}
