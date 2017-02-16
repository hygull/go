/*
	{	"Date of creation" => "Fri Dec 16 2016"	}
	{	"Aim of program"   => "Preorder order traversal in BT(Recursive implementation)" }
	{	"Coded by"         => "Rishikesh Agrawani"	}
	{	"Go version"	   => "1.7"	}
*/
package main

import "fmt"

/* To define a data structure that represents a Node of the binary tree */
type Node struct {
	data  int
	left  *Node
	right *Node
}

/* This function creates a new Node and returns it to the caller */
func nodeCreator(item int) *Node {
	newNode := new(Node)
	newNode.data = item
	newNode.left = nil
	newNode.right = nil
	return newNode
}

//This function traverses a binary tree in preorder
func preOrderTraversal(top *Node) {
	if top == nil {
		return
	}
	fmt.Print(top.data, "\t")
	preOrderTravesal(top.left)
	preOrderTravesal(top.right)
}

//Starter function
func main() {
	root := nodeCreator(56)
	root.left = nodeCreator(95)
	root.right = nodeCreator(34)
	root.left.left = nodeCreator(-12)
	root.right.left = nodeCreator(10)
	root.right.left.left = nodeCreator(23)
	root.right.right = nodeCreator(51)
	root.right.right.right = nodeCreator(-57)

	preOrderTraversal(root)
	fmt.Println("\n")
}

/* Diagram of the Binary Tree created by the above program :-

		56
	   /  \
	  95  34
	 /   /  \
  -12   10   51
       /       \
      23       -57

---------------------------------------------------------*/

/* OUTPUT(Preorder traversal) :-

56	95	-12	34	10	23	51	-57

*/
