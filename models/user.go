package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `gorm:"not null unique"`
	Username string `gorm:"primary_key unique"`
	Password string `gorm:"not null"`
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password, err = hashPassword(u.Password)
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
