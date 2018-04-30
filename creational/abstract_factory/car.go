package abstract_factory

import "fmt"

type Car interface {
	NumDoors() int
}

const (
	LuxuryCarType = iota
	FamilyCarType
)

type CarFactory struct{}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recongized\n", v)
	}
}
