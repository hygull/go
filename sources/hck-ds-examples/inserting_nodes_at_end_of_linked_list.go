/*
	{	"Date of creation" => "Sat Dec 10 20:28:10 IST 2016"	}
	{	"Aim of program"   => "To create a singly linked list and traversing through it after inserting nodes at end" }
	{	"Coded by"         => "Rishikesh Agrawani"	}
	{	"Go version"	   => "1.7"	}
*/
package main

import "fmt"

/* Defining the data structure that represents the Singly linked list's Node */
type Node struct {
	data int8  //int8 type of data (in DATA field)
	next *Node //next is pointer to next Node (in NEXT POINTER field)
}

/*This function will create a complete independent Node and returns it to the caller*/
func createNode(item int8) *Node {
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

func main() {
	var root *Node = nil //root:=nil  ->  /n.go:49: use of untyped nil

	node := createNode(34)             //Creating a Node with Node information as 12
	root = insertNodeAtEnd(root, node) //Calling function to insert Node at end

	node = createNode(-15)             //Creating a Node with Node information as -15
	root = insertNodeAtEnd(root, node) //Calling function to insert Node at end

	node = createNode(0)               //Creating a Node with Node information as 0
	root = insertNodeAtEnd(root, node) //Calling function to insert Node at end

	node = createNode(67)              //Creating a Node with Node information as 67
	root = insertNodeAtEnd(root, node) //Calling function to insert Node at end

	fmt.Println("The information available on each node of linkeds list are as follows:-")
	showNodes(root) //Calling
}

/*

The information available on each node of linkeds list are as follows:-
34	-15	0	67

*/
