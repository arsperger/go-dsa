package main

/**

Doubly linked List

arsperger 2021

*/

// importing fmt package
import (
	"fmt"
)

// Node struc
type Node struct {
	data         interface{}
	nextNode     *Node
	previousNode *Node
}

// LinkedList struct
type LinkedList struct {
	headNode *Node
}

//AddToHead method of LinkedList
func (list *LinkedList) AddToHead(data interface{}) {
	node := new(Node)
	node.data = data
	node.nextNode = nil
	if list.headNode != nil {
		node.nextNode = list.headNode     // add node to the head
		list.headNode.previousNode = node // point next node to the previous
	}

	list.headNode = node

}

//NodeWithValue returns node with value
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

//AddAfter insert Node with @data after the node with @nodeData
func (list *LinkedList) AddAfter(nodeData interface{}, data interface{}) {
	node := new(Node)
	node.data = data
	node.nextNode = nil

	var nodeWith *Node

	nodeWith = list.NodeWithValue(nodeData)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}

}

//LastNode returns last node from the list
func (list *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = list.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

//AddToEnd adds node to the end of the list
func (list *LinkedList) AddToEnd(data interface{}) {
	node := new(Node)
	node.data = data
	node.nextNode = nil

	var lastNode *Node

	lastNode = list.LastNode()

	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

//IterateList method of LinkedList
func (list *LinkedList) IterateList() {
	var node *Node
	for node = list.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.data)
	}
}

func main() {

	llist := new(LinkedList)

	llist.AddToHead(1)
	llist.AddToHead(3)
	llist.AddToEnd(5)

	llist.IterateList()

}
