package models

import (
	"errors"
	"fmt"

	bolow "github.com/boolow5/bolow/crypto"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int      `json:"-"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Profile  *Profile `json:"profile" orm:"rel(one)"`
	IsAdmin  bool     `json:"is_admin"`
}

func (this *User) TableName() string {
	return "user"
}

func (this *User) HashPassword() (err error) {
	if bolow.IsBcryptHash(this.Password) {
		return err
	}
	this.Password, err = bolow.HashPassword(this.Password)
	if err != nil {
		return err
	}
	return nil
}

func (this *User) Add() (bool, error) {
	err := this.HashPassword()
	if err != nil {
		return false, err
	}
	o.Begin()
	profile := Profile{}
	profile.FirstName = this.Username
	this.Profile = &profile
	counter, err := o.Insert(this.Profile)
	if err != nil || counter < 1 {
		o.Rollback()
		return false, err
	}
	counter, err = o.Insert(this)
	if err != nil || counter < 1 {
		o.Rollback()
		return false, err
	}
	o.Commit()
	return true, nil
}

func (this *User) Update() (bool, error) {
	if this.ID < 1 {
		return false, errors.New("Invalid user id")
	}
	err := this.HashPassword()
	if err != nil {
		return false, err
	}
	counter, err := o.Update(this)
	if err != nil || counter < 1 {
		return false, err
	}
	return true, nil
}

func (this *User) Delete() (bool, error) {
	if this.ID < 1 {
		return false, errors.New("Invalid user id")
	}
	counter, err := o.Delete(this)
	if err != nil || counter < 1 {
		return false, err
	}
	return true, nil
}

func (this *User) IsValid() bool {
	if this.Username == "" || this.Password == "" {
		return false
	}
	return true
}

func (this *User) Exists() bool {
	counter, err := o.QueryTable(this).Filter("Username__exact", this.Username).Count()
	if err != nil {
		return false
	}
	if counter < 1 {
		return false
	}
	return true
}

// Authentication

func (this *User) Authenticate() (map[string]interface{}, error) {
	fmt.Println("Authenticating user")
	if len(this.Username) < 1 && len(this.Password) < 1 {
		return nil, errors.New("username and password are required")
	}
	user := User{}
	o.QueryTable(&user).Filter("Username__exact", this.Username).One(&user)
	if user.ID < 1 {
		return nil, errors.New("incorrect username or password")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(this.Password))
	if err != nil {
		return nil, errors.New("Either username or password is incorrect")
	}

	o.QueryTable(user.Profile).Filter("User", user.ID).One(user.Profile)
	// authenticated
	user_data := map[string]interface{}{
		"user_id":       user.ID,
		"user_username": user.Username,
		"user_profile":  user.Profile,
	}
	return user_data, nil
}
