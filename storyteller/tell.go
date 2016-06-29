package storyteller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 5,
		},
	}

	return client
}

// Tell another node
func Tell(story Story) error {

	var (
		host = "http://localhost:8080"
		path = "/verses"
	)

	httpClient := createHTTPClient()

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(story)

	req, err := http.NewRequest("POST", host+path, b)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, err = httpClient.Do(req)
	if err != nil {
		return err
	}
	req.Body.Close()
	req.Close = true
	log.Println("Sent that something!")

	return nil
}
