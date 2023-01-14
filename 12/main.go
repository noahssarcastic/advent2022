package main

import (
	"fmt"

	coord "github.com/noahssarcastic/advent2022/12/coord"
)

type FifoQueue struct {
	queue []coord.Coord
}

func (q *FifoQueue) length() int {
	return len(q.queue)
}

func (q *FifoQueue) enqueue(c coord.Coord) {
	q.queue = append(q.queue, c)
}

func (q *FifoQueue) dequeue() coord.Coord {
	c := q.queue[0]
	q.queue = q.queue[1:]
	return c
}

func (q *FifoQueue) indexOf(c coord.Coord) int {
	for i, el := range q.queue {
		if coord.Equal(c, el) {
			return i
		}
	}
	return -1
}

func (q *FifoQueue) contains(c coord.Coord) bool {
	return q.indexOf(c) >= 0
}

func main() {
	inputFile := parseArgs()
	hm, start, end := parseInput(inputFile)

	visited := &FifoQueue{}
	queue := &FifoQueue{}
	depth := []int{}

	queue.enqueue(start)
	visited.enqueue(start)
	depth = append(depth, 0)

	for queue.length() > 0 {
		current := queue.dequeue()
		index := visited.indexOf(current)
		currentDepth := depth[index]

		// If current coord is the final destination...
		if coord.Equal(current, end) {
			fmt.Println(currentDepth)
		}

		for _, adj := range hm.Adjacent(current) {
			if !hm.InBounds(adj) || visited.contains(adj) {
				continue
			}
			if hm.IsTraversable(current, adj) {
				queue.enqueue(adj)
				visited.enqueue(adj)
				depth = append(depth, currentDepth+1)
			}
		}
	}
}
