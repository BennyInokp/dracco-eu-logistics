package packages

import (
	"fmt"
	"time"
)

type Package struct {
	ID               int       `json:"id"`               // Unique identifier for the Package
	DistributorID    int       `json:"distributorID"`    // ID of the distributor associated with the Package
	Origin           string    `json:"origin"`           // Origin of the Package
	Destination      string    `json:"destination"`      // Destination of the Package
	Type             string    `json:"type"`             // Type of the Package
	EstimatedValue   float64   `json:"estimatedValue"`   // Estimated value of the Package
	OfferPrice       float64   `json:"offerPrice"`       // Offered price for transporting the Package
	PublicationDate  time.Time `json:"publicationDate"`  // Date when the Package is published
	ModificationDate time.Time `json:"modificationDate"` // Date when the Package is modified
	DeadlineDate     time.Time `json:"deadlineDate"`     // Deadline for transporting the Package
	Dimensions       string    `json:"dimensions"`       // Dimensions (length, width, height) of the Package
	IsEuPalet        bool      `json:"IsEuPalet"`        // Flag indicating whether the Package is an EU pallet
	IsAccepted       bool      `json:"isAccepted"`       // Flag indicating whether the Package has been accepted
	IsCompleted      bool      `json:"isCompleted"`      // Flag indicating whether the Package has been completed
}

// NewPackage creates a new Package
func NewPackage(DistributorID int) *Package {
	return &Package{
		DistributorID: DistributorID,
	}
}

// SetOrigin sets the origin of the Package
func (p *Package) SetOrigin(origin string) *Package {
	p.Origin = origin
	return p
}

// SetDestination sets the destination of the Package
func (p *Package) SetDestination(destination string) *Package {
	p.Destination = destination
	return p
}

// SetType sets the type of the Package
func (p *Package) SetType(packageType string) *Package {
	p.Type = packageType
	return p
}

// SetEstimatedValue sets the estimated value of the Package
func (p *Package) SetEstimatedValue(value float64) *Package {
	p.EstimatedValue = value
	return p
}

// SetPublicationDate sets the publication date for the Package
func (p *Package) SetPublicationDate(date time.Time) *Package {
	p.PublicationDate = date
	return p
}

// SetModificationDate sets the modification date for the Package
func (p *Package) SetModificationDate(date time.Time) *Package {
	p.ModificationDate = date
	return p
}

// SetDeadlineDate sets the deadline date for the Package
func (p *Package) SetDeadlineDate(date time.Time) *Package {
	p.DeadlineDate = date
	return p
}

// SetDimensions sets the dimensions of the Package
func (p *Package) SetDimensions(dimensions string) *Package {
	p.Dimensions += dimensions
	return p
}

// SetEuPalet sets the IsEuPalet flag to true
func (p *Package) SetEuPalet() *Package {
	p.IsEuPalet = true

	//EU standard palet dimensions
	if p.IsEuPalet {
		p.Dimensions = "120;80"
	}
	return p
}

// SetIsAccepted sets the IsAccepted flag to true
func (p *Package) SetIsAccepted() *Package {
	p.IsAccepted = true
	return p
}

// SetIsCompleted sets the IsCompleted flag to true
func (p *Package) SetIsCompleted() *Package {
	p.IsCompleted = true
	return p
}

// SetOfferPrice sets the offered price for transporting the Package
func (p *Package) SetOfferPrice(price float64) *Package {
	p.OfferPrice = price
	return p
}

func (p *Package) SetDimensionValues(length, width, height float64) *Package {
	p.Dimensions = fmt.Sprintf("%.2f x %.2f x %.2f", length, width, height)
	return p
}

// // SetGeo sets the longitude and latitude
// func (parcel *Parcel) SetGeo() *Parcel {
// 	// Set the longitude and latitude of the pick up location
// 	// Example: "40.416775,-3.703790"
// 	pickupGeoCoord := HttpGeoHereRequest(parcel.PickupLocation.Address)
// 	parcel.PickupLocation.Longitude = pickupGeoCoord.Longitude
// 	parcel.PickupLocation.Latitude = pickupGeoCoord.Latitude
// 	parcel.PickupLocation.Address = pickupGeoCoord.Address
// 	parcel.PickupLocation.City = pickupGeoCoord.City
// 	parcel.PickupLocation.ZipCode = pickupGeoCoord.ZipCode

// 	// Set the longitude and latitude of the drop off location
// 	// Example: "40.416775,-3.703790"
// 	dropOffGeoCoord := HttpGeoHereRequest(parcel.DropOffDestination.Address)
// 	parcel.DropOffDestination.Longitude = dropOffGeoCoord.Longitude
// 	parcel.DropOffDestination.Latitude = dropOffGeoCoord.Latitude
// 	parcel.DropOffDestination.Address = dropOffGeoCoord.Address
// 	parcel.DropOffDestination.City = dropOffGeoCoord.City
// 	parcel.DropOffDestination.ZipCode = dropOffGeoCoord.ZipCode

// 	return parcel
// }

// // Create creates a new parcel and performs any necessary validations or business logic checks
// func (parcel *Parcel) Save() (*Parcel, []error) {
// 	// Perform any necessary validations or business logic checks before creating the parcel
// 	// For example, checking if the required fields are provided, validating dimensions, etc.
// 	errorslist := []error{}

// 	// Example validation: Check if the description field is empty
// 	if parcel.Description == "" {
// 		parcel.Description = "No description provided"
// 	}

// 	// Example validation: Check if the weight is zero or negative
// 	if parcel.Weight <= 0 {
// 		errorslist = append(errorslist, errors.New("weight must be greater than zero"))
// 	}

// 	// Example validation: Check if the dimensions are valid
// 	if parcel.Dimensions.Length <= 0 || parcel.Dimensions.Width <= 0 || parcel.Dimensions.Height <= 0 {
// 		errorslist = append(errorslist, errors.New("invalid dimensions"))
// 	}

// 	// Example validation: Check if the pickup location is provided
// 	if parcel.PickupLocation.Address == "" || parcel.PickupLocation.City == "" || parcel.PickupLocation.ZipCode == "" {
// 		errorslist = append(errorslist, errors.New("pickup location is required"))
// 	}

// 	//EU standard palet dimensions
// 	if parcel.IsEuPalet {
// 		parcel.Dimensions.Length = 120
// 		parcel.Dimensions.Width = 80
// 	}

// 	// If all validations pass, you can proceed with creating the parcel
// 	// Add your code here to save the parcel in the database or perform any other necessary operations

// 	return parcel, errorslist
// }
