package main

import (
	"log"

	"github.com/dracco-eu-logistics/app/models"
	"github.com/dracco-eu-logistics/app/utilities"
)

func main() {
	customer, errs := models.NewCustomer().
		SetCustomerID().
		SetCustomerName("Naoufal Dahouli").
		SetCustomerEmail("neufSal79@gmail.com").
		SetCustomerPhone("90000000").
		SetCustomerIsTransporter().
		SetCustomerPassword("9876543").
		Save()

	log.Println(utilities.PrettyPrintStruct(customer), errs)
}
