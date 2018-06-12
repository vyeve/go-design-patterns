package workers_pool

import (
	"fmt"
	"strings"
)

type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}

type PrefixSuffixWorker struct {
	id      int
	prefixS string
	suffixS string
}

func NewPrefixSuffixWorker(id int, prefixS string, suffixS string) WorkerLauncher {
	return &PrefixSuffixWorker{
		id:      id,
		prefixS: prefixS,
		suffixS: suffixS,
	}
}

func (w *PrefixSuffixWorker) LaunchWorker(in chan Request) {
	w.prefix(w.append(w.uppercase(in)))
}

func (w *PrefixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)

			if !ok {
				msg.handler(nil)
				continue
			}

			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()

	return out
}

func (w *PrefixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			uppercaseString, ok := msg.Data.(string)
			if !ok {
				msg.handler(nil)
				continue
			}

			msg.Data = fmt.Sprintf("%s%s", uppercaseString, w.suffixS)
			out <- msg
		}
		close(out)
	}()

	return out
}

func (w *PrefixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercaseStringWithSuffix, ok := msg.Data.(string)
			if !ok {
				msg.handler(nil)
				continue
			}
			msg.handler(fmt.Sprintf("%s%s", w.prefixS, uppercaseStringWithSuffix))
		}
	}()
}
