package model

import (
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

func (s Subscribe) AfterDelete(tx *gorm.DB) error {
	if s.IsDefault {
		return NewUndeleteable("this is default subscribe")
	}
	destroyEndpoint(s.Endpoint)
	destroyRelevant()
	return nil
}

func destroyEndpoint(endpoint string) {
	// TODO: destroy endpoint
	log.Debug("endpoint: %s", endpoint)
}

func destroyRelevant() {
	// TODO: destroy relevant
	log.Debug("destroyRelevant")
}

type SubscribeEntities struct {
	SubscribeID uint   `gorm:"index,not null"`
	EntityID    string `gorm:"index,not null"`
	UniqueKey   string `gorm:"index, unique,size:255"`
}

func (e *SubscribeEntities) AfterCreate(tx *gorm.DB) error {
	if err := createCoreSubscription(e.EntityID, e.UniqueKey); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func createCoreSubscription(entityID string, topic string) error {
	return coreClient.Subscribe(entityID, topic)
}

func NewUndeleteable(content string) error {
	return errors.Wrap(ErrUndeleteable, content)
}
