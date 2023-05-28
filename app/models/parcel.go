package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/dracco-eu-logistics/app/utilities"
)

type Parcel struct {
	ID                      string     `json:"id"`                      // Unique identifier for the parcel
	CompanyName             string     `json:"companyName"`             // Name of the company associated with the parcel
	CompanyID               string     `json:"companyID"`               // ID of the company associated with the parcel
	Description             string     `json:"description"`             // Description of the parcel
	Weight                  float64    `json:"weight"`                  // Weight of the parcel
	Dimensions              Dimensions `json:"dimensions"`              // Dimensions (length, width, height) of the parcel
	PickupLocation          Location   `json:"pickupLocation"`          // Location where the parcel is picked up
	PickUpDate              time.Time  `json:"pickUpDate"`              // Date when the parcel is picked up
	DropOffDate             time.Time  `json:"dropOffDate"`             // Date when the parcel is dropped off
	DropOffDestination      Location   `json:"destination"`             // Destination location for the parcel
	TransporterID           int64      `json:"transporterId"`           // ID of the transporter assigned to the parcel
	TransporterName         string     `json:"transporterName"`         // Name of the transporter assigned to the parcel
	TransporterEmail        string     `json:"transporterEmail"`        // Email of the transporter assigned to the parcel
	TransporterPhone        string     `json:"transporterPhone"`        // Phone number of the transporter assigned to the parcel
	TransporterLicensePlate string     `json:"transporterLicensePlate"` // License plate of the transporter assigned to the parcel
	IsAccepted              bool       `json:"isAccepted"`              // Flag indicating whether the parcel has been accepted
	IsCompleted             bool       `json:"isCompleted"`             // Flag indicating whether the parcel has been completed
	IsEuPalet               bool       `json:"IsEuPalet"`               // Flag indicating whether the parcel is an EU pallet
	OfferPrice              float64    `json:"offerPrice"`              // Offered price for transporting the parcel
	Country                 Country    `json:"country"`                 // Country where the parcel is situated
}

type Location struct {
	Address string `json:"address"` // Street address of the location
	City    string `json:"city"`    // City where the location is situated
	ZipCode string `json:"zipCode"` // Postal/ZIP code of the location
}

type Country struct {
	Name string `json:"name"` // Name of the country
	Code string `json:"code"` // ISO code representing the country
}

type Dimensions struct {
	Length float64 `json:"length"` // Length dimension of an object
	Width  float64 `json:"width"`  // Width dimension of an object
	Height float64 `json:"height"` // Height dimension of an object
}

func MD5ParcelID(parcelKey string) string {
	hash := md5.Sum([]byte(parcelKey))
	return hex.EncodeToString(hash[:])
}

func NewParcel(CompanyName, CompanyID string) *Parcel {
	return &Parcel{
		CompanyName: CompanyName,
		CompanyID:   CompanyID,
	}
}
func (parcel *Parcel) SetCountry() *Parcel {
	parcel.Country = Country{
		Name: "Spain",
		Code: "ES",
	}
	return parcel
}

func (parcel *Parcel) SetDescription(description string) *Parcel {
	parcel.Description = description
	return parcel
}

func (parcel *Parcel) SetParcelID() *Parcel {
	parcel.ID = MD5ParcelID(utilities.GenerateUniqueID("parcel"))
	return parcel
}

func (parcel *Parcel) SetPickUpLocation(PickUp Location) *Parcel {
	parcel.PickupLocation = PickUp
	return parcel
}

func (parcel *Parcel) SetDropOffDestination(Destination Location) *Parcel {
	parcel.DropOffDestination = Destination
	return parcel
}

func (parcel *Parcel) SetEuPalet() *Parcel {
	parcel.IsEuPalet = true
	return parcel
}

func (parcel *Parcel) SetDimensions(length, width, height float64) *Parcel {
	if parcel.IsEuPalet {
		parcel.Dimensions = Dimensions{
			Length: 120,
			Width:  80,
			Height: height,
		}
	} else {
		parcel.Dimensions = Dimensions{
			Length: length,
			Width:  width,
			Height: height,
		}
	}
	return parcel
}

