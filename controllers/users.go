package controllers

import "lenslocked.com/views"

func NewUsers() *Users {
	return &Users {
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}
