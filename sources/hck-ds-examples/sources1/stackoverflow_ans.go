/*Visit my gist available at https://gist.github.com/hygull/0fbc428dc77bef4a665b19d598f865d8 to view the clear code.Here you will also find a link for an online terminal where my code is available.Just try to understand the below code. Here you will get the idea for extracting structure's information(field's name their type & value). Just see the below code and its related output.*/

/*Problem link*/
/*http://stackoverflow.com/questions/24512112/golang-how-to-print-struct-variables-in-console*/

/*After submission*/
/*http://stackoverflow.com/questions/24512112/golang-how-to-print-struct-variables-in-console/41198275#41198275*/
package main

import "fmt"
import "reflect"

func main() {
	type Person struct {
		Name     string
		Age      int
		IsActive bool
	}
	person1 := Person{"Rishikesh", 24, true}
	p1 := reflect.ValueOf(&person1).Elem()

	fmt.Printf("%+v\n\n", person1)
	for i := 0; i < p1.NumField(); i++ {
		fieldName := p1.Type().Field(i).Name
		fmt.Printf("%T , %v\n", fieldName, fieldName)

		fieldType := p1.Field(i).Type()
		fmt.Printf("%T, %v\n", fieldType, fieldType)

		fieldValue := p1.Field(i).Interface()
		fmt.Printf("%T, %v\n\n", fieldValue, fieldValue)
	}
}

/*
   {Name:Rishikesh Age:24 IsActive:true}

   string , Name
   *reflect.rtype, string
   string, Rishikesh

   string , Age
   *reflect.rtype, int
   int, 24

   string , IsActive
   *reflect.rtype, bool
   bool, true
*/
