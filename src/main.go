package main

import (
	biHeap "bashnya_hw1/data-structures/binaryHeap"
	"bashnya_hw1/data-structures/queue"
	"fmt"
)

func main() {
	arr := []int{4, 5, 6, 3, 6}
	h := biHeap.Build(arr)

	h.Insert(7)
	h.Insert(9)
	h.Insert(8)

	h.Insert(1)
	h.Insert(7)
	h.Insert(10)
	h.Insert(11)

	h.Insert(13)
	h.Insert(12)
	h.Insert(100)

	fmt.Println(h)

	h.ExtractTop()
	fmt.Println(h)

	h.ExtractTop()
	fmt.Println(h)

	q := queue.NewQueue()
	q.Push(1)

	q.Push(2)
	fmt.Println(q)
	q.Pop()

	fmt.Println(q)

}
