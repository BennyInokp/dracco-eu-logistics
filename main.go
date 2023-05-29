package main

import (
	"log"

	"github.com/dracco-eu-logistics/app/models"
	"github.com/dracco-eu-logistics/app/utilities"
)

func main() {

	customer, errs := models.NewCustomer().SetCustomerID(001).
		SetCustomerName("Naoufal").
		SetCustomerEmail("neufal777@gmail.com").
		SetCustomerPhone("87897979").
		SetCustomerIsTransporter(false).
		SetCustomerIsCompany(false).
		SetCustomerPassword("123456").
		Validate()

	log.Println(utilities.PrettyPrintStruct(customer), errs)
}

// transporter, err := models.NewTransporter().
// 	SetTransporterID(001).
// 	SetTransporterName("NaoufL").
// 	SetTransporterEmail("neufalgmail.com").
// 	SetTransporterPhone("632834694").
// 	SetTransporterLicensePlate("0211DTJ").
// 	SetTransporterVehicleBrand("mercedes").
// 	SetTransporterVehicleModel("vito").
// 	SetTransporterVehicleMaxCapacity(2).
// 	SetTransporterTaxID("42397606G").
// 	SetTransporterDocuments([]models.Documents{
// 		{
// 			DocumentType:        "AUTONOMO",
// 			DocumentDescription: "Documento de último pago de cuota autonomo",
// 			DocumentURL:         "42397606g_autonomo.pdf",
// 		},
// 		{
// 			DocumentType:        "TARJETA_TRANSPORTE",
// 			DocumentDescription: "Documento de tarjeta de transporte",
// 			DocumentURL:         "42397606g_tarjeta_transporte.pdf",
// 		},
// 	}).
// 	Verify()

// log.Println(utilities.PrettyPrintStruct(transporter))
// log.Println(err)

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
