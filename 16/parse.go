package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const InputLineFormat = `Valve ([A-Z][A-Z]) has flow rate=(\d+); tunnels? leads? to valves? ((?:[A-Z][A-Z], )*[A-Z][A-Z])`

type parseData struct {
	label    string
	pressure int
	adjacent []string
}

func parse(input string) *graph {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	re, err := regexp.Compile(InputLineFormat)
	if err != nil {
		panic(err)
	}

	data := make([]parseData, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		match := re.FindAllStringSubmatch(scanner.Text(), 1)
		pressure, err := strconv.Atoi(match[0][2])
		if err != nil {
			panic(err)
		}
		data = append(data, parseData{
			label:    match[0][1],
			pressure: pressure,
			adjacent: strings.Split(match[0][3], ", "),
		})
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	idxMap := make(map[string]int)
	for i, el := range data {
		idxMap[el.label] = i
	}

	g := newGraph()
	for _, d := range data {
		adj := make([]int, 0)
		for _, adjacentNode := range d.adjacent {
			adj = append(adj, idxMap[adjacentNode])
		}
		g.addNode(d.pressure, d.label, adj)
	}
	return g
}
