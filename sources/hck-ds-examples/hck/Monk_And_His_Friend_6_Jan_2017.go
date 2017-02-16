package main

import "fmt"

func main() {
	var testcases, count, g, l, i, a, b uint
	fmt.Scanf("%d", &testcases)

	for i = 0; i < testcases; i++ {
		fmt.Scanf("%d%d", &a, &b)
		if a > b {
			g = a
			l = b
		} else {
			g = b
			l = a
		}
		count = 0

		for g != 0 {
			if (g & 1) != (l & 1) {
				count += 1
			}
			g = g >> 1
			l = l >> 1
		}
		fmt.Println(count)
	}
}

/* OUTPUT:-

4
1 4
2
3 3
0
5 1
1
8 7
4
*/
