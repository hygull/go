package main

import "fmt"

func main() {
	var n, codeStreak, codeStreakMaxAmongStrings, i, j, k, length uint
	//var jIndex uint
	//var lastStreak, maxStreak uint
	var s string

	fmt.Scanf("%d", &n)

	for i = 0; i < n; i++ { //No. of testcases
		fmt.Scanf("%s", &s)
		length = uint(len(s))

		for j = 0; j < length; j++ { //Iterating through string
			codeStreak = 0
			for k = j; k < length && string(s[k]) == "C"; k++ {
				codeStreak += 1
				j = k
			}

			if codeStreakMaxAmongStrings < codeStreak {
				codeStreakMaxAmongStrings = codeStreak
			}
			fmt.Println("Codestreak calculated(i:", i, ",j:", j, ") ==> ", codeStreak)
		}
		fmt.Println("Final Codestreak[", i, "] : ", codeStreakMaxAmongStrings)
	}

}
