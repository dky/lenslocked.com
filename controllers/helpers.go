package controllers 

import (
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		//panic(err)
		return err
	}

	dec := schema.NewDecoder()
	//form := SignupForm{}

	//if err := dec.Decode(&form, r.PostForm); err != nil {
	if err := dec.Decode(dst, r.PostForm); err != nil {
		//panic(err)
		return err
	}
	return nil
}
