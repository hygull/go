package main

import "fmt"

func main() {
	var n, codeStreak, codeStreakMaxAmongString, i, j, k, length uint
	//var jIndex uint
	var lastStreak uint //maxStreak uint
	var s string
	var tillLast, tillLast2 bool

	fmt.Scanf("%d", &n)

	for i = 0; i < n; i++ { //No. of testcases
		fmt.Scanf("%s", &s)
		length = uint(len(s))

		for j = 0; j < length; j++ { //Iterating through string
			codeStreak = 0
			for k = j; k < length && string(s[k]) == "C"; k++ {
				codeStreak += 1
				j = k
				if k == length-1 {
					tillLast = true
				}
			}

			if codeStreakMaxAmongString < codeStreak {
				codeStreakMaxAmongString = codeStreak
			}

			fmt.Println("Codestreak calculated(i:", i, ",j:", j, ") ==> ", codeStreak)

		}

		fmt.Println("Executing...")
		if tillLast2 {
			lastStreak += codeStreak
		} else {
			if tillLast {
				lastStreak = codeStreak
			}
		}

		tillLast2 = tillLast
		tillLast = false
		fmt.Println("Final Codestreak[", i, "] : ", codeStreakMaxAmongString)
	}
	fmt.Println(lastStreak)
}
