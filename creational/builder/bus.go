package creational

type BusBuilder struct {
	v VehicleProduct
}

func (bb *BusBuilder) SetWheels() BuildProcess {
	bb.v.Wheels = 4 * 2
	return bb
}

func (bb *BusBuilder) SetSeats() BuildProcess {
	bb.v.Seats = 30
	return bb
}

func (bb *BusBuilder) SetStructure() BuildProcess {
	bb.v.Structure = "Bus"
	return bb
}

func (bb *BusBuilder) GetVehicle() VehicleProduct {
	return bb.v
}
