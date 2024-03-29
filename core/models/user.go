package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User : user struct
type User struct {
	Model
	Email       string `gorm:"unique;not null"`
	Password    string
	FirstName   string
	LastName    string
	CompanyId   string
	Logo        string
	Description string
}

// NewUser : create new user
func NewUser(email, firstname, secondname, companyid, password, description string) User {
	user := new(User)
	user.Email = email
	user.Password = HashAndSaltPassword(password)
	user.FirstName = firstname
	user.LastName = secondname
	user.CompanyId = companyid
	user.Description = description

	return *user
}

//TODO REMOVE OU REVOIR
func HashAndSaltPassword(password string) (hashedPassword string) {

	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	hashedPassword = string(hashedPasswordByte)

	fmt.Println(hashedPassword)
	return
}

func CompareHashAndPassword(hashedPassword, password string) (result bool) {
	result = false

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	fmt.Println(err) // nil means it is a match

	if err == nil {
		result = true
	}

	return
}
