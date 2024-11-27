package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	rdb := connRdb()
	ctx := context.Background()
	err := rdb.Set(ctx, "session_id:user_id", "session_id", 5*time.Second).Err()
	if err != nil {
		panic(err)
	}
	session_id, err := rdb.Get(ctx, "session_id:user_id").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(session_id)
}

func connRdb() *redis.Client {
	//redis-cli
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return rdb
}
