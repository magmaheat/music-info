package repo

import (
	"github.com/magmaheat/music-info/internal/repo/pgdb"
	"github.com/magmaheat/music-info/pkg/postgres"
)

type SongRepo interface {
}

type Repositories struct {
	Song SongRepo
}

func New(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		Song: pgdb.NewSongRepo(pg),
	}
}
