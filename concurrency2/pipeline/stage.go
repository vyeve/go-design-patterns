/*
There are the properties of a pipeline stage:
- A stage consumes and returns the same type.
- A stage must be reified.
*/

package pipeline

import "fmt"

func multiply(values []int, multiplier int) []int {
	multipliedValues := make([]int, len(values))
	for i, v := range values {
		multipliedValues[i] = v * multiplier
	}
	return multipliedValues
}

func add(values []int, additive int) []int {
	addedValues := make([]int, len(values))
	for i, v := range values {
		addedValues[i] = v + additive
	}
	return addedValues
}

func example() {
	ints := []int{1, 2, 3, 4}

	for _, v := range multiply(add(multiply(ints, 2), 1), 2) {
		fmt.Println(v)

		// the output will be
		// 6
		// 10
		// 14
		// 18
	}
}
