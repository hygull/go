/*
	@Date of creation : 24 November 2016.
	@Aim of program   : To check the availability status of keys in MAP.
	@Go version       : 1.7.1 [go version command(on MAC )prints -> go version go1.7.1 darwin/amd64]
	@Coded by		  : Rishikesh Agrawani.
*/
package main

import "fmt"

func main() {
	detailsMap := make(map[string]string) //Defining a map

	detailsMap["name"] = "Rishikesh Agrawani" //Adding one key-value pair to map
	detailsMap["branch"] = "CSE"              //Adding another key-value pair to map

	val, isKeyPresent := detailsMap["name"] //Getting the value of key 'name', that exists
	fmt.Println(val, isKeyPresent)          //Printing the fetched value and its availability information

	val, isKeyPresent = detailsMap["branch"] //Getting the value of key 'branch', that exists
	fmt.Println(val, isKeyPresent)           //Printing the fetched value and its availability information

	val, isKeyPresent = detailsMap["age"] //Getting the value of key 'age', that does not exist
	fmt.Println(val, isKeyPresent)        //Printing the fetched value and its availability information
}

/*COMMAND:-
			go run map_key_availability_check.go
  OUPUT:-
			Rishikesh Agrawani true
			CSE true
			 false
*/
