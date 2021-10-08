package set

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

//returns true if item is in the set
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

//Union returns union of both sets
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
