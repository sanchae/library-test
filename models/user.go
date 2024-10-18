package models

import (
	"fmt"
	"net/http"
)

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type UserList struct {
	Users []User `json:"users"`
}

func(i *User) Bind(r *http.Request) error {
	if i.FirstName == "" {
		return fmt.Errorf("firstname is a required field")
	}
	if i.LastName == "" {
		return fmt.Errorf("lastname is a required field")
	}
	return nil
}

func (*UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}