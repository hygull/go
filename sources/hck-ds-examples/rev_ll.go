/*
	{
		"date_of_creation" => "14 Dec 2016, Wed",
		"aim_of_program"   => "To reverse a singly linked list using iterative method",
		"coded_by"         => "Rishikesh Agrawani",
		"memory"          => "I wrote code in C on 13 Dec's evening on a paper with diagram and implemented on 14 Dec's morning, executed it on office(1 run execution)"
		"Go_version"	   => "1.7",
	}
*/

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

/*This function will create a complete independent Node and returns it to the caller*/
func createNode(item int) *Node {
	newNode := new(Node) //Allocate a memory for Node
	newNode.data = item  //Store item's value in the allocated memory for Node
	newNode.next = nil   //Node's next pointer will be empty(As this is independent now)

	return newNode //Return the newly created node
}

/* This function inserts a node at end and returns the address of first node to the caller */
func insertNodeAtEnd(root *Node, node *Node) *Node {
	if root == nil { //If linked list is empty
		root = node //Set the node as first Node of the linked list
		return root //return root to the caller
	}
	temp := root
	for temp.next != nil { //Moving to the last Node
		temp = temp.next
	}
	temp.next = node //Linking the new Node with the last Node in the linked list
	return root
}

/* This function displays all the node's information of singly linked list*/
func showNodes(node *Node) {
	if node == nil { //If linked list is empty
		fmt.Println("Singly linked lis is empty")
		return
	}
	temp := node
	for temp != nil { //Iterating through all the nodes of linked list
		fmt.Print(temp.data, "\t") //Printing data available on each node
		temp = temp.next           //Forwarding pointer to the next available node
	}
	fmt.Println() //Newline
}

//A function that reverses the linked list and returns the pointer to first node
func reverseSinglyLinkedList(node *Node) *Node {
	//If linked list is empty
	if node == nil {
		return nil
	}
	//If only 1 node is there
	if node.next == nil {
		return node
	}
	//If more than 2 nodes are available in the linked list
	var forw, back *Node

	for node != nil {
		forw = node.next //Saving the address of next node
		node.next = back //Setting the current node's next so that it could point to the previous node
		back = node      //back pointer will point to a Node poited by pointer node
		node = forw      //node will point to the Node pointed by forw
	}
	node = back //At last node will poit to nil(None or NULL), back will point to last node of original linked list
	//Now it becomes the first node of reversed linked list, So we have to set it
	return node //Returning the address of first node of reversed linked list
}

func main() {
	//1st testcase
	var root *Node = nil //root:=nil  ->  /n.go:49: use of untyped nil
	intsSlice := []int{12, -34, 56, 0, -23, 75, 21}

	for _, num := range intsSlice {
		node := createNode(num)            //Creating a Node with Node information as num's value
		root = insertNodeAtEnd(root, node) //Calling function to insert Node at end
	}

	fmt.Println("The information available on each node of linked list are as follows:-")
	showNodes(root) //Calling

	fmt.Println("Reversing the linked list...")
	root = reverseSinglyLinkedList(root)
	if root == nil {
		fmt.Println("linked list is empty")
		return
	}
	fmt.Println("The information available on each node of reversed linked list are as follows:-")
	showNodes(root) //Calling
}

/* OUTPUT:-

The information available on each node of linked list are as follows:-
12	-34	56	0	-23	75	21
Reversing the linked list...
The information available on each node of reversed linked list are as follows:-
21	75	-23	0	56	-34	12
*/

/********************************** OTHER TRIALS ***********************************/
/* []int{12,-34,56,	0,-23,75,21}
12	-34	56	0	-23	75	21
Reversing the linked list...
The information available on each node of reversed linked list are as follows:-
21	75	-23	0	56	-34	12
*/

/* []int{}
The information available on each node of linked list are as follows:-
Singly linked lis is empty
Reversing the linked list...
linked list is empty
*/

/* []int{12}
The information available on each node of linked list are as follows:-
12
Reversing the linked list...
The information available on each node of reversed linked list are as follows:-
12
*/
