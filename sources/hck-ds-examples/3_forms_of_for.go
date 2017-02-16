/*
	{	"Date of creation" => "07 Dec 2016 (Started after 06:45 am)"	}
	{	"Aim of program"   => "To use 3 different forms of for loop"	}
	{	"Coded by"         => "Rishikesh Agrawani"	}
	{	"Go version"		=>	"1.7"	}
*/
package main

import "fmt"

func main() {
	//1st form...................
	i := 1
	for i <= 10 {
		fmt.Printf("%d ", i)
		i++
	}
	println() //To add newline...
	//2nd form...................
	j := 1
	for {
		if j <= 10 {
			fmt.Printf("%d ", j)
			j += 1
			continue
		}
		break
	}
	println()
	//3rd form...................
	for k := 1; k <= 10; k += 1 {
		fmt.Print(k, " ")
	}
	println()
}
