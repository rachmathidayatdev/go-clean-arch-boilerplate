package library

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

//SessionInit function
func SessionInit() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
		// Addr:     "localhost:6379",
		Password: fmt.Sprintf("%s", os.Getenv("REDIS_PASS")), // no password set
		DB:       0,                                          // use default DB
	})

	_, err := client.Ping().Result()

	if err != nil {
		fmt.Println(err)
	}

	return client
} /* end */
