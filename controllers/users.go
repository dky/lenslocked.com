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
	Name    string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type LoginForm struct {
	Email string `schema:"email"`
	Password string `schema:"password"`
}

//func NewUsers() *Users {
func NewUsers(us models.UserService) *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		UserService: us,
	}
}

type Users struct {
	NewView *views.View
	LoginView *views.View
	models.UserService
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

//Post signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	/*
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Password is", form.Password)
	*/
	form := SignupForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := &models.User{
		Name: form.Name,
		Email: form.Email,
		Password: form.Password,
	}

	if err := u.UserService.Create(user); err != nil {
		panic(err)
	}
	fmt.Println(w, user)
}

func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := u.UserService.Authenticate(form.Email, form.Password)
	if user != nil {
		cookie := &http.Cookie {
			Name: "email",
			Value: form.Email,
		}
		http.SetCookie(w, cookie)
	} else {
		//fmt.Fprintln(w, user)
		fmt.Fprintln(w, "Invalid login credentials.")
	}
}

func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email")
	if err != nil {
		fmt.Fprintln(w, "Error retreiving the cookie:", err)
	} else {
		fmt.Fprintln(w, "Email is:", cookie.Value)
	}
}
