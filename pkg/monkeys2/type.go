package monkeys2

const (
	None = -1
	Add  = 0
	Sub  = 1
	Mul  = 2
	Div  = 3
)

type Monkeys map[string]Monkey
type Monkey struct {
	number, op int
	op1, op2   string
}
