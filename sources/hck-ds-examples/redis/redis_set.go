package main

import "gopkg.in/redis.v3"
import "fmt"
import "time"

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println("Connection test : ", pong)
	fmt.Println("Error  : ", err)

	err = client.Set("language", "Golang", 1*time.Minute).Err()
	if err != nil {
		fmt.Println("Didn't set the key named 'language' on REDIS...", err)
		return
	}
	fmt.Println("Successfully set the key named 'lang' on REDIS")

	fmt.Println("Now run -> get lang....on the redis-cli terminal to get the key")
	/*
		127.0.0.1:6379> get lang
		"Golang"
		127.0.0.1:6379>
	*/
}
