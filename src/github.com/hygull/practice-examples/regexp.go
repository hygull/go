package main

import "regexp"
import "fmt"

func main() {
	var regExp, sourceStr string
	fmt.Print("Enter a rege_exp : ")
	fmt.Scan(&regExp)
	fmt.Print("Enter original source string : ")
	fmt.Scan(&sourceStr)

	r := regexp.MustCompile(regExp)
	if r.MatchString(sourceStr) {
		fmt.Println("Matched...")
	} else {
		fmt.Println("Didn't match...")
	}

	//regExp = `[[A-Z]+[1-2]+]*`   // `ADF1GTH3DRT4`
	//regExp = `(^([A-Z]+)(\s*)$)*`
	regExp = `\w+(\s+\w+)*`
	fmt.Print("Enter a string eg.  'ABC DEF FGH' , 'LET US C' etc.....  to match ", regExp, " :  ")
	//fmt.Scan(&sourceStr)
	sourceStr = `AEDF jhGDF65`
	r = regexp.MustCompile(regExp)
	if r.MatchString(sourceStr) {
		fmt.Println("Matched...")
	} else {
		fmt.Println("Didn't match...")
	}
}
