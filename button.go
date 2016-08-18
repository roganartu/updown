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

	b := []byte(`{"up":false}`)
	n, err := readExpected(req.Body, b)
	if err != nil {
		return JSONError{
			msg:    err.Error(),
			status: http.StatusBadRequest,
		}
	}

	processor.Input(b[:n])
	w.Write(okResp)
	return nil
}

// readExpected reads the body up to a max of len(b) bytes.
//
// This is to prevent huge bodies being sent and processed.
func readExpected(body io.ReadCloser, b []byte) (int, error) {
	total := 0
	for {
		n, err := body.Read(b)
		total += n

		if total > len(b) {
			return 0, ErrInvalidBody
		}

		if err != nil {
			if err != io.EOF {
				return 0, ErrInvalidBody
			}

			if total == 0 {
				return 0, ErrMissingBody
			}

			if total < len(b)-1 {
				return 0, ErrInvalidBody
			}

			break
		}
	}
	return total, nil
}
