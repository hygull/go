package main

import "strconv"
import "fmt"

func main() {
	var strf1 string
	var f1 float64
	//var f2 float32;

	fmt.Print("Enter a float number to be parsed <It is being taken as string format> : ")
	fmt.Scan(&strf1) //23, 45.45, 67.894

	f1, err := strconv.ParseFloat(strf1, 64)
	if err != nil {
		fmt.Println("Please enter a proper float number...Retry")
		return
	} else {
		fmt.Println("The entered float as string successfully converted to a float number of type float64 in Golang...and it is ", f1)
	}

}
