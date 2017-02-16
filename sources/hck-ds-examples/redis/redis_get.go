package main

import "gopkg.in/redis.v3"
import "fmt"

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println("Connection test : ", pong)
	fmt.Println("Error  : ", err)
	/*
		I have set the following keys from terminal ->
					 name="Rishikesh"
					 count=5

	*/
	var e error
	var myName string

	myName, e = client.Get("name").Result()
	if e == redis.Nil {
		fmt.Println("key -> name does not exist")
		return
	} else if e != nil {
		panic(e)
		return
	} else {
		fmt.Println("name : ", myName)
	}

	var countStr string
	countStr, e = client.Get("count").Result()
	if e != nil {
		panic(e)
		return
	}
	fmt.Println("name : ", countStr)

}
