/*
	CodedOn : 14 Jan 2017.
	CodedBy : Rishikesh Agrawani.
	Aim     : To display the list of sets having the numbers whose sum should be equal to a given number.
*/
package main

import "fmt"

func main() {
	var n, m, rows, count, i, j, sum uint
	fmt.Scanf("%d%d", &n, &m) //Desired sum, Numeber of unique elements in the slice/list

	a := make([]uint, m)
	rows = 1
	for i = 0; i < m; i++ {
		fmt.Scanf("%d", &a[i]) //Inputting elements of slice
		rows *= 2
	}

	setIndex := 1
	for i = 0; i < rows; i++ {
		sum = 0
		var b []uint
		for j = 0; j < m; j++ {
			if (i>>j)&1 == 1 { //Checking the ON status of bits in after shifting
				sum += a[j]
				b = append(b, a[j]) //Appending to new slice
			}
		}
		if sum == n {
			fmt.Println("SET", setIndex, " => ", b)
			setIndex += 1
			count += 1
		}
	}
	fmt.Printf("Number of valid sets : %d", count)
}

/* The first 2 lines shows the I/P */

/* 	First RUN:-
6 3
3 3 6
SET 1  =>  [3 3]
SET 2  =>  [6]
Number of valid sets : 2
*/

/*
6 10
1 2 3 4 5 6 7 8 9 10
SET 1  =>  [1 2 3]
SET 2  =>  [2 4]
SET 3  =>  [1 5]
SET 4  =>  [6]
Number of valid sets : 4
*/

/*
10 10
1 2 3 4 5 6 7 8 9 10
SET 1  =>  [1 2 3 4]
SET 2  =>  [2 3 5]
SET 3  =>  [1 4 5]
SET 4  =>  [1 3 6]
SET 5  =>  [4 6]
SET 6  =>  [1 2 7]
SET 7  =>  [3 7]
SET 8  =>  [2 8]
SET 9  =>  [1 9]
SET 10  =>  [10]
Number of valid sets : 10
*/
