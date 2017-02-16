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
	newQueueNode := new(QueueNode)
	newQueueNode.spidersPosition = position
	newQueueNode.power = power
	newQueueNode.next = nil

	if rear == nil {
		rear = newQueueNode
	} else {
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
	var max_power_index int
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
	//Dequeue the node with maximum power
	var tempBack *QueueNode = nil
	temp = newFront

	for temp.spidersPosition != max_power_index {
		tempBack = temp
		temp = temp.next
	}
	numOfNodesInQueue -= 1
	if temp != nil {
		if tempBack != nil {
			tempBack.next = temp.next
			temp.next = nil
		} else {
			newFront = temp.next
		}

	}

	return newFront, max_power_index
}

func main() {
	var no_of_spiders, no_of_selection, power int
	var front, rear, dequeuedNodesPtr *QueueNode

	fmt.Scanf("%d%d", &no_of_spiders, &no_of_selection) //Acc. to problem variables are ---> N, X
	for i := 0; i < no_of_spiders; i++ {                //Spiders are in a Queue
		fmt.Scanf("%d", &power)
		rear = enqueue(rear, power, i+1) //Adding, Power of ith spider,postion of spider into Queue's node
		if front == nil {
			front = rear
		}
	}

	var max_power_index int
	/*The for loop that will produce the desired o/p */
	for i := 0; i < no_of_selection; i++ {
		if no_of_selection > numOfNodesInQueue { //Dequeue all spiders
			dequeuedNodesPtr = front
			front = nil
		} else { //Dequeue 1st num_of_selection spiders
			j := 1
			var temp *QueueNode
			dequeuedNodesPtr = front
			for j < no_of_selection {
				front = front.next
				j += 1
			}
			temp = front
			front = front.next
			temp.next = nil //Dividing the list in 2 parts
		}
		//Finding max power spider
		dequeuedNodesPtr, max_power_index = getMaxPowerIndex(dequeuedNodesPtr)
		fmt.Print(max_power_index, " ")
		//Again enqueue the qequeued nodes
		if front == nil {
			front = dequeuedNodesPtr
		} else {
			rear.next = dequeuedNodesPtr
			rear = front
			for rear.next != nil {
				rear = rear.next
			}
		}
	}

}
