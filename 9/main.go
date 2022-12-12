package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Knot struct {
	x, y int
}

type Direction struct {
	xOffset, yOffset int
}

var directions = map[byte]Direction{
	'U': {0, 1},
	'R': {1, 0},
	'D': {0, -1},
	'L': {-1, 0},
}

func moveKnot(k *Knot, direction byte) {
	k.x += directions[direction].xOffset
	k.y += directions[direction].yOffset
}

func moveKnotTo(k *Knot, x, y int) {
	k.x = x
	k.y = y
}

func touching(k1, k2 *Knot) bool {
	return math.Abs(float64(k1.x-k2.x)) <= 1 &&
		math.Abs(float64(k1.y-k2.y)) <= 1
}

func move(visits map[string]int, head, tail *Knot, direction byte, steps int) {
	printPositions(head, tail)
	fmt.Println()

	for i := 0; i < steps; i++ {
		oldHead := Knot{head.x, head.y}
		moveKnot(head, direction)
		if touching(head, tail) {
			fmt.Println("still touching; don't move tail")
		} else if head.x == tail.x || head.y == tail.y {
			fmt.Println("same row/col; move in same direction")
			moveKnot(tail, direction)
		} else {
			fmt.Println("head too far; diagonal move")
			moveKnotTo(tail, oldHead.x, oldHead.y)
		}
		visits[getKey(tail)] += 1
		printPositions(head, tail)
	}
	fmt.Println()
}

func printPositions(head, tail *Knot) {
	fmt.Printf("H(%v,%v) T(%v,%v)\n", head.x, head.y, tail.x, tail.y)
}

func getKey(k *Knot) string {
	return fmt.Sprintf("%v:%v", k.x, k.y)
}

func main() {
	f, _ := os.Open("input2.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	head := Knot{0, 0}
	tail := Knot{0, 0}
	visits := map[string]int{"0:0": 1}

	for scanner.Scan() {
		line := scanner.Text()
		update := strings.Split(line, " ")
		direction := update[0][0]
		steps, _ := strconv.Atoi(update[1])
		fmt.Printf("Move %v, %v times\n", directions[direction], steps)
		move(visits, &head, &tail, direction, steps)
	}

	printPositions(&head, &tail)
	fmt.Println(len(visits))
}
