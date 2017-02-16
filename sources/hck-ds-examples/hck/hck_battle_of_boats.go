package main

import "fmt"

func main() {
	matrix := make([][]uint8, 6)
	var maxMoves, playerId uint8

	for i := 0; i < 6; i++ {
		a := make([]uint8, 7)
		for j := 0; j < 7; j++ {
			fmt.Scanf("%d", &a[j])
		}
		matrix[i] = a
	}
	fmt.Println(matrix)
	fmt.Scanf("%d", &maxMoves)
	fmt.Scanf("%d", &playerId)
}

/*
x coordinates increases form top to bottom
y cor
*/
