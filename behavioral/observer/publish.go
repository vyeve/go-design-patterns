package observer

import "fmt"

type Observer interface {
	Notify(string)
}

type Publisher struct {
	ObserverList []Observer
}

func (s *Publisher) AddObserver(o Observer) {
	s.ObserverList = append(s.ObserverList, o)
}

func (s *Publisher) RemoveObserver(o Observer) {
	var indexToRemove int
	for i, observer := range s.ObserverList {
		if observer == o {
			indexToRemove = i
			break
		}
	}
	s.ObserverList = append(s.ObserverList[:indexToRemove], s.ObserverList[indexToRemove+1:]...)
}

func (s *Publisher) NotifyObservers(m string) {
	fmt.Printf("Publisher received message '%s' to notify observers\n", m)
	for _, observer := range s.ObserverList {
		observer.Notify(m)
	}
}
