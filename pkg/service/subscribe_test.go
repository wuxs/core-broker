package service

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tkeel-io/core-broker/pkg/model"
	"github.com/tkeel-io/core-broker/pkg/pagination"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestDB(t *testing.T) {
	os.Setenv("DSN", "root:a3fks=ixmeb82a@tcp(192.168.123.9:31815)/core-broker?charset=utf8mb4&parseTime=True&loc=Local")
	model.Setup()

	page := pagination.Page{
		Num:          1,
		Size:         20,
		OrderBy:      "",
		IsDescending: false,
		KeyWords:     "",
		SearchKey:    "",
		Total:        0,
	}

	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(35)}, UserID: "usr-33737945c2b718db4c309d633d2f"}
	validateSubscribeResult := model.DB().First(&subscribe)
	if validateSubscribeResult.RowsAffected == 0 {
		err := errors.Wrap(validateSubscribeResult.Error, "subscribe and user ID mismatch")
		fmt.Println("err:", err)
	}

	var records []model.SubscribeEntities
	result := model.Paginate(&records, page, model.SubscribeEntities{SubscribeID: subscribe.ID})
	if result.Error != nil {
		fmt.Println("err:", result.Error)
	}

	entitiesIDs := make([]string, 0, len(records))
	for i := range records {
		entitiesIDs = append(entitiesIDs, records[i].EntityID)
	}

}
