package main

import "fmt"
import "strconv"

func main() {
	var strNum string
	var intNum int
	fmt.Print("Enter a string as integer number :\t")
	fmt.Scan(&strNum)
	intNum, err := strconv.Atoi(strNum)

	if err != nil {
		fmt.Println("This is not a valid integer number in string form")
	} else {
		fmt.Println("You entered ", intNum)
	}
}
