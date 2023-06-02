package models

import (
	"errors"
	"log"

	"github.com/dracco-eu-logistics/app/utilities"
)

type Transporter struct {
	TransporterID    int64  `json:"id"`
	TransporterName  string `json:"transporterName"`
	TransporterEmail string `json:"transporterEmail"`
	TransporterPhone string `json:"transporterPhone"`

	Vehicle    Vehicle     `json:"vehicle"`
	IsVerified bool        `json:"isVerified"`
	TaxID      string      `json:"taxId"`
	Documents  []Documents `json:"documents"`
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
	transporter.TransporterID = id
	return transporter
}

func (transporter *Transporter) SetTransporterName(name string) *Transporter {
	transporter.TransporterName = name
	return transporter
}

func (transporter *Transporter) SetTransporterEmail(email string) *Transporter {
	transporter.TransporterEmail = email
	return transporter
}

func (transporter *Transporter) SetTransporterPhone(phone string) *Transporter {
	transporter.TransporterPhone = phone
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

func (transporter *Transporter) Verify() (*Transporter, []error) {
	// Verify the transporter basic information for signup
	// name, email
	errorslist := []error{}
	log.Println("Verifying transporter...")

	if transporter.TransporterName == "" {
		errorslist = append(errorslist, errors.New("missing name"))
	}
	if transporter.TransporterEmail != "" {
		// Validate email
		if !utilities.ValidateEmail(transporter.TransporterEmail) {
			errorslist = append(errorslist, errors.New("invalid email"))
		}
	} else {
		errorslist = append(errorslist, errors.New("missing email"))
	}

	return transporter, errorslist

}

func (Transporter *Transporter) Save() (int64, error) {
	// Insert into database the following values
	log.Println("Inserting transporter into database...")
	return 0, nil
}
