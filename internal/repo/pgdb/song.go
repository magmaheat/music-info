package pgdb

import "github.com/magmaheat/music-info/pkg/postgres"

type SongRepo struct {
	*postgres.Postgres
}

func NewSongRepo(pg *postgres.Postgres) *SongRepo {
	return nil
}
