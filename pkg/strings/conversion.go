package strings

import (
	"strconv"
	"strings"
)

func ParseRange(text string) (int, int) {
	var parts = strings.Split(text, "-")
	if len(parts) == 1 {
		var value, _ = strconv.Atoi(parts[0])
		return value, value
	}
	var v1, _ = strconv.Atoi(parts[0])
	var v2, _ = strconv.Atoi(parts[1])
	if v1 < v2 {
		return v1, v2
	}
	return v2, v1
}

func ToIntSlice(input []string) []int {
	res := make([]int, len(input))
	for i, v := range input {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}
