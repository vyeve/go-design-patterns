package abstract_factory

import "fmt"

type Motorbike interface {
	GetMotorbikeType() int
}

const (
	SportMotorbikeType = iota
	CruiseMotorbikeType
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized\n", v)
	}
}
