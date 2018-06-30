package dbpackage

import (
	"sync"
	"github.com/jinzhu/gorm"
)

type dbS struct {
	sync.Mutex
	master *gorm.DB
}

var dbMap = make(map[Service]*dbS)

type MysqlConfig struct {
	IP        string
	Port      int
	DB        string
	User      string
	Password  string
	MaxIdle   int
	MaxActive int
}