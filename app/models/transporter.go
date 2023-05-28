package models

import "log"

type Transporter struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	PhoneNumber string      `json:"phoneNumber"`
	Vehicle     Vehicle     `json:"vehicle"`
	IsVerified  bool        `json:"isVerified"`
	TaxID       string      `json:"taxId"`
	Documents   []Documents `json:"documents"`
}
type Documents struct {
	DocumentType        string `json:"documentType"`
	DocumentDescription string `json:"documentDescription"`
	DocumentURL         string `json:"documentUrl"`
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

func (transporter *Transporter) SetTransporterID(id int64) *Transporter {
	transporter.ID = id
	return transporter
}

func (transporter *Transporter) SetTransporterName(name string) *Transporter {
	transporter.Name = name
	return transporter
}

func (transporter *Transporter) SetTransporterEmail(email string) *Transporter {
	transporter.Email = email
	return transporter
}

func (transporter *Transporter) SetTransporterPhone(phone string) *Transporter {
	transporter.PhoneNumber = phone
	return transporter
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

func (transporter *Transporter) SetTransporterIsVerified(isVerified string) *Transporter {
	transporter.IsVerified = true
	return transporter
}

func (transporter *Transporter) SetTransporterTaxID(taxID string) *Transporter {
	transporter.TaxID = taxID
	return transporter
}

func (transporter *Transporter) SetTransporterDocuments(documents []Documents) *Transporter {
	transporter.Documents = documents
	return transporter
}

func (transporter *Transporter) Verify() *Transporter {
	// Verify the transporter
	log.Println("Verifying transporter...")
	if transporter.Name != "" {
		transporter.Insert()
	}
	return transporter
}

func (Transporter *Transporter) Insert() (int64, error) {
	// Insert into database the following values
	log.Println("Inserting transporter into database...")
	return 0, nil
}
