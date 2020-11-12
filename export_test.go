package gormx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	_ "gorm.io/gorm/dialects/mysql"
	_ "gorm.io/gorm/dialects/postgres"
	_ "gorm.io/gorm/dialects/sqlite"

	"c5x.io/chassix"
)

func TestDBs(t *testing.T) {
	//defer CloseAllDB()
	// given
	chassix.Init()
	assert.Equal(t, "root:@tcp(database:3306)/test?parseTime=true", datasource.Databases[0].DSN)
	dbCfg := datasource.Databases
	assert.NotEmpty(t, dbCfg)
	// when
	dbs := DBs()
	assert.NotEmpty(t, dbs)
	assert.NotNil(t, dbs[0])
	assert.Nil(t, dbs[0].DB().Ping())

	assert.NotNil(t, DB())
}
