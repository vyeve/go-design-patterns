package main

import (
	"fmt"
	"sync"

	"github.com/go-design-patterns/concurrency/workers_pool"
)

func main() {
	bufferSize := 100
	dispatcher := workers_pool.NewDispatcher(bufferSize)

	workers := 3

	for i := 0; i < workers; i++ {
		w := workers_pool.NewPrefixSuffixWorker(
			i,
			fmt.Sprintf("WorkerID: %d ->", i),
			" World",
		)
		dispatcher.LaunchWorker(w)
	}

	requests := 200

	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := workers_pool.NewStringRequest("(Msg_id: %d) -> Hello", i, &wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()
	wg.Wait()
}
