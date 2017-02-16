/*
	{
		"date_of_creation" => "15 Dec 2016, Thurs(evening)",
		"aim_of_program"   => "Doubly linked list(Inserting nodes at end and traversing in both directions)",
		"coded_by"         => "Rishikesh Agrawani",
		"memory"           => "Thought of the day, Type assertion",
		"resources"		   => "{ https://tour.golang.org/methods/15 (A tour of Go : Type assertions) },
							   { http://stackoverflow.com/questions/14289256/cannot-convert-data-type-interface-to-type-string-need-type-assertion }",
		"Go_version"	   => "1.7",
	}
*/
package main

import "fmt"

/* Defining the Node of doubly linked list */
type DllNode struct {
	name     string   //Name of person
	age      int8     //Age of person
	isActive bool     //Person's activity status
	prev     *DllNode //It will point to the previous node if any, otherwise it will be <nil>
	next     *DllNode //It will point to the next node if any, otherwise it will be <nil>
}

/* A function that creates node and returns it to the caller */
func nodeCreator(dataList ...interface{}) *DllNode {
	newNode := new(DllNode) //Storage allocation for DllNode

	for i, data := range dataList {
		if i == 0 {
			newNode.name = data.(string) //type assertion...otherwise -> annot use data (type interface {}) as type string in assignment: need type assertion
		} else if i == 1 {
			newNode.age = data.(int8) //type assertion...otherwise -> cannot use data (type interface {}) as type int8 in assignment: need type assertion
		} else {
			newNode.isActive = data.(bool) //type assertion...otherwise -> cannot use data (type interface {}) as type bool in assignment: need type assertion
		}
	}
	newNode.prev, newNode.next = nil, nil //Initially independent node's prev & next both pointers will point to <nil>
	return newNode
}

/* A function that inserts node at the end of the doubly linked list */
/* No need to traverse from front like singly linked list */
func nodeInsertorAtEnd(last *DllNode, node *DllNode) *DllNode {
	if last == nil {
		last = node
	} else {
		last.next = node
		node.prev = last
		last = node
	}
	return last
}

/* A function that displays the information of nodes from end (one by one)*/
func nodesDisplayerInBackwardDirection(last *DllNode) {
	for last == nil {
		fmt.Println("Doubly linked list is empty")
		return
	}
	for last != nil {
		fmt.Println("NODE -> ", last.name, ", ", last.age, ", ", last.isActive, "\n")
		last = last.prev
	}
}

/* A function that displays the information of nodes from start (one by one) */
func nodesDisplayerInForwardDirection(root *DllNode) {
	for root == nil {
		fmt.Println("Doubly linked list is empty")
		return
	}
	for root != nil {
		fmt.Println("NODE -> ", root.name, ", ", root.age, ", ", root.isActive, "\n")
		root = root.next
	}
}

func main() {
	var start, end, node *DllNode //Default values for start,end & node pointers is <nil>

	/* Defining data structure for storing details of any person*/
	type Person struct {
		name     string
		age      int8
		isActive bool
	}
	/* Creating 5 records related to person */
	person1 := Person{"Hemkesh", 22, true}
	person2 := Person{"Smarika", 19, true}
	person3 := Person{"Malinikesh", 20, false}
	person4 := Person{"Surendra", 24, true}
	person5 := Person{"Rishikesh", 24, false}

	/* Putting all the records in a slice */
	persons := []Person{person1, person2, person3, person4, person5}

	/*
		Iterating through slice, creating node and storing details in that node, inserting node(s)
		at the end of the doubly linked list
	*/
	for i, person := range persons {
		node = nodeCreator(person.name, person.age, person.isActive)
		end = nodeInsertorAtEnd(end, node)
		/*
		   start will always point to the beginning node in the list, as our intention is to insert
		   nodes at end so the 1st inserted node will be in the beginning after insertion(s).
		*/
		if i == 0 {
			start = end
		}
	}
	fmt.Println("Displaying nodes information of Doubly linked list( Forward direction ):\n")
	nodesDisplayerInForwardDirection(start)

	fmt.Println("\nDisplaying nodes information of Doubly linked list( Backward direction ):\n")
	nodesDisplayerInBackwardDirection(end)
}

/*OUTPUT :-

Displaying nodes information of Doubly linked list( Forward direction ):

NODE ->  Hemkesh ,  22 ,  true

NODE ->  Smarika ,  19 ,  true

NODE ->  Malinikesh ,  20 ,  false

NODE ->  Surendra ,  24 ,  true

NODE ->  Rishikesh ,  24 ,  false


Displaying nodes information of Doubly linked list( Backward direction ):

NODE ->  Rishikesh ,  24 ,  false

NODE ->  Surendra ,  24 ,  true

NODE ->  Malinikesh ,  20 ,  false

NODE ->  Smarika ,  19 ,  true

NODE ->  Hemkesh ,  22 ,  true

*/
