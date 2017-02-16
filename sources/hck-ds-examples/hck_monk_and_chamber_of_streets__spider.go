/*
	@Date of creation : 31 Dec 2016.
	@Aim of program   : Monk and chamber of streets.
	@Go version       : 1.7.1 [go version command(on MAC )prints -> go version go1.7.1 darwin/amd64]
	@Coded by		  : Rishikesh Agrawani.
*/

package main

import "fmt"

//A structure that represents the node of Queue
type QueueNode struct {
	spidersPosition int
	power           int
	next            *QueueNode
}

var numOfNodesInQueue = 0

//A function that adds a binary tree's node into Queue
func enqueue(rear *QueueNode, power int, position int) *QueueNode {
	fmt.Println("Got rear as ==> ", rear)
	newQueueNode := new(QueueNode)
	newQueueNode.spidersPosition = position
	newQueueNode.power = power
	newQueueNode.next = nil

	if rear == nil {
		fmt.Println("fitting2....................................")
		rear = newQueueNode
	} else {
		fmt.Println("fitting....................................")
		rear.next = newQueueNode
		rear = rear.next
	}
	numOfNodesInQueue += 1
	return rear
}

//A function that removes a binary tree node from Queue
func dequeue(front *QueueNode) *QueueNode {
	if front != nil {
		var dequeuedNode *QueueNode
		dequeuedNode = front
		front = front.next
		dequeuedNode.next = nil

		numOfNodesInQueue -= 1
	}
	return front
}

func getMaxPowerIndex(newFront *QueueNode) (*QueueNode, int) {
	if newFront == nil {
		return nil, -1
	}
	var max_power_index int
	fmt.Println("max(init):", max_power_index)
	var max_power int = -1
	temp := newFront

	for temp != nil {

		if temp.power > max_power {
			max_power = temp.power
			max_power_index = temp.spidersPosition
		}
		if temp.power > 0 {
			temp.power = temp.power - 1
		}
		temp = temp.next
	}
	fmt.Println("Max Pow Index,Max Pow", max_power_index, max_power)
	//Dequeue the node with maximum power
	var tempBack *QueueNode = nil
	temp = newFront
	for temp.spidersPosition != max_power_index {
		fmt.Println(tempBack, temp)
		tempBack = temp
		temp = temp.next
	}
	numOfNodesInQueue -= 1
	fmt.Println("AAAA:", temp, tempBack)
	if temp != nil {
		tempBack.next = temp.next
		temp.next = nil
	}

	return newFront, max_power_index
}

func showNodes(top *QueueNode) {
	for top != nil {
		fmt.Println("POWER:", top.power)
		top = top.next
	}
}

func main() {
	var no_of_spiders, no_of_selection, power int
	var front, rear, dequeuedNodesPtr *QueueNode

	fmt.Scanf("%d%d", &no_of_spiders, &no_of_selection) //Acc. to problem variables are ---> N, X
	for i := 0; i < no_of_spiders; i++ {                //Spiders are in a Queue
		fmt.Scanf("%d", &power)
		fmt.Println("Now=>", rear)
		rear = enqueue(rear, power, i+1) //Adding, Power of ith spider,postion of spider into Queue's node
		fmt.Println("Now2=>", rear)
		if front == nil {
			front = rear
		}
		fmt.Println("Now2=>", rear)
		fmt.Println("--->>", front, rear, front.next)
		showNodes(front)
	}
	fmt.Println("Crossed the boundary...")
	/*The for loop that will produce the desired o/p */

	var max_power_index int
	for i := 0; i < no_of_selection; i++ {
		fmt.Println("::::>>", no_of_selection, numOfNodesInQueue)
		if no_of_selection > numOfNodesInQueue { //Dequeue all spiders
			fmt.Println("go_get****************1")
			showNodes(front)
			dequeuedNodesPtr = front
			fmt.Println("Tiiing...")
			front, rear = nil, nil
		} else { //Dequeue 1st num_of_selection spiders
			fmt.Println("go_get****************2")
			showNodes(front)
			fmt.Println("Tiiing2...")
			j := 1
			var temp *QueueNode
			dequeuedNodesPtr = front
			for j < no_of_selection {
				fmt.Println("traversal data=>", front.power)
				front = front.next
				fmt.Println("got")

				j += 1
			}
			fmt.Println("coming...")
			temp = front
			front = front.next
			temp.next = nil //Dividing the list in 2 parts
		}
		fmt.Println("coming <...>")
		fmt.Println("...M...")
		showNodes(dequeuedNodesPtr)
		//Finding max power spider
		dequeuedNodesPtr, max_power_index = getMaxPowerIndex(dequeuedNodesPtr)

		fmt.Println("max index:", max_power_index)
		//Again enqueue the qequeued nodes
		fmt.Println("rear==>", rear)
		if front == nil {
			fmt.Println("Getting...")
			front = dequeuedNodesPtr
			rear = front
			for rear.next != nil {
				rear = rear.next
			}
		} else {
			fmt.Println("Getting2")
			rear.next = dequeuedNodesPtr
			rear = front
			for rear.next != nil {
				rear = rear.next
			}
		}
	}

}
