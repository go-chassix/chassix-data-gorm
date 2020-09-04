package gormx

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"

	"c5x.io/chassix"
)

func TestDBs(t *testing.T) {
	//defer CloseAllDB()
	// given
	chassix.LoadConfig()
	assert.Equal(t, "root:@tcp(database:3306)/test?parseTime=true", datasource.Databases[0].DSN)
	dbCfg := Databases()
	assert.NotEmpty(t, dbCfg)
	// when
	dbs, _ := DBs()
	// then
	assert.NotNil(t, dbs[0])
	assert.Nil(t, dbs[0].DB().Ping())
}
