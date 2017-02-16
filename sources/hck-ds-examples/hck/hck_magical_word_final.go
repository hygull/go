/*
https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/magical-word/

https://www.hackerearth.com/submission/6669500/
30 points

/*
2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, and 97.

101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197 and 199
*/

/*
A very special program for me that successfully compiled (On 4 jan 2017, 03:05 p.m)and gave
the specified ANSWER given on Hackerearth(not for all test cases),it is remaining,
I have to test it for all the test cases.

See the o/p at very bottom...

Execution(successful): 03:34 p.m
*/

package main

import "fmt"

//A function that gives a map that maps prime ASCII values of alphabets to their corresponding alphabets
func getAlphabeticPrimesMap() map[uint8]string {

	alphaPrimesMap := make(map[uint8]string) //Defining map from uint8 to string
	for _, pair := range [][]uint8{[]uint8{67, 89}, []uint8{97, 113}} {
		start := pair[0]
		end := pair[1]
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

	for i = 0; i < t; i++ { /**/
		fmt.Scanf("%d", &n) //Length of input string
		fmt.Scanf("%s", &s) //String
		intsSlice := []uint8(s)

		/*To generate magical word*/
		for index, num := range intsSlice { /***/
			_, exists := alphaPrimesMap[num] //If the character's ASCII is already prime
			if exists {
				continue
			}
			a := num - 1
			b := num + 1
			for true { /****/
				_, ok1 := alphaPrimesMap[a] //To check if key named a exists in map
				_, ok2 := alphaPrimesMap[b] //To check if key named b exists in map

				if ok1 && ok2 { //If both keys exist
					if a < b {
						intsSlice[index] = a
					} else {
						intsSlice[index] = b
					}
					break
				} else {
					if ok1 { //If only key named a exists
						intsSlice[index] = a
						break
					} else {
						if ok2 {
							intsSlice[index] = b //If only key named a exists
							break
						}
					}
				}
				a -= 1
				b += 1
			} /****/
		} /***/
		fmt.Println(string(intsSlice))
	} /**/
}

/* INPUT:-

1
6
AFREEN
*/

/*OUTPUT:-

CGSCCO
*/
