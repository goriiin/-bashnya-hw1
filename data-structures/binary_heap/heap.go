package binary_heap

import (
	"sort"
)

type Heap struct {
	data []int
}

// NewHeap создать пустую кучу
func NewHeap() Heap {
	return Heap{
		make([]int, 0),
	}
}

// Build «построить кучу» - разместить элементы в массиве так, чтобы сформировалась куча
func Build(arr []int) Heap {
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	return Heap{
		arr,
	}
}

// Insert вставка
func (h *Heap) Insert(elem int) {
	h.data = append(h.data, elem)

	h.siftUp(len(h.data) - 1)
}

// ExtractTop удалить элемент из верха кучи
func (h *Heap) ExtractTop() {
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.siftDown(0)
}

// siftUp поднять на уровень
func (h *Heap) siftUp(index int) {
	parent := findParent(index)
	for index > 0 && h.data[index] > h.data[findParent(index)] {
		swap(&h.data[index], &h.data[parent])
		h.siftUp(parent)
	}
}

// siftDown понизить на уровень
func (h *Heap) siftDown(index int) {
	leftChild := findLeftChild(index)
	rightChild := findRightChild(index)
	size := len(h.data)
	largest := index

	if leftChild < size && h.data[leftChild] > h.data[largest] {
		largest = leftChild
	}

	if rightChild < size && h.data[rightChild] > h.data[largest] {
		largest = rightChild
	}

	if largest != index {
		swap(&h.data[index], &h.data[largest])
		h.siftDown(largest)
	}
}

func findParent(index int) int {
	if index <= 0 {
		return -1
	}

	return (index - 1) / 2
}

func findLeftChild(index int) int {
	if index < 0 {
		return -1
	}

	return index*2 + 1
}

func findRightChild(index int) int {
	if index < 0 {
		return -1
	}

	return index*2 + 2
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
