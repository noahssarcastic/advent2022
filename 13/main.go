package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	"golang.org/x/exp/slices"
)

// func rightOrder(first, second Node) bool {
// 	maybeLess := first.Less(second)
// 	if less, err := maybeLess.Get(); err != nil {
// 		return true
// 	} else {
// 		return less
// 	}
// }

type ByPacket []ListNode

func (a ByPacket) Len() int      { return len(a) }
func (a ByPacket) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPacket) Less(i, j int) bool {
	// If equal Swap or do nothing; it doesn't matter.
	less, _ := a[i].Less(a[j]).Get()
	return less
}

func main() {
	f, err := os.Open("input_final.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	packets := make([]ListNode, 0)
	for scanner.Scan() {
		// Scan two, then skip one line.
		packets = append(packets, parsePacket(scanner.Text()))
		scanner.Scan()
		packets = append(packets, parsePacket(scanner.Text()))
		scanner.Scan()
	}
	packets = append(packets, parsePacket("[[2]]"))
	packets = append(packets, parsePacket("[[6]]"))
	sort.Sort(ByPacket(packets))
	first := slices.IndexFunc(packets, func(ln ListNode) bool {
		return ln.String() == "[[2]]"
	})
	second := slices.IndexFunc(packets, func(ln ListNode) bool {
		return ln.String() == "[[6]]"
	})
	fmt.Println((first + 1) * (second + 1))
}
