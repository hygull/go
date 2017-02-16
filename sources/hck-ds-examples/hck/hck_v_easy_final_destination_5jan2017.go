/*
 coded_on : 5 january 2016.
 https://www.hackerearth.com/practice/basic-programming/implementation/basics-of-implementation/practice-problems/algorithm/final-destination-cakewalk/
*/
package main

import "fmt"

func main() {
	var s string
	var x, y int
	fmt.Scanf("%s", &s)

	//bytes:=byte(s)
	for _, c := range s {
		dir := string(c)
		if dir == "L" { //L...x will decrease by 1
			x -= 1
		} else {
			if dir == "R" { //R...x will increase by 1
				x += 1
			} else {
				if dir == "U" { //U...y will increase by 1
					y += 1
				} else { //D...y will decrease by 1
					y -= 1
				}
			}
		}
	}
	fmt.Print(x, " ", y)
}
