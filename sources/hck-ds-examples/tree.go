/*
	{	"date_of_creation" : "Sun Dec 11 2016"	},
	{	"aim_of_program"   : "To create a binary tree data structure and traverse through it" },
	{	"coded_by"         : "Rishikesh Agrawani"	},
	{	"Go_version"	   : "1.7"	},
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

/* This function gives you a chance to create a binary tree as you want by inserting nodes at choosen place*/
func insertNodeOnTreeAccToChoice(root *Node, node *Node) *Node {
	/* If tree is empty */
	if root == nil {
		root = node
		fmt.Println("[...] Initially tree is empty.So Node is inserted as the root of the tree and... INFO(root) : ", root.data)
		return root
	}
	/* If tree is not empty */
	current := root
	println("Current : ", current)
	if current.left != nil && current.right != nil { //If left & right both childs exist for the current Node(pointed by current)
		fmt.Println("[...] The currently visited Node has 2 childs")
		showNodesData(current)
		fmt.Print("\n[...] Select 1/2 to visit left/right child of Node -> INFO(node) : ", current.data, " --> ")
		direction := selectVisitDirection()
		visitTree(current, node, direction)
	} else if current.left == nil && current.right == nil { //If only left child exists
		fmt.Println("[...] The currently visited Node has no any child.")
		showNodesData(current)
		fmt.Print("\n[...] Select 1/2 to insert Node as left/right child of Node -> INFO(node) : ", current.data, " --> ")
		direction := selectVisitDirection()
		fmt.Println(direction, " selected")
		if direction == 1 {
			current.left = node
			fmt.Println("[...] Node inserted as left child of the current Node")
		} else {
			current.right = node
			fmt.Println("[...] Node inserted as right child of the current Node")
		}
	} else if current.left != nil { //If only left child exists
		fmt.Println("[...] The currently visited Node has only left child.")
		showNodesData(current)
		fmt.Print("\n[...] Select 1/2 to visit left subtree OR insert Node as right child of Node -> INFO(node) : ", current.data, " --> ")
		direction := selectVisitDirection()
		fmt.Println(direction, " selected")
		if direction == 2 {
			(current.right) = node
			fmt.Println("[...] Node inserted as right child of the current Node")
			return root
		}

		insertNodeOnTreeAccToChoice(current.left, node)
	} else { //If only right child is there
		fmt.Print("[...] The currently visited Node has only right child.")
		showNodesData(current)
		fmt.Print("\n[...] Select 1/2 to visit right subtree OR insert Node as left child of Node -> INFO(node) : ", current.data, " --> ")
		direction := selectVisitDirection()
		fmt.Println(direction, " selected")
		if direction == 1 {
			current.left = node
			fmt.Println("[...] Node inserted as left child of the current Node")
			return root
		}

		insertNodeOnTreeAccToChoice(current.right, node)
	}
	return root
}

func selectVisitDirection() int {
	var direction, d int
	var err error //Enter 1 to visit left child and 2 to visit right one
	d, err = fmt.Scanf("%d", &direction)
	//fmt.Printf("%T%T..%d%d...%v..%v..%v..%v\n", d, direction, d, direction, err, err != nil, !(d > 0 && d < 3), err != nil || !(d > 0 && d < 3))
	for (err != nil) || !(d > 0 && d < 3) {
		fmt.Println("[...] Either Improper input entered in place of integer Or you have entered an integer other than 1/2")
		fmt.Print("[...] Lets enter again : ")
		d, err = fmt.Scanf("%d", &direction)
	}
	return direction
}

func showNodesData(current *Node) {

	fmt.Printf("[...] INFO(node):%d \n", current.data)
	if current.left != nil {
		fmt.Printf("      INFO(node.left) : %d ", (current.left).data)
	}
	if current.right != nil {
		fmt.Printf("... INFO(node.right) : %d\n", (current.right).data)
	}
}

func visitTree(current *Node, node *Node, direction int) *Node {
	var n *Node
	if direction == 1 {
		fmt.Println("[...] Moving left....")
		n = insertNodeOnTreeAccToChoice(current.left, node)
		fmt.Println("[...] Node inserted")
	} else {
		fmt.Println("[...] Moving right...")
		n = insertNodeOnTreeAccToChoice(current.right, node)
		fmt.Println("[...] Node inserted")
	}
	return n
}

