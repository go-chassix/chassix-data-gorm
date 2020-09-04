package gormx

import (
	"sync"
	"time"

	"github.com/jinzhu/gorm"

	"c5x.io/chassix"
	"c5x.io/logx"

	"c5x.io/data-gormx/internal"
)

func init() {
	chassix.Register(&chassix.Module{Name: chassix.ModuleDataGorm, ConfigPtr: datasource})
}

var (
	multiDBSource *internal.MultiDBSource
	initOnce      sync.Once
)

func initMultiDBSource() {
	initOnce.Do(func() {
		multiCfg := datasource.Databases
		multiDBSource = new(internal.MultiDBSource)
		multiDBSource.Lock.Lock()
		defer multiDBSource.Lock.Unlock()
		for _, v := range multiCfg {
			multiDBSource.DBs = append(multiDBSource.DBs, mustConnectDB(v))
		}
	})
}

func mustConnectDB(dbCfg *DatabaseConfig) *gorm.DB {
	log := logx.New().Service("chassix").Category("gorm")
	dialect := dbCfg.Dialect
	if "" == dialect {
		dialect = "mysql"
	}
	db, err := gorm.Open(dialect, dbCfg.DSN)
	if err != nil {
		log.Fatalln(err)
	}
	db.LogMode(dbCfg.ShowSQL)

	if dbCfg.MaxIdle > 0 {
		db.DB().SetMaxIdleConns(dbCfg.MaxIdle)
	}
	if dbCfg.MaxOpen > 0 && dbCfg.MaxOpen > dbCfg.MaxIdle {
		db.DB().SetMaxOpenConns(100)
	}
	if dbCfg.MaxLifetime > 0 {
		db.DB().SetConnMaxLifetime(time.Duration(dbCfg.MaxLifetime) * time.Second)
	}
	return db
}
