package dbpackage

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"fmt"
	"github.com/DemoLiang/gopackage/logpackage"
	"sync"
)

var cacheMap = make(map[Service]*DemoCacheS)

//init pool
func InitRedisPool(config *RedisConfig) (*redis.Pool,error) {
	p := &redis.Pool{
		MaxIdle:config.MaxIdle,
		MaxActive:config.MaxActive,
		IdleTimeout:time.Duration(config.IdleTimeout)*time.Second,
		Dial: func() (redis.Conn, error) {
			t:=time.Duration(config.ConnTimeout)*time.Second
			c,err := redis.DialTimeout("tcp",fmt.Sprintf("%s:%d",config.Ip,config.Port),t,t,t)
			if err != nil{
				demoLog.Error("%v",err)
				return nil,err
			}
			if config.Password != ""{
				if _,err := c.Do("auth",config.Password);err!=nil{
					c.Close()
					demoLog.Error("%v",err)
					return nil,err
				}
			}
			return c,err
		},
	}
	return p, nil
}

//init master cache
func masterCache(service Service,env string,config *RedisConfig){
	var onceLock sync.Once
	for {
		if p,err := InitRedisPool(config);err!=nil{
			cacheMap[service].Mutex.Lock()
			cacheMap[service].masterCache = p
			cacheMap[service].Mutex.Unlock()
		}else {
			demoLog.Error("%v",err)
		}
		onceLock.Do(func() {
			//TODO sync other
		})
		break
	}
}