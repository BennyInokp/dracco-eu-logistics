package main

import (
	"log"

	"github.com/dracco-eu-logistics/app/models"
	"github.com/dracco-eu-logistics/app/utilities"
)

func main() {
	transporter := models.NewTransporter().
		SetTransporterID(001).
		SetTransporterName("").
		SetTransporterEmail("neufal@gmail.com").
		SetTransporterPhone("632834694").
		SetTransporterLicensePlate("0211DTJ").
		SetTransporterVehicleBrand("mercedes").
		SetTransporterVehicleModel("Vito").
		SetTransporterVehicleMaxCapacity(0).
		SetTransporterTaxID("42397606G").
		SetTransporterDocuments([]models.Documents{
			{
				DocumentType:        "AUTONOMO",
				DocumentDescription: "Documento de último pago de cuota autonomo",
				DocumentURL:         "42397606g_autonomo.pdf",
			},
			{
				DocumentType:        "TARJETA_TRANSPORTE",
				DocumentDescription: "Documento de tarjeta de transporte",
				DocumentURL:         "42397606g_tarjeta_transporte.pdf",
			},
		}).
		Verify()

	log.Println(utilities.PrettyPrintStruct(transporter))
	// parcel, err := models.NewParcel("Dracco NF OÜ", "0001").
	// 	SetParcelID().
	// 	SetDescription("").
	// 	SetCountry().
	// 	SetWeight(100).
	// 	SetPickUpLocation(models.Location{
	// 		Address: "amazon 45",
	// 		City:    "El Prat",
	// 		ZipCode: "08001",
	// 	}).
	// 	SetDropOffDestination(models.Location{
	// 		Address: "amazon 45",
	// 		City:    "El Prat",
	// 		ZipCode: "08001",
	// 	}).
	// 	SetEuPalet().
	// 	SetTransporterID(001).
	// 	SetDimensions(0, 0, 150).
	// 	SetPickUpDate(time.Now()).
	// 	SetDropOffDate(time.Now().AddDate(0, 0, 1)).
	// 	SetIsAccepted().
	// 	SetOfferPrice(450.76).Create()

	// fmt.Println(parcel, err)

}
