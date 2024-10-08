package service

import (
	"context"
	"github.com/magmaheat/music-info/internal/entity"
	"github.com/magmaheat/music-info/internal/repo"
)

type MusicLibraryService struct {
	library repo.SongRepo
}

func NewMusicLibrary(repo repo.SongRepo) *MusicLibraryService {
	return &MusicLibraryService{library: repo}
}

func (m *MusicLibraryService) GetInfoLibrary(ctx context.Context, input entity.InfoLibrary) ([]entity.Song, error) {
	songs, err := m.library.GetInfoLibrary(ctx, input)
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func (m *MusicLibraryService) CreateSong() {}

func (m *MusicLibraryService) GetSong() {}

func (m *MusicLibraryService) UpdateSong() {}
