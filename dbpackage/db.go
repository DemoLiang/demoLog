package dbpackage

import (
	"fmt"
	"github.com/DemoLiang/gopackage/logpackage"
	"github.com/jinzhu/gorm"
	"sync"
)

//init db pool
func InitDBPool(service Service, env string, mysqlConfig MysqlConfig, installTable []interface{}) {
	var onceLock sync.Once
	for {
		if db, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/gorm?charset=utf8&parseTime=True&loc=Local",
			mysqlConfig.User, mysqlConfig.Password, mysqlConfig.IP, mysqlConfig.Port)); err == nil {
			db.AutoMigrate(installTable)
			dbMap[service].Mutex.Lock()
			dbMap[service].master = db
			dbMap[service].Mutex.Unlock()
		} else {
			demoLog.Error("init db error:%v", err)
		}
		onceLock.Do(func() {
			//TODO sync for other
		})
		break
	}
}
