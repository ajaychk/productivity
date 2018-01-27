package controllers

import (
	r "github.com/revel/revel"
)

// App is main controller
type App struct {
	*r.Controller
}

// Index - start page of site
func (c App) Index() r.Result {
	return c.Render()
}

// AddTask - add a task
func (c App) AddTask() r.Result {
	return c.Render()
}

// UpdateTask - update a task
func (c App) UpdateTask() r.Result {
	return c.Render()
}

// DeleteTask - delete a task
func (c App) DeleteTask() r.Result {
	return c.Render()
}
