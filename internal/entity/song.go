package entity

import "errors"

type Song struct {
	Name        string  `json:"song"`
	Group       string  `json:"group"`
	Text        string  `json:"text"`
	Genre       string  `json:"genre"`
	ReleaseYear int     `json:"release_year"`
	Duration    float32 `json:"duration"`
}

type InfoLibrary struct {
	StartReleaseYear int     `json:"start_release_year"`
	EndReleaseYear   int     `json:"end_release_year"`
	Genre            string  `json:"genre"`
	StartDuration    float32 `json:"start_duration"`
	EndDuration      float32 `json:"end_duration"`
	Offset           int     `json:"offset"`
	Limit            int     `json:"limit"`
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
