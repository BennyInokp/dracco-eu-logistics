package models

import (
	"log"
)

type Transporter struct {
	Vehicle Vehicle `json:"vehicle"`
}

type Vehicle struct {
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	LicensePlate string `json:"licensePlate"`
	Maxcapacity  int64  `json:"maxcapacity"`
}

func NewTransporter() *Transporter {
	return &Transporter{}
}

func (transporter *Transporter) SetTransporterLicensePlate(LicensePlate string) *Transporter {
	transporter.Vehicle.LicensePlate = LicensePlate
	return transporter
}

func (transporter *Transporter) SetTransporterVehicleBrand(brand string) *Transporter {
	transporter.Vehicle.Brand = brand
	return transporter
}

func (transporter *Transporter) SetTransporterVehicleModel(model string) *Transporter {
	transporter.Vehicle.Model = model
	return transporter
}

func (transporter *Transporter) SetTransporterVehicleMaxCapacity(maxcapacity int64) *Transporter {
	transporter.Vehicle.Maxcapacity = maxcapacity
	return transporter
}

func (transporter *Transporter) Validate() (*Transporter, []error) {
	// Verify the transporter basic information for signup
	// name, email
	errorslist := []error{}
	log.Println("Verifying transporter...")

	return transporter, errorslist

}

func (Transporter *Transporter) Save() (int64, error) {
	// Insert into database the following values
	log.Println("Inserting transporter into database...")
	return 0, nil
}
