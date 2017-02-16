package main

import "fmt"

func main() {
	//Initializing a slice of integers
	intsSlice := []int{12, 67, 98, -56, 0, 23, -4}

	/*Displaying all the elements of integer slice*/
	for _, num := range intsSlice {
		fmt.Println(num)
	}
}
