package main

import (
	"time"

	"github.com/dracco-eu-logistics/app/models"
)

func main() {
	parcel := models.NewParcel("Dracco NF OÜ", "0001").
		SetParcelID().
		SetDescription("").
		SetCountry().
		SetWeight(100).
		SetPickUpLocation(models.Location{
			Address: "amazon 45",
			City:    "El Prat",
			ZipCode: "08001",
		}).
		SetDropOffDestination(models.Location{
			Address: "amazon 45",
			City:    "El Prat",
			ZipCode: "08001",
		}).
		SetEuPalet().
		SetTransporter(models.Transporter{
			ID:          1,
			Name:        "John Doe",
			Email:       "neufal888@gmail.com",
			PhoneNumber: "632987654",
			Vehicle: models.Vehicle{
				Brand:        "Mercedes",
				Model:        "Sprinter",
				LicensePlate: "1234ABC",
				Maxcapacity:  1000,
			},
		}).
		SetDimensions(0, 0, 150).
		SetPickUpDate(time.Now()).
		SetDropOffDate(time.Now().AddDate(0, 0, 1)).
		SetIsAccepted().
		SetOfferPrice(450.76)

	_, _ = parcel.Create()

	// fmt.Printf("Parcel Information: %+v \n", utilities.PrettyPrintStruct(parcel))
}
