/*
The most famous interpreter is probably SQL. It's defined as a special-purpose programming language for managing
data held in relational databases. SQL is quite complex and big but, all in all, is a set of words and
operators that allow us to perform operations such as insert, select, or delete.
Another typical example is musical notation. It's a language itself and the interpreter is the musician who
knows the connection between a note and its representation on the instrument they are playing.
*/

package interpreter

import (
	"strconv"
	"strings"
)

const (
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)

func Calculate(o string) (int, error) {
	stack := polishNotationStack{}
	operators := strings.Fields(o)

	for _, operator := range operators {
		if isOperator(operator) {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := getOperationFunc(operator)
			res := mathFunc(left, right)
			stack.Push(res)
		} else {
			val, err := strconv.Atoi(operator)
			if err != nil {
				return 0, err
			}
			stack.Push(val)
		}
	}
	return stack.Pop(), nil
}

type polishNotationStack []int

func (p *polishNotationStack) Push(s int) {
	*p = append(*p, s)
}

func (p *polishNotationStack) Pop() int {
	length := len(*p)

	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}
	return 0
}

func isOperator(o string) bool {
	switch o {
	case SUM, SUB, MUL, DIV:
		return true
	}
	return false
}

func getOperationFunc(o string) func(a, b int) int {
	switch o {
	case SUM:
		return func(a, b int) int {
			return a + b
		}
	case SUB:
		return func(a, b int) int {
			return a - b
		}
	case MUL:
		return func(a, b int) int {
			return a * b
		}
	case DIV:
		return func(a, b int) int {
			return a / b
		}
	}
	return nil
}
