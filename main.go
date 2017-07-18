package main

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
	"github.com/gorilla/mux"
)

var (
	homeView *views.View
	contactView *views.View
	faqsView *views.View
	signupView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// call the contactView object render function and pass w to it.
	must(contactView.Render(w, nil))
}

func faqs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqsView.Render(w, nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signupView.Render(w, nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We could not find the page you "+
		"were looking for :(</h1>"+
		"<p>Please email us if you keep being sent to an "+
		"invalid page.</p>")
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err := nil {
		panic(err)
	}
}


func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqsView = views.NewView("bootstrap_improved", "views/faqs.gohtml")
	signupView = views.NewView("bootstrap", "views/signup.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faqs", faqs)
	r.HandleFunc("/signup", signup)
	http.ListenAndServe(":3000", r)
}
