/*
	{
		"date_of_creation" => "14 Dec 2016, Wed",
		"aim_of_program"   => "To print sum of all digits of each number of a given list of positive integers using recursion",
		"coded_by"         => "Rishikesh Agrawani",
		"memory"           => "1 submit execution, recursion"
		"Go_version"	   => "1.7",
	}
*/
package main

import "fmt"

/* A function that returns a number which is equal to sum of all the digits in number n */
func getSumOfDigits(n uint) uint {
	if n > 0 {
		return n%10 + getSumOfDigits(n/10)
	}
	return n
}

/* Starter function */
func main() {
	positiveIntsSlice := []uint{1235, 47891, 96354, 0, 2345457, 342846975} /* slice of unsigned integers */

	for _, num := range positiveIntsSlice {
		fmt.Println("The sum of digits of ", num, " is : ", getSumOfDigits(num))
	}
}

/*

The sum of digits of  1235  is  11
The sum of digits of  47891  is  29
The sum of digits of  96354  is  27
The sum of digits of  0  is  0
The sum of digits of  2345457  is  30
The sum of digits of  342846975  is  48

*/
