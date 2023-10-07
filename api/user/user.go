package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/dracco-eu-logistics/api/utilities"
)

type User struct {
	UserID           int         `json:"UserID"`
	UserName         string      `json:"UserName"`
	UserEmail        string      `json:"UserEmail"`
	UserPhone        string      `json:"UserPhone"`
	UserAddress      string      `json:"UserAddress"`
	UserType         string      `json:"UserType"`
	UserVatNumber    string      `json:"UserVatNumber"`
	IsVerified       bool        `json:"IsVerified"`
	ModificationDate time.Time   `json:"ModificationDate"`
	Password         string      `json:"Password"`
	Documents        []Documents `json:"Documents"`
}

type Documents struct {
	DocumentID          int    `json:"DocumentID"`
	UserID              int    `json:"UserID"`
	OrderID             int64  `json:"OrderID"`
	DocumentType        string `json:"DocumentType"`
	DocumentDescription string `json:"DocumentDescription"`
	DocumentName        string `json:"DocumentName"`
}

func NewUser() *User {
	return &User{
		UserID: SetUserID(),
	}
}

func SetUserID() int {
	return utilities.RandomNumber()
}

func (u *User) SetUserIsVerified() *User {
	u.IsVerified = true
	return u
}

func (u *User) SetUserName(name string) *User {
	u.UserName = name
	return u
}

func (u *User) SetUserEmail(email string) *User {
	u.UserEmail = email
	return u
}

func (u *User) SetUserPhone(phone string) *User {
	u.UserPhone = phone
	return u
}

func (u *User) SetUserAddress(address string) *User {
	u.UserAddress = address
	return u
}

func (u *User) SetUserType(userType string) *User {
	u.UserType = userType
	return u
}

func (u *User) SetUserVatNumber(vatNumber string) *User {
	u.UserVatNumber = vatNumber
	return u
}

func (u *User) SetUserPassword(password string) *User {
	u.Password = password
	return u
}

func (u *User) SetDocuments(documents []Documents) *User {
	u.Documents = documents
	return u
}

func (u *User) Validate() (*User, []error) {
	errorslist := []error{}

	if u.UserName == "" {
		errorslist = append(errorslist, errors.New("name is required"))
	}

	if u.UserEmail == "" || !utilities.ValidateEmail(u.UserEmail) {
		errorslist = append(errorslist, errors.New("email is required or not valid"))
	}

	if u.UserPhone == "" {
		errorslist = append(errorslist, errors.New("phone is required"))
	}
	if u.Password == "" {
		errorslist = append(errorslist, errors.New("password is required"))
	}

	if len(errorslist) > 0 {
		return nil, errorslist
	}

	fmt.Println("User validated correctly!")
	return u, errorslist
}
