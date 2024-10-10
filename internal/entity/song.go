package entity

import (
	"errors"
	"fmt"
	"strings"
)

type Song struct {
	Id          string  `json:"id"`
	Name        string  `json:"song"`
	Group       string  `json:"group"`
	Text        string  `json:"text"`
	Genre       string  `json:"genre"`
	ReleaseDate string  `json:"release_date"`
	Duration    float32 `json:"duration"`
	Link        string  `json:"link"`
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

type SongDetail struct {
	Id          string `json:"id"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func (s *SongDetail) FormatText(offset, limit int) {
	if s.Text == "" || offset < 1 || limit < 0 {
		return
	}

	verseList := strings.Split(s.Text, "\n\n")

	result := ""

	for idx, item := range verseList {
		if offset == idx+1 || offset+limit > idx+1 {
			if result == "" {
				result += fmt.Sprint(item, '\n')
			} else {
				result += fmt.Sprint('\n', item, '\n')
			}
		}
	}

	s.Text = result
}

type InfoLibrary struct {
	StartReleaseData string  `json:"start_release_date"`
	EndReleaseYear   string  `json:"end_release_date"`
	Genre            string  `json:"genre"`
	StartDuration    float32 `json:"start_duration"`
	EndDuration      float32 `json:"end_duration"`
	Offset           int     `json:"offset"`
	Limit            int     `json:"limit"`
}
