package models

type MinQueue struct {
	queue []PartnerRecord
	size  int
}

func NewMinQueue() MinQueue {
	return MinQueue{
		queue: make([]PartnerRecord, 0),
		size:  -1,
	}
}

func (q *MinQueue) Parent(i int) int {
	return (i - 1) / 2
}

func (q *MinQueue) LeftChild(i int) int {
	return (2 * i) + 1
}

func (q *MinQueue) RightChild(i int) int {
	return (2 * i) + 2
}

func (q *MinQueue) ShiftUp(i int) {
	p := q.Parent(i)
	for i > 0 && q.queue[p].CostPerGB > q.queue[i].CostPerGB {
		q.Swap(p, i)
		i = p
		p = q.Parent(i)
	}
}

func (q *MinQueue) ShiftDown(i int) {
	maxIndex := i
	l := q.LeftChild(i)
	if l <= q.size && q.queue[l].CostPerGB < q.queue[maxIndex].CostPerGB {
		maxIndex = l
	}
	r := q.RightChild(i)
	if r <= q.size && q.queue[r].CostPerGB < q.queue[maxIndex].CostPerGB {
		maxIndex = r
	}
	if i != maxIndex {
		q.Swap(i, maxIndex)
		q.ShiftDown(maxIndex)
	}
}

func (q *MinQueue) Swap(i int, j int) {
	temp := q.queue[i]
	q.queue[i] = q.queue[j]
	q.queue[j] = temp
}

func (q *MinQueue) ExtractMin() *PartnerRecord {
	if q.size > -1 {
		record := q.queue[0]
		q.queue[0] = q.queue[q.size]
		q.size = q.size - 1
		q.ShiftDown(0)
		return &record
	}
	return nil
}

func (q *MinQueue) GetMin() *PartnerRecord {
	if q.size > -1 {
		return &q.queue[0]
	}
	return nil
}

func (q *MinQueue) Insert(record PartnerRecord) {
	q.size = q.size + 1
	q.queue = append(q.queue, record)
	q.ShiftUp(q.size)
}

func (q *MinQueue) Size() int {
	return q.size
}
