package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func main() {
	var ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "10.20.152.240:6379",
		Password: "V3l}yGP:*%KY",
		DB:       0,
	})
	res, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	val, err := client.Get(ctx, "tokenSecret").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("tokenSecret:", val)
}
