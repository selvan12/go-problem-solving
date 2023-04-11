package main

import (
	"container/list"
	"fmt"
)

func showElements(list *list.List) {
	for element := list.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}

func queue() {
	queue := list.New()

	queue.PushBack(9)
	queue.PushBack(8)
	queue.PushBack(7)
	queue.PushBack(6)
	showElements(queue)

	fmt.Println("Remove back element")
	queue.Remove(queue.Back())
	showElements(queue)

	fmt.Println("Remove front element")
	queue.Remove(queue.Front())
	showElements(queue)
}

/*
Output:
9
8
7
6
Remove back element
9
8
7
Remove front element
8
7
*/
