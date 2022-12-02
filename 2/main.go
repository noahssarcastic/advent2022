package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Shape struct {
	shape int
	beats int
	loses int
}

// Shape.shape options
const (
	// Rock adds 1 to the total score.
	rock = iota + 1
	// Paper adds 2 to the total score.
	paper = iota + 1
	// Scissors adds 3 to the total score.
	scissors = iota + 1
)

// Map a secret code to the shape your opponent will play.
var shapeCode = map[byte]Shape{
	'A': {
		shape: rock,
		beats: scissors,
		loses: paper,
	},
	'B': {
		shape: paper,
		beats: rock,
		loses: scissors,
	},
	'C': {
		shape: scissors,
		beats: paper,
		loses: rock,
	},
}

// Result options
const (
	// A loss add 0 points to the total score.
	loss = iota * 3
	// A tie adds 3 points to the total score.
	tie = iota * 3
	// A win adds 6 points to the total score.
	win = iota * 3
)

// Map a secret code to the result of the round.
var resultCode = map[byte]int{
	'X': loss,
	'Y': tie,
	'Z': win,
}

// Get the shape to play to achieve the given result.
func shape(op Shape, result int) int {
	if result == tie {
		return op.shape
	} else if result == loss {
		return op.beats
	} else {
		return op.loses
	}
}

// Calculate the round score.
func score(op Shape, me int) int {
	if op.shape == me {
		return tie
	} else if op.loses == me {
		return win
	} else {
		return loss
	}
}

func main() {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		op := shapeCode[line[0]]
		result := resultCode[line[2]]

		me := shape(op, result)

		total += score(op, me) + me
	}
	check(scanner.Err())

	fmt.Printf("My total score was %v.", total)
}
