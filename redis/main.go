package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type User struct {
	Name   string
	Age    int64
	Gender string
}

func main() {
	// TODO

	data := User{
		Name:   "andi",
		Age:    time.Now().UnixNano(),
		Gender: "Men",
	}
	j, _ := json.Marshal(data)
	var redisHost = "localhost:6379"
	var redisPassword = "endi"

	rdb := newRedisClient(redisHost, redisPassword)
	fmt.Println("redis client initialized")

	key := data.Name
	ttl := time.Duration(60) * time.Second

	// store data using SET command
	op1 := rdb.Set(context.Background(), key, j, ttl)
	if err := op1.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
	log.Println("set operation success")

	// get data
	op2 := rdb.Get(context.Background(), key)
	if err := op2.Err(); err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}
	res, err := op2.Result()
	if err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}
	log.Println("get operation success. result:", res)

	// ...

}

// ...
func newRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}
