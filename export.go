package gormx

import (
	"gorm.io/gorm"

	"c5x.io/data-gorm/internal"
)

//DB get the default(first) *Db connection
func DB() *gorm.DB {
	if dbs := DBs(); dbs == nil || len(dbs) == 0 {
		return nil
	} else {
		return dbs[0]
	}
}

//DBs get all database connections
func DBs() []*gorm.DB {
	if initMultiDBSource(); 0 == multiDBSource.Size() {
		return nil
	}
	return multiDBSource.DBs
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
