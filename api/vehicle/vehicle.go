package vehicle

type Vehicle struct {
	VehicleID     int     `json:"vehicle_id"`
	TransporterID int     `json:"transporter_id"`
	LicensePlate  string  `json:"plate"`
	Brand         string  `json:"brand"`
	Model         string  `json:"model"`
	CapacityKg    float64 `json:"capacity_kg"`
}

func NewVehicle(TransporterID int) *Vehicle {
	return &Vehicle{
		TransporterID: TransporterID,
	}
}

func (v *Vehicle) SetVehicleID(VehicleID int) *Vehicle {
	v.VehicleID = VehicleID
	return v
}

func (v *Vehicle) SetLicensePlate(plate string) *Vehicle {
	v.LicensePlate = plate
	return v
}

func (v *Vehicle) SetBrand(brand string) *Vehicle {
	v.Brand = brand
	return v
}

func (v *Vehicle) SetModel(model string) *Vehicle {
	v.Model = model
	return v
}

func (v *Vehicle) SetCapacityKg(capacityKg float64) *Vehicle {
	v.CapacityKg = capacityKg
	return v
}
