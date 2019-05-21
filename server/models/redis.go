package models

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

// InitializeCache creates redis client
func InitializeCache() *redis.Client {
	dbNumber, err := strconv.Atoi(os.Getenv("CACHE_DB"))
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("CACHE_ADDRS"),
		Password: os.Getenv("CACHE_PASSWORD"), // no password set
		DB:       dbNumber,                    // use default DB
	})

	_, err = client.Ping().Result()
	if err != nil {
		fmt.Println("here")
		log.Fatal(err)
	}

	return client
}

// RunCacheSeeds Sets default response to redis
func RunCacheSeeds(c *redis.Client) {
	c.FlushAll()
	c.Set("-74.0060:40.7128", string(DefaultLocation), 0)
}
