/*5 jan 2016...
https://www.hackerearth.com/practice/basic-programming/implementation/basics-of-implementation/practice-problems/algorithm/string-sum/

Special : A very fast executed submission on HAckerearth...
*/

package main

import "fmt"

func main() {
	var sum uint32
	var s string

	fmt.Scanf("%s", &s)
	for _, c := range s {
		sum += uint32(c % 96)
	}
	fmt.Print(sum)
}
