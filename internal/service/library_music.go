package service

import "github.com/magmaheat/music-info/internal/repo"

type LibraryMusicService struct {
	library repo.SongRepo
}

func NewLibraryMusic(repo repo.SongRepo) *LibraryMusicService {
	return &LibraryMusicService{library: repo}
}

func (l *LibraryMusicService) CreateSong() {}

func (l *LibraryMusicService) GetSong() {}

func (l *LibraryMusicService) UpdateSong() {}
