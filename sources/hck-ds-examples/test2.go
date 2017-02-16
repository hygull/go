// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// var a, b int
// 	// var t string
// 	// c, err := fmt.Scanf("%d%d", &a)
// 	// fmt.Scanf("%s", &t)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// fmt.Println(c, b, t)
// 	d := []string{"a", "q", "l"}
// 	s := "Rishikesh"
// 	s1 := "AppleDqpldasfglwweq"
// 	s2 := s1
// 	//AppleDqpldasfglwweq
// 	fmt.Println(strings.Replace(s, "i", "", -1))

// 	fmt.Println(strings.Replace(s2, "a", "", -1))
// 	fmt.Println("==>", s2)
// 	fmt.Println(strings.Replace(s2, "q", "", -1))
// 	fmt.Println("==>", s2)

// 	for _, w := range d {
// 		s1 = strings.Replace(s1, w, "", -1)
// 		fmt.Println(w, s1)
// 	}
// 	fmt.Println(s, s1)
// 	//fmt.Println(c, b)
// }
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d%d", &a, &b)
	fmt.Println(a, "\t", b)
	fmt.Scanf("%d", &c)

	// fmt.Print("Enter number of inputs")
	fmt.Println(c)
}
