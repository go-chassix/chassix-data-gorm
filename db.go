package gormx

import (
	"log"
	"sync"

	"c5x.io/chassix"
	"c5x.io/data-gorm/internal"
	"gorm.io/gorm"
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
		multiDBSource.DBs = make([]*gorm.DB, 0)
		for _, v := range multiCfg {
			db, err := driver[v.Dialect].Connect(v)
			if err != nil {
				log.Fatalf("connect db failed: %s\n", err)
				continue
			}
			multiDBSource.DBs = append(multiDBSource.DBs, db)
		}
	})
}

type DatabaseProvider interface {
	Connect(config *DatabaseConfig) (*gorm.DB, error)
}
