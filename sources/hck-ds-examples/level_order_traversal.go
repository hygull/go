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
	return newNode
}

//A function that adds a binary tree's node into Queue
func enqueue(rear **QueueNode, front **QueueNode, btNode *BtNode) **QueueNode {
	fmt.Println("hen....", rear)
	//fmt.Print((*rear).btNodePtr.data)
	newQueueNode := nodeCreatorForQueue(btNode)
	fmt.Println("hen....")
	if rear == nil {
		fmt.Println("hen..*..")
		//(*rear).btNodePtr = btNode
		rear = &newQueueNode
		front = rear
	}

	// newQueueNode.next = *rear
	// *rear = newQueueNode
	fmt.Println("hen...*.")
	(*rear).next = newQueueNode
	fmt.Println("hen....")
	*rear = newQueueNode
	return rear
}

//A function that removes a binary tree node from Queue
func dequeue(front *QueueNode, rear **QueueNode) (*QueueNode, *QueueNode) {
	fmt.Print("uhhhhhh", front, ",,,,,", front.btNodePtr.data)
	if front == nil {
		return nil, nil
	}
	var dequeuedNode *QueueNode
	dequeuedNode = front
	dequeuedNode.next = nil
	front = front.next
	fmt.Print("bbggg..", front)
	if front == nil {
		*rear = front
	}
	//fmt.Println("....>", *rear, front)....OK...<nil>....<nil>
	//fmt.Print("....>>", dequeuedNode, dequeuedNode.btNodePtr.data)....add,<nil>..OK
	fmt.Print("hhhhjj", *rear)
	return front, dequeuedNode
}

//A function that traverses binary tree in level order(like in BFS)
func levelOrderTraversal(root *BtNode, front **QueueNode, rear **QueueNode) []int {
	listOfNodes := []int{}
	var dequeuedNode *QueueNode

	if root == nil {
		return listOfNodes
	}
	*front = nodeCreatorForQueue(root)
	*rear = *front
	//fmt.Print(*front)...OK
	//fmt.Print(*rear)....OK
	for *front != nil {
		*front, dequeuedNode = dequeue(*front, rear) //dequeue 1 node
		if dequeuedNode == nil {                     //If Queue becomes empty
			return listOfNodes
		}

		fmt.Printf("%d...\t", dequeuedNode.btNodePtr.data)             //Print the information stored at dequeued node
		listOfNodes = append(listOfNodes, dequeuedNode.btNodePtr.data) /* Append the information stored at dequeued node to list
		of informations stored at processed nodes */

		//If there is a left child of current node, then enqueue it
		if dequeuedNode.btNodePtr.left != nil { // *front.btNodePtr.left != nil --> front.btNodePtr undefined (type **QueueNode has no field or method btNodePtr)
			//node := nodeCreatorForQueue(dequeuedNode.btNodePtr.left)
			rear = enqueue(rear, front, dequeuedNode.btNodePtr.left) //enqueue the left child

			fmt.Print("...>", rear, *rear)
			// fmt.Print("..>", front, *front)
			// fmt.Print(".....>", dequeuedNode.btNodePtr.left)
			// fmt.Print("adfggg")
			// fmt.Print("\t", dequeuedNode.btNodePtr.left.data)
		}

		//If there is a right child of current node, then enqueue it
		if dequeuedNode.btNodePtr.right != nil {
			// fmt.Print("76y6777")
			rear = enqueue(rear, front, dequeuedNode.btNodePtr.right) //enqueue the right child
			// fmt.Print("\t", dequeuedNode.btNodePtr.right.data)
		}
		fmt.Print("...", *front)
	}
	return listOfNodes
}

//Starter function
func main() {
	var root *BtNode

	root = nodeCreatorForBt(12)
	root.left = nodeCreatorForBt(14)
	root.right = nodeCreatorForBt(76)

	var front, rear *QueueNode
	// rear = nodeCreatorForQueue(root)
	// front = rear //As there is only 1 node in Queue, So front and rear both will point to the same
	//fmt.Print(front, rear)
	fmt.Println(levelOrderTraversal(root, &front, &rear))
}
