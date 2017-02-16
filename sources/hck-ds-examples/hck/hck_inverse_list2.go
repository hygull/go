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

		reader := bufio.NewReader(os.Stdin)
		nStr, _ := reader.ReadString('\n')
		n, _ = strconv.Atoi(nStr)

		numsStr, _ := reader.ReadString('\n')
		numsStrList := strings.Fields(numsStr)

		a := make([]int, n)
		for j = 0; j < n; j++ {
			a[j], _ = strconv.Atoi(numsStrList[j])
		}
		fmt.Println(a)

		isInverse := true
		for j = 0; j < n; j++ {
			if a[a[j]-1]-1 != j {
				isInverse = false
			}
		}

		if !isInverse {
			fmt.Println("not inverse")
		} else {
			fmt.Println("inverse")
		}
		reader.ReadString('\n')
	}
}
