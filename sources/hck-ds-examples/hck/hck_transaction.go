package main

import "fmt"

func main() {
	var testcases, queries, i, j int

	fmt.Scanf("%d%d", &testcases, &queries)
	tcosts := make([]int, testcases)
	for i = 0; i < testcases; i++ {
		fmt.Scanf("%d", &tcosts[i])
	}

	for i = 0; i < queries; i++ {
		var m, n int
		fmt.Scanf("%d%d", &m, &n)

		j = 0
		for true {
			if n < 0 || j == testcases {
				fmt.Println(-1)
				break
			} else {
				if tcosts[j] >= m {
					n--
					if n == 0 {
						fmt.Println(tcosts[j])
						break
					}
				}
			}
			j++
		}
	}
}
