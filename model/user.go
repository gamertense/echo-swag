package model

import (
	"errors"
	"fmt"
)

// Error example
var (
	ErrNameInvalid = errors.New("name is empty")
)

// User example
type User struct {
	ID    int    `json:"id" example:"1" format:"int64"`
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

// AddUser example
type AddUser struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

// UpdateUser example
type UpdateUser struct {
	Name  string `json:"name" example:"user name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

// Validation example
func (a AddUser) Validation() error {
	switch {
	case len(a.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// UserOne example
func UserOne(id int) (User, error) {
	for _, v := range users {
		if id == v.ID {
			return v, nil
		}
	}
	return User{}, nil
}

// UsersAll example
func UsersAll() []User {
	return users
}

// Insert example
func (u User) Insert() (int, error) {
	userMaxID++
	u.ID = userMaxID
	users = append(users, u)
	return userMaxID, nil
}

// Delete example
func Delete(id int) error {
	for k, v := range users {
		if id == v.ID {
			users = append(users[:k], users[k+1:]...)
			return nil
		}
	}
	return fmt.Errorf("account id=%d is not found", id)
}

// Update example
func (u User) Update() error {
	for k, v := range users {
		if u.ID == v.ID {
			users[k].Name = u.Name
			return nil
		}
	}
	return fmt.Errorf("account id=%d is not found", u.ID)
}

var userMaxID = 3
var users = []User{
	{ID: 1, Name: "user_1", Email: "Test"},
	{ID: 2, Name: "user_2", Email: "Test"},
	{ID: 3, Name: "user_3", Email: "Test"},
}
