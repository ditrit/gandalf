package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	RoleID   uint
	Role     Role
}

func NewUser(name, email, password string, role Role) *User {
	user := new(User)
	user.Name = name
	user.Email = email
	user.Password = HashAndSaltPassword(password)
	user.Role = role

	return user
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
