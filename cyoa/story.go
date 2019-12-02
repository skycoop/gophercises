package cyoa

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Arc struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type ArcMap map[string]Arc

type Story struct {
	Arcs ArcMap
}

func (s Story) GetArc(name string) (Arc, error) {
	arc, exists := s.Arcs[name]
	if !exists {
		return Arc{}, fmt.Errorf("arc with name \"%s\" does not exist", name)
	}
	return arc, nil
}

func LoadStory(filename string) (*Story, error) {
	log.WithField("filename", filename).Debug("Loading story file")
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open story file: %w", err)
	}

	var arcs ArcMap
	err = json.NewDecoder(f).Decode(&arcs)
	if err != nil {
		return nil, fmt.Errorf("failed to load story from json: %w", err)
	}
	return &Story{Arcs: arcs}, nil
}
