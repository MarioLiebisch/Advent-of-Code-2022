package numbers

func Median(values []int) int {
	length := len(values)
	if length == 0 {
		return 0
	}
	return values[length/2]
}

func Mean(values []int) int {
	length := len(values)
	if length == 0 {
		return 0
	}
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum / length
}

func SummarizeGroups(array []int) []int {
	var ret []int
	var sum int = 0
	for _, num := range array {
		if num == 0 {
			ret = append(ret, sum)
			sum = 0
		} else {
			sum += num
		}
	}
	if sum != 0 {
		ret = append(ret, sum)
	}
	return ret
}
