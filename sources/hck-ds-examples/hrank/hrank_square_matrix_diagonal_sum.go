package main

/*
https://www.hackerrank.com/challenges/diagonal-difference?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=30-day-campaign
Solved...
*/
import "fmt"

func main() {
	var n, i, j, primarySum, secondarySum int
	var a [][]int
	fmt.Scanf("%d", &n)

	for i = 0; i < n; i++ {
		var arr []int
		for j = 0; j < n; j++ {
			var n int
			fmt.Scanf("%d", &n)
			arr = append(arr, n)
		}
		a = append(a, arr)

		primarySum += a[i][i]
		secondarySum += a[i][n-1-i]
	}
	d := primarySum - secondarySum
	if d < 0 {
		fmt.Printf("%d", -d)
	} else {
		fmt.Printf("%d", d)
	}
}
