package main

import "fmt"

func main() {
	var questions, maxDifficultyLevel, maxMarks int //Number of questions, Maximum difficulty level
	var questionsSkipped bool = false
	var difficultyLevel []int

	fmt.Scanf("%d", &questions, &maxDifficultyLevel)

	//Taking the input for difficulty level of each questions
	for i := 0; i < questions; i++ {
		var d int
		fmt.Scanf("%d", &d)
		difficultyLevel = append(difficultyLevel, d)

		if difficultyLevel[i] < maxDifficultyLevel {
			maxMarks += 1
		} else {
			if questionsSkipped {
				break
			}
			questionsSkipped = true
		}
	}

	// for i := 0; i < questions; i++ {
	// 	if difficultyLevel[i] < maxDifficultyLevel {
	// 		maxMarks += 1
	// 	} else {
	// 		if questionsSkipped {
	// 			break
	// 		}
	// 		questionsSkipped = true
	// 	}
	// }
	fmt.Println(maxMarks)
}

/*
package main

import "fmt"

func main() {
	var questions, maxDifficultyLevel, maxMarks int //Number of questions, Maximum difficulty level
	var questionsSkipped bool = false

	fmt.Scanf("%d%d", &questions, &maxDifficultyLevel)

	//Taking the input for difficulty level of each questions
	for i := 0; i < questions; i++ {
		var d int
		fmt.Scanf("%d", &d)
		if d >= maxDifficultyLevel {
			if questionsSkipped {
				break
			}
			questionsSkipped = true
		}
		maxMarks += 1
	}
	fmt.Println(maxMarks)
}

*/

/*
package main

import "fmt"

func main() {
	var questions, maxDifficultyLevel, maxMarks int //Number of questions, Maximum difficulty level
	var questionsSkipped bool = false

	fmt.Scanf("%d%d", &questions, &maxDifficultyLevel)
    difficultyLevel :=make([]int,questions,questions)//Creating a slice

	//Taking the input for difficulty level of each questions
	for i := 0; i < questions; i++ {
		fmt.Scanf("%d",&difficultyLevel[i] )

		if difficultyLevel[i] <= maxDifficultyLevel {
			maxMarks += 1
		} else {
			if questionsSkipped {
				break
			}
			questionsSkipped = true
		}
	}
	difficultyLevel=nil
	fmt.Println(maxMarks)
}

*/
