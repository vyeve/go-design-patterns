/*
The Abstract Factory design pattern is a new layer of grouping to achieve a bigger composite
object, which is used through its interfaces. The idea behind grouping objects in
families and grouping families is to have big factories that can be interchangeable and can grow more
easily.
*/

package abstract_factory

type Vehicle interface {
	NumWheels() int
	NumSeats() int
}
