package main

import "fmt"

/*

Stack Simple Go

C like implementation but with slices

Methods:

Push
Pop
isEmpty
isFull

arsperger 2021

*/

type st struct {
	max   int
	top   int
	stack []interface{}
}

//StackInit - init stack
func StackInit(n int) *st {
	var s = new(st)
	s.max = n
	s.top = -1
	s.stack = make([]interface{}, s.max)
	return s
}

//isEmpty - returns true is stack is empty
func (s *st) isEmpty() bool {
	return s.top < 0
}

//isFull - return true if stack is full
func (s *st) isFull() bool {
	return s.top == s.max
}

//Push insert item into the stack
func (s *st) Push(d interface{}) bool {
	if s.isFull() {
		return false
	}
	s.top++
	s.stack[s.top] = d
	return true
}

//Pop returns item from the stack
func (s *st) Pop() interface{} {
	if s.isEmpty() {
		return nil
	}
	var v interface{}
	v = s.stack[s.top]
	s.stack[s.top] = nil
	s.top--
	return v

}

func main() {

	var s *st

	s = StackInit(5)

	s.Push(1)
	s.Push(2)
	s.Push(3)

	for i := 0; i < 3; i++ {
		fmt.Printf("stack val: %d\n", s.Pop())
	}

}
