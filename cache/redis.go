package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
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
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.addr"), viper.GetString("redis.port")),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		log.Printf("cannot ping postgres database: %s\n", err.Error())
		return err
	}

	log.Println("database redis: Connected!")
	rdb = redisClient
	return nil
}
