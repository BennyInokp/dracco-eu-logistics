package models

import (
	"errors"
	"time"

	"github.com/dracco-eu-logistics/app/utilities"
)

// Parcel represents a parcel
type Parcel struct {
	ID        int   `json:"id"`        // Unique identifier for the parcel
	CompanyID int64 `json:"companyID"` // ID of the company associated with the parcel

	Description        string     `json:"description"`      // Description of the parcel
	Weight             float64    `json:"weight"`           // Weight of the parcel
	Dimensions         Dimensions `json:"dimensions"`       // Dimensions (length, width, height) of the parcel
	PickupLocation     Location   `json:"pickupLocation"`   // Location where the parcel is picked up
	PickUpDate         time.Time  `json:"pickUpDate"`       // Date when the parcel is picked up
	DropOffDate        time.Time  `json:"dropOffDate"`      // Date when the parcel is dropped off
	DropOffDestination Location   `json:"destination"`      // Destination location for the parcel
	IsEuPalet          bool       `json:"IsEuPalet"`        // Flag indicating whether the parcel is an EU pallet
	OfferPrice         float64    `json:"offerPrice"`       // Offered price for transporting the parcel
	Country            Country    `json:"country"`          // Country where the parcel is situated
	TransporterID      int64      `json:"transporterId"`    // ID of the transporter assigned to the parcel
	IsAccepted         bool       `json:"isAccepted"`       // Flag indicating whether the parcel has been accepted
	IsCompleted        bool       `json:"isCompleted"`      // Flag indicating whether the parcel has been completed
	Distance           float64    `json:"distance"`         // Distance between the pick up and drop off locations
	UpdatedLocations   []Location `json:"updatedLocations"` // List of updated locations for the parcel
}

// location represents a location
type Location struct {
	Address   string  `json:"address"`   // Street address of the location
	City      string  `json:"city"`      // City where the location is situated
	ZipCode   string  `json:"zipCode"`   // Postal/ZIP code of the location
	Longitude float64 `json:"longitude"` // Longitude of the location
	Latitude  float64 `json:"latitude"`  // Latitude of the location
}

// Country represents a country
type Country struct {
	Name string `json:"name"` // Name of the country
	Code string `json:"code"` // ISO code representing the country
}

// Dimensions represents the dimensions of an object
type Dimensions struct {
	Length float64 `json:"length"` // Length dimension of an object
	Width  float64 `json:"width"`  // Width dimension of an object
	Height float64 `json:"height"` // Height dimension of an object
}

// MD5ParcelID generates an MD5 hash of the parcel key
// func MD5ParcelID(parcelKey string) string {
// 	hash := md5.Sum([]byte(parcelKey))
// 	return hex.EncodeToString(hash[:])
// }

// NewParcel creates a new parcel
func NewParcel(CompanyID int64) *Parcel {
	return &Parcel{
		CompanyID: CompanyID,
	}
}

// setcountry sets the country of the parcel
func (parcel *Parcel) SetCountry() *Parcel {
	parcel.Country = Country{
		Name: "Spain",
		Code: "ES",
	}
	return parcel
}

// SetDescription sets the description of the parcel
func (parcel *Parcel) SetDescription(description string) *Parcel {
	parcel.Description = description
	return parcel
}

// SetParcelID sets the ID of the parcel
func (parcel *Parcel) SetParcelID() *Parcel {
	parcel.ID = utilities.RandomNumber()
	return parcel
}

// SetPickUpLocation sets the pick up location for the parcel
func (parcel *Parcel) SetPickUpLocation(PickUp Location) *Parcel {
	parcel.PickupLocation = PickUp
	return parcel
}

// SetDropOffDestination sets the drop off destination for the parcel
func (parcel *Parcel) SetDropOffDestination(Destination Location) *Parcel {
	parcel.DropOffDestination = Destination
	return parcel
}

// setEuPalet sets the IsEuPalet flag to true
func (parcel *Parcel) SetEuPalet() *Parcel {
	parcel.IsEuPalet = true
	return parcel
}

// SetDimensions sets the dimensions of the parcel
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

// SetWeight sets the weight of the parcel
func (parcel *Parcel) SetWeight(weight float64) *Parcel {
	parcel.Weight = weight
	return parcel
}

// SetPickUpDate sets the pick up date for the parcel
func (parcel *Parcel) SetPickUpDate(date time.Time) *Parcel {
	parcel.PickUpDate = date
	return parcel
}

// SetDropOffDate sets the drop off date for the parcel
func (parcel *Parcel) SetDropOffDate(date time.Time) *Parcel {
	parcel.DropOffDate = date
	return parcel
}

// SetIsAccepted sets the IsAccepted flag to true
func (parcel *Parcel) SetIsAccepted() *Parcel {
	parcel.IsAccepted = true
	return parcel
}

// SetIsCompleted sets the IsCompleted flag to true
func (parcel *Parcel) SetIsCompleted() *Parcel {
	parcel.IsCompleted = true
	return parcel
}

// SetOfferPrice sets the offered price for transporting the parcel
func (parcel *Parcel) SetOfferPrice(price float64) *Parcel {
	parcel.OfferPrice = price
	return parcel
}

// SetTransporterID sets the ID of the transporter assigned to the parcel
func (parcel *Parcel) SetTransporterID(ID int64) *Parcel {
	parcel.TransporterID = ID
	return parcel
}

// SetDistance sets the distance between the pick up and drop off locations
func (parcel *Parcel) SetDistance(distance float64) *Parcel {
	parcel.Distance = distance
	return parcel
}

// SetGeo sets the longitude and latitude
func (parcel *Parcel) SetGeo() *Parcel {
	// Set the longitude and latitude of the pick up location
	// Example: "40.416775,-3.703790"
	pickupGeoCoord := HttpGeoHereRequest(parcel.PickupLocation.Address)
	parcel.PickupLocation.Longitude = pickupGeoCoord.Longitude
	parcel.PickupLocation.Latitude = pickupGeoCoord.Latitude
	parcel.PickupLocation.Address = pickupGeoCoord.Address
	parcel.PickupLocation.City = pickupGeoCoord.City
	parcel.PickupLocation.ZipCode = pickupGeoCoord.ZipCode

	// Set the longitude and latitude of the drop off location
	// Example: "40.416775,-3.703790"
	dropOffGeoCoord := HttpGeoHereRequest(parcel.DropOffDestination.Address)
	parcel.DropOffDestination.Longitude = dropOffGeoCoord.Longitude
	parcel.DropOffDestination.Latitude = dropOffGeoCoord.Latitude
	parcel.DropOffDestination.Address = dropOffGeoCoord.Address
	parcel.DropOffDestination.City = dropOffGeoCoord.City
	parcel.DropOffDestination.ZipCode = dropOffGeoCoord.ZipCode

	return parcel
}

// Create creates a new parcel and performs any necessary validations or business logic checks
func (parcel *Parcel) Save() (*Parcel, []error) {
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

	//EU standard palet dimensions
	if parcel.IsEuPalet {
		parcel.Dimensions.Length = 120
		parcel.Dimensions.Width = 80
	}

	// If all validations pass, you can proceed with creating the parcel
	// Add your code here to save the parcel in the database or perform any other necessary operations

	return parcel, errorslist
}
