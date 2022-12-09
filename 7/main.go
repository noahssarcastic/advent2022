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
	name     string
	nodeType int
	size     int
	children []*Node
	parent   *Node
}

func getChild(node *Node, childName string) *Node {
	for _, child := range node.children {
		if child.name == childName {
			return child
		}
	}
	return nil
}

func cd(currentNode *Node, rootNode *Node, exec Execution) *Node {
	if exec.arguments[0] == ".." {
		return currentNode.parent
	} else if exec.arguments[0] == "/" {
		return rootNode
	} else {
		return getChild(currentNode, exec.arguments[0])
	}
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

func calculateDirSizes(node *Node) {
	totalSize := 0
	for _, child := range node.children {
		if child.size < 0 {
			calculateDirSizes(child)
		}
		totalSize += child.size
	}
	node.size = totalSize
}

// Search tree

func countDirsUnder(node *Node) int {
	if node.nodeType == file {
		return 0
	}

	sum := 0
	for _, child := range node.children {
		sum += countDirsUnder(child)
	}

	if node.size <= 100000 {
		sum += node.size
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func smallestToDelete(node *Node, requiredSpace int) int {
	// Base case; files are leaf nodes
	if node.nodeType == file {
		return math.MaxInt
	}

	smallest := math.MaxInt
	for _, child := range node.children {
		smallest = min(smallest, smallestToDelete(child, requiredSpace))
	}

	if requiredSpace <= node.size && node.size < smallest {
		return node.size
	}
	return smallest
}

func main() {
	f, _ := os.Open("input2.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	execs := parseOutput(scanner)

	rootNode := createFileSystem(execs)
	calculateDirSizes(rootNode)

	fmt.Printf(
		"The sum of dirs with a size <= 100k is %v\n",
		countDirsUnder(rootNode))

	remainingSpace := 70000000 - rootNode.size
	requiredSpace := 30000000 - remainingSpace
	fmt.Printf(
		"The smallest dir eligible for deletion is %v\n",
		smallestToDelete(rootNode, requiredSpace))
}
