package dbpackage

import (
	"sync"
	"github.com/garyburd/redigo/redis"
)

type RedisConfig struct {
	Ip          string `json:"ip"`
	Port        int    `json:"port"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	Retry       int    `json:"retry"`
	TTL         int    `json:"ttl"`
	Prefix      string `json:"prefix"`
	MaxIdle     int    `json:"max_idle"`
	MaxActive   int    `json:"max_active"`
	IdleTimeout int    `json:"idle_timeout"`
	ConnTimeout int    `json:"conn_timeout"`
}


type DemoCacheS struct {
	sync.Mutex
	masterCache *redis.Pool
}

type Service int