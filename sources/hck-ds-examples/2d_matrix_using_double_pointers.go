/*
	#Date of creation : 14 Jan 2016.
	#Coded by         : Rishikesh Agrawani.
	#Aim              : To implement integer matrix multiplication using a single triple pointer.
*/

package main

import "fmt"

func main() {

	a := make([]**int, 3)

	var rows, columns int

	fmt.Scanf("%d%d", &rows, &columns)

	for i := 0; i < 3; i++ {
		a[i] = new(make([]*int, rows))

	}

	for k := 0; k < 3; k++ {
		for i := 0; i < 2; i++ {
			a[k][i] = make([]int, columns)
			for j := 0; j < columns; j++ {
				fmt.Scanf("%d", &a[k][i][j])
			}
		}
	}

	for i := 0; i < rows; i++ {
		a[3][i][j] = 0
		for j := 0; j < columns; j++ {
			a[2][i][j] += a[0][i][k] * a[1][k][j]
			fmt.Println(a[2][i][j], "\t")
		}
		fmt.Println()
	}
}
