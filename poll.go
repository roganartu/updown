package main

import (
	"encoding/json"
	"math/big"
	"net/http"
)

type Status struct {
	Count *big.Int `json:"count"`
}

func pollHandler(w http.ResponseWriter, req *http.Request) error {
	req.Body.Close()

	if req.Method != http.MethodGet {
		return ErrBadRequestType
	}

	resp := &Status{
		Count: count,
	}

	return json.NewEncoder(w).Encode(resp)
}
