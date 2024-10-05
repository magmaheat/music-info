package entity

import "errors"

type Song struct {
	Name        string  `json:"song"`
	Group       string  `json:"group"`
	Text        string  `json:"text"`
	Genre       string  `json:"genre"`
	ReleaseYear string  `json:"release_year"`
	Duration    float32 `json:"duration"`
}

func (s *Song) Validate() error {
	if s.Group == "" {
		return errors.New("group is required")
	}

	if s.Name == "" {
		return errors.New("song is required")
	}

	return nil
}
