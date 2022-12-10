package cpu

import "aoc2022/pkg/numbers"

type CPU struct {
	x      int
	cycle  int
	active []ProgramPart

	signal_strength int
	screen          [240]rune
}

func NewCPU() *CPU {
	var cpu CPU
	cpu.x = 1
	cpu.cycle = 0
	for i := 0; i < 240; i++ {
		cpu.screen[i] = '.'
	}
	return &cpu
}

func (cpu *CPU) Run(p Program, threads int) {
	pc := 0
	for {
		// Append the next instruction, if there's room
		if pc < len(p) && len(cpu.active) < threads {
			cpu.active = append(cpu.active, p[pc])
			pc++
		} else if len(cpu.active) == 0 {
			// We're out of instructions, quit
			break
		}

		// Update screen
		pos := cpu.cycle % 240
		x := pos % 40
		if numbers.Abs(x-cpu.x) <= 1 {
			cpu.screen[pos] = '#'
		} else {
			cpu.screen[pos] = '.'
		}

		// Simulate cycles/update counter
		cpu.cycle++
		if cpu.cycle == 20 || (cpu.cycle-20)%40 == 0 {
			cpu.signal_strength += cpu.cycle * cpu.x
		}

		// Continue all active instructions
		for i := 0; i < len(cpu.active); {
			cpu.active[i].op[0](cpu, cpu.active[i].params)
			cpu.active[i].op = cpu.active[i].op[1:]
			if len(cpu.active[i].op) == 0 {
				cpu.active = append(cpu.active[:i], cpu.active[i+1:]...)
			} else {
				i++
			}
		}
	}
}

func (cpu *CPU) GetSignalStrength() int {
	return cpu.signal_strength
}

func (cpu *CPU) GetScreen() string {
	var ret []rune
	for i, r := range cpu.screen {
		if i%40 == 0 {
			ret = append(ret, '\n')
		}
		ret = append(ret, r)
	}
	return string(ret)
}
