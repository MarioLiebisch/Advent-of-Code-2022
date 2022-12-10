package cpu

import (
	"aoc2022/pkg/io"
	"strconv"
	"strings"
)

type InstructionPart func(cpu *CPU, params []int)
type Instruction []InstructionPart

func noop(cpu *CPU, params []int) {}

var ops = map[string]Instruction{
	"noop": {noop},
	"addx": {noop, func(cpu *CPU, params []int) {
		cpu.x += params[0]
	}},
}

type ProgramPart struct {
	op     Instruction
	params []int
}

type Program []ProgramPart

func ReadProgram(file string) Program {
	var p Program
	for _, line := range io.ReadLines(file) {
		data := strings.Split(line, " ")
		var params []int
		for i := 1; i < len(data); i++ {
			value, _ := strconv.Atoi(data[i])
			params = append(params, value)
		}
		p = append(p, ProgramPart{ops[data[0]], params})
	}
	return p
}
