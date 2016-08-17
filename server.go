package main

import (
	"fmt"
	"net/http"

	"github.com/roganartu/orbit"
)

const (
	jsonErrString = `{"status":"error",message":"%s","status":%d}`
	BUFFER_SIZE   = 1000000
)

var (
	processor *orbit.Loop
)

type JSONHandler struct {
	h func(http.ResponseWriter, *http.Request) error
}

func init() {
	processor = orbit.New(
		BUFFER_SIZE,
		nil, // receiver
		nil, // journaller
		nil, // replicator
		nil, // unmarshaller
		nil, // executor
	)
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
