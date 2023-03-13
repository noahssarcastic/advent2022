package main

import (
	"bufio"
	"fmt"
	"os"
)

func rightOrder(first, second Node) bool {
	maybeLess := first.Less(second)
	if less, err := maybeLess.Get(); err != nil {
		return true
	} else {
		return less
	}
}

func main() {
	f, _ := os.Open("input_final.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	pairIdx := 1
	correct := 0
	for scanner.Scan() {
		first := scanner.Text()
		scanner.Scan()
		second := scanner.Text()
		scanner.Scan()
		if rightOrder(parsePacket(first), parsePacket(second)) {
			correct += pairIdx
		}
		pairIdx++
	}
	fmt.Println(correct == 6187)
}
