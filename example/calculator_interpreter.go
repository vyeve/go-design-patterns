package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-design-patterns/behavioral/interpreter"
)

func main() {
	stack := interpreter.PolishNotationStack{}
	operators := strings.Fields("3 4 sum 2 sub")

	for _, operator := range operators {
		if operator == interpreter.SUM || operator == interpreter.SUB {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := interpreter.OperatorFactory(operator, left, right)
			res := interpreter.Value(mathFunc.Read())
			stack.Push(&res)
		} else {
			val, err := strconv.Atoi(operator)
			if err != nil {
				panic(err)
			}

			temp := interpreter.Value(val)
			stack.Push(&temp)
		}
	}

	fmt.Println(int(stack.Pop().Read()))
}
