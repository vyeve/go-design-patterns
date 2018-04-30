/*
The Builder pattern helps to construct complex objects without directly instantiating their struct,
or writing the logic they require.
*/

package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) Construct() {
	m.builder.SetSeats().SetStructure().SetWheels()
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}
