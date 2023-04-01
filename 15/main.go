package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime/pprof"
	"strconv"
	"time"
)

var input = flag.String("f", "input.txt", "input file")
var cpuprofile = flag.String("profile", "", "write cpu profile to file")

var MaxCoord = 4000000

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

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
	ss := make([]sensor, 0)
	for scanner.Scan() {
		match := re.FindAllStringSubmatch(scanner.Text(), 1)
		s := fromSlice(match[0][1:3])
		b := fromSlice(match[0][3:5])
		ss = append(ss, sensor{s, manhattan(s, b)})
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	start := time.Now()
	for i, s := range ss {
		fmt.Printf("sensor %d (%v)\n", i, time.Since(start))
		pts := perimeter(s.loc, s.d+1)
		for _, pt := range pts {
			if pt.x < 0 || pt.x > MaxCoord || pt.y < 0 || pt.y > MaxCoord {
				continue
			}
			if !inRange(ss, pt) {
				fmt.Println(tuningFreq(pt))
				return
			}
		}
	}
}

func inRange(sensors []sensor, pt point) bool {
	for _, s := range sensors {
		d := manhattan(pt, s.loc)
		inRange := d <= s.d
		if inRange {
			return true
		}
	}
	return false
}

type sensor struct {
	loc point
	d   int
}

// Point

type point struct {
	x, y int
}

func equal(a, b point) bool {
	return a.x == b.x && a.y == b.y
}

func fromSlice(ss []string) point {
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

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func manhattan(a, b point) int {
	return absDiff(a.x, b.x) + absDiff(a.y, b.y)
}

func tuningFreq(pt point) int {
	return pt.x*4000000 + pt.y
}

func perimeter(origin point, rad int) (pts []point) {
	// perimeter walk clock-wise
	var (
		up    = point{origin.x, origin.y + rad}
		right = point{origin.x + rad, origin.y}
		down  = point{origin.x, origin.y - rad}
		left  = point{origin.x - rad, origin.y}
	)
	var pt = up
	for !equal(pt, right) {
		pts = append(pts, pt)
		pt.x++
		pt.y--
	}
	for !equal(pt, down) {
		pts = append(pts, pt)
		pt.x--
		pt.y--
	}
	for !equal(pt, left) {
		pts = append(pts, pt)
		pt.x--
		pt.y++
	}
	for !equal(pt, up) {
		pts = append(pts, pt)
		pt.x++
		pt.y++
	}
	return pts
}
