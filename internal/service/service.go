package service

import (
	"context"
	"github.com/magmaheat/music-info/internal/entity"
	"github.com/magmaheat/music-info/internal/repo"
)

type MusicLibrary interface {
	GetInfoLibrary(ctx context.Context, input entity.InfoLibrary) ([]entity.Song, error)
	CreateSong()
	GetSong()
	UpdateSong()
}

type Services struct {
	MusicLibrary MusicLibrary
}

func NewServices(repo *repo.Repositories) *Services {
	return &Services{MusicLibrary: NewMusicLibrary(repo.Song)}
}
