package main

import "fmt"

func main() {
	fixIncrementOnRow := 1
	var n int
	for i := 0; i < 10; i++ {
		n = 21 + i*21
		for j := 1; j < 11; j++ {
			if j == 1 {
				fmt.Print(n, "\t")
				continue
			}
			n += fixIncrementOnRow
			fmt.Print(n, "\t")
		}
		fixIncrementOnRow += 1
		fmt.Println()
	}
}
