package models

type Transporter struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phoneNumber"`
	Vehicle     Vehicle `json:"vehicle"`
}

type Vehicle struct {
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	LicensePlate string `json:"licensePlate"`
	Maxcapacity  int64  `json:"maxcapacity"`
}
