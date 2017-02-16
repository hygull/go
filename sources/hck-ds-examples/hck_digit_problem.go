/*
	999932736841900368
	2 january 2017
	3 january 2017(Before 11:04 am), I solved it with full interest
	https://www.hackerearth.com/practice/basic-programming/implementation/basics-of-implementation/practice-problems/algorithm/digit-problem/description/

	::Waiting, because of wrong test cases
	Now solved on 3 jan 2017, wed
*/

/*
	{
		"date_of_creation" => "3 jan 2017, Wednesday",
		"problem_name"   => "Digit problem on Hackerearth",
		"coded_by"         => "Rishikesh Agrawani",
		"resources"		   => "{ 	https://www.hackerearth.com/practice/basic-programming/implementation/basics-of-implementation/practice-problems/algorithm/digit-problem/description/}",
		"Go_version"	   => "1.7",
	}
*/
package main

import "fmt"

func main() {
	var n, places, r int
	var a []int
	fmt.Scanf("%d%d", &n, &places)

	for n != 0 {
		r = n % 10
		a = append(a, r)
		n = n / 10
	}
	for i := len(a) - 1; i > -1; i-- {
		if (places != 0) && (a[i] != 9) {
			fmt.Print(9)
			places -= 1
			continue
		}
		fmt.Print(a[i])
	}
	fmt.Println()
}

// package main

// import "fmt"

// func main() {
// 	var n, places, r int
// 	var a []int
// 	fmt.Scanf("%d%d", &n, &places)

// 	for n != 0 {
// 		r = n % 10
// 		a = append(a, r)
// 		n = n / 10
// 	}
// 	for i := len(a) - 1; i > -1; i-- {
// 		if places != 0 {
// 			fmt.Print(9)
// 			places -= 1
// 			continue
// 		}
// 		fmt.Print(a[i]) //Focusing only on printing the digits, not the original number after replacing digits
// 	}
// 	fmt.Println()
// }
