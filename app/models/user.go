package models

import "time"

// User represents a user of productivity
type User struct {
	ID             string
	FirstName      string
	LastName       string
	Username       string
	Password       []byte
	HashedPassword []byte
	Timestamp      time.Time
}

// Register - registers a user
func (u *User) Register(fname, lname, uname, password, email string) error {
	return nil
}

// Login - logs a user into the site
func (u *User) Login(uname, password string) error {
	return nil
}

// Logout - logs a user out of site
func (u *User) Logout() error {
	return nil
}
