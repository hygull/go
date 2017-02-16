/*
   #Date of creation : 8 Jan 2016.
   #Aim of program   : To print binary combinations.
   #Coded by         : Rishikesh Agrawani.
*/

package main

import "fmt"

func main() {
	var n, r, i, j uint
	fmt.Print("Enter the number of binary variables: ")
	fmt.Scanf("%d", &n)
	r = 1
	for i = 1; i <= n; i++ {
		r *= 2

	}
	fmt.Println("Columns => ", n, "\nRows    => ", r)

	for i = 0; i < r; i++ {
		for j = 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				fmt.Print("0 ")
			} else {
				fmt.Print("1 ")
			}
		}
		fmt.Println()
	}
}

/*1st RUN

Enter the number of binary variables:  3
Columns =>  3
Rows    =>  8
1 1 1
0 1 1
1 0 1
0 0 1
1 1 0
0 1 0
1 0 0
0 0 0
*/

/*2nd RUN

Enter the number of binary variables:  4
Columns =>  4
Rows    =>  16
1 1 1 1
0 1 1 1
1 0 1 1
0 0 1 1
1 1 0 1
0 1 0 1
1 0 0 1
0 0 0 1
1 1 1 0
0 1 1 0
1 0 1 0
0 0 1 0
1 1 0 0
0 1 0 0
1 0 0 0
0 0 0 0
*/
