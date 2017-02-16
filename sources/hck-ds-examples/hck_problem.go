/*
	Date of creation : 02 January 2017.
	Aim of program   : The great Kian problem on Hackerearth.
	Coded by         : Rishikesh Agrawani.
	link             : https://www.hackerearth.com/practice/basic-programming/implementation/basics-of-implementation/practice-problems/algorithm/the-great-kian/
*/

package main

import "fmt"

func main() {
	var n, sum int
	//numsSlice := make([]int, n, n)   => [] , NZEC, Runtime error
	fmt.Scanf("%d", &n)
	numsSlice := make([]int, n, n) //Creating slice of integers

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numsSlice[i])
	}

	for i := 0; i < 3; i++ { //We have to print 3 numbers
		j := i
		sum = 0
		for true {
			if j > n-1 {
				break
			}
			sum += numsSlice[j]
			j += 3
		}
		fmt.Print(sum, " ")
	}
}

/*
INPUT:-
			5
			1 2 3 4 5
*/

/*
OUTPUT:-
			5 7 3

*/
