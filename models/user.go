package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	Firstname string `json:first_name`
	Lastname  string `json:last_name`
	Email     string `json:email`
	//we use byte because we need to encrypt this, we put . because we don't what passowrd to be returned
	Password []byte `json:"-"`
	Phone    string `json:"phone"`
}

func (user *User) SetPassword(password string) {
	hashedpass, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedpass
}

func (user *User) ComparePass(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
