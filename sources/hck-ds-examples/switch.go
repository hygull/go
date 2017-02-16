/*
	{
		"date_of_creation" => "24 Dec 2016, Saturday",
		"aim_of_program"   => "Using switch cases in Golang"
		"coded_by"         => "Rishikesh Agrawani",
		"Go_version"	   => "1.7",
	}
*/
package main

import "fmt"
import "reflect"

//switch example 1
func PrintNums(i int) {
	switch i {
	case 1, 10:
		fmt.Println(i, "\t--> One/Ten...\n")
	case 2:
		fmt.Println(i, "\t--> Two...\n")
	case 3, 6, 9:
		fmt.Println(i, "\t--> Three/Six/Nine ...\n")
	default:
		fmt.Println(i, "\t--> Either greater than 3 or less than 1\n")
	}
}

//switch example 2
func PrintFloats(f float64) {
	switch f {
	case 23.1, 87.56:
		fmt.Println(f, "\t--> Twenty three point one/Eighty seven point fifty six\n")
	case 23.98 + 26.02:
		fmt.Println(f, "\t--> Fifty...\n")
	case 38.2:
		fmt.Println(f, "\t--> Thirty eight point two...\n")
	default:
		fmt.Println(f, "\t--> This number is not the number that we want...\n")
	}
}

//switch example 3
func PrintStrings(s string) {
	switch s {
	case "hygull", "rob", "ken", "robert":
		fmt.Println(s, "\t--> A person...\n")
	case "tiger":
		fmt.Println(s, "\t--> An animal...\n")
	case "apple":
		fmt.Println(s, "\t--> A fruit...\n")
	default:
		fmt.Println(s, "\t--> This is not the thing that we want...\n")
	}
}

//switch example 4
func CheckTypeOfData(data interface{}) {
	switch typ := data.(type) {
	case []int:
		fmt.Println(data.([]int), " is of type []int...\n")
	case []string:
		fmt.Println(data.([]string), " is of type []string...\n")
	case []float64:
		fmt.Println(data.([]float64), " is of type []float64...\n")
	case []int32:
		fmt.Println(data.([]int32), " is of type []int32...\n")
	default:
		fmt.Println(typ, " is of type", reflect.TypeOf(data), "... but we are expecting...[]int,[]string,[]float64,[]int32...\n")
	}
}

//switch example 5(using fallthrough)
func UseFallthrough(n int) {
	switch {
	case n == 1: //the expression should be a bool value(in this case)
		fmt.Println(n, "One...\n")
	case n == 2:
		fmt.Println(n, "Two...\n")
		fallthrough //forwards control to next case
	case n == 3:
		fmt.Println(n, "Three/Two...\n")
		fallthrough //forwards control to the next case (It will execute for n=2 or n=3)
	case n == 4:
		fmt.Println(n, "Four/Three/Two\n") //In case of fallthruogh (It will execute for n=3 or n=4 or n=2)
	default:
		fmt.Println(n, "We don't want this number...")
	}
}

func main() {
	for _, num := range []int{1, 0, -5, 9, 3, -11, 2, 5} {
		PrintNums(num)
	}

	CheckTypeOfData([]int{23, 56, 9, -2})
	CheckTypeOfData([]string{"Glue", "Fish", "Hygull", "Rob"})
	CheckTypeOfData([]float32{4.3, 5.8, -45.2})
	CheckTypeOfData([]float64{45, 89.76, 9.33})
	CheckTypeOfData([]int16{7, 9, 12, -5, 39})
	CheckTypeOfData([]int32{34, 65, 21, 39, 87})

	for _, floatNum := range []float64{50.0, 56.9, 23.1, 38.2, 50, 32.215} {
		PrintFloats(floatNum)
	}

	for _, str := range []string{"rob", "hen", "tiger", "rose", "apple"} {
		PrintStrings(str)
	}

	for _, n := range []int{4, 5, 2, 7, 1, 9, 3} {
		UseFallthrough(n)
	}
}

/* RUN:-

admins-MacBook-Pro-3:GoFiles admin$ go run switch.go
1 	--> One/Ten...

0 	--> Either greater than 3 or less than 1

-5 	--> Either greater than 3 or less than 1

9 	--> Three/Six/Nine ...

3 	--> Three/Six/Nine ...

-11 	--> Either greater than 3 or less than 1

2 	--> Two...

5 	--> Either greater than 3 or less than 1

[23 56 9 -2]  is of type []int...

[Glue Fish Hygull Rob]  is of type []string...

[4.3 5.8 -45.2]  is of type []float32 ... but we are expecting...[]int,[]string,[]float64,[]int32...

[45 89.76 9.33]  is of type []float64...

[7 9 12 -5 39]  is of type []int16 ... but we are expecting...[]int,[]string,[]float64,[]int32...

[34 65 21 39 87]  is of type []int32...

50 	--> Fifty...

56.9 	--> This number is not the number that we want...

23.1 	--> Twenty three point one/Eighty seven point fifty six

38.2 	--> Thirty eight point two...

50 	--> Fifty...

32.215 	--> This number is not the number that we want...

rob 	--> A person...

hen 	--> This is not the thing that we want...

tiger 	--> An animal...

rose 	--> This is not the thing that we want...

apple 	--> A fruit...

4 Four/Three/Two

5 We don't want this number...
2 Two...

2 Three/Two...

2 Four/Three/Two

7 We don't want this number...
1 One...

9 We don't want this number...
3 Three/Two...

3 Four/Three/Two

*/
