package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node interface {
	Less(other Node) MaybeBool
	String() string
}

type ListNode []Node

func (a ListNode) compareLists(b ListNode) MaybeBool {
	for i := 0; i < min(len(a), len(b)); i++ {
		switch justLess := a[i].Less(b[i]).(type) {
		case Just:
			return justLess
		}
	}
	if len(a) == len(b) {
		return Nothing{}
	} else if len(a) < len(b) {
		return Just{true}
	} else {
		return Just{false}
	}
}

func (a ListNode) Less(b Node) MaybeBool {
	switch bTyped := b.(type) {
	case ListNode:
		return a.compareLists(bTyped)
	case IntNode:
		return a.compareLists(ListNode{bTyped})
	default:
		panic("Unsupported node type.")
	}
}

func (l ListNode) String() string {
	strs := make([]string, 0, len(l))
	for i := range l {
		strs = append(strs, l[i].String())
	}
	return fmt.Sprintf("[%s]", strings.Join(strs, ","))
}

type IntNode int

func (a IntNode) compareInts(b IntNode) MaybeBool {
	if a == b {
		return Nothing{}
	}
	return Just{a < b}
}

func (a IntNode) Less(b Node) MaybeBool {
	switch n := b.(type) {
	case ListNode:
		return ListNode{a}.compareLists(n)
	case IntNode:
		return a.compareInts(n)
	default:
		panic("Unsupported node type.")
	}
}

func (i IntNode) String() string {
	return strconv.Itoa(int(i))
}