func showNodes(current *Node) {
	if current == nil {
		return
	}
	fmt.Printf("%d\t", current.data)
	showNodes(current.left)
	showNodes(current.right)
}

/* This is main() function that starts the application , it is an entry point of execution */
func main() {
	var root *Node = nil //Initially Tree will be empty

	dataItems := []int{12, 45, 78, 98, 23, -45}
	for index, item := range dataItems {
		fmt.Println("[...] ", (index + 1), "Creating a Node with data(of type int) : ", item)
		node := nodeCreator(item)
		fmt.Println("Inserting ...", item)
		root = insertNodeOnTreeAccToChoice(root, node)
		fmt.Println("RetRoot:", root) //
	}
	fmt.Println("\n[...] Visiting Binary tree...")
	showNodes(root)
}

/*
[...]  1 Creating a Node with data(of type int) :  12
Inserting ... 12
[...] Initially tree is empty.So Node is inserted as the root of the tree and... INFO(root) :  12
RetRoot: &{12 <nil> <nil>}
[...]  2 Creating a Node with data(of type int) :  45
Inserting ... 45
Current :  0xc42007e000
[...] The currently visited Node has no any child.
[...] INFO(node):12

[...] Select 1/2 to insert Node as left/right child of Node -> INFO(node) : 12 --> 2
2  selected
[...] Node inserted as right child of the current Node
RetRoot: &{12 <nil> 0xc42007e040}
[...]  3 Creating a Node with data(of type int) :  78
Inserting ... 78
Current :  0xc42007e000
[...] The currently visited Node has only right child.[...] INFO(node):12
... INFO(node.right) : 45

[...] Select 1/2 to visit right subtree OR insert Node as left child of Node -> INFO(node) : 12 --> 1
1  selected
[...] Node inserted as left child of the current Node
RetRoot: &{12 0xc42007e080 0xc42007e040}
[...]  4 Creating a Node with data(of type int) :  98
Inserting ... 98
Current :  0xc42007e000
[...] The currently visited Node has 2 childs
[...] INFO(node):12
      INFO(node.left) : 78 ... INFO(node.right) : 45

[...] Select 1/2 to visit left/right child of Node -> INFO(node) : 12 --> 2
[...] Moving right...
Current :  0xc42007e040
[...] The currently visited Node has no any child.
[...] INFO(node):45

[...] Select 1/2 to insert Node as left/right child of Node -> INFO(node) : 45 --> 2
2  selected
[...] Node inserted as right child of the current Node
[...] Node inserted
RetRoot: &{12 0xc42007e080 0xc42007e040}
[...]  5 Creating a Node with data(of type int) :  23
Inserting ... 23
Current :  0xc42007e000
[...] The currently visited Node has 2 childs
[...] INFO(node):12
      INFO(node.left) : 78 ... INFO(node.right) : 45

[...] Select 1/2 to visit left/right child of Node -> INFO(node) : 12 --> 2
[...] Moving right...
Current :  0xc42007e040
[...] The currently visited Node has only right child.[...] INFO(node):45
... INFO(node.right) : 98

[...] Select 1/2 to visit right subtree OR insert Node as left child of Node -> INFO(node) : 45 --> 1
1  selected
[...] Node inserted as left child of the current Node
[...] Node inserted
RetRoot: &{12 0xc42007e080 0xc42007e040}
[...]  6 Creating a Node with data(of type int) :  -45
Inserting ... -45
Current :  0xc42007e000
[...] The currently visited Node has 2 childs
[...] INFO(node):12
      INFO(node.left) : 78 ... INFO(node.right) : 45

[...] Select 1/2 to visit left/right child of Node -> INFO(node) : 12 --> 1
[...] Moving left....
Current :  0xc42007e080
[...] The currently visited Node has no any child.
[...] INFO(node):78

[...] Select 1/2 to insert Node as left/right child of Node -> INFO(node) : 78 --> 1
1  selected
[...] Node inserted as left child of the current Node
[...] Node inserted
RetRoot: &{12 0xc42007e080 0xc42007e040}

[...] Visiting Binary tree...
12	78	-45   45   23   98

*/
