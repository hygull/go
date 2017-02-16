/*
*	Date of creation : 20/10/2016.
* 	Aim of program   : To use maps to represent it as JSON object.
* 	Coded by		 : Rishikesh Agrawani.
 */
package main

import "encoding/json"
import "fmt"

func main() {
	fmt.Println("Users data without (omitempty)\n")
	//Defining structure locally
	type User struct {
		Name   string `json:"name"`
		Age    int8   `json:"age,omitempty"`
		Branch string `json:"branch,omitempty"`
	}
	usersMap := make(map[string]User)
	usersMap["User1"] = User{"Rishikesh Agrawani", 0, ""}
	usersMap["User2"] = User{"Hemkesh Agrawani", 22, "CBZ"}

	usersBytes, _ := json.MarshalIndent(usersMap, "Users| ", "\t")
	fmt.Println(string(usersBytes))

	//Defining structure locally
	fmt.Println("Books data without(omitempty)\n")
	type Book struct {
		Name   string  `json:"name"`
		Price  float32 `json:"price"`
		Author string  `json:"author`
	}
	booksMap := make(map[string]Book)
	booksMap["Book1"] = Book{"Have a taste of Golang", 350, "Rishikesh Agrawani"}
	booksMap["Book2"] = Book{"Living a life in critical situation", 360, "Hemkesh Agrawani"}
	booksMap["Book3"] = Book{"", 0, "Malinikesh Agrawani"}
	booksBytes, _ := json.MarshalIndent(booksMap, "Books| ", "\t\t")
	fmt.Println(string(booksBytes))
}
