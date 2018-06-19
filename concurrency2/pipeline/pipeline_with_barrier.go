package pipeline

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var buffSize = runtime.NumCPU()

func exampleOfPipelineWithLongFunction() {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	t := time.Now()
	for n := range generator2(data...) {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
	for n := range pow(generator2(data...)) {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
	fmt.Println(time.Since(t))

	// Output can be:
	// 1 2 3 4 5 6 7
	// 25 36 1 16 9 4 49
	// 2.003011531s

}

// This function is a long time function (hard calculation)
func longFunc(num int) int {
	time.Sleep(time.Second * 2)
	return num * num
}

// This function converts slice to channel
func generator2(nums ...int) <-chan int {
	res := make(chan int, buffSize)
	go func() {
		for _, n := range nums {
			res <- n
		}
		close(res)
	}()
	return res
}

// This is function takes a channel of data for calculation and calls a Long function
func pow(data <-chan int) <-chan int {
	res := make(chan int, buffSize)
	var wg sync.WaitGroup
	go func() {
		for d := range data {
			wg.Add(1)
			go func(n int) {
				res <- longFunc(n)
				wg.Done()
			}(d)
		}
		wg.Wait()
		close(res)
	}()
	return res
}
