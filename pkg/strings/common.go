package strings

func GetCommonChars(a, b string) string {
	var ret []rune
	for _, ca := range a {
		for i, cb := range b {
			if ca == cb {
				var found = false
				for _, cr := range ret {
					if cr == ca {
						found = true
						break
					}
				}
				if !found {
					ret = append(ret, ca)
				}
				b = b[0:i] + b[i+1:]
				break
			}
		}
	}
	return string(ret)
}

func HasDuplicateChars(text string) bool {
	return len(GetUniqueChars(text)) != len(text)
}

func GetUniqueChars(text string) string {
	counts := map[rune]int{}
	for _, r := range text {
		counts[r]++
	}
	var res []rune
	for r, c := range counts {
		if c == 1 {
			res = append(res, r)
		}
	}
	return string(res)
}
