/**

Simple Priority Queue

Slice - based

arsperger 2021
*/

package main

import "fmt"

// Queue is array of nodes
type Queue []*Node

// Node with priorit is exported
type Node struct {
	priority int
	data     interface{}
}

// New is exported
func (n *Node) New(p int, v interface{}) {
	n.priority = p
	n.data = v
	return
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

// Enqueue adds the order to the queue
func (q *Queue) Enqueue(n *Node) {
	if q.isEmpty() {
		*q = append(*q, n)
	} else {
		added := false
		for i, inQueue := range *q {
			if n.priority > inQueue.priority {
				fmt.Printf("reorder: %d with %d\n", inQueue.priority, n.priority)
				*q = append((*q)[:i], append(Queue{n}, (*q)[i:]...)...)
				added = true
				break
			}
		}
		if !added {
			*q = append(*q, n)
		}
	}
}

// Dequeue removes item from the queue and prints it
func (q *Queue) Dequeue() {
	if q.isEmpty() {
		return
	}
	removed := (*q)[0]
	fmt.Printf("removed element: %v", removed)
	*q = (*q)[1:]
}

// Peak shows the first item in the queue w/o removing it
func (q *Queue) Peak() {
	if q.isEmpty() {
		return
	}
	fmt.Println((*q)[len(*q)-1])
}

// Dispaly prints the queue
func (q *Queue) Dispaly() {
	if q.isEmpty() {
		fmt.Println("Empty queue")
	}
	for _, v := range *q {
		fmt.Println(v)
	}
}

func main() {

	data1 := new(Node)
	data2 := new(Node)
	data3 := new(Node)
	data4 := new(Node)

	data1.New(1, 10)
	data2.New(1, 20)
	data3.New(1, 30)
	data4.New(4, 40)

	var q *Queue
	q = &Queue{}

	q.Enqueue(data1)
	q.Enqueue(data2)
	q.Enqueue(data3)
	q.Enqueue(data4)

	q.Dispaly()

	q.Dequeue()

	fmt.Println()

	q.Dispaly()

}
