package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func isMarker(buffer []byte, distinct int) bool {
	if len(buffer) < distinct {
		return false
	}
	for _, character := range buffer {
		if bytes.Count(buffer, []byte{character}) > 1 {
			return false
		}
	}
	return true
}

func findMarker(stream []byte, distinct int) int {
	buffer := make([]byte, 0, distinct)
	for i, character := range stream {
		if len(buffer) == distinct {
			buffer = buffer[1:]
		}
		buffer = append(buffer, character)
		if isMarker(buffer, distinct) {
			return i + 1
		}
	}
	return -1
}

func main() {
	f, _ := os.Open("input2.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(findMarker(scanner.Bytes(), 14))
	}
}
