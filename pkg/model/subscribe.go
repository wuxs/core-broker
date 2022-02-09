package model

import (
	"github.com/tkeel-io/core-broker/pkg/util"
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

type SubscribeEntities struct {
	SubscribeID uint   `gorm:"index"`
	EntityID    string `gorm:"index"`
	UniqueKey   string `gorm:"index, unique,size:255"`
}
