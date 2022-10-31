package db

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

// Init - redis init
func Init() error {
	var err error
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	dsn := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	rdb = redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       db,                          // use default DB
	})

	_, err = rdb.Ping(rdb.Context()).Result()
	if err != nil {
		return errors.New("Redis Connection Error")
	}

	return nil
}

// RedisManager - return db connection
func RedisManager() *redis.Client {
	return rdb
}
