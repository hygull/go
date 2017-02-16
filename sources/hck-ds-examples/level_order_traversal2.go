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
	fmt.Println("Start front ==> ", front)
	if front == nil {
		return nil, nil
	}
	fmt.Println("dequeeing")
	var dequeuedNode *QueueNode
	dequeuedNode = front
	fmt.Println(front, front.next)
	front = front.next

	fmt.Println("this front now...", front)
	dequeuedNode.next = nil

	return front, dequeuedNode
}

func showNodes(top *QueueNode) {
	for top != nil {
		fmt.Println(top.btNodePtr.data)
		top = top.next
	}

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
		fmt.Println("deq=>", dequeuedNode.btNodePtr.data)
		fmt.Print("=>")
		showNodes(front)
		listOfNodes = append(listOfNodes, dequeuedNode.btNodePtr.data) /* Append the information stored at dequeued node to list*/

		fmt.Println("list => ", listOfNodes)
		//If there is a left child of current node, then enqueue it
		if dequeuedNode.btNodePtr.left != nil { // *front.btNodePtr.left != nil --> front.btNodePtr undefined (type **QueueNode has no field or method btNodePtr)
			//node := nodeCreatorForQueue(dequeuedNode.btNodePtr.left)
			fmt.Println("L===>")
			rear = enqueue(rear, dequeuedNode.btNodePtr.left) //enqueue the left child
			if front == nil {
				front = rear
			}
			// fmt.Print("...>", rear, *rear)
			// fmt.Print("..>", front, *front)
			// fmt.Print(".....>", dequeuedNode.btNodePtr.left)
			// fmt.Print("adfggg")
			// fmt.Print("\t", dequeuedNode.btNodePtr.left.data)
		}

		//If there is a right child of current node, then enqueue it
		if dequeuedNode.btNodePtr.right != nil {
			fmt.Println("R===>")

			rear = enqueue(rear, dequeuedNode.btNodePtr.right) //enqueue the right child
			if front == nil {
				front = rear
			}
			// fmt.Print("\t", dequeuedNode.btNodePtr.right.data)
		}
		// fmt.Print("...", *front)
	}
	return listOfNodes
}

//Starter function
func main() {
	var root *BtNode

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

	// rear = nodeCreatorForQueue(root)
	// front = rear //As there is only 1 node in Queue, So front and rear both will point to the same
	//fmt.Print(front, rear)
	fmt.Println(levelOrderTraversal(root))
}

/* TEST CASE 1:-
					12
				   /  \
                 14    76
                /  \   / \
			  -89  23 11  79
			  /      \    / \
			 75      34  45  -5
*/
