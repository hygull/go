/*
	{
		"date_of_creation" => "14 Dec 2016, Wed",
		"aim_of_program"   => "To print the front most digit of integers",
		"coded_by"         => "Rishikesh Agrawani",
		"memory"           => "1 submit execution, tail recursion"
		"Go_version"	   => "1.7",
	}
*/

package main

import "fmt"

func getFrontDigit(num int) int {
	if num/10 == 0 { //If number has only 1 digit then return it
		return num
	}
	return getFrontDigit(num / 10) /*Passing the number formed after eliminating the last digit*/
}

func main() {
	numsSlice := []int{123, -765, 321, 231, -5165, -34, 569}

	for _, num := range numsSlice {
		n := getFrontDigit(num) //In case of -34 the value of n will be -34 (for -ve integers)
		if n < 0 {
			n = -n
		}
		fmt.Println("The front most digit of ", num, " is\t: ", n)
	}
}

/* OUTPUT:-

The front most digit of  123  is	:  1
The front most digit of  -765  is	:  7
The front most digit of  321  is	:  3
The front most digit of  231  is	:  2
The front most digit of  -5165  is	:  5
The front most digit of  -34  is	:  3
The front most digit of  569  is	:  5

*/
