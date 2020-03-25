package main

import "fmt"

// Node class
type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}
type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty &&
				node.nextNode.property == secondProperty {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}
func (linkedList *LinkedList) AddToHead(property int) {
	node := &Node{property: property, nextNode: nil}
	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
		linkedList.headNode.previousNode = node
	}
	linkedList.headNode = node
}
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var nodeWithValue *Node
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWithValue = node
			break

		}
	}
	return nodeWithValue
}
func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	node := &Node{property: property, nextNode: nil, previousNode: nil}
	nodeWithValue := linkedList.NodeWithValue(nodeProperty)
	if nodeWithValue != nil {
		nextNode := nodeWithValue.nextNode
		node.nextNode = nextNode
		node.previousNode = nodeWithValue

		nodeWithValue.nextNode = node
		nextNode.previousNode = node
	}
}
func (linkedList *LinkedList) LastNode() *Node {
	var lastNode *Node
	for node := linkedList.headNode; node.nextNode != nil; node = node.nextNode {
		lastNode = node.nextNode
	}
	return lastNode
}

func (linkedList *LinkedList) AddToEnd(property int) {
	node := &Node{property: property, nextNode: nil, previousNode: nil}
	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

// main method
func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	fmt.Println(linkedList.headNode.property)
	var node *Node
	node = linkedList.NodeBetweenValues(1, 5)
	fmt.Println(node.property)
}
