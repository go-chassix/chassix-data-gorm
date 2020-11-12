package gormx

import (
	"testing"

	"github.com/stretchr/testify/assert"

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

	assert.NotNil(t, DB())
}
