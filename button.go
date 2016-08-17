package main

import (
	"net/http"
)

type Press struct {
	Up bool `json:"up"`
}

var (
	okResp = []byte(`{"status":"success"}`)
)

func buttonHandler(w http.ResponseWriter, req *http.Request) error {
	if req.Method != http.MethodPost {
		return ErrBadRequestType
	}

	w.Write(okResp)
	return nil
}
