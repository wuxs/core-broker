package model

import (
	"github.com/tkeel-io/core-broker/pkg/pagination"
	"os"
	"sync"

	"github.com/tkeel-io/kit/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsnFromOSEnvKey = "DSN"
	defaultMySQLDSN = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
)

type WhereOptions func() (query interface{}, args interface{})

var _once sync.Once
var db *gorm.DB

func SetUp(dsn string) error {
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db = connection
	return db.AutoMigrate(&Subscribe{}, &SubscribeEntities{})
}

func DB() *gorm.DB {
	if db == nil {
		_once.Do(func() {
			dsn := defaultMySQLDSN
			dsn = os.Getenv(dsnFromOSEnvKey)
			if err := SetUp(dsn); err != nil {
				log.Error("SetUp DB Error: ", err)
				return
			}
		})
	}
	return db
}

func Paginate(find interface{}, page pagination.Page, where interface{}, args ...interface{}) *gorm.DB {
	conditions := page.SearchCondition()
	if conditions == nil {
		return DB().Where(where, args...).Limit(int(page.Limit())).Offset(int(page.Offset())).Find(find)
	}
	return DB().Where(where, args...).Where(conditions).Limit(int(page.Limit())).Offset(int(page.Offset())).Find(find)
}

func ListAll(find interface{}, where interface{}, args ...interface{}) *gorm.DB {
	return DB().Where(where, args...).Find(find)
}

func Count(count *int64, where interface{}, args ...interface{}) *gorm.DB {
	return DB().Where(where, args).Count(count)
}
