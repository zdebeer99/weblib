package weblib

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"net/http"
)

// BindForms binds a go structure to a html form
// Uses gorilla.schema
func BindForm(req *http.Request, model interface{}) {
	decoder := schema.NewDecoder()
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	err = decoder.Decode(model, req.PostForm)
	if err != nil {
		panic(err)
	}
}

func BindJson(req *http.Request, model interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&model)
	if err != nil {
		panic(err)
	}
}
