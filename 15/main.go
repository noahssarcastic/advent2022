package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	"golang.org/x/exp/slices"
)

var input = flag.String("f", "input.txt", "input file")

var MaxCoord = 4000000

func main() {
	flag.Parse()

	re, err := regexp.Compile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(f)
	var (
		sensors = make([]point, 0)
		beacons = make([]point, 0)
		ranges  = make([]int, 0)
	)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindAllStringSubmatch(line, 1)
		sensor := pointFromSlice(match[0][1:3])
		beacon := pointFromSlice(match[0][3:5])
		sensors = append(sensors, sensor)
		beacons = append(beacons, beacon)
		ranges = append(ranges, manhattan(sensor, beacon))
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(tuningFreq(*findSignal(beacons, sensors, ranges)))
}

func findSignal(beacons, sensors []point, ranges []int) *point {
	for y := 0; y < MaxCoord; y++ {
		for x := 0; x < MaxCoord; x++ {
			pt := point{x, y}
			if isSignal(beacons, sensors, ranges, pt) {
				return &pt
			}
		}
	}
	return nil
}

func isSignal(beacons, sensors []point, ranges []int, pt point) bool {
	return !isBeacon(beacons, pt) && !inRange(sensors, ranges, pt)
}

func isBeacon(beacons []point, pt point) bool {
	idx := slices.IndexFunc(beacons, func(i point) bool {
		return equal(pt, i)
	})
	return idx > 0
}

func inRange(sensors []point, ranges []int, pt point) bool {
	for i, s := range sensors {
		d := manhattan(pt, s)
		inRange := d <= ranges[i]
		if inRange {
			return true
		}
	}
	return false
}

// Point

type point struct {
	x, y int
}

func equal(a, b point) bool {
	return a.x == b.x && a.y == b.y
}

func pointFromSlice(ss []string) point {
	x, err := strconv.Atoi(ss[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(ss[1])
	if err != nil {
		panic(err)
	}
	return point{x, y}
}

// Utils

func abs(i int) int {
	return int(math.Abs(float64(i)))
}

func manhattan(a, b point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func tuningFreq(pt point) int {
	return pt.x*4000000 + pt.y
}
