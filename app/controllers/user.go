package controllers

import (
	r "github.com/revel/revel"
)

// User represents a user of this website
type User struct {
	App
}

// Register displays a register form to register a new user
func (u User) Register() r.Result {
	return u.Render()
}

// SaveNewUser saves a new user credential's
func (u User) SaveNewUser() r.Result {
	return u.Render()
}

// Login checks user's credential and on success user is logged in
func (u User) Login() r.Result {
	return u.Render()
}

// Logout logs a user out of the server
func (u User) Logout() r.Result {
	return u.Render()
}
