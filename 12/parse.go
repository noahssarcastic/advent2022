package main

import (
	"bufio"
	"bytes"
	"flag"
	"os"

	coord "github.com/noahssarcastic/advent2022/12/coord"
	hmap "github.com/noahssarcastic/advent2022/12/hmap"
)

func parseArgs() string {
	inputFile := flag.String("input", "input_final.txt", "The heightmap input file.")
	flag.Parse()
	return *inputFile
}

func parseInput(inputFile string) (hm hmap.Heightmap, start, end coord.Coord) {
	f, _ := os.Open(inputFile)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		b := scanner.Bytes()
		if i := bytes.IndexByte(b, 'S'); i >= 0 {
			start = coord.Coord{i, len(hm)}
			b = bytes.Replace(b, []byte{'S'}, []byte{'a'}, 1)
		}
		if i := bytes.IndexByte(b, 'E'); i >= 0 {
			end = coord.Coord{i, len(hm)}
			b = bytes.Replace(b, []byte{'E'}, []byte{'z'}, 1)
		}
		hm = append(hm, b)
	}
	return hm, start, end
}

func findStarts(hm hmap.Heightmap) (starts []coord.Coord) {
	for j := 0; j < hm.Height(); j++ {
		for i := 0; i < hm.Width(); i++ {
			if hm.Get(coord.Coord{i, j}) == int('a') {
				starts = append(starts, coord.Coord{i, j})
			}
		}
	}
	return starts
}
