package monkeys2

func (ms *Monkeys) Resolve(name string) int {
	m := (*ms)[name]
	// Monkey has a number
	if m.number != None {
		return m.number
	}
	// Calculate this monkey's number and cache/return it
	switch m.op {
	case Add: // result = op1 + op2
		m.number = ms.Resolve(m.op1) + ms.Resolve(m.op2)
	case Sub: // result = op1 - op2
		m.number = ms.Resolve(m.op1) - ms.Resolve(m.op2)
	case Mul: // result = op1 * op2
		m.number = ms.Resolve(m.op1) * ms.Resolve(m.op2)
	case Div: // result = op1 / op2
		m.number = ms.Resolve(m.op1) / ms.Resolve(m.op2)
	}
	m.op = None
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
			// Solve the second operand the regular way
			other := ms.Resolve(m.op2)
			if n == "root" {
				// This is root, so first operand is the same
				return other
			} else {
				// Get the expected result for this monkey
				result := ms.ResolveReversed(n)
				// Reverse the operation and return the first operand
				switch m.op {
				case Add: // result = x + other
					return result - other
				case Sub: // result = x - other
					return result + other
				case Mul: // result = x * other
					return result / other
				case Div: // result = x / other
					return result * other
				}
			}
			break
		} else if m.op2 == name {
			// Solve the first operand the regular way
			other := ms.Resolve(m.op1)
			if n == "root" {
				// This is root, so second operand is the same
				return other
			} else {
				// Get the expected result for this monkey
				result := ms.ResolveReversed(n)
				// Reverse the operation and return the second operand
				switch m.op {
				case Add: // result = other + x
					return result - other
				case Sub: // result = other - x
					return other - result
				case Mul: // result = other * x
					return result / other
				case Div: // result = other / x
					return other / result
				}
			}
			break
		}
	}
	return 0
}
