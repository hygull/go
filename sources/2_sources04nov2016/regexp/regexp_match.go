package main

import "fmt"
import "regexp"

func main() {
	var regExpression, stringToBeMatched string

	fmt.Print("Enter a regular expression     : ")
	_, err := fmt.Scan(&regExpression)
	if err != nil {
		fmt.Println("[input:error] Error while reading string input")
		return
	}
	fmt.Print("Enter a string to be macthed   : ")
	_, err = fmt.Scan(&stringToBeMatched)
	if err != nil {
		fmt.Println("[input:error] Error while reading string input")
		return
	}

	matched, err := regexpIsValid(regExpression, stringToBeMatched)
	if err != nil {
		fmt.Println("\n[regexp:error] Error while matching regular expression with entered string...")
		fmt.Println(err, "[regexp:error]Check regular expression and retry...")
		return
	}
	if matched {
		fmt.Println("\n[output] The string '", stringToBeMatched, `' matched with '`, regExpression, "'\n")
	} else {
		fmt.Println("\n[output] The string '", stringToBeMatched, `' didn't match with '`, regExpression, "'\n")
	}
}

func regexpIsValid(regExpression string, stringToBeMatched string) (bool, error) {
	_matched_, _err_ := regexp.MatchString(regExpression, stringToBeMatched)
	return _matched_, _err_
}
