package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

	re, err := regexp.Compile(
		`Valve ([A-Z][A-Z]) has flow rate=(\d+); tunnels? leads? to valves? ((?:[A-Z][A-Z], )*[A-Z][A-Z])`,
	)
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
	for scanner.Scan() {
		match := re.FindAllStringSubmatch(scanner.Text(), 1)
		fmt.Println(match[0][1:])
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
