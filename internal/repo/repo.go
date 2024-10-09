package repo

import (
	"context"
	"github.com/magmaheat/music-info/internal/entity"
	"github.com/magmaheat/music-info/internal/repo/pgdb"
	"github.com/magmaheat/music-info/pkg/postgres"
)

type SongRepo interface {
	GetInfoLibrary(ctx context.Context, input entity.InfoLibrary) ([]entity.Song, error)
	GetSongDetail(ctx context.Context, song, group string) (entity.SongDetail, error)
	DeleteSong(ctx context.Context, id string) error
	UpdateSong(ctx context.Context, song entity.Song) (entity.Song, error)
	AddSong(ctx context.Context, song entity.Song) (string, error)
}

type Repositories struct {
	Song SongRepo
}

func New(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		Song: pgdb.NewSongRepo(pg),
	}
}
