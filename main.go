package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

type Person struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
}

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ping)

	personID := uuid.NewString()
	jsonString, _ := json.Marshal(&Person{
		Id:         personID,
		Name:       "John Smith",
		Age:        31,
		Occupation: "DB Engineer",
	})

	personKey := fmt.Sprintf("person:%s", personID)
	err = client.Set(context.Background(), personKey, jsonString, 1000*time.Second).Err()
	if err != nil {
		fmt.Println("Failed to set value", err.Error())
		return
	} else {
		fmt.Println("Set Success")
	}

	val, err := client.Get(context.Background(), personKey).Result()
	if err != nil {
		fmt.Println("Failed to get value", err.Error())
		return
	}

	fmt.Printf("value - %v  (type: %T)\n", val, val)
}
