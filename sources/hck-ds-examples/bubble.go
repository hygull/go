/*
	{
		"date_of_creation" => "21 Dec 2016,Friday",
		"aim_of_program"   => "Bubble sort",
		"coded_by"         => "Rishikesh Agrawani",
		"Go_version"	   => "1.7",
	}
*/
package main

import "fmt"
import "reflect"

func bubbleSort(intsList []int, size int) {

	fmt.Printf("%T , %v", reflect.TypeOf(size), reflect.TypeOf(size))
	for i := size - 2; i >= 0; i-- {
		swapped := false
		for j := 0; j <= i; j++ {
			if intsList[j] > intsList[j+1] {
				intsList[j] = intsList[j] + intsList[j+1]
				intsList[j+1] = intsList[j] - intsList[j+1]
				intsList[j] = intsList[j] - intsList[j+1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	intsSlice := []int{34, 12, 56, 43, 1, 9, 93, -4, 93, -9}

	fmt.Println(intsSlice)

	bubbleSort(intsSlice, len(intsSlice))

	fmt.Println(intsSlice)
}
