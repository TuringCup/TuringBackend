package cache

import (
	"time"

	"github.com/TuringCup/TuringBackend/config"
	redigo "github.com/gomodule/redigo/redis"
)

var pool *redigo.Pool

func InitCache() {
	address := config.Conf.Redis.Host + ":" + config.Conf.Redis.Port
	pwd := config.Conf.Redis.Password
	pool = &redigo.Pool{
		MaxIdle:     5,
		IdleTimeout: 360 * time.Second,
		MaxActive:   8,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", address)
			if err != nil {
				return nil, err
			}
			if pwd != "" {
				if _, err := c.Do("AUTH", pwd); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
