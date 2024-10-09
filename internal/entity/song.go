package entity

import (
	"errors"
	"time"
)

type Song struct {
	Id          string    `json:"id"`
	Name        string    `json:"song"`
	Group       string    `json:"group"`
	Text        string    `json:"text"`
	Genre       string    `json:"genre"`
	ReleaseData time.Time `json:"release_data"`
	Duration    float32   `json:"duration"`
	Link        string    `json:"link"`
}

type SongDetail struct {
	Id          string    `json:"id"`
	ReleaseDate time.Time `json:"release_date"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}

type InfoLibrary struct {
	StartReleaseData time.Time `json:"start_release_year"`
	EndReleaseYear   time.Time `json:"end_release_year"`
	Genre            string    `json:"genre"`
	StartDuration    float32   `json:"start_duration"`
	EndDuration      float32   `json:"end_duration"`
	Offset           int       `json:"offset"`
	Limit            int       `json:"limit"`
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
