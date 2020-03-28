package main

import (
	"fmt"
	"strconv"
)

type Element struct {
	elementValue int
}

func (e *Element) String() string {
	return strconv.Itoa(e.elementValue)
}

type Stack struct {
	elements     []*Element
	elementCount int
}

// NewStack returns a new stack.
func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

// Push adds a node to the stack.
func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.elementCount], element)
	stack.elementCount = stack.elementCount + 1
}
func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}
	length := len(stack.elements)
	element := stack.elements[length-1]
	//stack.elementCount = stack.elementCount - 1
	if length > 1 {
		stack.elements = stack.elements[:length-1]
	} else {
		stack.elements = stack.elements[0:]
	}
	stack.elementCount = len(stack.elements)
	return element
}
func main() {
	var stack *Stack = &Stack{}
	stack.New()
	var element1 *Element = &Element{3}
	var element2 *Element = &Element{5}
	var element3 *Element = &Element{7}
	var element4 *Element = &Element{9}
	stack.Push(element1)
	stack.Push(element2)
	stack.Push(element3)
	stack.Push(element4)
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}
