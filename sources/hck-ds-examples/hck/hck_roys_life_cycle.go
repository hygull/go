package main

import "fmt"

func main() {
	var n, codeStreak, codeStreakPrev, i, j, k, length, jIndex uint
	var lastStreak, maxStreak, maxStreak2 uint
	var s string
	//var fromFirst bool
	fmt.Scanf("%d", &n)

	for i = 0; i < n; i++ {
		fmt.Scanf("%s", &s)
		length = uint(len(s))
		tillLast := false
		for j = 0; j < length; j++ { //uint,byte

			codeStreak = 0
			jIndex = j

			if jIndex == 0 && string(s[jIndex]) == "C" && i > 0 {
				streaksSum := codeStreak + lastStreak
				fmt.Println("code===>", codeStreak)

				fmt.Println("i:", i)
				fmt.Println("streaksSum:", streaksSum)

				if tillLast {
					maxStreak += streaksSum
					fmt.Println("maxStreak:", maxStreak)
				} else {
					if maxStreak2 < maxStreak {
						maxStreak2 = maxStreak
					}
					maxStreak += codeStreak
				}
			}

			for k = j; k < length && string(s[k]) == "C"; k++ { //string(s[k]) == "C"  &&  k < 1  ==> panic: runtime error: index out of range,
				//in case==> CCCCEEEESSSSEECCCC

				if k == length-1 {
					tillLast = true
					fmt.Println("i:", i, "true...................")
				}
				codeStreak += 1
				j = k
			}

			// if codeStreakPrev < codeStreak {
			// 	codeStreakPrev = codeStreak
			// }
			if k-1 == length-1 && string(s[k-1]) == "C" {
				lastStreak = codeStreak
			}
		}
	}
	if codeStreakPrev > maxStreak {
		maxStreak = codeStreakPrev
	}
	fmt.Println(codeStreakPrev, maxStreak, maxStreak2)
}

/* To get all the codestreak

package main

import "fmt"

func main() {
	var n, codeStreak, codeStreakPrev, i, j, k uint
	var s string

	fmt.Scanf("%d", &n)

	for i = 0; i < n; i++ {
		fmt.Scanf("%s", &s)
		fmt.Println(s, []byte(s))
		for j = 0; j < 18; j++ { //uint,byte
			codeStreak = 0
			for k = j; k < 18 && string(s[k]) == "C"; k++ { //string(s[k]) == "C"  &&  k < 1  ==> panic: runtime error: index out of range,
				//in case==> CCCCEEEESSSSEECCCC
				fmt.Println("Got...", s[k])
				codeStreak += 1
				j = k
				fmt.Println("Now...j=", j)
			}

			if codeStreakPrev < codeStreak {
				codeStreakPrev = codeStreak
			}
			fmt.Println("Codestreak==>", codeStreakPrev)
		}
		fmt.Printf("%d", codeStreakPrev)
	}

}

*/

/* ...LastCodeStreak
package main

import "fmt"

func main() {
	var n, codeStreak, codeStreakPrev, i, j, k uint
	var lastStreak uint //firstStreak, maxStreak uint
	var s string

	fmt.Scanf("%d", &n)

	for i = 0; i < n; i++ {
		fmt.Scanf("%s", &s)
		fmt.Println(s, []byte(s))
		for j = 0; j < 18; j++ { //uint,byte

			codeStreak = 0
			for k = j; k < 18 && string(s[k]) == "C"; k++ { //string(s[k]) == "C"  &&  k < 1  ==> panic: runtime error: index out of range,
				//in case==> CCCCEEEESSSSEECCCC
				fmt.Println("Got...", s[k])
				codeStreak += 1
				j = k
				fmt.Println("Now...j=", j)
			}

			if codeStreakPrev < codeStreak {
				codeStreakPrev = codeStreak
			}
			if string(s[17]) == "C" {
				lastStreak = codeStreak
				fmt.Println("LastCodeStreak", lastStreak)
			}
			fmt.Println("Codestreak==>", codeStreakPrev)
		}

		fmt.Printf("%d", codeStreakPrev)
	}

}


*/

/* 6 testcases passed
package main

import "fmt"

func main() {
	var n, codeStreak, codeStreakPrev, i, j, k,length, jIndex uint
	var lastStreak, maxStreak uint
	var s string

	fmt.Scanf("%d", &n)

	for i = 0; i < n; i++ {
		fmt.Scanf("%s", &s)
		length = uint(len(s))
		for j = 0; j < length; j++ { //uint,byte

			codeStreak = 0
			jIndex = j
			for k = j; k < length&& string(s[k]) == "C"; k++ { //string(s[k]) == "C"  &&  k < 1  ==> panic: runtime error: index out of range,
				//in case==> CCCCEEEESSSSEECCCC
				codeStreak += 1
				j = k
			}

			if jIndex == 0 && string(s[jIndex]) == "C" && i > 0 {
				streaksSum := codeStreak + lastStreak
				if streaksSum > maxStreak {
					maxStreak = streaksSum
				}
			}

			if codeStreakPrev < codeStreak {
				codeStreakPrev = codeStreak
			}
			if k-1 == length-1 && string(s[k-1]) == "C" {
				lastStreak = codeStreak
			}
		}
	}
	if codeStreakPrev > maxStreak {
		maxStreak = codeStreakPrev
	}
	fmt.Println(codeStreakPrev, maxStreak)
}
*/
