package main

import (
	"io"
	"net/http"
)

type Press struct {
	Up bool `json:"up"`
}

var (
	okResp = []byte(`{"status":"success"}`)
)

func buttonHandler(w http.ResponseWriter, req *http.Request) error {
	defer req.Body.Close()

	if req.Method != http.MethodPost {
		return ErrBadRequestType
	}

	b := make([]byte, 20)
	limitReader := io.LimitReader(req.Body, 20)
	n, err := limitReader.Read(b)
	if err != nil && err != io.EOF {
		return JSONError{
			msg:    "Invalid request",
			status: http.StatusBadRequest,
		}
	}

	processor.Input(b[0:n])
	w.Write(okResp)
	return nil
}
