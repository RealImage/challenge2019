package models

type MaxQueue struct {
	queue []Input
	size  int
}

func NewMaxQueue() MaxQueue {
	q := MaxQueue{
		queue: make([]Input, 0),
		size:  -1,
	}
	return q
}

func (q *MaxQueue) Parent(i int) int {
	return (i - 1) / 2
}

func (q *MaxQueue) LeftChild(i int) int {
	return (2 * i) + 1
}

func (q *MaxQueue) RightChild(i int) int {
	return (2 * i) + 2
}

func (q *MaxQueue) Swap(i int, j int) {
	temp := q.queue[i]
	q.queue[i] = q.queue[j]
	q.queue[j] = temp
}

func (q *MaxQueue) ShiftUp(i int) {
	p := q.Parent(i)
	for i > 0 && q.queue[p].Volume < q.queue[i].Volume {
		q.Swap(p, i)
		i = p
		p = q.Parent(i)
	}
}

func (q *MaxQueue) ShiftDown(i int) {
	maxIndex := i
	l := q.LeftChild(i)
	if l <= q.size && q.queue[l].Volume > q.queue[maxIndex].Volume {
		maxIndex = l
	}
	r := q.RightChild(i)
	if r <= q.size && q.queue[r].Volume > q.queue[maxIndex].Volume {
		maxIndex = r
	}
	if i != maxIndex {
		q.Swap(i, maxIndex)
		q.ShiftDown(maxIndex)
	}
}

func (q *MaxQueue) Insert(input Input) {
	q.size = q.size + 1
	q.queue = append(q.queue, input)
	q.ShiftUp(q.size)
}

func (q *MaxQueue) ExtractMax() *Input {
	if q.size > -1 {
		record := q.queue[0]
		q.queue[0] = q.queue[q.size]
		q.size = q.size - 1
		q.ShiftDown(0)
		return &record
	}
	return nil
}

func (q *MaxQueue) GetMax() *Input {
	if q.size > -1 {
		return &q.queue[0]
	}
	return nil
}

func (q *MaxQueue) Size() int {
	return q.size
}
