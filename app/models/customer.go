package models

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/dracco-eu-logistics/app/utilities"
)

const (
	DATABASE_CONNECTION_STRING = "root:nenufer555@tcp(localhost:3306)/dracco_eu"
)

type Documents struct {
	CustomerID          int    `json:"customerID"`
	DocumentType        string `json:"documentType"`
	DocumentDescription string `json:"documentDescription"`
	DocumentURL         string `json:"documentUrl"`
}

type Customer struct {
	CustomerID    int         `json:"customerID"`
	CustomerName  string      `json:"customerName"`
	CustomerEmail string      `json:"customerEmail"`
	CustomerPhone string      `json:"customerPhone"`
	TaxID         string      `json:"taxId"`
	Documents     []Documents `json:"documents"`
	IsVerified    bool        `json:"isVerified"`

	// kind of customer
	Password      string `json:"password"`
	IsTransporter bool   `json:"isTransporter"`
	IsWarehouse   bool   `json:"isWarehouse"`

	// only for transporters
	Transporter Transporter `json:"transporter"`
}

// NewCustomer creates a new customer
func NewCustomer() *Customer {
	return &Customer{}
}

// SetCustomerID sets the customer ID
func (customer *Customer) SetCustomerID() *Customer {
	customer.CustomerID = utilities.RandomNumber()
	return customer
}

func (customer *Customer) SetCustomerIsVerified() *Customer {
	customer.IsVerified = true
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

func (customer *Customer) SetCustomerIsTransporter() *Customer {
	customer.IsTransporter = true
	return customer
}

func (customer *Customer) SetCustomerIsWarehouse() *Customer {
	customer.IsWarehouse = true
	return customer
}

func (customer *Customer) SetCustomerPassword(password string) *Customer {
	customer.Password = password
	return customer
}

func (customer *Customer) SetTaxID(taxID string) *Customer {
	customer.TaxID = taxID
	return customer
}

func (customer *Customer) SetDocuments(documents []Documents) *Customer {
	customer.Documents = documents
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

	if !customer.IsWarehouse && !customer.IsTransporter {
		errorslist = append(errorslist, errors.New("group requires (transporter or company)"))
	}

	if len(errorslist) > 0 {
		return nil, errorslist
	}

	fmt.Println("Customer validated correctly!")
	return customer, errorslist
}

func (customer *Customer) Save() (*Customer, error) {
	// MySQL database connection information
	dsn := DATABASE_CONNECTION_STRING

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return customer, err
	}
	defer db.Close()

	// Check if the email already exists in the table
	emailExistsQuery := "SELECT COUNT(*) FROM customers WHERE customer_email = ?"
	var count int
	err = db.QueryRow(emailExistsQuery, customer.CustomerEmail).Scan(&count)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return customer, err
	}

	if count > 0 {
		return customer, errors.New("user with the same email already exists")
	}

	// Prepare the SQL statement
	stmt, err := db.Prepare("INSERT INTO customers (customer_id, customer_name, customer_email, customer_phone, tax_id, verified, password, is_transporter, is_warehouse) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Failed to prepare SQL statement:", err)
		return customer, err
	}
	defer stmt.Close()

	// Execute the prepared statement with the sample data
	_, err = stmt.Exec(customer.CustomerID, customer.CustomerName, customer.CustomerEmail, customer.CustomerPhone, "", false, customer.Password, customer.IsTransporter, customer.IsWarehouse)
	if err != nil {
		fmt.Println("Failed to insert data into the table:", err)
		return customer, err
	}

	fmt.Println("Data inserted successfully!")
	return customer, nil
}