func (parcel *Parcel) SetWeight(weight float64) *Parcel {
	parcel.Weight = weight
	return parcel
}

func (parcel *Parcel) SetTransporterID(id int64) *Parcel {
	parcel.TransporterID = id
	return parcel
}

func (parcel *Parcel) SetTransporterName(name string) *Parcel {
	parcel.TransporterName = name
	return parcel
}

func (parcel *Parcel) SetTransporterEmail(email string) *Parcel {
	parcel.TransporterEmail = email
	return parcel
}

func (parcel *Parcel) SetTransporterPhone(phone string) *Parcel {
	parcel.TransporterPhone = phone
	return parcel
}

func (parcel *Parcel) SetTransporterLicensePlate(LicensePlate string) *Parcel {
	parcel.TransporterLicensePlate = LicensePlate
	return parcel
}

func (parcel *Parcel) SetTransporter(transporter Transporter) *Parcel {
	return parcel.
		SetTransporterID(transporter.ID).
		SetTransporterName(transporter.Name).
		SetTransporterEmail(transporter.Email).
		SetTransporterPhone(transporter.PhoneNumber).
		SetTransporterLicensePlate(transporter.Vehicle.LicensePlate)
}

func (parcel *Parcel) SetPickUpDate(date time.Time) *Parcel {
	parcel.PickUpDate = date
	return parcel
}

func (parcel *Parcel) SetDropOffDate(date time.Time) *Parcel {
	parcel.DropOffDate = date
	return parcel
}

func (parcel *Parcel) SetIsAccepted() *Parcel {
	parcel.IsAccepted = true
	return parcel
}

func (parcel *Parcel) SetIsCompleted() *Parcel {
	parcel.IsCompleted = true
	return parcel
}

func (parcel *Parcel) SetOfferPrice(price float64) *Parcel {
	parcel.OfferPrice = price
	return parcel
}

// Create creates a new parcel and performs any necessary validations or business logic checks
func (parcel *Parcel) Create() (*Parcel, []error) {
	// Perform any necessary validations or business logic checks before creating the parcel
	// For example, checking if the required fields are provided, validating dimensions, etc.
	errorslist := []error{}
	// Example validation: Check if the description field is empty
	if parcel.Description == "" {
		parcel.Description = "No description provided"
	}

	// Example validation: Check if the weight is zero or negative
	if parcel.Weight <= 0 {
		errorslist = append(errorslist, errors.New("weight must be greater than zero"))
	}

	// Example validation: Check if the dimensions are valid
	if parcel.Dimensions.Length <= 0 || parcel.Dimensions.Width <= 0 || parcel.Dimensions.Height <= 0 {
		errorslist = append(errorslist, errors.New("invalid dimensions"))
	}

	// Example validation: Check if the pickup location is provided
	if parcel.PickupLocation.Address == "" || parcel.PickupLocation.City == "" || parcel.PickupLocation.ZipCode == "" {
		errorslist = append(errorslist, errors.New("pickup location is required"))
	}

	// Example validation: Check if the transporter details are provided
	if parcel.TransporterID <= 0 || parcel.TransporterName == "" || parcel.TransporterEmail == "" || parcel.TransporterPhone == "" {

		errorslist = append(errorslist, errors.New("transporter details are required"))
	}

	//EU standard palet dimensions
	if parcel.IsEuPalet {
		parcel.Dimensions.Length = 120
		parcel.Dimensions.Width = 80
	}

	// Example validation: Check if any other required fields are provided and perform additional checks as per your application's requirements

	// If all validations pass, you can proceed with creating the parcel
	// Add your code here to save the parcel in the database or perform any other necessary operations

	for _, error := range errorslist {
		fmt.Println(error)
	}
	return &Parcel{}, errorslist
}
