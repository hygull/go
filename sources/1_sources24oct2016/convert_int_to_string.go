package main

import "fmt"
import "strconv"

func main() {
	var strNum string
	var intNum int
	fmt.Print("Enter an integer number :\t")
	fmt.Scan(&intNum)
	strNum = strconv.Itoa(intNum) //If you will try to convert invalid string like "abc" , "123fv" etc then

	if strNum == "0" {
		var remainingInputCharsSequeceInStrInput string //To prevent unneccessary message...in case of wrong input
		//it will consume the rest entered string for example
		//if you will enter "abc" then remainingInputCharsSequeceInStrInput will contain "bc"
		fmt.Println("This is not a valid integer")
		fmt.Scan(&remainingInputCharsSequeceInStrInput) //This is for consuming the characeters from 1th index to onwards
		//In case of wrong input
		return
	} else {
		fmt.Println("You entered an integer <a valid integer> ie. ", strNum)
	}
}
