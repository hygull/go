/*
	{	"Date of creation" => "Fri Dec 16 2016"	}
	{	"Aim of program"   => "Inorder order traversal in BT(Non recursive implementation)" }
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

//This function traverses a binary tree in inorder
func inOrderTraversal(top *Node) {
	if top == nil {
		return
	}
	inOrderTraversal(top.left)
	fmt.Print(top.data, "\t")
	inOrderTraversal(top.right)
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
	root.right.left.right = nodeCreator(-47)

	inOrderTraversal(root)
	fmt.Println("\n")
}

/* Diagram of the Binary Tree created by the above program :-
              	    56
              	   /  \
              	  95   34
              	 /    /  \
              -12    10   51
                    /  \    \
                  23   -47  -57
---------------------------------------------------------*/

/* OUTPUT(Preorder traversal) :-

-12	95	56	23	10	-47	34	51	-57

*/
