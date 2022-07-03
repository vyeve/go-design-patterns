package abstract_factory

import "fmt"

type Motorbike interface {
	Vehicle
	GetMotorbikeType() int
}

const (
	SportMotorbikeType = iota
	CruiseMotorbikeType
)

var (
	_ Motorbike = (*SportMotorbike)(nil)
	_ Motorbike = (*CruiseMotorbike)(nil)
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
