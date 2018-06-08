package singleton_concurrency

import "sync"

type singletonM struct {
	count int
	sync.RWMutex
}

var instanceM singletonM

func GetInstanceM() *singletonM {
	return &instanceM
}

func (s *singletonM) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

func (s *singletonM) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
