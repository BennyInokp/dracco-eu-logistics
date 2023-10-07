package transporter

import (
	"errors"
	"log"

	"github.com/dracco-eu-logistics/api/vehicle"
)

type Transporter struct {
	TransporterID int             `json:"transporter_id"`
	UserID        int             `json:"user_id"`
	Vehicle       vehicle.Vehicle `json:"vehicle"`
}

func NewTransporter(UserID int) *Transporter {
	return &Transporter{
		UserID: UserID,
	}
}

func (t *Transporter) SetTransporterID(TransporterID int) *Transporter {
	t.TransporterID = TransporterID
	return t
}

func (t *Transporter) SetTransporterVehicle(vehicle vehicle.Vehicle) *Transporter {
	t.Vehicle = vehicle
	return t
}

func (t *Transporter) Validate() (*Transporter, []error) {
	errorslist := []error{}
	log.Println("Verifying transporter...")

	// Perform validation checks here if needed
	if t.UserID <= 0 {
		errorslist = append(errorslist, errors.New("user_id is required"))
	}

	return t, errorslist
}

func (t *Transporter) Save() (int64, error) {
	// Insert the transporter into the database
	log.Println("Inserting transporter into the database...")

	// Perform the database insertion logic here and return the inserted ID (if available) and any error

	// For example:
	// db.Exec("INSERT INTO dbo.transporters (user_id) VALUES (?)", t.UserID)
	// transporterID := // Get the inserted ID
	// return transporterID, nil

	return 0, nil // Replace with actual logic
}
