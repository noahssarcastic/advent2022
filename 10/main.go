package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Execution struct {
	cyclesLeft  int
	instruction Instruction
	args        []string
	register    map[string]float64
}

func NewExecution() *Execution {
	return &Execution{
		register: map[string]float64{"x": 1}}
}

func (exec *Execution) start() { exec.instruction.start(exec) }
func (exec *Execution) each()  { exec.instruction.each(exec) }
func (exec *Execution) stop()  { exec.instruction.stop(exec) }

func (exec *Execution) setInstruction(line string) {
	instructionSet := map[string]Instruction{
		"noop": NullInstruction{},
		"addx": AddX{},
	}
	exec.instruction = instructionSet[strings.Split(line, " ")[0]]
	exec.args = strings.Split(line, " ")[1:]
}

type Instruction interface {
	start(exec *Execution)
	each(exec *Execution)
	stop(exec *Execution)
}

type NullInstruction struct{}

func (n NullInstruction) start(exec *Execution) { exec.cyclesLeft = 1 }
func (n NullInstruction) each(exec *Execution)  { exec.cyclesLeft-- }
func (n NullInstruction) stop(exec *Execution)  {}

type AddX struct{}

func (n AddX) start(exec *Execution) { exec.cyclesLeft = 2 }
func (n AddX) each(exec *Execution)  { exec.cyclesLeft-- }
func (n AddX) stop(exec *Execution) {
	x, _ := strconv.ParseFloat(exec.args[0], 64)
	exec.register["x"] += x
}

func main() {
	f, _ := os.Open("input3.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	instructions := make([]string, 0)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	// Run program
	exec := NewExecution()
	counter := 0.0
	for cycle, i := 1, 0; i < len(instructions); cycle++ {
		if (cycle-20)%40 == 0 {
			counter += float64(cycle) * exec.register["x"]
		}
		if exec.cyclesLeft < 1 {
			exec.setInstruction(instructions[i])
			exec.start()
		}
		exec.each()
		if exec.cyclesLeft < 1 {
			exec.stop()
			i++
		}
	}
	fmt.Println(counter)
}
