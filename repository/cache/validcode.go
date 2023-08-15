package cache

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

func GenerateValidCode() (code string, err error) {
	rand.Seed(time.Now().Unix())
	code = strconv.Itoa(rand.Int() % 1000000)
	c := pool.Get()
	defer c.Close()
	key := "validcode_" + code
	_, err = c.Do("SET", key, "1")
	_, err = c.Do("EXPIRE", key, "300")
	return code, err
}

func CheckValidCode(code string) (err error) {
	c := pool.Get()
	defer c.Close()
	key := "validcode_" + code
	exist, err := redis.Bool(c.Do("EXISTS", key))
	if exist {
		c.Do("DEL", key)
		return nil
	} else {
		err = errors.New("Wrong Valid Code")
		return
	}
}
