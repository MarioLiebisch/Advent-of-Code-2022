package monkeys2

func (ms *Monkeys) Resolve(name string) int {
	m := (*ms)[name]
	// Monkey has a number
	if m.number != None {
		return m.number
	}
	// Calculate this monkey's number and cache/return it
	switch m.op {
	case Add:
		m.number = ms.Resolve(m.op1) + ms.Resolve(m.op2)
		m.op = None
		return m.number
	case Sub:
		m.number = ms.Resolve(m.op1) - ms.Resolve(m.op2)
		m.op = None
		return m.number
	case Mul:
		m.number = ms.Resolve(m.op1) * ms.Resolve(m.op2)
		m.op = None
		return m.number
	case Div:
		m.number = ms.Resolve(m.op1) / ms.Resolve(m.op2)
		m.op = None
		return m.number
	}
	// Shouldn't ever happen
	return m.number
}

func (ms *Monkeys) ResolveReversed(name string) int {
	// Find the monkey waiting for the given name/monkey
	for n, m := range *ms {
		// Skip non-computed monkeys
		if m.number != None {
			continue
		}
		if m.op1 == name {
			// Solve the first operand the regular way
			other := ms.Resolve(m.op2)
			if n == "root" {
				// This is root, expected result is the other value
				return other
			} else {
				// Get the expected result for this monkey
				result := ms.ResolveReversed(n)
				// Reverse the operation and return the other operand
				switch m.op {
				case Add:
					return result - other
				case Sub:
					return result + other
				case Mul:
					return result / other
				case Div:
					return result * other
				}
			}
			break
		} else if m.op2 == name {
			// Solve the first operand the regular way
			other := ms.Resolve(m.op1)
			if n == "root" {
				// This is root, expected result is the other value
				return other
			} else {
				// Get the expected result for this monkey
				result := ms.ResolveReversed(n)
				// Reverse the operation and return the other operand
				switch m.op {
				case Add:
					return result - other
				case Sub:
					return other - result
				case Mul:
					return result / other
				case Div:
					return other / result
				}
			}
			break
		}
	}
	return 0
}
