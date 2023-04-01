package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
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
		s := pointFromSlice(match[0][1:3])
		b := pointFromSlice(match[0][3:5])
		ss = append(ss, sensor{s, manhattan(s, b)})
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	start := time.Now()
	var pt point
	for y := 0; y <= MaxCoord; y++ {
		for x := 0; x <= MaxCoord; x++ {
			pt = point{float64(x), float64(y)}
			inRange := false
			for _, s := range ss {
				if manhattan(pt, s.loc) <= s.d {
					inRange = true
					break
				}
			}
			if !inRange {
				fmt.Println(tuningFreq(pt))
				return
			}
		}
		if y > 0 && y%1000 == 0 {
			fmt.Printf("Finished row %d (%v)\n", y, time.Since(start))
			return
		}
	}
}

type sensor struct {
	loc point
	d   float64
}

// Point

type point struct {
	x, y float64
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
	return point{float64(x), float64(y)}
}

// Utils

func manhattan(a, b point) float64 {
	return math.Abs(a.x-b.x) + math.Abs(a.y-b.y)
}

func tuningFreq(pt point) float64 {
	return pt.x*4000000 + pt.y
}
