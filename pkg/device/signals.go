package device

import "aoc2022/pkg/strings"

func FindStartOfPacket(signal string, length int) int {
	l := len(signal)
	for i := 0; i < l-length; i++ {
		if !strings.HasDuplicateChars(signal[i : i+length]) {
			return i + length
		}
	}
	return 0 // Shouldn't happen
}
