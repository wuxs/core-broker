package model

import (
	"github.com/tkeel-io/core-broker/pkg/util"
	"github.com/tkeel-io/kit/log"
	"gorm.io/gorm"
)

type Subscribe struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	UserID      string `gorm:"index"`
	Endpoint    string `gorm:"index"`
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
	log.Debug("endpoint: %s", endpoint)
}

func destroyRelevant() {
	log.Debug("destroyRelevant")
}

type SubscribeEntities struct {
	SubscribeID uint   `gorm:"index,not null"`
	EntityID    string `gorm:"index,not null"`
	UniqueKey   string `gorm:"index, unique,size:255"`
}
