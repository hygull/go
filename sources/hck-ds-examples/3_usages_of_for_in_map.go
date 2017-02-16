/*
	{	"Date of creation" 	=> 	"07 Dec 2016 (Started after 07:35 am)"	}
	{	"Aim of program"   	=> 	"To use for loop on maps in 3 ways"	}
	{	"Coded by"         	=> 	"Rishikesh Agrawani"	}
	{	"Go version"		=>	"1.7"	}
*/
package main

import "fmt"

func main() {
	myDetailsMap := make(map[string]interface{})
	myDetailsMap["Name"] = "Rishikesh Agrawani"
	myDetailsMap["Age"] = 24
	myDetailsMap["Salary"] = 150000.5

	fmt.Println(myDetailsMap)
	//Printing key and values both......
	fmt.Println("\nKeys and values...")
	for key, value := range myDetailsMap {
		fmt.Printf("%s %v\n", key, value)
	}

	//Printing keys only................
	fmt.Println("\nKeys...")
	for key := range myDetailsMap {
		fmt.Printf("%s\n", key)
	}

	//Printing values only...............
	fmt.Println("\nValues...")
	for _, value := range myDetailsMap { //Using blank identifier _
		fmt.Println(value)
	}
}

/* OUTPUT:-

map[Age:24 Salary:150000.5 Name:Rishikesh Agrawani]

Keys and values...
Salary 150000.5
Name Rishikesh Agrawani
Age 24

Keys...
Name
Age
Salary

Values...
Rishikesh Agrawani
24
150000.5

*/
