/*
 */
package weblib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var plainContentType = []string{"text/plain; charset=utf-8"}
var htmlContentType = []string{"text/html; charset=utf-8"}
var jsonContentType = []string{"application/json; charset=utf-8"}

func writeContentType(wr http.ResponseWriter, value []string) {
	header := wr.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
	header["Content-Type"] = value
}

func WriteJson(wr http.ResponseWriter, model interface{}) error {
	wr.WriteHeader(http.StatusOK)
	return writeJson(wr, model)
}

func writeJson(wr http.ResponseWriter, model interface{}) error {
	writeContentType(wr, jsonContentType)
	return json.NewEncoder(wr).Encode(model)
}

func WriteString(wr http.ResponseWriter, format string, data ...interface{}) {
	writeContentType(wr, plainContentType)
	if len(data) > 0 {
		fmt.Fprintf(wr, format, data...)
	} else {
		io.WriteString(wr, format)
	}
}

func WriteResponse(wr http.ResponseWriter, data interface{}, err error) {
	resp := NewResp(err)
	resp.Data = data
	err1 := WriteJson(wr, resp)
	if err1 != nil {
		panic(err1)
	}
}

// type JadeRenderer struct {
// 	jadeEngine *gojade.Engine
// }
//
// func (this *JadeRenderer) Render(c *Context, view string, model interface{}) {
// 	m := ViewModel{Model: model, User: c.User}
// 	m.Html = &Html{&m, c}
// 	writeContentType(c.ResponseWriter(), htmlContentType)
// 	this.jadeEngine.RenderFileW(c.Response(), view, m)
// }
//
// func NewJadeRender(viewpath string) *JadeRenderer {
// 	e := &JadeRenderer{gojade.New()}
// 	e.jadeEngine.ViewPath = viewpath
// 	return e
// }
