package main

import (
	"net/http"
)

type Press struct {
	Up bool `json:"up"`
}

func buttonHandler(w http.ResponseWriter, req *http.Request) error {
	if req.Method != http.MethodPost {
		return ErrBadRequestType
	}

	return nil
}
