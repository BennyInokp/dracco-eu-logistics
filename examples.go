package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dracco-eu-logistics/app/models"
	"github.com/dracco-eu-logistics/app/utilities"
)

func NewCustomerExample() {
	customer, errs := models.NewCustomer().SetCustomerID(001).
		SetCustomerName("Naoufal").
		SetCustomerEmail("neufal777@gmail.com").
		SetCustomerPhone("87897979").
		SetCustomerIsTransporter(false).
		SetCustomerIsCompany(true).
		SetCustomerPassword("123456").
		Validate()

	log.Println(utilities.PrettyPrintStruct(customer), errs)
}

func NewParcelExample() {
	parcel, err := models.NewParcel(0001).
		SetParcelID().
		SetDescription("").
		SetCountry().
		SetWeight(100).
		SetPickUpLocation(models.Location{
			Address: "Rda. de Sant Antoni, 59, 08011 Barcelona",
		}).
		SetDropOffDestination(models.Location{
			Address: "calle de providencia 33 manresa 08241",
		}).
		SetGeo().
		SetDistance(100).
		SetEuPalet().
		SetDimensions(0, 0, 150).
		SetPickUpDate(time.Now()).
		SetDropOffDate(time.Now().AddDate(0, 0, 1)).
		SetOfferPrice(450.76).

		// Once transporter is accepted, we can set the transporter ID
		SetTransporterID(929).
		SetIsAccepted().

		// Save the parcel
		Save()

	fmt.Println(utilities.PrettyPrintStruct(parcel), err)
}

func NewTransporterExample() {
	transporter, err := models.NewTransporter().
		SetTransporterID(001).
		SetTransporterName("NaoufL").
		SetTransporterEmail("neufalgmail.com").
		SetTransporterPhone("632834694").
		SetTransporterLicensePlate("0211DTJ").
		SetTransporterVehicleBrand("mercedes").
		SetTransporterVehicleModel("vito").
		SetTransporterVehicleMaxCapacity(2).
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
	log.Println(err)

}
