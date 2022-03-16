package model

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithoutDBConnectionAndDBName(t *testing.T) {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	connection, dbName := withoutDBConnectionAndDBName(dsn)

	slashIndex := strings.LastIndex(dsn, "/")
	_dsn := dsn[:slashIndex+1]
	items := strings.Split(dsn[slashIndex+1:], "?")
	expectDBNName := items[0]
	expectConnection := fmt.Sprintf("%s?%s", _dsn, items[1])
	assert.Equal(t, expectDBNName, dbName)
	assert.Equal(t, expectConnection, connection)
}
