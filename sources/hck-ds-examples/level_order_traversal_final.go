/*
	@Date of creation : 31 Dec 2016.
	@Aim of program   : Level order traversal of binary tree.
	@Go version       : 1.7.1 [go version command(on MAC )prints -> go version go1.7.1 darwin/amd64]
	@Coded by		  : Rishikesh Agrawani.
*/

package main

import "fmt"

//A structure that represents the node of Queue
type QueueNode struct {
	btNodePtr *BtNode
	next      *QueueNode
}

//A structure that represents the node of binary tree
type BtNode struct {
	data  int
	left  *BtNode
	right *BtNode
}

//A function that creates a node for binary tree
func nodeCreatorForBt(item int) *BtNode {
	newNode := new(BtNode)
	newNode.data = item
	newNode.left = nil  //This line is not required, as nil is the default value for pointers....
	newNode.right = nil //It's just for clear understanding, what actually happening behind the scene.

	return newNode
}

//A  function that creates a node for Queue
func nodeCreatorForQueue(binaryTreeNodePtr *BtNode) *QueueNode {
	newNode := new(QueueNode)
	newNode.btNodePtr = binaryTreeNodePtr
	newNode.next = nil
	return newNode
}

//A function that adds a binary tree's node into Queue
func enqueue(rear *QueueNode, btNode *BtNode) *QueueNode {
	newQueueNode := nodeCreatorForQueue(btNode)
	if rear == nil {
		rear = newQueueNode
	} else {
		rear.next = newQueueNode
		rear = rear.next
	}

	return rear
}

//A function that removes a binary tree node from Queue
func dequeue(front *QueueNode) (*QueueNode, *QueueNode) {
	if front == nil {
		return nil, nil
	}

	var dequeuedNode *QueueNode
	dequeuedNode = front
	front = front.next
	dequeuedNode.next = nil

	return front, dequeuedNode
}

//A function that traverses binary tree in level order(like in BFS)
func levelOrderTraversal(root *BtNode) []int {
	listOfNodes := []int{}
	var dequeuedNode *QueueNode
	var front, rear *QueueNode

	if root == nil {
		return listOfNodes
	}
	rear = enqueue(rear, root)
	front = rear
	for front != nil {
		front, dequeuedNode = dequeue(front) //dequeue 1 node
		if dequeuedNode == nil {             //If Queue becomes empty
			return listOfNodes
		}
		listOfNodes = append(listOfNodes, dequeuedNode.btNodePtr.data) /* Append the information stored at dequeued node to list*/

		//If there is a left child of current node, then enqueue it
		if dequeuedNode.btNodePtr.left != nil {
			rear = enqueue(rear, dequeuedNode.btNodePtr.left) //enqueue the left child
			if front == nil {
				front = rear
			}
		}

		//If there is a right child of current node, then enqueue it
		if dequeuedNode.btNodePtr.right != nil {
			rear = enqueue(rear, dequeuedNode.btNodePtr.right) //enqueue the right child
			if front == nil {
				front = rear
			}
		}
	}
	return listOfNodes
}

//Starter function
func main() {
	var root *BtNode

	/* TEST CASE 1 */
	root = nodeCreatorForBt(12)
	root.left = nodeCreatorForBt(14)
	root.right = nodeCreatorForBt(76)
	root.left.left = nodeCreatorForBt(-89)
	root.left.left.left = nodeCreatorForBt(75)
	root.left.right = nodeCreatorForBt(23)
	root.left.right.right = nodeCreatorForBt(34)
	root.right.left = nodeCreatorForBt(11)
	root.right.right = nodeCreatorForBt(79)
	root.right.right.left = nodeCreatorForBt(79)
	root.right.right.right = nodeCreatorForBt(-5)

	fmt.Println(levelOrderTraversal(root))

	/* TEST CASE 2 */
	root = nil
	root = nodeCreatorForBt(34)
	root.left = nodeCreatorForBt(-31)
	root.right = nodeCreatorForBt(0)
	fmt.Println(levelOrderTraversal(root))

	/* TEST CASE 3 */
	root = nil
	fmt.Println(levelOrderTraversal(root))
}

/* TEST CASE 1:-
					12<---root
				   /  \
                 14    76
                /  \   / \
			  -89  23 11  79
			  /      \    / \
			 75      34  45  -5


OUTPUT :-
			[12 14 76 -89 23 11 79 75 34 79 -5]
--------------------------------------------------------------------*/

/* TEST CASE 2:-

				 34<---root
			     /\
			  -31  0


OUTPUT:-
			[34 -31 0]
--------------------------------------------------------------------*/

/*
	root--->nil

OUTPUT:-
			[]
*/
