/**

Basic Queue implementation

FIFO

Slice based

arsperger 2021

Enqueue
Dequeue
Peak

*/

package main

import "fmt"

// Queue struct w/ underlying slice (which is in turn based on array)
type Queue struct {
	queue []interface{}
}

// Enqueue add item to the queue
func (q *Queue) Enqueue(v interface{}) {
	q.queue = append(q.queue, v)
}

// isEmpty returns true if queue is empty
func (q *Queue) isEmpty() bool {
	return len(q.queue) == 0
}

// Deque remove iterm from the queue
func (q *Queue) Dequeue() interface{} {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return nil
	}
	removed := q.queue[0]
	q.queue = q.queue[1:]
	return removed
}

// Peak returns the first item from the queue w/o removin it
func (q *Queue) Peak() interface{} {
	if q.isEmpty() {
		return nil
	}
	return q.queue[0]
}

func main() {

	myQueue := new(Queue)

	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)

	fmt.Println(myQueue)

	fmt.Println(myQueue.Dequeue())

	fmt.Println(myQueue)

	fmt.Println(myQueue.Peak())
}
