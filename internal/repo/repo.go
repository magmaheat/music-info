package repo

import (
	"context"
	"github.com/magmaheat/music-info/internal/entity"
	"github.com/magmaheat/music-info/internal/repo/pgdb"
	"github.com/magmaheat/music-info/pkg/postgres"
)

type SongRepo interface {
	GetInfoLibrary(ctx context.Context, input entity.InfoLibrary) ([]entity.Song, error)
}

type Repositories struct {
	Song SongRepo
}

func New(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		Song: pgdb.NewSongRepo(pg),
	}
}
