//https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/roy-and-profile-picture/
package main

import "fmt"

func main() {
	var l, n, i uint16
	fmt.Scanf("%d", &l) //Length of side of valid square shaped photo
	fmt.Scanf("%d", &n) //Number of photos

	for i = 0; i < n; i++ {
		var w, h uint16
		fmt.Scanf("%d%d", &w, &h)
		if w < l || h < l {
			fmt.Println("UPLOAD ANOTHER")
		} else {
			if w == h { //(180,180), (240,240)
				fmt.Println("ACCEPTED")
			} else {
				fmt.Println("CROP IT")
			}
		}
	}
}

/*
admins-MacBook-Pro-3:hck admin$ go run hck_roy_and_profile_pic.go
180
3
640 480
CROP IT
120 300
UPLOAD ANOTHER
180 180
ACCEPTED

admins-MacBook-Pro-3:hck admin$ date
Wed Jan  4 11:08:36 IST 2017
admins-MacBook-Pro-3:hck admin$
*/
