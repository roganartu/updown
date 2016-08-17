package main

import (
	"net/http"
)

var (
	ErrBadRequestType = JSONError{
		msg:    "Invalid HTTP request type",
		status: http.StatusMethodNotAllowed,
	}
	ErrMissingBody = JSONError{
		msg:    "Request missing body",
		status: http.StatusBadRequest,
	}
	ErrInvalidBody = JSONError{
		msg:    "Request body invalid",
		status: http.StatusBadRequest,
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
