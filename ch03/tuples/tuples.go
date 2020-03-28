package main

import "fmt"

func h(x int, y int) int {
	return x * y
}
func g(l int, m int) (x int, y int) {
	x = 2 * l
	y = 4 * m
	return
}
func main() {
	fmt.Println(h(g(1, 2)))
}
