/*
	{	"Date of creation" => "Sat Dec 17 2016"	}
	{	"Aim of program"   => "Preorder order traversal in BT(Non Recursive implementation using Stack)" }
	{	"Coded by"         => "Rishikesh Agrawani"	}
	{   "special"          => "Linked list implementation of Stack,Thought of the day"}
	{	"Go version"	   => "1.7"	}
*/
package main

import "fmt"

/* A stack that will hold the addresses of binary tree's nodes */
type StackNode struct {
	btNodePtr *BtNode    //A pointer to the node of binary tree
	next      *StackNode //A pointer to next node in the linkedlist
}

/* To define a data structure that represents a Node of the binary tree */
type BtNode struct {
	data  int
	left  *BtNode
	right *BtNode
}

/* This function creates a new Node For Stack returns it to the caller */
func nodeCreatorForStack(binaryTreeNodePtr *BtNode) *StackNode {
	newNode := new(StackNode)
	newNode.btNodePtr = binaryTreeNodePtr
	newNode.next = nil
	return newNode
}

/* This function pushes a binary tree's node into Stack*/
func Push(start *StackNode, node *BtNode) *StackNode {
	//Node creator code for stack, that stores a binary tree's node in data field
	newStackNode := new(StackNode)
	newStackNode.btNodePtr = node
	newStackNode.next = nil

	//insert as a starting node
	if start == nil {
		start = newStackNode
		return start
	}
	//Insert at starting and maintaining the links
	newStackNode.next = start
	start = newStackNode
	return start
}

/* This function pops the binary tree's node from stack */
func Pop(start **StackNode) *StackNode {
	var node *StackNode
	if *start != nil {
		node = *start
		*start = (*start).next
		return node
	}
	return nil
}

//A function that displays the data stored on right child of processed nodes(Current status, It will vary based on Push & Pop operations)
func showNodes(top *StackNode) {
	var stackNodesValues []int //A slice that will contain the data stored on each node of the stack
	for top != nil {
		stackNodesValues = append(stackNodesValues, top.btNodePtr.data)
		top = top.next
	}
	fmt.Println("Stack : ", stackNodesValues)
}

/* This function creates a new Node For Binary tree and returns it to the caller */
func nodeCreatorForBt(item int) *BtNode {
	newNode := new(BtNode)
	newNode.data = item
	newNode.left = nil
	newNode.right = nil
	return newNode
}

//Traversing a binary tree in preorder
func preOrderTraversal(start **BtNode, top **StackNode) []int {
	btNodesValues := []int{} //A slice that will contain the data stored on the processed nodes
	btNodePtr := *start
	for btNodePtr != nil { //binary tree's current node pointer is not nil...
		showNodes(*top)
		btNodesValues = append(btNodesValues, btNodePtr.data)
		fmt.Println("List of values from processed nodes : ", btNodesValues)
		if btNodePtr.right != nil {
			*top = Push(*top, btNodePtr.right)
		}
		if btNodePtr.left != nil {
			btNodePtr = btNodePtr.left
		} else { //If there is no left child of the current node then pop a binary tree's node from stack
			poppedStackNode := Pop(top)
			if poppedStackNode == nil { //If there's no node on stack then return to main()
				return btNodesValues
			} else {
				btNodePtr = poppedStackNode.btNodePtr //Set the current node to popped one
			}
		}
	}
	return btNodesValues //[]int{}...empty slice
}

//Starter function
func main() {

	root := nodeCreatorForBt(50)
	root.left = nodeCreatorForBt(35)
	root.right = nodeCreatorForBt(14)
	root.left.left = nodeCreatorForBt(-31)
	root.right.left = nodeCreatorForBt(11)
	root.right.right = nodeCreatorForBt(51)
	root.left.left.left = nodeCreatorForBt(67)
	root.right.left.left = nodeCreatorForBt(23)
	root.right.left.right = nodeCreatorForBt(-47)
	root.right.right.right = nodeCreatorForBt(-31)

	var top *StackNode //Top pointer of Stack

	processedNodesValues := preOrderTraversal(&root, &top)
	fmt.Println("\nThe list of values stored on the nodes(In preorder taversal of binary tree) : ", processedNodesValues)
}

/* Diagram of the Binary Tree created by the above program :-
              	    50
              	   /  \
              	  35   14
              	 /    /  \
               -31   11   51
               /    /  \    \
             67    23   -47  -31
---------------------------------------------------------*/

/* OUTPUT :- preorder traversal of binary tree

Stack :  []
List of values from processed nodes :  [50]
Stack :  [14]
List of values from processed nodes :  [50 35]
Stack :  [14]
List of values from processed nodes :  [50 35 -31]
Stack :  [14]
List of values from processed nodes :  [50 35 -31 67]
Stack :  []
List of values from processed nodes :  [50 35 -31 67 14]
Stack :  [51]
List of values from processed nodes :  [50 35 -31 67 14 11]
Stack :  [-47 51]
List of values from processed nodes :  [50 35 -31 67 14 11 23]
Stack :  [51]
List of values from processed nodes :  [50 35 -31 67 14 11 23 -47]
Stack :  []
List of values from processed nodes :  [50 35 -31 67 14 11 23 -47 51]
Stack :  []
List of values from processed nodes :  [50 35 -31 67 14 11 23 -47 51 -31]

The list of values stored on the nodes(In preorder taversal of binary tree) :  [50 35 -31 67 14 11 23 -47 51 -31]

*/
