package main

import (
	"net/http"
)

var (
	ErrBadRequestType = JSONError{
		msg:    "Invalid HTTP request type",
		status: http.StatusMethodNotAllowed,
	}
)

type JSONError struct {
	msg    string
	status int
}

func (j JSONError) Error() string {
	return j.msg
}

func (j JSONError) Status() int {
	return j.status
}
