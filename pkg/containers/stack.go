package containers

type RuneStack []rune

func (s *RuneStack) Push(v rune) {
	*s = append(*s, v)
}

func (s *RuneStack) Pop() rune {
	var i = len(*s) - 1
	var r = (*s)[i]
	*s = (*s)[0:i]
	return r
}

func (s *RuneStack) Top() rune {
	var l = len(*s)
	if l == 0 {
		return ' '
	}
	return (*s)[l-1]
}

func (s *RuneStack) PushN(v []rune) {
	*s = append(*s, v...)
}

func (s *RuneStack) PopN(n int) []rune {
	var l = len(*s)
	var r = (*s)[l-n:]
	*s = (*s)[0 : l-n]
	return r
}
