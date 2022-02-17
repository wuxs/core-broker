package model

import (
	"fmt"
	"github.com/tkeel-io/core-broker/pkg/util"
	"github.com/tkeel-io/kit/log"
	"gorm.io/gorm"
)

type Subscribe struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	UserID      string `gorm:"index, not null"`
	TenantID    string `gorm:"index, not null"`
	Endpoint    string `gorm:"index, not null"`
}

func (s *Subscribe) BeforeCreate(tx *gorm.DB) error {
	if s.Endpoint == "" {
		s.Endpoint = util.GenerateSubscribeEndpoint()
	}
	return nil
}

func (s Subscribe) AfterDelete(tx *gorm.DB) error {
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

func (receiver *SubscribeEntities) BeforeCreate(tx *gorm.DB) error {
	if receiver.UniqueKey == "" {
		receiver.UniqueKey = fmt.Sprintf("%d:%s", receiver.SubscribeID, receiver.EntityID)
	}
	return nil
}

type SubscribeUsers struct {
	SubscribeID uint   `gorm:"index,not null"`
	UserID      string `gorm:"index,not null"`
	UniqueKey   string `gorm:"index, unique,size:255"`
}

func (receiver *SubscribeUsers) BeforeCreate(tx *gorm.DB) error {
	if receiver.UniqueKey == "" {
		receiver.UniqueKey = fmt.Sprintf("%d:%s", receiver.SubscribeID, receiver.UserID)
	}
	return nil
}
