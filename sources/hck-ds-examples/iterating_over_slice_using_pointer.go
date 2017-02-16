/*
	{
		"date_of_creation" => "21 Dec 2016,Friday",
		"aim_of_program"   => "Iterating over arrays using pointer & printing the elememnts",
		"coded_by"         => "Rishikesh Agrawani",
		"Go_version"	   => "1.7",
	}
*/
package main

import "unsafe"
import "fmt"

func bubbleSort(intsArr *int, size int) {
	i := 0
	for i < size {

		fmt.Println(*intsArr, "is at address ", intsArr)

		intsArr = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(intsArr)) + unsafe.Sizeof(int(1))))

		i++
	}
}

func main() {
	intsArr := []int{34, 12, 56, 43, 1, 9, 93, -4, 93, -9}
	bubbleSort(&intsArr[0], len(intsArr))
}
