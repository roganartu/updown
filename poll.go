package main

import (
	"encoding/json"
	"math/big"
	"net/http"
)

type Status struct {
	Value           *big.Int `json:"value"`
	TotalClickCount *big.Int `json:"total_click_count"`
}

func pollHandler(w http.ResponseWriter, req *http.Request) error {
	req.Body.Close()

	if req.Method != http.MethodGet {
		return ErrBadRequestType
	}

	resp := &Status{
		Value:           count,
		TotalClickCount: stats.TotalRequests,
	}

	return json.NewEncoder(w).Encode(resp)
}
