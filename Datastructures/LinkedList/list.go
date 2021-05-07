package main

/*

Linkedlist get 1th and last element in O(1)
Search O(n)

Exported Methods:

- AddtoHead
- AddtoEnd
- IterateList
- LastNode
- NodeWithValue
- AddAfter

---|Head|->---|Node|->---|Node|->---NULL

arsperger 2021

*/

import "fmt"

// Node is the list Node
// it contains properties and pointer to the next Node
type Node struct {
	data     interface{}
	nextNode *Node
}

// LinkedList cotains pointer to the Node.
// we can iterate throu the list by traversin
// from headNode to the NextNode
type LinkedList struct {
	headNode *Node
}

// AddtoHead adds node to head of the Linkedlist
func (list *LinkedList) AddtoHead(data interface{}) {
	var node = new(Node)
	node.data = data
	if list.headNode != nil {
		node.nextNode = list.headNode
	}
	list.headNode = node
}

//IterateList iterates over LinkedList
func (list *LinkedList) IterateList() {
	var node *Node
	for node = list.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.data)
	}
}

// CountNodes returns number of Nodes in the List
func (list *LinkedList) CountNodes() int {
	var node *Node
	node = list.headNode
	count := 0
	if node == nil {
		return count
	}
	for ; node != nil; node = node.nextNode {
		count += 1
	}
	return count
}

// DeleteHead removes head Node
func (list *LinkedList) DeleteHead() {
	var node *Node
	node = list.headNode
	if node == nil {
		fmt.Println("List is EMPTY")
		return
	}
	if node.nextNode != nil {
		fmt.Println("REMOVE HEAD!")
		list.headNode = node.nextNode
	}
	return
}

// DeleteEnd remove Node from the end of the List
func (list *LinkedList) DeleteEnd() {
	var node *Node
	var lastNode *Node

	lastNode = list.LastNode()
	if lastNode == nil {
		return
	}

	for node = list.headNode; node != lastNode; node = node.nextNode {
		if node.nextNode == lastNode {
			node.nextNode = nil
			return
		}
	}
	return
}

//DeleteAll removes all Nodes from the List
func (list *LinkedList) DeleteAll() {
	fmt.Println("Warning: remove all nodes!")
	list.headNode = nil
}

//LastNode returns the last Node
func (list *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	node = list.headNode

	if node == nil {
		return nil
	}

	for ; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

//AddToEnd adds the node to the end
func (list *LinkedList) AddToEnd(data interface{}) bool {
	var node = new(Node)
	node.data = data
	node.nextNode = nil
	var lastNode *Node
	lastNode = list.LastNode()

	if lastNode != nil {
		lastNode.nextNode = node
		return true
	}
	return false
}

//NodeWithValue returns Node with given data
func (list *LinkedList) NodeWithValue(data interface{}) *Node {
	var node *Node
	var nodeWith *Node
	for node = list.headNode; node != nil; node = node.nextNode {
		if node.data == data {
			nodeWith = node
			break
		}
	}

	return nodeWith
}

//AddAfter add Node in the Middle of the List (after *Node)
func (list *LinkedList) AddAfter(nodeData interface{}, data interface{}) bool {
	node := new(Node)
	node.data = data
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = list.NodeWithValue(nodeData)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
		return true
	}
	return false
}

func main() {

	var list = new(LinkedList)

	list.AddtoHead(1)
	list.AddtoHead(2)
	list.AddtoHead(3)
	list.AddtoHead(4)
	list.AddtoHead(5)
	//list.AddToEnd(300)
	//list.AddAfter(1, 7)
	//list.IterateList()

	fmt.Println("Iterate:")
	list.IterateList()

	//var last Node
	//last = *list.LastNode()
	//fmt.Println("last node:")
	//fmt.Println(last.data)

	/*
		var last Node
		last = *list.LastNode()
		fmt.Printf("last node: %d\n", last.data)
	*/

	fmt.Printf("Nodes in list: %d\n", list.CountNodes())

	//list.AddAfter(4, 7)
	//list.AddAfter(7, 10)

	list.DeleteEnd()
	fmt.Println("Iterate:")
	list.IterateList()

	//fmt.Println(list.headNode.nextNode)

	//list.DeleteHead()

	//list.IterateList()
	fmt.Printf("Nodes in list: %d\n", list.CountNodes())
	//list.DeleteAll()
	list.IterateList()

}
