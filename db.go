package gormx

import (
	"errors"
	"sync"
	"time"

	"github.com/jinzhu/gorm"

	"c5x.io/chassix"
	"c5x.io/chassix/config"
	"c5x.io/logx"
)

func init() {
	chassix.Register(&chassix.Module{Name: config.KeyDatasourceConfig, ConfigPtr: datasource})
}

type MultiDBSource struct {
	lock sync.RWMutex
	dbs  []*gorm.DB
}

var (
	ErrNoDatabaseConfiguration = errors.New("there isn't any database setting in the configuration file")
)

var (
	multiDBSource *MultiDBSource
	initOnce      sync.Once
)

func initMultiDBSource() {
	initOnce.Do(func() {
		multiCfg := datasource.Databases
		multiDBSource = new(MultiDBSource)
		multiDBSource.lock.Lock()
		defer multiDBSource.lock.Unlock()
		for _, v := range multiCfg {
			multiDBSource.dbs = append(multiDBSource.dbs, mustConnectDB(v))
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

//DB get the default(first) *Db connection
func DB() (*gorm.DB, error) {
	if dbs, err := DBs(); nil != err {
		return nil, err
	} else {
		return dbs[0], nil
	}
}

//DBs get all database connections
func DBs() ([]*gorm.DB, error) {
	if initMultiDBSource(); 0 == multiDBSource.Size() {
		return nil, ErrNoDatabaseConfiguration
	}
	return multiDBSource.dbs, nil
}

//Close close all db connection
func CloseAllDB() error {
	if 0 == multiDBSource.Size() {
		return ErrNoDatabaseConfiguration
	}
	for _, v := range multiDBSource.dbs {
		if err := v.Close(); nil != err {
			return err
		}
	}
	return nil
}

//Size get db connection size
func (s MultiDBSource) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.dbs)
}
