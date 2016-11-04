package main

import "fmt"

func main() {
	numsSlice := make([]int, 2, 3)
	fmt.Printf("%v %v %v\n", numsSlice, len(numsSlice), cap(numsSlice))

	numsSlice = append(numsSlice, 3, 4)
	fmt.Printf("%v %v %v\n", numsSlice, len(numsSlice), cap(numsSlice))

	numsSlice = append(numsSlice, 45, 56)
	fmt.Printf("%v %v %v\n", numsSlice, len(numsSlice), cap(numsSlice))
}
