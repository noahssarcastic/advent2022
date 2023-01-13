package main

import (
	"bufio"
	"bytes"
	"flag"
	"os"

	"github.com/noahssarcastic/advent2022/12/coord"
)

func parseArgs() string {
	inputFile := flag.String("input", "input_final.txt", "The heightmap input file.")
	flag.Parse()
	return *inputFile
}

func parseInput(inputFile string) *Map {
	f, _ := os.Open(inputFile)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	m := &Map{}
	for scanner.Scan() {
		b := scanner.Bytes()
		if i := bytes.IndexByte(b, 'S'); i >= 0 {
			m.start = &coord.Coord{i, len(m.hm)}
			b = bytes.Replace(b, []byte{'S'}, []byte{'a'}, 1)
		}
		if i := bytes.IndexByte(b, 'E'); i >= 0 {
			m.end = &coord.Coord{i, len(m.hm)}
			b = bytes.Replace(b, []byte{'E'}, []byte{'z'}, 1)
		}
		m.hm = append(m.hm, b)
	}
	return m
}
