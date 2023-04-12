package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var input = flag.String("f", "input.txt", "input file")
var prof = flag.String("profile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *prof != "" {
		f, err := os.Create(*prof)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	g := parse(*input)
	currIdx := 0
	tick := 30
	total := 0
	visited := newQueue()
	visited.enqueue(0)
	for tick > 0 {
		maxIdx, maxPotential, maxLen := -1, -1, 0
		for otherIdx, other := range g.nodes {
			if visited.contains(otherIdx) {
				continue
			}
			pathLen := shortestPath(g, currIdx, otherIdx)
			potential := (tick - (pathLen + 1)) * other.pressure
			if potential > maxPotential {
				maxIdx, maxPotential, maxLen = otherIdx, potential, pathLen
			}
		}
		tick -= (maxLen + 1)
		total += maxPotential
		visited.enqueue(maxIdx)
		fmt.Printf("start: %v; walked %v to %v; total: %v\n", currIdx, maxLen, maxIdx, total)
		currIdx = maxIdx
	}
	fmt.Println(total)
}

func shortestPath(g *graph, start, end int) int {
	var (
		scanQueue = newQueue()
		visited   = newQueue()
		depth     = make([]int, 0, g.v())
	)
	scanQueue.enqueue(start)
	visited.enqueue(start)
	depth = append(depth, 0)

	var curr int
	for scanQueue.length() > 0 {
		curr = scanQueue.dequeue()
		if curr == end {
			return depth[visited.indexOf(curr)]
		}
		for _, adj := range g.get(curr).adjacent {
			if visited.contains(adj) {
				continue
			}
			scanQueue.enqueue(adj)
			visited.enqueue(adj)
			depth = append(depth, depth[visited.indexOf(curr)]+1)
		}
	}
	return -1
}
