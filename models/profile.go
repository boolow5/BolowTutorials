package models

import (
	"errors"
	"fmt"
)

type Profile struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	ImageURL   string `json:"image_url"`
	User       *User  `json:"user" orm:"reverse(one)"`
	Status     bool   `json:"status"`
}

func (this *Profile) TableName() string {
	return "profile"
}

func (this *Profile) FullName() string {
	return fmt.Sprintf("%s %s %s", this.FirstName, this.MiddleName, this.LastName)
}

func (this *Profile) Add() (bool, error) {
	counter, err := o.Insert(this)
	if err != nil {
		return false, err
	}
	if counter < 1 {
		return false, nil
	}
	return true, nil
}

func (this *Profile) Update() (bool, error) {
	if this.ID < 1 {
		return false, errors.New("Invalid Profile id")
	}
	counter, err := o.Update(this)
	if err != nil || counter < 1 {
		return false, err
	}
	return true, nil
}

func (this *Profile) Delete() (bool, error) {
	if this.ID < 1 {
		return false, errors.New("Invalid Profile id")
	}
	counter, err := o.Delete(this)
	if err != nil || counter < 1 {
		return false, err
	}
	return true, nil
}
