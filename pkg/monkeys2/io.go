package monkeys2

import (
	"aoc2022/pkg/io"
	"strconv"
	"strings"
)

func ReadMonkeys(file string) Monkeys {
	res := make(Monkeys)
	for _, line := range io.ReadLines(file) {
		tmp := strings.Split(line, ": ")
		name := tmp[0]
		var m Monkey
		if tmp[1][0] >= '0' && tmp[1][0] <= '9' {
			m.number, _ = strconv.Atoi(tmp[1])
			m.op = None
		} else {
			m.number = None
			tmp = strings.Split(tmp[1], " ")
			m.op1 = tmp[0]
			m.op2 = tmp[2]
			switch tmp[1] {
			case "+":
				m.op = Add
			case "-":
				m.op = Sub
			case "*":
				m.op = Mul
			case "/":
				m.op = Div
			default: // Shouldn't happen
				m.op = None
			}
		}
		res[name] = m
	}
	return res
}
