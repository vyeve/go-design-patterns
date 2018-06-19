package tee_channel

import (
	"fmt"
)

func ExampleTee() {
	done := make(chan interface{})
	defer close(done)
	out1, out2 := tee(done, repeat(done, "hello", 3))
	for val1 := range out1 {
		fmt.Printf("out1: %v; out2: %v\n", val1, <-out2)
	}

	//Output:
	//out1: hello; out2: hello
	//out1: hello; out2: hello
	//out1: hello; out2: hello
}

func repeat(done <-chan interface{}, msg string, times int) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		for i := 0; i < times; i++ {
			select {
			case <-done:
				return
			case out <- msg:
			}
		}
	}()
	return out
}
