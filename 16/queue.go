package main

type queue struct {
	queue []int
}

func newQueue() *queue {
	return &queue{make([]int, 0)}
}

func (q *queue) length() int {
	return len(q.queue)
}

func (q *queue) enqueue(i int) {
	q.queue = append(q.queue, i)
}

func (q *queue) dequeue() int {
	i := q.queue[0]
	newQueue := make([]int, len(q.queue)-1)
	copy(newQueue, q.queue[1:])
	q.queue = newQueue
	return i
}

func (q *queue) indexOf(i int) int {
	for idx, el := range q.queue {
		if el == i {
			return idx
		}
	}
	return -1
}

func (q *queue) contains(i int) bool {
	return q.indexOf(i) >= 0
}
