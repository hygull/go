package main

import "fmt"

func main() {
	var testcases, no_of_drivers, i, j int
	fmt.Scanf("%d", &testcases)

	for i = 0; i < testcases; i++ {
		fmt.Scanf("%d", &no_of_drivers)
		var sight int = 0

		var position int = 0

		a := make([]int, no_of_drivers)
		//fmt.Println(a)
		for j = 0; j < no_of_drivers; j++ {
			fmt.Scanf("%d", &a[j])
		}
		//fmt.Println(a)
		for j = 0; j < no_of_drivers; j++ {
			var b, f int
			var frontdrivers int = 0
			var backdrivers int = 0
			//fmt.Println("Came")
			for f = j - 1; f >= 0; f-- {
				//fmt.Println(f, j)
				//fmt.Println("Comparing", f, a[f], "and", j, a[j])
				if a[f] < a[j] {
					frontdrivers += 1
				}
			}
			//fmt.Println("jth drivers front", frontdrivers)
			for b = j + 1; b < no_of_drivers; b++ {
				if a[b] < a[j] {
					backdrivers += 1
				}
			}
			//fmt.Println("jth drivers back", backdrivers)
			var new_sight int = (frontdrivers + backdrivers) * (j + 1)
			if new_sight > sight {
				sight = new_sight
				position = j + 1
			}
		}
		fmt.Println(position)
	}
}
