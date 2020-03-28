package main

import "fmt"

type Set struct {
	integerMap map[int]bool
}

func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}
func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}
func (set *Set) ContainsElement(element int) bool {
	_, exists := set.integerMap[element]
	return exists
}
func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}
func (set *Set) Intersect(anotherSet *Set) *Set {
	intersectSet := &Set{}
	intersectSet.New()
	for value := range set.integerMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}
func (set *Set) Union(anotherSet *Set) *Set {
	unionSet := &Set{}
	unionSet.New()
	for value := range set.integerMap {
		unionSet.AddElement(value)
	}
	for value := range anotherSet.integerMap {
		unionSet.AddElement(value)
	}
	return unionSet
}

// main method
func main() {
	var set = &Set{}
	set.New()
	set.AddElement(1)
	set.AddElement(2)
	fmt.Println("initial set", set)
	fmt.Println(set.ContainsElement(1))
	var anotherSet = &Set{}
	anotherSet.New()
	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)
	fmt.Println(set.Intersect(anotherSet))
	fmt.Println(set.Union(anotherSet))
}
