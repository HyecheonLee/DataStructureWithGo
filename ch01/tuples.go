package main

import "fmt"

func powerSeries(a int) (square int, cube int) {
	return a * a, a * a * a
}
func main() {
	var square int
	var cube int
	square, cube = powerSeries(3)
	fmt.Println("Square: ", square, "Cube", cube)
}