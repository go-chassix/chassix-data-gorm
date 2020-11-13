package gormx

import (
	"c5x.io/chassix"
	"c5x.io/data-gorm/internal"
	"gorm.io/gorm"
	"sync"
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
		multiDBSource.DBs = make([]*gorm.DB, len(multiCfg))
		//for _, v := range multiCfg {
		//	multiDBSource.DBs = append(multiDBSource.DBs, mustConnectDB(v))
		//}
	})
}

//
//func mustConnectDB(dbCfg *DatabaseConfig) *gorm.DB {
//	log := logx.New().Service("chassix").Category("gorm")
//	dialect := dbCfg.Dialect
//	if "" == dialect {
//		dialect = "mysql"
//	}
//
//	var sqlDB *sql.DB
//	var err error
//	if sqlDB, err = sql.Open(dialect, dbCfg.DSN); err != nil {
//		//todo
//		log.Fatal(err)
//	}
//	var db *gorm.DB
//	if dialect == "postgres" {
//		db, err = gorm.Open(postgres.New(postgres.Config{
//			Conn: sqlDB,
//		}), &gorm.Config{
//			Logger: logger.New(
//				lg.New(os.Stdout, "\r\n", lg.LstdFlags),
//				logger.Config{
//					SlowThreshold: dbCfg.Logger.SlowThreshold,
//					Colorful:      dbCfg.Logger.Colorful,
//					LogLevel:      dbCfg.Logger.Level,
//				},
//			),
//		})
//	}
//	if err != nil {
//		log.Fatalln(err)
//	}
//	if dbCfg.MaxIdle > 0 {
//		sqlDB.SetMaxIdleConns(dbCfg.MaxIdle)
//	}
//	if dbCfg.MaxOpen > 0 && dbCfg.MaxOpen > dbCfg.MaxIdle {
//		sqlDB.SetMaxOpenConns(100)
//	}
//	if dbCfg.MaxLifetime > 0 {
//		sqlDB.SetConnMaxLifetime(time.Duration(dbCfg.MaxLifetime) * time.Second)
//	}
//	return db
//}
