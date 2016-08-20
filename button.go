package main

import (
	"io"
	"net/http"
	"time"
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

	processor.Input(&Message{
		Client: &Client{
			IP:        req.RemoteAddr,
			UserAgent: req.UserAgent(),
		},
		Body:     b[0:n],
		Received: time.Now(),
	})
	w.Write(okResp)
	return nil
}
