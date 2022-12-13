package distress

import "sort"

func (m *Message) VerifyOrder() int {
	var res int
	for i := 0; i < len(m.packages)/2; i++ {
		if Compare(m.packages[i*2], m.packages[i*2+1]) < 1 {
			res += i + 1
		}
	}
	return res
}

// Sort.Interface for Packages
func (ps Packages) Len() int           { return len(ps) }
func (ps Packages) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }
func (ps Packages) Less(i, j int) bool { return Compare(ps[i], ps[j]) == -1 }

func (m *Message) Sort() {
	sort.Sort(m.packages)
}
