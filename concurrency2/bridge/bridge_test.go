package bridge

import "fmt"

func ExampleBridge() {
	done := make(chan interface{})
	defer close(done)
	for v := range bridge(done, genVals()) {
		fmt.Printf("%v ", v)
	}
	// Output:
	// 0 1 2 3 4 5 6 7 8 9
}

func genVals() <-chan <-chan interface{} {
	chanStream := make(chan (<-chan interface{}))
	go func() {
		defer close(chanStream)
		for i := 0; i < 10; i++ {
			stream := make(chan interface{}, 1)
			stream <- i
			close(stream)
			chanStream <- stream
		}
	}()
	return chanStream
}
