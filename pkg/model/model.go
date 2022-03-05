package model

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/tkeel-io/core-broker/pkg/core"
	"github.com/tkeel-io/core-broker/pkg/pagination"
	"github.com/tkeel-io/kit/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	// schema like: "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsnFromOSEnvKey = "DSN"
	amqpServer      = "AMQP_SERVER"
)

type WhereOptions func() (query interface{}, args interface{})

var (
	_once      sync.Once
	db         *gorm.DB
	coreClient *core.Client

	AMQPServerAddr = "amqp://localhost:3172"
)

func Setup() error {
	var err error
	coreClient, err = core.NewCoreClient()
	if err != nil {
		log.Fatal(err)
	}

	amqpServerStr := os.Getenv(amqpServer)
	if amqpServerStr != "" {
		AMQPServerAddr = amqpServerStr
	}

	dsn := os.Getenv(dsnFromOSEnvKey)
	slashIndex := strings.LastIndex(dsn, "/")
	_dsn := dsn[:slashIndex+1]
	items := strings.Split(dsn[slashIndex+1:], "?")
	dbName := items[0]
	fmt.Println(_dsn)
	fmt.Println(dbName)
	_dsn = fmt.Sprintf("%s?%s", _dsn, items[1])
	db, err := gorm.Open(mysql.Open(_dsn), nil)
	if err != nil {
		log.Fatal(err)
	}

	createSQL := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4;",
		dbName,
	)

	err = db.Exec(createSQL).Error
	if err != nil {
		log.Fatal(err)
	}

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = connection
	return db.AutoMigrate(&Subscribe{}, &SubscribeEntities{})
}

func MakeAMQPAddress(endpoint string) string {
	return AMQPServerAddr + "/" + endpoint
}

func DB() *gorm.DB {
	if db == nil {
		_once.Do(func() {
			if err := Setup(); err != nil {
				log.Error("Setup DB Error: ", err)
				return
			}
		})
	}
	return db
}

func Paginate(find interface{}, page pagination.Page, where interface{}, args ...interface{}) *gorm.DB {
	conditions, fields := page.SearchCondition()

	if conditions != nil && fields == nil {
		return DB().Where(where, args...).Where(conditions).Limit(int(page.Limit())).Offset(int(page.Offset())).Find(find)
	}

	if conditions != nil && fields != nil {
		return DB().Select(fields).Where(where, args...).Where(conditions).Limit(int(page.Limit())).Offset(int(page.Offset())).Find(find)
	}

	if conditions == nil && fields != nil {
		return DB().Select(fields).Where(where, args...).Limit(int(page.Limit())).Offset(int(page.Offset())).Find(find)
	}

	return DB().Where(where, args...).Limit(int(page.Limit())).Offset(int(page.Offset())).Find(find)
}

func ListAll(find interface{}, where interface{}, args ...interface{}) *gorm.DB {
	return DB().Where(where, args...).Find(find)
}

func Count(count *int64, model interface{}, where interface{}, args ...interface{}) *gorm.DB {
	return DB().Model(model).Where(where, args).Count(count)
}
