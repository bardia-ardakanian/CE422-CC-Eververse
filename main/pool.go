package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
)

var pool *redis.Pool

func initRedis() {
	// init redis connection pool
	initPool()
}

func initPool() {
	pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", RedisHost+":"+RedisPort)
			if err != nil {
				log.Printf("[ERROR] Fail init Redis: %s\n", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
}

func set(key string, val string) error {
	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	log.Printf(" - [TRY] Set (key %s, val %s)", key, val)
	_, err := conn.Do("SET", key, val, "EX", KeyExp)
	if err != nil {
		log.Printf(" - [ERROR] Fail set (Key %s, Val %s, EX %v, Error %s\n", key, val, KeyExp, err.Error())
		return err
	}

	log.Printf(" - [INFO] SET %v:%v\n", key, val)
	return nil
}

func get(key string) (string, error) {
	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", key))
	if err != nil {
		log.Printf(" - [ERROR] Fail get key %s, Error %s\n", key, err.Error())
		return "", err
	}

	log.Printf(" - [INFO] GET %v\n", s)
	return s, nil
}
