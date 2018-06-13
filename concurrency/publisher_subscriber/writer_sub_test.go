package publisher_subscriber

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

type mockWriter struct {
	testingFunc func(string)
}

func (m *mockWriter) Write(p []byte) (int, error) {
	m.testingFunc(string(p))
	return len(p), nil
}

func TestWriter(t *testing.T) {
	sub := NewWriterSubscriber(0, nil)
	msg := "Hello"

	var wg sync.WaitGroup
	wg.Add(1)

	stdoutPrinter := sub.(*writerSubscriber)
	stdoutPrinter.Writer = &mockWriter{
		testingFunc: func(res string) {
			if !strings.Contains(res, msg) {
				t.Fatal(fmt.Errorf("incorrect string: %s", res))
			}
			wg.Done()
		},
	}

	if err := sub.Notify(msg); err != nil {
		wg.Done()
		t.Error(err)
	}
	wg.Wait()
	sub.Close()
}
