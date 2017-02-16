/*
	{
		"cretaed_after" : "Sat Dec 10 23:14:26 IST 2016"
		"aim_of_program" : "To print X symbol using * characters"
		"coded_by" : "Rishikesh Agrawani"
	}
*/
package main

import "fmt"

func main() {
	var n int //var n int8 ... will generate an error
	fmt.Print("Enter a number between 3-10 :\t")
	fmt.Scanf("%d", &n)

	if n > 2 && n < 11 { // 3 <= n <= 10
		for i := 0; i < n; i++ { //Outer for loop takes care about rows
			for j := 0; j < n; j++ { //Inner for loop takes care about columns
				if i == j || (i+j) == (n-1) { //Based on matrix concept we got these conditions
					fmt.Print("*")
				} else {
					fmt.Print(" ") //In case when conditions don't match
				}
			}
			fmt.Println() //Newline
		}
	} else {
		fmt.Println("\nThe number should be in inclusive range [3-10]")
	}
}

/*

admins-MacBook-Pro-3:GoFiles admin$ go run X.go
Enter a number between 3-10 :	5
*   *
 * *
  *
 * *
*   *

admins-MacBook-Pro-3:GoFiles admin$ go run X.go
Enter a number between 3-10 :	10
*        *
 *      *
  *    *
   *  *
    **
    **
   *  *
  *    *
 *      *
*        *

*/
