/*
  Creation date  : 30/12/2016.
  Problem's link : https://www.hackerearth.com/practice/data-structures/queues/basics-of-queues/tutorial/
  Aim of program : To perform enqueue & dequeue operations on queue.
  Coded by       : Rishikesh Agrawani.
*/

package main

import "fmt"

//Defining a node for Queue
type QueueNode struct {
	x    int8
	next *QueueNode
}

var totalItemsInQueue int8

func Enqueue(rear *QueueNode, item int8) *QueueNode {
	newNode := new(QueueNode)
	newNode.x = item
	if rear == nil {
		rear = newNode
	} else {
		rear.next = newNode
		rear = rear.next
	}
	totalItemsInQueue += 1
	return rear
}

func Dequeue(front *QueueNode) (*QueueNode, *QueueNode) {
	var dequeuedNode *QueueNode
	if front != nil {
		dequeuedNode = front
		front = front.next
		dequeuedNode.next = nil
		totalItemsInQueue -= 1
	}

	return front, dequeuedNode
}
func main() {
	var n, item int8
	var queueOperationType string //E or D => E for Enqueue & D for Dequeue
	var rear, front, dequeuedNode *QueueNode

	fmt.Scanf("%d", &n) //storing a value into n
	hash := make(map[int8][]int8)
	var i int8
	for i = 0; i < n; i++ {
		fmt.Scanf("%s", &queueOperationType)
		if queueOperationType == "E" { //E => Enqueue
			fmt.Scanf("%d", &item)
			rear = Enqueue(rear, item)
			if front == nil {
				front = rear
			}
			hash[int8(i)] = []int8{totalItemsInQueue}
		} else { //D => Dequeue
			front, dequeuedNode = Dequeue(front)
			arr := []int8{}
			if front == nil {
				rear = front
			}
			if dequeuedNode == nil {
				arr = append(arr, -1, totalItemsInQueue)
			} else {
				arr = append(arr, dequeuedNode.x, totalItemsInQueue)
			}
			hash[int8(i)] = arr
		}
	}
	//Displaying the result
	for i = 0; i < n; i++ {
		for index, num := range hash[int8(i)] {
			if index == 0 {
				fmt.Printf("%d", num)
			} else {
				fmt.Printf(" %d", num)
			}
		}
		fmt.Println()
	}
}
