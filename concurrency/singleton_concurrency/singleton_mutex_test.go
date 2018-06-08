package singleton_concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestStartInstanceM(*testing.T) {
	singleton := GetInstanceM()
	singleton2 := GetInstanceM()

	n := 5000

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
}
