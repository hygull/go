package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var testcases, i, j, n int
	//var str uint8

	fmt.Scanf("%d", &testcases)
	for i = 0; i < testcases; i++ {
		// fmt.Println("Iteration (", i, ")")

		fmt.Scanf("%d", &n)
		fmt.Println(n)
		// fmt.Println("INPUT", f)
		reader := bufio.NewReader(os.Stdin)
		numsStr, _ := reader.ReadString('\n')
		fmt.Println(numsStr)
		numsStrList := strings.Fields(numsStr)
		// fmt.Println(numsStrList)
		a := make([]int, n)
		for j = 0; j < n; j++ {
			a[j], _ = strconv.Atoi(numsStrList[j])
		}
		fmt.Println(a)

		isInverse := true
		// fmt.Println("N:", n)
		for j = 0; j < n; j++ {
			//b[a[a[j]-1]-1] = a[j]
			// fmt.Println(a[j]-1, "...", "j:", j)
			if a[a[j]-1]-1 != j {
				isInverse = false
			}
		}
		if !isInverse {
			fmt.Println("not inverse")
		} else {
			fmt.Println("inverse")
		}
		fmt.Println(n, "...")
	}
}
