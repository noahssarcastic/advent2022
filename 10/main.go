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

// Run at the beginning of an instruction execution.
func (exec *Execution) start() { exec.instruction.start(exec) }

// Run every cycle an instruction is active.
func (exec *Execution) each() { exec.instruction.each(exec) }

// Run at the end of an instruction execution.
func (exec *Execution) stop() { exec.instruction.stop(exec) }

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

func drawPixel(cycle int, exec *Execution) string {
	drawCursor := float64((cycle - 1) % 40)
	spritePos := exec.register["x"]
	if spritePos-1 <= drawCursor && drawCursor <= spritePos+1 {
		return "#"
	} else {
		return "."
	}
}

func main() {
	f, _ := os.Open("input3.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	instructions := make([]string, 0)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	exec := NewExecution()
	crtLine := ""
	for cycle, i := 1, 0; i < len(instructions); cycle++ {
		crtLine += drawPixel(cycle, exec)
		if len(crtLine) >= 40 {
			fmt.Println(crtLine)
			crtLine = ""
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
}
