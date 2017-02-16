package main

import "fmt"

func main() {
	/*Getting number of rows & columns*/
	var rows, columns, i, j int8

	fmt.Scanf("%d%d", &rows, &columns)
	a := make([][]int8, rows)
	for i = 0; i < rows; i++ {
		a[i] = make([]int8, columns)
	}

	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}
	for i = 0; i < columns; i++ {
		for j = 0; j < rows; j++ {
			fmt.Print(a[j][i], " ")
		}
		fmt.Println()
	}
}
