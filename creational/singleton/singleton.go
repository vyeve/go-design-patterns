/*
The Singleton pattern provides you with a single instance of an object,
and guarantee that there are no duplicates.
*/

package creational

//Singleton interface provides all methods that are supported by our
//singleton instance
type Singleton interface {
	//increment counter by 1
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

//GetInstance method returns singleton instance
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
