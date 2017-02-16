/*
	{	"Date of creation" => "Fri Dec 16 2016"	}
	{	"Aim of program"   => "Postorder order traversal in BT(Recursive implementation)" }
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

//This function traverses a binary tree in postorder
func postOrderTraversal(top *Node) {
	if top == nil {
		return
	}
	postOrderTraversal(top.left)
	postOrderTraversal(top.right)
	fmt.Print(top.data, "\t")
}

//Starter function
func main() {
	root := nodeCreator(89)
	root.left = nodeCreator(35)
	root.right = nodeCreator(14)
	root.left.left = nodeCreator(-12)
	root.right.left = nodeCreator(10)
	root.right.right = nodeCreator(51)
	root.left.left.left = nodeCreator(67)
	root.right.left.left = nodeCreator(23)
	root.right.left.right = nodeCreator(-47)
	root.right.right.right = nodeCreator(-39)

	postOrderTraversal(root)
	fmt.Println("\n")
}

/* Diagram of the Binary Tree created by the above program :-
              	    89
              	   /  \
              	  35   14
              	 /    /  \
              -12    10   51
               /    /  \    \
             67    23   -47  -39
---------------------------------------------------------*/

/* OUTPUT(Postorder traversal) :-

67	-12	35	23	-47	10	-39	51	14	89

*/
