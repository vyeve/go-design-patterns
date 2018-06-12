package workers_pool

import (
	"fmt"
	"regexp"
	"sync"
	"testing"
)

func Test_Dispatcher(t *testing.T) {
	bufferSize := 100
	dispatcher := NewDispatcher(bufferSize)

	workers := 3

	for i := 0; i < workers; i++ {
		w := NewPrefixSuffixWorker(i, fmt.Sprintf("WorkerID: %d -> ", i), "World")
		dispatcher.LaunchWorker(w)
	}

	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := Request{
			Data: fmt.Sprintf("(Msg_id: %d) -> Hello", i),
			handler: func(i interface{}) {
				s, ok := i.(string)

				if !ok {
					t.Fail()
				}
				ok, err := regexp.Match(`WorkerID: \d+ -> \(MSG_ID: \d+\) -> [A-Z]*\sWorld]`, []byte(s))
				if !ok || err != nil {
					t.Fail()
				}
			},
		}
		dispatcher.MakeRequest(req)
	}
}
