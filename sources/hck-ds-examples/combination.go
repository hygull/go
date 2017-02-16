/*
	{
		"date_of_creation" => "15 Dec 2016, Wed",
		"aim_of_program"   => "To show binary combinations according to the number of binary variables",
		"coded_by"         => "Rishikesh Agrawani",
		"memory"           => "iteration"
		"Go_version"	   => "1.7",
	}
*/
package main

import (
	"fmt"
)

/* This function calculates the Pow(base,pow) and returns it to the caller */
func pow(base int, pow int) int {
	rows := 1
	for i := 0; i < pow; i++ {
		rows = rows * base
	}
	return rows
}

/* This function lists all the binary combinations according to the number of variables passed */
func showBinaryCombinationsFor(cols int) {
	rows := pow(2, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			shiftCount := uint(j) //Shift count should be an unsigned integer, So i>>j will not work
			if 1&(i>>shiftCount) == 1 {
				fmt.Print(1, "\t")
			} else {
				fmt.Print(0, "\t")
			}
		}
		fmt.Println("\n")
	}
}

func main() {
	numsArr := []int{3, 4, 5} //List of integers each denoting number of binary variables in any particular operation

	for _, num := range numsArr {
		showBinaryCombinationsFor(num) //Calling a function that displays the binary combinations
	}
}

/*
0	0	0

1	0	0

0	1	0

1	1	0

0	0	1

1	0	1

0	1	1

1	1	1

0	0	0	0

1	0	0	0

0	1	0	0

1	1	0	0

0	0	1	0

1	0	1	0

0	1	1	0

1	1	1	0

0	0	0	1

1	0	0	1

0	1	0	1

1	1	0	1

0	0	1	1

1	0	1	1

0	1	1	1

1	1	1	1

0	0	0	0	0

1	0	0	0	0

0	1	0	0	0

1	1	0	0	0

0	0	1	0	0

1	0	1	0	0

0	1	1	0	0

1	1	1	0	0

0	0	0	1	0

1	0	0	1	0

0	1	0	1	0

1	1	0	1	0

0	0	1	1	0

1	0	1	1	0

0	1	1	1	0

1	1	1	1	0

0	0	0	0	1

1	0	0	0	1

0	1	0	0	1

1	1	0	0	1

0	0	1	0	1

1	0	1	0	1

0	1	1	0	1

1	1	1	0	1

0	0	0	1	1

1	0	0	1	1

0	1	0	1	1

1	1	0	1	1

0	0	1	1	1

1	0	1	1	1

0	1	1	1	1

1	1	1	1	1

*/
