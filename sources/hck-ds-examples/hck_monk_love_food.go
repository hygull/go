/*
  Creation date  : 29/12/2016.
  Problem's link : https://www.hackerearth.com/practice/data-structures/stacks/basics-of-stacks/tutorial/
  Aim of program : To help the manager to list out the costs for food packages.
  Coded by       : rishikesh Agrawani.
*/

package main

import "fmt"

type FoodPackage struct {
	costOfPackage int
	next          *FoodPackage
}

func CustomerQuery(pileTop *FoodPackage) (*FoodPackage, *FoodPackage) { //Type-1 Query
	if pileTop == nil {
		return nil, nil
	}

	topPackage := pileTop
	pileTop = pileTop.next
	topPackage.next = nil
	return pileTop, topPackage
}

func ChefQuery(pileTop *FoodPackage, cost int) *FoodPackage { //Type-2 Query
	newFoodPack := new(FoodPackage)
	newFoodPack.costOfPackage = cost

	if pileTop == nil {
		pileTop = newFoodPack //New food package added to pile of food packages
	} else {
		newFoodPack.next = pileTop
		pileTop = newFoodPack
	}
	return pileTop
}

func main() {
	var queries, queryType, cost int
	var pileTop, poppedFoodPack *FoodPackage
	var costsSlice []interface{}

	fmt.Scanf("%d", &queries) //No. of queries

	for i := 0; i < queries; i++ {
		fmt.Scanf("%d", &queryType) //1 or 2
		if queryType == 1 {
			pileTop, poppedFoodPack = CustomerQuery(pileTop)
			if poppedFoodPack == nil {
				costsSlice = append(costsSlice, "No Food") //A message if no food package is available
			} else {
				costsSlice = append(costsSlice, poppedFoodPack.costOfPackage)
			}
		} else {
			fmt.Scanf("%d", &cost)
			pileTop = ChefQuery(pileTop, cost)
		}
	}
	for _, cost := range costsSlice {
		fmt.Println(cost)
	}
}
