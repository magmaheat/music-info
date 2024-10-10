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

func (m *MusicLibraryService) GetSongDetail(ctx context.Context, song, group string, offset, limit int) (entity.SongDetail, error) {
	songDetail, err := m.library.GetSongDetail(ctx, song, group)
	if err != nil {
		return entity.SongDetail{}, err
	}

	songDetail.FormatText(offset, limit)

	return songDetail, nil
}

func (m *MusicLibraryService) UpdateSong(ctx context.Context, input entity.Song) (entity.Song, error) {
	song, err := m.library.UpdateSong(ctx, input)
	if err != nil {
		return entity.Song{}, err
	}

	return song, nil
}

func (m *MusicLibraryService) AddSong(ctx context.Context, input entity.Song) (string, error) {
	id, err := m.library.AddSong(ctx, input)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (m *MusicLibraryService) DeleteSong(ctx context.Context, id string) error {
	return m.DeleteSong(ctx, id)
}
