/*

Set implementation

Set - static (immutable) has only query operations

Mutable or dynamic Set has Operations:

addItem
deleteItem
containsItem
intersect
union

arsperger 2021

*/

package main

import (
	"fmt"
)

// Set is a set truct
type Set struct {
	setMap map[interface{}]bool
}

//New init a new set
func (s *Set) New() {
	s.setMap = make(map[interface{}]bool)
}

//AddItem adds new item to the set
func (s *Set) AddItem(v interface{}) {
	if !s.containsItem(v) {
		s.setMap[v] = true
	}
}

//containsItem returns true if item is in the set
func (s *Set) containsItem(v interface{}) bool {
	_, ok := s.setMap[v]
	return ok
}

//DeleteItem removes item from the set
func (s *Set) DeleteItem(v interface{}) bool {
	if s.containsItem(v) {
		delete(s.setMap, v)
		return true
	}
	return false
}

//Intersect search for intersections and returns new set
func (s *Set) Intersect(secondSet *Set) *Set {
	newSet := new(Set)

	for value, _ := range s.setMap {
		if secondSet.containsItem(value) {
			newSet.AddItem(value)
		}
	}
	return newSet
}

//Union returns union ob both sets
func (s *Set) Union(secondSet *Set) *Set {
	newSet := new(Set)

	for value, _ := range s.setMap {
		newSet.AddItem(value)
	}

	for value, _ := range secondSet.setMap {
		newSet.AddItem(value)
	}

	return newSet
}

func main() {

	//var testSet *Set
	testSet := new(Set)

	testSet.New()

	testSet.AddItem(1)
	testSet.AddItem(2)
	testSet.AddItem(2)

	fmt.Printf("added to set")

}
