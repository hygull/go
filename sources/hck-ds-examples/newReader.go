/*
	{
		"cretaed_after" : "Sat Dec 10 10:21:51 IST 2016"
		"aim_of_program" : "Taking input from stdin in different ways"
		"coded_by" : "Rishikesh Agrawani"
	}
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	newReader := bufio.NewReader(os.Stdin) //os.Stdin-->newReader

	fmt.Print("Enter your full name :\t")
	name, err := newReader.ReadString('\n')

	if err == nil {
		fmt.Println("You name is ", name)
	} else {
		fmt.Println("Error while reading your name")
	}
}
