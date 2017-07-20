package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Email string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers() *Users {
	return &Users {
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

//Post signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, form)
	/*
	fmt.Fprintln(w, "This is a temporary response. Check back later.")
	fmt.Fprintln(w, r.PostForm["email"])
	fmt.Fprintln(w, r.PostForm["password"])
	*/
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Password is", form.Password)
}
