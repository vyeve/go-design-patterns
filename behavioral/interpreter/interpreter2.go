package interpreter

type Interpreter interface {
	Read() int
}

type Value int

func (v *Value) Read() int {
	return int(*v)
}

type operationSum struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationSum) Read() int {
	return a.Left.Read() + a.Right.Read()
}

type operationSubstract struct {
	Left  Interpreter
	Right Interpreter
}

func (s *operationSubstract) Read() int {
	return s.Left.Read() - s.Right.Read()
}

func OperatorFactory(o string, left, right Interpreter) Interpreter {
	switch o {
	case SUM:
		return &operationSum{
			Left:  left,
			Right: right,
		}
	case SUB:
		return &operationSubstract{
			Left:  left,
			Right: right,
		}
	}
	return nil
}

type PolishNotationStack []Interpreter

func (p *PolishNotationStack) Push(s Interpreter) {
	*p = append(*p, s)
}

func (p *PolishNotationStack) Pop() Interpreter {
	length := len(*p)

	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}
	return nil
}
