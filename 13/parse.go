package main

import (
	"fmt"
	"strconv"
)

func parseInt(s string) (i IntNode, offset int) {
	for offset = 1; offset <= len(s); offset++ {
		temp, err := strconv.Atoi(s[:offset])
		if err == nil {
			i = IntNode(temp)
		} else {
			break
		}
	}
	return i, offset - 1
}

func parseList(s string) (list ListNode, offset int) {
	for offset = 0; offset < len(s); {
		cursor := s[offset]
		_, err := strconv.Atoi(string(cursor))
		if isInt := err == nil; isInt {
			node, newOffset := parseInt(s[offset:])
			list = append(list, node)
			offset += newOffset
		} else if cursor == '[' {
			offset++
			node, newOffset := parseList(s[offset:])
			list = append(list, node)
			offset += newOffset
		} else if cursor == ',' {
			offset++
			continue
		} else if cursor == ']' {
			return list, offset + 1
		} else {
			panic(fmt.Sprintf("Parse error: got byte '%v'\n", cursor))
		}
	}
	return list, offset
}

func parsePacket(packet string) ListNode {
	list, _ := parseList(packet[1 : len(packet)-1])
	return list
}
