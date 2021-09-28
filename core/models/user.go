package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User : user struct
type User struct {
	gorm.Model
	Name      string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string
	FirstName string
	LastName  string
	CompanyId string
}

// NewUser : create new user
func NewUser(name, email, firstname, secondname, companyid, password string) User {
	user := new(User)
	user.Name = name
	user.Email = email
	user.Password = HashAndSaltPassword(password)
	user.FirstName = firstname
	user.LastName = secondname
	user.CompanyId = companyid

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
