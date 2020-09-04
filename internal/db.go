package internal

import (
	"errors"
	"github.com/jinzhu/gorm"
	"sync"
)

type MultiDBSource struct {
	Lock sync.RWMutex
	DBs  []*gorm.DB
}

var (
	ErrNoDatabaseConfiguration = errors.New("there isn't any database setting in the configuration file")
)

//Size get db connection size
func (s MultiDBSource) Size() int {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.DBs)
}
