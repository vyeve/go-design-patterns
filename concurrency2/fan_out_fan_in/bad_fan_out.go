package fan_out_fan_in

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func exampleOfBadSearch() {
	max := 50000000
	done := make(chan interface{})
	defer close(done)

	t := time.Now()
	fmt.Println("Primes:")
	numbers := randIntStream(done, max)
	for prime := range take(done, primeFinder(done, numbers), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v\n", time.Since(t))

	// Output:
	// Primes:
	//	22466923
	//  ...
	//	18044753
	//	Search took: 22.458129071s
}

func randI(max int) int {
	return rand.Intn(max)
}

func randIntStream(done <-chan interface{}, max int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case out <- randI(max):
			}
		}
	}()
	return out
}

func primeFinder(done <-chan interface{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case n := <-in:
				if isPrime(n) {
					out <- n
				}
			}
		}
	}()
	return out
}

func isPrime(n int) bool {
	if n < 4 {
		return true
	}
	for i := n - 1; i > 1; i-- {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func take(done <-chan interface{}, in <-chan int, count int) <-chan int {
	out := make(chan int, count)

	go func() {
		unique := make(map[int]bool, count)
		defer close(out)
		bound := 0
		for {
			select {
			case <-done:
				return
			case n := <-in:
				if !unique[n] {
					unique[n] = true
					out <- n
					bound++
					if bound == count {
						return
					}
				}

			}
		}
	}()
	return out
}
