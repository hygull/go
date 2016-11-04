package main

import "fmt"

import "github.com/13x/jsoncfgo"

func main() {
	var Users jsoncfgo.Obj
	Users = jsoncfgo.Load("./json/users.json")
	user1 := Users.OptionalObject("joesample")
	fmt.Println("Firstname : ", user1["firstname"])
}
