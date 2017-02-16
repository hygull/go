/*
	{	"Date of creation" => "Wed Dec 21 2016"	}
	{	"Aim of program"   => "4 ways of creating slice, printing their length and capacity" }
	{	"Coded by"         => "Rishikesh Agrawani"	}
	{	"Resource"		   => "https://blog.golang.org/slices"}
	{	"Go version"	   => "1.7"	}
*/
package main

import "fmt"

func main() {

	//1st way of creating slice of integers
	var evenNumbersSlice = []int{24, 56, 0, 42, 98}

	//2nd way of creating slice of integers
	var evenNumbersSlice2 []int = []int{34, 56, 88, 32, 0}

	//3rd way of creating slice of integers (direct way)
	evenNumbersSlice3 := []int{34, 0, 86, 72, 58}

	//4th way of creating slice [using make() built-in]
	var evenNumbersSlice4 []int
	evenNumbersSlice4 = make([]int, 6, 10) /* 1st arg:type of slice,  2nd arg:length of slice, 3rd arg:capacity of slice & this is optional */
	evenNumbersSlice4[0] = 28              /* Assigning a value 28 to the first element of array pointed by slice */
	evenNumbersSlice4[1] = 64
	evenNumbersSlice4[2] = 32

	//Printing contents of slices
	fmt.Println(evenNumbersSlice)
	fmt.Println(evenNumbersSlice2)
	fmt.Println(evenNumbersSlice3)
	fmt.Println(evenNumbersSlice4) /* 4th to 6th elements will have the zero values of int type*/

	fmt.Println("Length of 4th slice   : ", len(evenNumbersSlice4))
	fmt.Println("Capacity of 4th slcie : ", cap(evenNumbersSlice4))
}

/* OUTPUT:-

[24 56 0 42 98]
[34 56 88 32 0]
[34 0 86 72 58]
[28 64 32 0 0 0]
Length of 4th slice   :  6
Capacity of 4th slcie :  10

*/
