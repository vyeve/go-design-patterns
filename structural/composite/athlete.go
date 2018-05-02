/*
The Composite design pattern favors composition (commonly defined as a 'has a' relationship)
over inheritance (an 'is a' relationship.
In the Composite design pattern, you will create hierarchies and trees of objects. Objects have different
objects with their own fields and methods inside them. This approach is very powerfull and solves many
problems of inheritance and multiple inheritances.
*/

package composite

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func Swim() {
	fmt.Println("Swimming!")
}

type (
	Swimmer interface {
		Swim()
	}
	Trainer interface {
		Train()
	}
)

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}
