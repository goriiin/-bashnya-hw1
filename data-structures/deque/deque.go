package deque

type Deque struct {
	data []int
}

func NewQueue() *Deque {
	return &Deque{
		make([]int, 0, 10),
	}
}

// PushFront - добавить в начало
func (d *Deque) PushFront(val int) {
	d.data = append(d.data[1:], d.data...)
	d.data[0] = val
}

// PushBack - добавить в конец
func (d *Deque) PushBack(val int) {
	d.data = append(d.data, val)
}

// PopBack - удалить с конца
func (d *Deque) PopBack() {
	d.data = d.data[:len(d.data)-1]
}

// PopFront - удалить с начала
func (d *Deque) PopFront() {
	d.data = d.data[1:]
}
