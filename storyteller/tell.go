package storyteller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Storyteller ...
type Storyteller struct {
	Client *http.Client
}

// NewStoryteller ...
func NewStoryteller() Storyteller {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 5,
		},
	}

	return Storyteller{client}
}

// Tell another node
func (st Storyteller) Tell(story Story) error {

	var (
		host = "http://localhost:8080"
		path = "/verses"
	)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(story)

	req, err := http.NewRequest("POST", host+path, b)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, err = st.Client.Do(req)
	if err != nil {
		return err
	}
	req.Body.Close()
	req.Close = true
	log.Println("Sent that something!")

	return nil
}
