package encryption

import (
	"aoc2022/pkg/io"
	"strconv"
)

func ReadFile(file string) *File {
	var res File
	for i, line := range io.ReadLines(file) {
		v, _ := strconv.Atoi(line)
		res = append(res, entry{value: v, origin: i})
	}
	return &res
}
