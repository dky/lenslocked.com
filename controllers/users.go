package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

type Users struct {
	NewView *views.View
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
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "This is a temporary response. Check back later.")
	fmt.Fprintln(w, r.PostForm["email"])
	fmt.Fprintln(w, r.PostForm["password"])
}
