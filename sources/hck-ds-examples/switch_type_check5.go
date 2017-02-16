/*
	{
		"date_of_creation" => "23 Dec 2016, Friday",
		"aim_of_program"   => "To sort a slice of integers & floats using ds package",
		"coded_by"         => "Rishikesh Agrawani",
		"Go_version"	   => "1.7.1",
	}
*/
package main

import (
	"ds"
	"fmt"
)

func main() {
	//Creating a slice of integers...int8/int32/int64 are also allowed
	wholeNums := []int{16, 13, 34, 65, 78, 43, 1, 98, 0}
	fmt.Println("Before applying bubble sort : ", wholeNums)
	ds.BubbleSort(wholeNums)
	fmt.Println("After  applying bubble sort : ", wholeNums)
	fmt.Println()

	//Creating a slice of float32
	floatsList := []float32{16.5999, 13.21, -34.459, 65.99, 78.09, 23.32, 1.43, 56, 0.9}
	fmt.Println("Before applying bubble sort : ", floatsList)
	ds.BubbleSort(floatsList)
	fmt.Println("After  applying bubble sort : ", floatsList)
	fmt.Println()

	//Creating a slice of float64
	floatsList2 := []float64{16.5999, 0.0, 1653.21, 3454544.459, -0.8, 655675.99, 78.09, 23.32, 1.43, 98, -0.9}
	fmt.Println("Before applying bubble sort : ", floatsList2)
	ds.BubbleSort(floatsList2)
	fmt.Println("After  applying bubble sort : ", floatsList2)
	fmt.Println()

	//Creating an empty slice of float64
	floatsList3 := []float64{} //Enmpty slice
	fmt.Println("Before applying bubble sort : ", floatsList3)
	ds.BubbleSort(floatsList3)
	fmt.Println("After  applying bubble sort : ", floatsList3)
	fmt.Println()

	//Creating a slice of strings
	names := []string{"hygull", "rob", "robert", "ken"}
	fmt.Println("Before applying bubble sort : ", names)
	ds.BubbleSort(names)
	fmt.Println("After  applying bubble sort : ", names)
	fmt.Println()
}
