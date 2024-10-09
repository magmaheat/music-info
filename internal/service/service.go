package service

import (
	"context"
	"github.com/magmaheat/music-info/internal/entity"
	"github.com/magmaheat/music-info/internal/repo"
)

type MusicLibrary interface {
	GetInfoLibrary(ctx context.Context, input entity.InfoLibrary) ([]entity.Song, error)
	GetSongDetail(ctx context.Context, song, group string, offset, limit int) (entity.SongDetail, error)
	DeleteSong(ctx context.Context, id string) error
	UpdateSong(ctx context.Context, input entity.Song) (entity.Song, error)
	AddSong(ctx context.Context, input entity.Song) (string, error)
}

type Services struct {
	MusicLibrary MusicLibrary
}

func NewServices(repo *repo.Repositories) *Services {
	return &Services{MusicLibrary: NewMusicLibrary(repo.Song)}
}
