package data

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var __redis_conn *redis.Conn

func RedisInit() error {
	redis_conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return err
	}
	__redis_conn = &redis_conn
	_, err = (*__redis_conn).Do("AUTH", "101921")
	return err
}

func RedisClose() {
	fmt.Println("RedisClose")
	(*__redis_conn).Close()
}

func RedisGetPassword(username string) (string, error) {
	fmt.Println("RedisGetPassword ", username)
	res, err := (*__redis_conn).Do("hget", "plateform_simulator", username)
	if res == nil {
		return "", err
	}
	return redis.String(res, err)
}

func RedisSetPassword(username string, password string) error {
	fmt.Println("RedisSetPassword", username, password)
	_, err := (*__redis_conn).Do("hset", "plateform_simulator", username, password)
	return err
}
