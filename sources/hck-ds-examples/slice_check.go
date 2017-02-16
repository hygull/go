/*
After 01:18 a.m
2 january 2016
Great though while solving a problem realated to an array.
https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/mark-the-answer-1/
*/
package main

import "fmt"

func main() {
	var questions, maxDifficultyLevel, maxMarks int //Number of questions, Maximum difficulty level
	var questionsSkipped bool = false
	var difficultyLevel []int

	fmt.Scanf("%d%d", &questions, &maxDifficultyLevel)

	//Taking the input for difficulty level of each questions
	for i := 0; i < questions; i++ {
		var d int
		fmt.Scanf("%d", &d) //
		difficultyLevel = append(difficultyLevel, d)

		if difficultyLevel[i] <= maxDifficultyLevel {
			maxMarks += 1
		} else {
			if questionsSkipped {
				break
			}
			questionsSkipped = true
		}
	}
	fmt.Println(maxMarks)
}
