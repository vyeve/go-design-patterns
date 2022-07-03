package abstract_factory

import "fmt"

type Car interface {
	Vehicle
	NumDoors() int
}

const (
	LuxuryCarType = iota
	FamilyCarType
)

var (
	_ Car = (*LuxuryCar)(nil)
	_ Car = (*FamilyCar)(nil)
)

type CarFactory struct{}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized\n", v)
	}
}
