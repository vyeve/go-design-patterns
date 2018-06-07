package observer

import (
	"fmt"
	"testing"
)

type TestObserver struct {
	ID      int
	Message string
}

func (p *TestObserver) Notify(m string) {
	fmt.Printf("Observer %d: message '%s' received\n", p.ID, m)
	p.Message = m
}

func TestSubject(t *testing.T) {
	testObserver1 := &TestObserver{1, ""}
	testObserver2 := &TestObserver{2, ""}
	testObserver3 := &TestObserver{3, ""}
	publisher := Publisher{}

	t.Run("AddObserver", func(t *testing.T) {
		publisher.AddObserver(testObserver1)
		publisher.AddObserver(testObserver2)
		publisher.AddObserver(testObserver3)

		if len(publisher.ObserverList) != 3 {
			t.Fail()
		}
	})

	t.Run("RemoveObserver", func(t *testing.T) {
		publisher.RemoveObserver(testObserver2)

		if len(publisher.ObserverList) != 2 {
			t.Errorf("The size of the observer list is not the expected. 3 != %d\n", len(publisher.ObserverList))
		}

		for _, observer := range publisher.ObserverList {
			testObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
			}
			if testObserver.ID == 2 {
				t.Fail()
			}
		}
	})

	t.Run("Notify", func(t *testing.T) {
		if len(publisher.ObserverList) == 0 {
			t.Error("the list is empty. Nothing to test\n")
		}
		for _, observer := range publisher.ObserverList {
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}
			if printObserver.Message != "" {
				t.Errorf("The observer's Message field weren't empty: %s\n", printObserver.Message)
			}
		}

		message := "Hello World!"
		publisher.NotifyObservers(message)

		for _, observer := range publisher.ObserverList {
			printObserver, ok := observer.(*TestObserver)

			if !ok {
				t.Fail()
				break
			}

			if printObserver.Message != message {
				t.Errorf("expected message on observer %d was not expected: '%s' != '%s'\n",
					printObserver.ID, printObserver.Message, message)
			}
		}
	})
}
