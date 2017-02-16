/*
2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, and 97.

101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197 and 199
*/

/*
A very special program for me that successfully compiled (On 4 jan 2017, 03:05 p.m)and gave
the specified ANSWER given on Hackerearth(not for all test cases),it is remaining,
I have to test it for all the test cases.

See the o/p at very bottom...
*/
package main

import "fmt"

func getAlphabeticPrimesMap() map[uint8]string {

	alphaPrimesMap := make(map[uint8]string)
	println("go...")
	for _, pair := range [][]uint8{[]uint8{67, 89}, []uint8{97, 113}} {
		start := pair[0]
		end := pair[1]
		println("go...", start, end)
		var num uint8
		for num = start; num <= end; num++ {
			/*Check for prime*/
			isPrime := true
			var i uint8
			for i = 2; i <= (num / 2); i++ {
				if num%i == 0 {
					isPrime = false
				}
			}
			if isPrime {
				alphaPrimesMap[num] = string(num) //eg. map => {67 : A, 71 : E}
			}
		}
	}
	return alphaPrimesMap
}
func main() {
	var t, i, n uint8 //Testcases
	var s string      //Input string that will be coverted to dhananjay's magical word
	fmt.Scanf("%d", &t)
	alphaPrimesMap := getAlphabeticPrimesMap()

	fmt.Println(alphaPrimesMap)
	for i = 0; i < t; i++ { /**/
		fmt.Scanf("%d", &n) //Length of input string
		fmt.Scanf("%s", &s) //String
		intsSlice := []uint8(s)

		/*To generate magical word*/
		for index, num := range intsSlice { /***/
			fmt.Println("extracting ", num, " for checking...")
			// test := 1
			a := num - 1
			b := num + 1
			for true { /****/

				fmt.Println("A=>", a, ", B=>", b)
				_, ok1 := alphaPrimesMap[a]
				_, ok2 := alphaPrimesMap[b]

				if ok1 && ok2 {
					if a < b {
						intsSlice[index] = a
					} else {
						intsSlice[index] = b
					}
					break
				} else {
					if ok1 {
						intsSlice[index] = a
						break
					} else {
						if ok2 {
							intsSlice[index] = b
							break
						}
					}
				}
				a -= 1
				b += 1
				//test += 1
			} /****/
		} /***/
		fmt.Println(string(intsSlice))
	} /**/
}

/* My Compilation & Execution o/p:-
1
go...
go... 67 89
go... 97 113
map[71:G 73:I 101:e 103:g 97:a 107:k 109:m 113:q 67:C 79:O 83:S 89:Y]
6
AFREEN
extracting  65  for checking...
A=> 64 , B=> 66
A=> 63 , B=> 67
extracting  70  for checking...
A=> 69 , B=> 71
extracting  82  for checking...
A=> 81 , B=> 83
extracting  69  for checking...
A=> 68 , B=> 70
A=> 67 , B=> 71
extracting  69  for checking...
A=> 68 , B=> 70
A=> 67 , B=> 71
extracting  78  for checking...
A=> 77 , B=> 79
CGSCCO

*/

/*Online Judge(INPUT):-
		1
		6
		AFREEN

OUTPUT:-
		CGSCCO
*/
