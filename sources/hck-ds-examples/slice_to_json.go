/*
	{	"Date of creation" => "07 Dec 2016 (Started after 09:51 am)"	}
	{	"Aim of program"   => "Marshalling JSON object"	}
	{	"Coded by"         => "Rishikesh Agrawani"	}
	{	"Go version"	   => "1.7"	}
*/

package main

import "fmt"
import "encoding/json"

type User struct {
	Name string `json:"name"`
	Age  int8   `json:"age`
}
type Users struct {
	Users []User
}

func main() {
	users := []User{User{"Golang", 24}, User{"Python", 23}, User{"Rishikesh", 24}}

	usersSlice, err := json.MarshalIndent(Users{users}, "...", "\t")
	if err != nil {
		fmt.Println("Error...")
	}
	fmt.Println(string(usersSlice))
}

/* OUTPUT FOR ==> admins-MacBook-Pro-3:GoFiles admin$ go run slice_to_json.go

{
...	"Users": [
...		{
...			"name": "Golang",
...			"Age": 24
...		},
...		{
...			"name": "Python",
...			"Age": 23
...		},
...		{
...			"name": "Rishikesh",
...			"Age": 24
...		}
...	]
...}

*/
