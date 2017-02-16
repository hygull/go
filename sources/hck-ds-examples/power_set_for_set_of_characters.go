/*
   #Date of creation : 9 Jan 2016.
   #Aim of program   : To print power set of a set of characters.
   #Coded by         : Rishikesh Agrawani.
*/
package main

import "fmt"

func main() {
	var n, r, i, j uint
	fmt.Print("Enter the number of binary variables(for which you want the binary combinations): ")
	fmt.Scanf("%d", &n)
	fmt.Print("\nEnter ", n, " binary variables name( name should be only 1 character long) separated by space: ")

	a := make([]string, n)
	r = 1
	for i = 0; i < n; i++ {
		fmt.Scanf("%s", &a[i])
		r *= 2

	}
	fmt.Println("\nColumns => ", n, "\nRows    => ", r)

	for i = 0; i < r; i++ {
		for j = 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				fmt.Print(a[j], " ")
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
}

/*1st RUN:

Enter the number of binary variables(for which you want the binary combinations):  4

Enter 4 binary variables name( name should be only 1 character long) separated by space:  a b c d

Columns =>  4
Rows    =>  16
- - - -
a - - -
- b - -
a b - -
- - c -
a - c -
- b c -
a b c -
- - - d
a - - d
- b - d
a b - d
- - c d
a - c d
- b c d
a b c d
*/

/*2nd RUN:
Enter the number of binary variables(for which you want the binary combinations):
Enter 5 binary variables name( name should be only 1 character long) separated by space:
Columns =>  5
Rows    =>  32
- - - - -
p - - - -
- q - - -
p q - - -
- - r - -
p - r - -
- q r - -
p q r - -
- - - s -
p - - s -
- q - s -
p q - s -
- - r s -
p - r s -
- q r s -
p q r s -
- - - - t
p - - - t
- q - - t
p q - - t
- - r - t
p - r - t
- q r - t
p q r - t
- - - s t
p - - s t
- q - s t
p q - s t
- - r s t
p - r s t
- q r s t
p q r s t
*/
