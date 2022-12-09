package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Parse
type Execution struct {
	command   string
	arguments []string
	output    []string
}

func isCommand(s string) bool {
	return s[0] == '$'
}

func parseOutput(scanner *bufio.Scanner) []Execution {
	execs := make([]Execution, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if isCommand(line) {
			command := strings.Split(line, " ")
			execs = append(execs, Execution{
				command:   command[1],
				arguments: command[2:],
			})
		} else {
			execs[len(execs)-1].output = append(execs[len(execs)-1].output, line)
		}
	}
	return execs
}

// Build tree

const (
	file      = iota
	directory = iota
)

type Node struct {
	name           string
	nodeType, size int
	children       []*Node
	parent         *Node
}

func child(node *Node, childName string) *Node {
	for _, child := range node.children {
		if child.name == childName {
			return child
		}
	}
	return nil
}

func cd(currentNode *Node, rootNode *Node, exec Execution) *Node {
	if exec.arguments[0] == ".." {
		currentNode = currentNode.parent
	} else if exec.arguments[0] == "/" {
		currentNode = rootNode
	} else {
		currentNode = child(currentNode, exec.arguments[0])
	}
	return currentNode
}

func ls(currentNode *Node, exec Execution) {
	for _, line := range exec.output {
		output := strings.Split(line, " ")
		nodeName := output[1]
		var newNode *Node = nil
		if output[0] == "dir" {
			newNode = &Node{nodeType: directory, name: nodeName, size: -1, parent: currentNode}
		} else {
			size, _ := strconv.Atoi(output[0])
			newNode = &Node{nodeType: file, name: nodeName, size: size, parent: currentNode}
		}
		currentNode.children = append(
			currentNode.children,
			newNode)

	}
}

func createFileSystem(execs []Execution) *Node {
	root := &Node{nodeType: directory, name: "/", size: -1}
	currentNode := root
	for _, e := range execs {
		if e.command == "cd" {
			currentNode = cd(currentNode, root, e)
		} else if e.command == "ls" {
			ls(currentNode, e)
		}
	}
	return root
}

func traverse(node *Node) {
	totalSize := 0
	for _, child := range node.children {
		if child.size < 0 {
			traverse(child)
		}
		totalSize += child.size
	}
	node.size = totalSize
}

// Search tree

func findCandidateDirectories(node *Node) int {
	sum := 0
	for _, child := range node.children {
		sum += findCandidateDirectories(child)
	}
	if node.nodeType == directory && node.size <= 100000 {
		sum += node.size
	}
	return sum
}

func findSmallestToDelete(node *Node, requiredSpace int) int {
	smallest := math.MaxInt
	if node.nodeType == file {
		return smallest
	}
	for _, child := range node.children {
		childSmallest := findSmallestToDelete(child, requiredSpace)
		if childSmallest < smallest {
			smallest = childSmallest
		}
	}
	if requiredSpace <= node.size && node.size < smallest {
		smallest = node.size
	}
	return smallest
}

func main() {
	f, _ := os.Open("input2.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	execs := parseOutput(scanner)
	rootNode := createFileSystem(execs)
	traverse(rootNode)
	fmt.Printf("The sum of dirs with a size <= 100k is %v\n", findCandidateDirectories(rootNode))
	remainingSpace := 70000000 - rootNode.size
	requiredSpace := 30000000 - remainingSpace
	fmt.Printf("The smallest dirs eligible for deletion is %v\n", findSmallestToDelete(rootNode, requiredSpace))
}
