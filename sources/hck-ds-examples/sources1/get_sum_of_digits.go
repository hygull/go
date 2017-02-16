/*
  {
    "date_of_creation" => "19 Dec 2016, Mon",
    "aim_of_program"   => "To get sum of all the digits in positive numbers",
    "coded_by"         => "Rishikesh Agrawani",
    "Go_version"       => "1.7",
  }
*/
package main

import "fmt"

func main() {
	intNumsArr := [5]int{1346, 678844, 575632, 2433434, 1}
	for _, num := range intNumsArr {
		sum := 0
		fmt.Print("The sum of all digits of ", num, " is\t: ")
		for num != 0 {
			sum += num % 10
			num /= 10
		}
		fmt.Println(sum)
	}
}

/*

The sum of all digits of 1346 is  : 14
The sum of all digits of 678844 is  : 37
The sum of all digits of 575632 is  : 28
The sum of all digits of 2433434 is : 23
The sum of all digits of 1 is : 1

*/
