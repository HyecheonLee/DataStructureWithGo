package main

import "fmt"

//Iprocess interface
type IProcess interface {
	process()
}

//Adapter struct
type Adapter struct {
	adaptee Adaptee
}

//adapter class method process
func (adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

// Adaptee struct
type Adaptee struct {
	adapterType int
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}
func main() {
	var processor IProcess = Adapter{}
	processor.process()
}
