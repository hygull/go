/*
  Date of creation : 10/12/2016.
  Aim of program   : To use Fields() to split the string into a list words.
  Coded by         : Rishikesh Agrawani.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	str := "We should always respect the time"
	listOfWords := strings.Fields(str)
	fmt.Print(str, " -> ", listOfWords) //Displaying a  slice containing all the words

	str1 := "Nice @ character is of 8 bits in C"
	listOfWords = strings.Fields(str)
	fmt.Print(str1, " -> ", listOfWords) //Displaying a  slice containing all the words
}
