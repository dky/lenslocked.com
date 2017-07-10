package main

import (
	"fmt"
	//0
	//"html/template"
	"net/http"

	"lenslocked.com/views"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	//fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	/* 3
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
	*/
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	/*
		fmt.Fprint(w, "To get in touch, please send an email "+
			"to <a href=\"mailto:support@lenslocked.com\">"+
			"support@lenslocked.com</a>.")
	*/
	/* 3
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
	*/
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
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

/* 1
var homeTemplate *template.Template
var contactTemplate *template.Template
*/

var homeView *views.View
var contactView *views.View

func main() {
	/*
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")
	*/
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
/* 2
	var err error
	homeTemplate, err = template.ParseFiles(
		"views/home.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}

	contactTemplate, err = template.ParseFiles(
		"views/contact.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}
*/

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
