/*Visit my gist available at https://gist.github.com/hygull/0fbc428dc77bef4a665b19d598f865d8 to view the clear code.Here you will also find a link for an online terminal where my code is available.Just try to understand the below code. Here you will get the idea for extracting structure's information(field's name their type & value). Just see the below code and its related output.*/

/*Problem link*/
/*http://stackoverflow.com/questions/24512112/golang-how-to-print-struct-variables-in-console*/

/*After submission*/
/*http://stackoverflow.com/questions/24512112/golang-how-to-print-struct-variables-in-console/41198275#41198275*/
package main

import "fmt"
import "reflect"

func main() {
	type Book struct {
		Id    int
		Name  string
		Title string
	}

	book := Book{1, "Let us C", "Enjoy programming wit practice"}
	e := reflect.ValueOf(&book).Elem()

	for i := 0; i < e.NumField(); i++ {
		fieldName := e.Type().Field(i).Name
		fmt.Printf("%v\n", fieldName)
	}
}

/*
Id
Name
Title
*/
