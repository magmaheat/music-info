package pgdb

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/magmaheat/music-info/internal/entity"
	"github.com/magmaheat/music-info/pkg/postgres"
	log "github.com/sirupsen/logrus"
)

type SongRepo struct {
	*postgres.Postgres
}

func NewSongRepo(pg *postgres.Postgres) *SongRepo {
	return &SongRepo{
		Postgres: pg,
	}
}

func (s *SongRepo) GetInfoLibrary(ctx context.Context, input entity.InfoLibrary) ([]entity.Song, error) {
	const fn = "pgdb - music_library - GetInfoLibrary"

	queryBuilder := s.Builder.Select("song_name, group_name, text_song, genre, release_year, duration").From("music")

	if input.StartReleaseYear > 0 {
		queryBuilder = queryBuilder.Where(squirrel.GtOrEq{"release_year": input.StartReleaseYear})
	}

	if input.EndReleaseYear > 0 {
		queryBuilder = queryBuilder.Where(squirrel.LtOrEq{"release_year": input.EndReleaseYear})
	}

	if input.Genre != "" {
		queryBuilder = queryBuilder.Where(squirrel.Eq{"genre": input.Genre})
	}

	if input.StartDuration > 0 {
		queryBuilder = queryBuilder.Where(squirrel.GtOrEq{"duration": input.StartDuration})
	}

	if input.EndDuration > 0 {
		queryBuilder = queryBuilder.Where(squirrel.LtOrEq{"duration": input.EndDuration})
	}

	if input.Offset > 0 {
		queryBuilder = queryBuilder.Offset(uint64(input.Offset))
	}

	if input.Limit == 0 {
		input.Limit = 10
	}

	queryBuilder = queryBuilder.Limit(uint64(input.Limit))

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		log.Errorf("%s - queryBuilder.ToSql: %v", fn, err)
		return nil, err
	}

	rows, err := s.Pool.Query(ctx, query, args...)
	if err != nil {
		log.Errorf("%s - s.Pool.Query: %v", fn, err)
		return nil, err
	}

	var songs []entity.Song
	for rows.Next() {
		var song entity.Song
		if err := rows.Scan(&song.Name, &song.Group, &song.Text, &song.Genre, &song.ReleaseYear, &song.Duration); err != nil {
			log.Errorf("%s - rows.Scan: %v", fn, err)
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}
