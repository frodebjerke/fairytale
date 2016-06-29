package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/frodebjerke/fairytale/storyteller"
)

type reply struct {
	Data   storyteller.Story `json:"data"`
	Status string            `json:"status"`
}

// ReceiveDataHandler accepts data
func ReceiveDataHandler(stories storyteller.Stories) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var payload storyteller.Story
		err := decoder.Decode(&payload)
		if err != nil {
			panic("helvette")
		}
		r.Body.Close()

		stories.Add(payload)

		response := reply{payload, "ok"}
		json.NewEncoder(w).Encode(response)
	}
}
