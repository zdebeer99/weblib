package weblib

import (
	"net/http"
	"net/url"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	RequestContext map[string]interface{}
}

func New(req http.ResponseWriter, res *http.Request) *Context {
	return NewContext(req, res)
}

func NewContext(req http.ResponseWriter, res *http.Request) *Context {
	return &Context{req, res, make(map[string]interface{})}
}

func (this *Context) HttpError(errormessage string, code int) {
	this.ResponseWriter.WriteHeader(code)
	WriteString(this.ResponseWriter, errormessage)
}

//Return a String to the client.
func (this *Context) ViewString(format string, data ...interface{}) {
	this.ResponseWriter.WriteHeader(http.StatusOK)
	WriteString(this.ResponseWriter, format, data)
}

// View Render a template to html.
// By default gojade rendering engine is used, this can be customized.
// func (this *Context) View(view string, model interface{}) {
// 	this.ResponseWriter.WriteHeader(http.StatusOK)
// 	this.app.RenderEngine.Render(this, view, model)
// }

// File return a file from a path
func (this *Context) File(filePath string) {
	http.ServeFile(this.ResponseWriter, this.Request, filePath)
}

// View Render a template to html.
// By default gojade rendering engine is used, this can be customized.
func (this *Context) Json(model interface{}) {
	if err := WriteJson(this.ResponseWriter, model); err != nil {
		panic(err)
	}
}

func (this *Context) Redirect(path string) {
	http.Redirect(this.ResponseWriter, this.Request, path, http.StatusSeeOther) //303 Redirect
	//http.Redirect(this.Response(), this.Request(), path, http.StatusTemporaryRedirect)//307 Redirect
}

// BindForms binds a go structure to a html form
// Uses gorilla.schema
func (this *Context) BindForm(model interface{}) {
	BindForm(this.Request, model)
}

func (this *Context) BindJson(model interface{}) {
	BindJson(this.Request, model)
}

func (this *Context) FormValues() url.Values {
	this.Request.ParseForm()
	return this.Request.Form
}

func (this *Context) PostFormValues() url.Values {
	this.Request.ParseForm()
	return this.Request.PostForm
}
