package sensors

import (
	"aoc2022/pkg/io"
	"aoc2022/pkg/numbers"
	"strconv"
	"strings"
)

func Read(file string) *Tunnels {
	var res Tunnels
	first := true
	for _, line := range io.ReadLines(file) {
		data := strings.Split(line, "=")
		var s Sensor
		s.X, _ = strconv.Atoi(strings.Split(data[1], ",")[0])
		s.Y, _ = strconv.Atoi(strings.Split(data[2], ":")[0])
		s.bX, _ = strconv.Atoi(strings.Split(data[3], ",")[0])
		s.bY, _ = strconv.Atoi(data[4])
		s.R = numbers.Abs(s.X-s.bX) + numbers.Abs(s.Y-s.bY)
		res.sensors = append(res.sensors, s)
		if first || s.X < s.left {
			s.left = s.X
		}
		if first || s.X > s.right {
			s.right = s.X
		}
		if first || s.Y < s.top {
			s.top = s.Y
		}
		if first || s.Y > s.bottom {
			s.bottom = s.Y
		}
		first = false
	}
	return &res
}
