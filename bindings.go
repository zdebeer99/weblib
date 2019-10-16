package weblib

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
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

func BindJson(req *http.Request, model interface{}) error {
	decoder := json.NewDecoder(req.Body)
	return decoder.Decode(&model)
}
