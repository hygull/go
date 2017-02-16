/*
	{
		"cretaed_after" : "Mon Dec 12 21:13:14 IST 2016"
		"aim_of_program": "Creating slices in 3 different ways"
		"coded_by" 		: "Rishikesh Agrawani"
	}
*/
package main

import "fmt"

func main() {
	/* 1ST way of creating slice of integers */
	intsSlice := []int{5, 12, -24, 0, 38, 19, 455, 1729, 1992, -67}

	/* 2ND way of creating slice of integers */
	evensSlice := make([]int, 0) //2nd parameter denotes the length of slice,
	//here my intention is to create a slice with no elements
	evensSlice = append(evensSlice, 12, 34, 58, 34, 90, 46, 32, 66, 82)

	evensSlice2 := make([]int, 5)                     // Initial length is 5 for this slice
	evensSlice2 = append(evensSlice2, 12, 34, 58, 34) // The first 5 elements will be initialized
	//zero values for integers

	/* 3RD way of creating slice of integers */
	var oddsSlice []int

	oddsSlice = append(oddsSlice, 5)
	oddsSlice = append(oddsSlice, 11)
	oddsSlice = append(oddsSlice, 17, 23, 55, 97, 9)

	/* Printing all 3 slices*/
	fmt.Println("First slice  : ", intsSlice)
	fmt.Println("Second slice : ", evensSlice)
	fmt.Println("Third slice  : ", evensSlice2)
	fmt.Println("Fourth slice : ", oddsSlice)
}

/*OUTPUT:-

First slice  :  [5 12 -24 0 38 19 455 1729 1992 -67]
Second slice :  [12 34 58 34 90 46 32 66 82]
Third slice  :  [0 0 0 0 0 12 34 58 34]
Fourth slice :  [5 11 17 23 55 97 9]

*/
