package main

import "fmt"

func main() {
	var n, steps, i, j, testcases uint

	fmt.Scanf("%d", &testcases)

	for i = 0; i < testcases; i++ {
		fmt.Scanf("%d%d", &n, &steps)
		a := make([]uint, n)
		for j = 0; j < n; j++ {
			var val uint
			fmt.Scanf("%d", &val)
			fmt.Println("index:", (j+steps)%n)
			a[(j+steps)%n] = val
		}
		// for j = 0; j < n; j++ {
		// 	fmt.Print(a[j], " ")
		// }
		fmt.Println()
	}

	/*
			    initially:-
			   	  0   1   2   3   4
			   	+---+---+---+---+---+
			   	| 5 | 6 | 7 | 9 | 3 |
			   	+---+---+---+---+---+

		--->

				1 step rotation:-
			   	+---+---+---+---+---+
			   	| 3 | 5 | 6 | 7 | 9 |
			   	+---+---+---+---+---+


	*/
}
