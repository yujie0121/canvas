package redis

import (
	"gopkg.in/redis.v5"
	"log"
	"time"
)

var client *redis.Client

func init() {
	newClient()
}

func newClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	log.Print(pong, err)
	// Output: PONG <nil>
}

func GetValue(key string) string {
	val, err := client.Get(key).Result()
	if err == redis.Nil {
		log.Print(key + " does not exists")
	} else if err != nil {
		panic(err)
	}
	return val
}

func SetValue(key string, value string, exp time.Duration) {
	err := client.Set(key, value, exp).Err()
	if err != nil {
		panic(err)
	}

}