/*
https://www.hackerearth.com/practice/basic-programming/implementation/basics-of-implementation/practice-problems/algorithm/matrix-symmetry/
Best submission...
https://www.hackerearth.com/submission/6698922/
*/

package main

import "fmt"

func main() {
	var i, j, testcases, rows, r, c uint
	fmt.Scanf("%d", &testcases)

	/*
		4
		*.**
		.***
		.***
		*.**

	*/

	for i = 0; i < testcases; i++ { //Testcases...
		fmt.Scanf("%d", &rows) //Number of rows in square matrix
		m := make([]string, rows)

		for j = 0; j < rows; j++ {
			var s string
			fmt.Scanf("%s", &s)
			m[j] = s //m = append(m, s) ==> panic: runtime error: index out of range
		}
		/* Checking for horizontal symmetricity */
		isHorizontal := true
		for r = 0; r < (rows / 2); r++ {
			for c = 0; c < rows; c++ {
				if m[r][c] != m[rows-1-r][c] {
					isHorizontal = false
					break
				}
			}
			if !isHorizontal {
				break
			}
		}
		/* Checking for vertical symmetricity */
		isVertical := true
		for r = 0; r < rows; r++ {
			for c = 0; c < (rows / 2); c++ {
				if m[r][c] != m[r][rows-1-c] {
					isVertical = false
					break
				}
			}
			if !isVertical {
				break
			}
		}

		if isVertical && isHorizontal {
			fmt.Println("BOTH")
		} else {
			if isHorizontal {
				fmt.Println("HORIZONTAL")
			} else {
				if isVertical {
					fmt.Println("VERTICAL")
				} else {
					fmt.Println("NO")
				}
			}
		}

	} //Iterates upto no. of test cases
}

// package main

// import "fmt"
// import "github.com/fatih/color"

// func main() {
// 	var i, j, testcases, rows, r, c uint
// 	fmt.Scanf("%d", &testcases)

// 	/*
// 		4
// 		*.**
// 		.***
// 		.***
// 		*.**

// 	*/

// 	for i = 0; i < testcases; i++ { //Testcases...
// 		fmt.Scanf("%d", &rows) //Number of rows in square matrix
// 		m := make([]string, rows)

// 		for j = 0; j < rows; j++ {
// 			var s string
// 			fmt.Scanf("%s", &s)
// 			m[j] = s //m = append(m, s) ==> panic: runtime error: index out of range
// 		}
// 		fmt.Println(m)
// 		/*...Checking for symmetricity...*/
// 		color.Blue("Checking for symmetricity...")
// 		isSymmetric := true
// 		for r = 0; r < rows; r++ {
// 			for c = 0; c < rows; c++ {
// 				fmt.Println("==> ", m[r][c], m[c][r])
// 				if (r != c) && m[r][c] != m[c][r] {
// 					isSymmetric = false
// 					r = rows //To come out from outer for loop
// 					break
// 				}
// 			}
// 		}
// 		/*...Symmetricity check ends here...*/

// 		/*...Checking for horizontal & vertical symmetricity...*/

// 		/* Checking for horizontal symmetricity */
// 		if isSymmetric {
// 			color.Green("Entered matrix is Symmetric.")
// 		} else {
// 			color.Green("Entered matrix is not Symmetric.")
// 		}

// 		color.Blue("Now checking for Horizontal Symmetricity...")
// 		isHorizontal := true
// 		for r = 0; r < (rows / 2); r++ {
// 			for c = 0; c < rows; c++ {
// 				if m[r][c] != m[rows-1-r][c] {
// 					isHorizontal = false
// 					break
// 				}
// 			}
// 			if !isHorizontal {
// 				break
// 			}
// 		}
// 		if isHorizontal {
// 			color.Green("Entered Matrix is Horizontal")
// 		} else {
// 			color.Green("Entered Matrix is not Horizontal")
// 		}

// 		/* Checking for vertical symmetricity */
// 		color.Blue("Now checking for Vertical Symmetricity")
// 		isVertical := true
// 		for r = 0; r < rows; r++ {
// 			for c = 0; c < (rows / 2); c++ {
// 				if m[r][c] != m[r][rows-1-c] {
// 					isVertical = false
// 					break
// 				}
// 			}
// 			if !isVertical {
// 				break
// 			}
// 		}
// 		if isVertical {
// 			color.Green("Entered Matrix is Vertical")
// 		} else {
// 			color.Green("Entered Matrix is not Vertical")
// 		}

// 		if isVertical && isHorizontal {
// 			fmt.Println("BOTH")
// 		} else {
// 			if isHorizontal {
// 				fmt.Println("HORIZONTAL")
// 			} else {
// 				if isVertical {
// 					fmt.Println("VERTICAL")
// 				} else {
// 					fmt.Println("NO")
// 				}
// 			}
// 		}

// 	} //Iterates upto no. of test cases
// }
