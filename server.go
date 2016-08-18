package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/roganartu/orbit"
	log "github.com/roganartu/trunk"
)

const (
	jsonErrString = `{"status":"error",message":"%s","status":%d}`
	BUFFER_SIZE   = 1000000
)

var (
	processor *orbit.Loop
	logger    log.Logger
)

type JSONHandler struct {
	h func(http.ResponseWriter, *http.Request) error
}

func init() {
	processor = orbit.New(
		BUFFER_SIZE,
		nil, // receiver
		nil, // journaler
		nil, // replicator
		nil, // unmarshaller
		nil, // executor
	)
	processor.Start()

	logger = log.New(os.Stdout, "", 8192)
}

func main() {
	http.Handle("/api/v1/button", JSONHandler{buttonHandler})
	http.ListenAndServe(":8080", nil)
}

func (h JSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	w.Header().Set("Content-Type", "application/json")
	status := 200

	err := h.h(w, r)
	if err != nil {
		switch e := err.(type) {
		case *JSONError:
			w.Write([]byte(fmt.Sprintf(jsonErrString, e.Error(), e.Status())))
			status = e.Status()
		default:
			w.Write([]byte(fmt.Sprintf(jsonErrString, e.Error(), 500)))
			status = 500
		}
	}
	elapsed := time.Now().Sub(start)

	path := r.URL.Path
	if r.URL.RawQuery != "" {
		path += "?" + r.URL.RawQuery
	}
	logger.Infof("%s %s %d - %d.%06d\n",
		r.Method,
		path,
		status,
		int64(elapsed.Seconds()),
		int64(time.Nanosecond)*elapsed.Nanoseconds()/int64(time.Microsecond))
}
