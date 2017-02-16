/*
	{	"Date of creation" => "Fri Dec 30 2016"	}
	{	"Aim of program"   => "Deleting items from slice" }
	{	"Coded by"         => "Rishikesh Agrawani"	}
	{	"Go version"	   => "1.7"	}

package main

import "fmt"

func main() {
	names := []string{"Rishikesh", "Hemkesh", "Malinikesh", "Bhagya", "Chhoti", "Kirit"}
	fmt.Println(names)

	deletingIndexes := []int{1, 5}
	itemsDeleted := 0
	newSlice := []int{}

	for
	fmt.Println(names)
}

....*/

package main

import "fmt"

func main() {
	names := []string{"Rishikesh", "Hemkesh", "Malinikesh", "Bhagya", "Chhoti", "Kirit"}
	fmt.Println(names, len(names))
	//blankSlice := []int{}

	deletingIndexes := []int{1, 5}
	for _, delIndx := range deletingIndexes {
		fmt.Println("DelIndex => ", delIndx)
		for index, _ := range names {
			if index == delIndx {
				if index == len(names)-1 {
					names = append(names[:index])
					fmt.Println("=>", names, len(names))
				} else {
					names = append(names[:index], names[index+1:]...)
					fmt.Println("==>", index, names, len(names))
				}
			}
		}
	}
	fmt.Println(names, len(names))
}
