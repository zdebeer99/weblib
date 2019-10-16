/*
Package api provides standard objects used for communication between server and client side.
*/
package weblib

import (
	"fmt"
)

const (
	// OK response status. all is good!
	OK = 0
	// ERROR response status. something went wrong and could not complete task.
	ERROR = 1
	// WARNING response status. its okay but remember for next time!
	WARNING = 2
)

// Response data model
type Response struct {
	//StatusCode
	StatusCode byte `json:"statusCode"`
	//StatusText
	StatusText string `json:"statusText"`
	//Data
	Data interface{} `json:"data"`
}

func NewResp(err error) *Response {
	resp := &Response{}
	resp.Error(err)
	return resp
}

func (this *Response) Ok() *Response {
	this.StatusCode = OK
	this.StatusText = ""
	return this
}

func (this *Response) ErrorStr(msg string) *Response {
	this.StatusCode = ERROR
	this.StatusText = msg
	return this
}

func (this *Response) ErrorStrf(format string, a ...interface{}) *Response {
	this.StatusCode = ERROR
	this.StatusText = fmt.Sprintf(format, a...)
	return this
}

func (this *Response) Error(err error) *Response {
	if err != nil && this.StatusCode != ERROR {
		this.StatusCode = ERROR
		this.StatusText = err.Error()
	}
	return this
}

func (this *Response) IsOk() bool {
	return this.StatusCode != ERROR
}

func (this *Response) IsError() bool {
	return this.StatusCode == ERROR
}
