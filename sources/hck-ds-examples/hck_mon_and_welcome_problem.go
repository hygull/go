/*
2 jan 2016
https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/monk-and-welcome-problem/
*/
package main

import "fmt"

func main() {
	var n, i int32

	fmt.Scanf("%d", &n)
	a := make([]int32, n, n)
	b := make([]int32, n, n)

	for i = 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i = 0; i < n; i++ {
		fmt.Scanf("%d", &b[i])
		fmt.Printf("%d ", a[i]+b[i])
	}
}

/*INPUT:-
5
1 2 3 4 4
1 2 5 6 7
*/

/*OUTPUT:-
2 4 8 10 11
*/
