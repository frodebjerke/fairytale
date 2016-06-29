package handlers

import (
	"encoding/json"
	"net/http"
)

type envelope struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type reply struct {
	Data   envelope `json:"data"`
	Status string   `json:"status"`
}

// ReceiveDataHandler accepts data
func ReceiveDataHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload envelope
	err := decoder.Decode(&payload)
	if err != nil {
		panic("helvette")
	}
	response := reply{payload, "ok"}
	json.NewEncoder(w).Encode(response)
}
