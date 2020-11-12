package gormx

import (
	"database/sql"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"c5x.io/chassix"
	"c5x.io/logx"

	"c5x.io/data-gorm/internal"
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

	var sqlDB *sql.DB
	var err error
	if sqlDB, err = sql.Open(dialect, dbCfg.DSN); err != nil {
		//todo
		log.Fatal(err)
	}
	var db *gorm.DB
	if dialect == "postgres" {
		db, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{
			Logger: NewLogger(&dbCfg.Logger),
		})
	}
	if err != nil {
		log.Fatalln(err)
	}
	if dbCfg.MaxIdle > 0 {
		sqlDB.SetMaxIdleConns(dbCfg.MaxIdle)
	}
	if dbCfg.MaxOpen > 0 && dbCfg.MaxOpen > dbCfg.MaxIdle {
		sqlDB.SetMaxOpenConns(100)
	}
	if dbCfg.MaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(dbCfg.MaxLifetime) * time.Second)
	}
	return db
}
