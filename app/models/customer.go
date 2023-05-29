package models

import (
	"errors"
	"fmt"

	"github.com/dracco-eu-logistics/app/utilities"
)

type Customer struct {
	CustomerID    int64  `json:"customerID"`
	CustomerName  string `json:"customerName"`
	CustomerEmail string `json:"customerEmail"`
	CustomerPhone string `json:"customerPhone"`

	Password      string `json:"password"`
	IsTransporter bool   `json:"isTransporter"`
	IsCompany     bool   `json:"isCompany"`
}

// NewCustomer creates a new customer
func NewCustomer() *Customer {
	return &Customer{}
}

// SetCustomerID sets the customer ID
func (customer *Customer) SetCustomerID(id int64) *Customer {
	customer.CustomerID = id
	return customer
}

func (customer *Customer) SetCustomerName(name string) *Customer {
	customer.CustomerName = name
	return customer
}

func (customer *Customer) SetCustomerEmail(email string) *Customer {
	customer.CustomerEmail = email
	return customer
}

func (customer *Customer) SetCustomerPhone(phone string) *Customer {
	customer.CustomerPhone = phone
	return customer
}

func (customer *Customer) SetCustomerIsTransporter(isTransporter bool) *Customer {
	customer.IsTransporter = isTransporter
	return customer
}

func (customer *Customer) SetCustomerIsCompany(isCompany bool) *Customer {
	customer.IsCompany = isCompany
	return customer
}

func (customer *Customer) SetCustomerPassword(password string) *Customer {
	customer.Password = password
	return customer
}

func (customer *Customer) Validate() (*Customer, []error) {
	errorslist := []error{}

	if customer.CustomerName == "" {
		errorslist = append(errorslist, errors.New("name is required"))
	}

	if customer.CustomerEmail == "" || !utilities.ValidateEmail(customer.CustomerEmail) {
		errorslist = append(errorslist, errors.New("email is required or not valid"))
	}

	if customer.CustomerPhone == "" {
		errorslist = append(errorslist, errors.New("phone is required"))
	}
	if customer.Password == "" {
		errorslist = append(errorslist, errors.New("password is required"))
	}

	if !customer.IsCompany && !customer.IsTransporter {
		errorslist = append(errorslist, errors.New("group requires (transporter or company)"))
	}

	if len(errorslist) > 0 {
		return nil, errorslist
	}

	fmt.Println("Customer validated correctly!")
	return customer, errorslist
}

func (Customer *Customer) Save() (int64, error) {
	return 0, nil
}
