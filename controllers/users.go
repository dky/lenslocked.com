package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/models"
	"lenslocked.com/views"
)

/*
type Users struct {
	NewView *views.View
}
*/

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//func NewUsers() *Users {
func NewUsers(us models.UserService) *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		UserService: us,
	}
}

type Users struct {
	NewView *views.View
	models.UserService
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

//Post signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Password is", form.Password)
}
