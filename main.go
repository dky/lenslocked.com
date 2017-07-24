package main

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
	"lenslocked.com/controllers"
	"github.com/gorilla/mux"
)

var (
	HomeView *views.View
	ContactView *views.View
	FaqsView *views.View
	//signupView *views.View
)

/*
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
*/

/*
func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signupView.Render(w, nil))
}
*/

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

func main() {
	/*
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqsView = views.NewView("bootstrap_improved", "views/faqs.gohtml")
	//signupView = views.NewView("bootstrap", "views/signup.gohtml")
	r := mux.NewRouter()
	/*
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/faqs", faqs).Methods("GET")
	*/
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	/*
	r.HandleFunc("/signup", signup)
	*/
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}
