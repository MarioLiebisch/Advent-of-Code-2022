package distress

import "aoc2022/pkg/io"

func ReadPackage(input string) (*Package, int) {
	var res Package
	var pos int
	if input[0] == '[' {
		// Read all sub elements
		res.value = -1
		res.contents = make([]*Package, 0)
		for pos++; pos < len(input); pos++ {
			sub, offset := ReadPackage(input[pos:])
			if offset > 0 {
				res.contents = append(res.contents, sub)
				pos += offset
			}
			if input[pos] == ']' {
				pos++
				break
			}
		}
	} else {
		// Read a single integer
		ok := false
		for ; input[pos] >= '0' && input[pos] <= '9'; pos++ {
			res.value = res.value*10 + int(input[pos]-'0')
			ok = true
		}
		// Nothing was read
		if !ok {
			return nil, 0
		}
	}
	return &res, pos
}

func ReadSignal(file string) *Message {
	var res Message
	var p1, p2 *Package
	for i, line := range io.ReadLines(file) {
		// Skip the empty lines (always the third)
		switch i % 3 {
		case 0: // First message
			p1, _ = ReadPackage(line)
		case 1: // Second message
			p2, _ = ReadPackage(line)
			res.packages = append(res.packages, p1, p2)
		case 3: // Empty line
			continue
		}
	}
	return &res
}
