package model

import (
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
