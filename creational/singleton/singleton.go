/*
The Singleton pattern provides you with a single instance of an object,
and guarantee that there are no duplicates.
*/

package singleton

import "sync"

// Singleton interface provides all methods that are supported by our
// singleton instance
type Singleton interface {
	//increment counter by 1
	AddOne() int
	// return current count
	GetCount() int
}

type singleton struct {
	sync.RWMutex
	count int
}

var instance *singleton

// GetInstance method returns singleton instance
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

// AddOne increments counter
func (s *singleton) AddOne() int {
	s.Lock()
	defer s.Unlock()
	s.count++
	return s.count
}

// GetCount returns current count
func (s *singleton) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
