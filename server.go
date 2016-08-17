package main

import (
	"fmt"
	"net/http"
)

const (
	jsonErrString = `{"status":"error",message":"%s","status":%d}`
)

type JSONHandler struct {
	h func(http.ResponseWriter, *http.Request) error
}

func init() {
	// set up the buffers
}

func main() {
	http.Handle("/api/v1/button", JSONHandler{buttonHandler})
	http.ListenAndServe(":8080", nil)
}

func (h JSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := h.h(w, r)
	if err != nil {
		switch e := err.(type) {
		case *JSONError:
			w.Write([]byte(fmt.Sprintf(jsonErrString, e.Error(), e.Status())))
		default:
			w.Write([]byte(fmt.Sprintf(jsonErrString, e.Error(), 500)))
		}
	}
}
