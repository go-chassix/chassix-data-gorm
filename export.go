package gormx

import (
	"github.com/jinzhu/gorm"

	"c5x.io/data-gormx/internal"
)

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
		return nil, internal.ErrNoDatabaseConfiguration
	}
	return multiDBSource.DBs, nil
}

//Close close all db connection
func CloseAllDB() error {
	if 0 == multiDBSource.Size() {
		return internal.ErrNoDatabaseConfiguration
	}
	for _, v := range multiDBSource.DBs {
		if err := v.Close(); nil != err {
			return err
		}
	}
	return nil
}
