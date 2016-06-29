package storyteller

import (
	"log"
	"time"
)

// Story the form of a story
type Story struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Seen  int    `json:"seen"`
}

// StoryCrypt ...
type StoryCrypt struct {
	Story Story
	Rcvd  time.Time
	Ready bool
	Tries int
}

// Stories the storyteller type
type Stories struct {
	Line chan StoryCrypt
}

// New creates a storyteller
func New() Stories {
	stories := Stories{make(chan StoryCrypt, 10)}

	go func() {
		for {
			select {
			case sc := <-stories.Line:
				if sc.Ready {
					log.Println(sc.Story.Key, "ready!")
					err := Tell(sc.Story)
					if err != nil {
						sc.Ready = false
						sc.Tries = sc.Tries + 1
						stories.Line <- sc
					}
				} else if sc.Rcvd.Add(1 * time.Second).Before(time.Now()) {
					log.Println(sc.Story.Key, "timeout complete")
					sc.Ready = true
					stories.Line <- sc
				} else {
					stories.Line <- sc
				}
			}
		}
	}()

	return stories
}

// Add add a story to the stories list
func (s Stories) Add(story Story) {
	s.Line <- StoryCrypt{story, time.Now(), false, 0}
}
