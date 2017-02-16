/*
	{
		"date_of_creation" => "17 Dec 2016, Thurs",
		"aim_of_program"   => "Extracting structure's information using reflect package",
		"coded_by"         => "Rishikesh Agrawani",
		"Go_version"	   => "1.7",
	}
*/
package main

import "reflect"
import "fmt"

func main() {

	//Defining Brother structure
	type Brother struct {
		Name string
		Age  int
	}

	//Defining Myself structure
	type MySelf struct {
		Name         string //name string -> panic: reflect.Value.Interface: cannot return value obtained from unexported field or method
		Age          int    //age int -> panic: reflect.Value.Interface: cannot return value obtained from unexported field or method
		IsActive     bool   //And so on...
		DailyExpense float32
		BrotherInfo  Brother
	}

	brother := Brother{"Hemkesh", 22}
	myself := MySelf{"Rishikesh", 24, true, 200.50, brother}

	p1 := reflect.ValueOf(&myself).Elem()

	fmt.Printf("%+v\n\n", myself)

	for i := 0; i < p1.NumField(); i++ {
		fieldName := p1.Type().Field(i).Name
		fmt.Printf("%T , %v\n", fieldName, fieldName)

		fieldType := p1.Field(i).Type()
		fmt.Printf("%T, %v\n", fieldType, fieldType)

		fieldValue := p1.Field(i).Interface() //In case of the field names that starts with smallcase, only this line will show the above errors
		fmt.Printf("%T, %v\n\n", fieldValue, fieldValue)
	}
}

/* OUTPUT:-

{Name:Rishikesh Age:24 IsActive:true DailyExpense:200.5 BrotherInfo:{Name:Hemkesh Age:22}}

string , Name
*reflect.rtype, string
string, Rishikesh

string , Age
*reflect.rtype, int
int, 24

string , IsActive
*reflect.rtype, bool
bool, true

string , DailyExpense
*reflect.rtype, float32
float32, 200.5

string , BrotherInfo
*reflect.rtype, main.Brother
main.Brother, {Hemkesh 22}

*/
