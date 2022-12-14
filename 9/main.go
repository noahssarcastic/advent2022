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
	x, y float64
}

func copyKnot(k *Knot) *Knot {
	return &Knot{k.x, k.y}
}

type Direction struct {
	xOffset, yOffset float64
}

func moveKnot(k *Knot, direction Direction) {
	k.x += direction.xOffset
	k.y += direction.yOffset
}

func touching(k1, k2 *Knot) bool {
	touch := math.Abs(float64(k1.x-k2.x)) <= 1 &&
		math.Abs(float64(k1.y-k2.y)) <= 1
	return touch
}

func distance(k1, k2 *Knot) float64 {
	return math.Sqrt(math.Pow(k1.x-k2.x, 2) + math.Pow(k1.y-k2.y, 2))
}

func moveFollower(head, tail *Knot, direction Direction) {
	if touching(head, tail) {
		return
	}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			direction := Direction{float64(x), float64(y)}
			copy := copyKnot(tail)
			moveKnot(copy, direction)
			if distance(copy, head) == 1 {
				moveKnot(tail, direction)
				return
			}
		}
	}
	diagonal := Direction{
		math.Copysign(1, head.x-tail.x),
		math.Copysign(1, head.y-tail.y),
	}
	moveKnot(tail, diagonal)
}

func moveRope(knots []Knot, direction Direction, steps int, visits map[string]int) {
	for i := 0; i < steps; i++ {
		moveKnot(&knots[0], direction)
		for k := 1; k < len(knots); k++ {
			followerDirection := direction
			moveFollower(&knots[k-1], &knots[k], followerDirection)
		}
		visits[getKey(&knots[len(knots)-1])] += 1

		// Debug
		// printBoardState(knots)
		// fmt.Print("Press 'Enter' to continue...")
		// bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

var inputs = map[byte]Direction{
	'U': {0, 1},
	'R': {1, 0},
	'D': {0, -1},
	'L': {-1, 0},
}

func getKey(k *Knot) string {
	return fmt.Sprintf("%v:%v", k.x, k.y)
}

func main() {
	f, _ := os.Open("input2.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	numKnots := 10
	knots := make([]Knot, 0)
	for i := 0; i < numKnots; i++ {
		knots = append(knots, Knot{0, 0})
	}
	visits := map[string]int{"0:0": 1}
	for scanner.Scan() {
		line := scanner.Text()
		update := strings.Split(line, " ")
		direction := inputs[update[0][0]]
		steps, _ := strconv.Atoi(update[1])
		moveRope(knots, direction, steps, visits)
	}
	fmt.Println(len(visits))
}
