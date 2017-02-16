/*
	{	"Date of completion" => "Sat Dec 17 2016"	}
	{	"Aim of program"     => "Postorder order traversal of BT(Non Recursive implementation using Stack)" }
	{	"Coded by"           => "Rishikesh Agrawani"	}
	{   "special"            => "Linked list implementation of Stack,Thought of the day,
								 Very nice experience with the implementation of Postorder traversal in Go"}
	{	"Go version"	     => "1.7"	}
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
	data    int
	isRight bool //This is an extra information that is used to know about the node whether it was right child or not
	left    *BtNode
	right   *BtNode
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

var c = 0

//Traversing a binary tree in postorder
func postOrderTraversal(start **BtNode, top **StackNode) []int {
	btNodesValues := []int{} //A slice that will contain the data stored on the processed nodes
	btNodePtr := *start
	for btNodePtr != nil { //binary tree's current node pointer is not nil...
		showNodes(*top)
		for btNodePtr != nil {
			*top = Push(*top, btNodePtr) //Push current node into Stack
			showNodes(*top)
			if btNodePtr.right != nil { //If there is any right child of currently visited node
				btNodePtr.right.isRight = true     //Set its isRight field to true (This is just for knowing which node was right child)
				*top = Push(*top, btNodePtr.right) //Push right child of the curent node into Stack
				showNodes(*top)
			}
			btNodePtr = btNodePtr.left //Visit left subtree rooted at the current node
		}

		for true { //[Backtracking] If there is no left subtree, then Pop the node from stack until we don't find the node which was right child(isRight=true)
			poppedStackNode := Pop(top)

			if poppedStackNode == nil { //If there's no node on stack then return to main()
				return btNodesValues
			} else {

				if poppedStackNode.btNodePtr.isRight == true { //If popped node was right child
					poppedStackNode.btNodePtr.isRight = false //Set its isRight field to false
					btNodePtr = poppedStackNode.btNodePtr     //Set the current node to popped one
					break
				}

				btNodesValues = append(btNodesValues, poppedStackNode.btNodePtr.data) //Append the node in the list of processed nodes

				fmt.Println("List of values from processed nodes : ", btNodesValues)
			}
		}
	}
	return btNodesValues //[]int{}...empty slice
}

//Starter function
func main() {
	var root *BtNode

	root = nodeCreatorForBt(50)
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

	processedNodesValues := postOrderTraversal(&root, &top)
	fmt.Println("\nThe list of values stored on the nodes(In postorder taversal of binary tree) : ", processedNodesValues)
}

/* Diagram of the Binary Tree created by the above program :-
             root-->50
              	   /  \
              	  35   14
              	 /    /  \
               -31   11   51
               /    /  \    \
             67    23   -47  -31
---------------------------------------------------------*/

/* OUTPUT :- postorder traversal of binary tree  (67->-31->35->23->47->11->-31->51->14->50)

Stack :  []
Stack :  [50]
Stack :  [14 50]
Stack :  [35 14 50]
Stack :  [-31 35 14 50]
Stack :  [67 -31 35 14 50]
List of values from processed nodes :  [67]
List of values from processed nodes :  [67 -31]
List of values from processed nodes :  [67 -31 35]
Stack :  [50]
Stack :  [14 50]
Stack :  [51 14 50]
Stack :  [11 51 14 50]
Stack :  [-47 11 51 14 50]
Stack :  [23 -47 11 51 14 50]
List of values from processed nodes :  [67 -31 35 23]
Stack :  [11 51 14 50]
Stack :  [-47 11 51 14 50]
List of values from processed nodes :  [67 -31 35 23 -47]
List of values from processed nodes :  [67 -31 35 23 -47 11]
Stack :  [14 50]
Stack :  [51 14 50]
Stack :  [-31 51 14 50]
Stack :  [51 14 50]
Stack :  [-31 51 14 50]
List of values from processed nodes :  [67 -31 35 23 -47 11 -31]
List of values from processed nodes :  [67 -31 35 23 -47 11 -31 51]
List of values from processed nodes :  [67 -31 35 23 -47 11 -31 51 14]
List of values from processed nodes :  [67 -31 35 23 -47 11 -31 51 14 50]

The list of values stored on the nodes(In preorder taversal of binary tree) :  [67 -31 35 23 -47 11 -31 51 14 50]

*/

/* Other test cases:-
....(1) When there is only 1 node in the binary tree

         root-->50
               /  \
             nil   nil
    O/P:
        Stack :  []
		Stack :  [50]

		List of values from processed nodes :  [50]

		The list of values stored on the nodes(In preorder taversal of binary tree) :  [50]

....(2) When binary tree is empty

		root-->nil

	O/P:
		The list of values stored on the nodes(In preorder taversal of binary tree) :  []
*/
