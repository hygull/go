package main

import "io/ioutil"
import "fmt"
import "encoding/json"

func main() {

	type User struct {
		Firstname string
		//Lastname string  /*In this case it won't be visible in JSON data*/
	}

	usersMap := make(map[string]User)

	usersBytes, err := ioutil.ReadFile("./json/users.json")

	if err != nil {
		fmt.Println("Either the json structure is wrong or there's an error while reading it")
		return
	} else {
		err = json.Unmarshal(usersBytes, &usersMap)
		fmt.Printf("%v\n\n", usersMap)
	}
}
