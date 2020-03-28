package main

import "fmt"

type Queue []*Order

type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

func NewOrder(priority int, quantity int, product string, customerName string) *Order {
	return &Order{priority: priority, quantity: quantity, product: product, customerName: customerName}
}
func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		appended := false
		for i, addedOrder := range *queue {

			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}
	}
}

// main method
func main() {
	var queue Queue
	queue = make(Queue, 0)
	order1 := NewOrder(2, 20, "Computer", "Greg White")
	order2 := NewOrder(1, 10, "Monitor", "John Smith")
	order3 := NewOrder(3, 40, "Monitor2", "John Smith2")
	queue.Add(order1)
	queue.Add(order2)
	queue.Add(order3)
	for i := 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
