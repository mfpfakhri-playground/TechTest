package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

var rdb *redis.Client

// GetDB get db
func GetDB() *redis.Client {
	if rdb != nil {
		return rdb
	}
	createConnection()
	return rdb
}

// CreateConnection open new connection Postgres
func createConnection() error {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		log.Printf("could not ping postgres database: %s\n", err.Error())
		return err
	}

	log.Println("database redis: Connected!")
	rdb = redisClient
	return nil
}
