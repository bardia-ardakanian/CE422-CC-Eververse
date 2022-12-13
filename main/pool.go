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
			conn, err := redis.Dial("tcp", "redis:"+RedisPort)
			if err != nil {
				log.Printf("[ERROR] Fail init Redis: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
}

func ping(conn redis.Conn) {
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		log.Printf("[ERROR] Fail ping Redis conn: %s", err.Error())
		os.Exit(1)
	}
}

func set(key string, val string) error {
	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, val)
	if err != nil {
		log.Printf("[ERROR] Fail set (Key %s, Val %s, Error %s", key, val, err.Error())
		return err
	}

	return nil
}

func get(key string) (string, error) {
	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", key))
	if err != nil {
		log.Printf("[ERROR] Fail get key %s, Error %s", key, err.Error())
		return "", err
	}

	return s, nil
}
