package pgdb

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/magmaheat/music-info/internal/entity"
	"github.com/magmaheat/music-info/pkg/postgres"
	log "github.com/sirupsen/logrus"
	"time"
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
	const fn = "pgdb - song - GetInfoLibrary"

	queryBuilder := s.Builder.Select("id, song_name, group_name, text_song, genre, release_date, duration").From("music")

	layout := "2006.02.01"
	defaultDate, _ := time.Parse(layout, "01.01.1970")

	if input.StartReleaseData != defaultDate {
		queryBuilder = queryBuilder.Where(squirrel.GtOrEq{"release_date": input.StartReleaseData})
	}

	if input.EndReleaseYear != defaultDate {
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
		if err := rows.Scan(&song.Id, &song.Name, &song.Group, &song.Text, &song.Genre, &song.ReleaseData, &song.Duration); err != nil {
			log.Errorf("%s - rows.Scan: %v", fn, err)
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (s *SongRepo) GetSongDetail(ctx context.Context, song, group string) (entity.SongDetail, error) {
	sql, args, _ := s.Builder.
		Select("id, release_date, text_song, link").
		From("music").
		Where("song_name = ? AND group_name = ?", song, group).
		ToSql()

	var songDetail entity.SongDetail
	err := s.Pool.QueryRow(ctx, sql, args...).Scan(
		&songDetail.Id,
		&songDetail.ReleaseDate,
		&songDetail.Text,
		&songDetail.Link,
	)

	if err != nil {
		log.Errorf("pgdb - song - s.Pool.QueryRow: %v", err)
		return entity.SongDetail{}, err
	}

	return songDetail, nil
}

func (s *SongRepo) DeleteSong(ctx context.Context, id string) error {
	sql, args, _ := s.Builder.Delete("music").Where("id = ?", id).ToSql()

	_, err := s.Pool.Query(ctx, sql, args...)
	if err != nil {
		log.Errorf("pgdb - song - DeleteSong - s.Pool.QueryRow: %v", err)
		return err
	}

	return nil
}

func (s *SongRepo) UpdateSong(ctx context.Context, song entity.Song) (entity.Song, error) {
	queryBuilder := s.Builder.Update("music").Where("id = ?", song.Id)

	if song.Name != "" {
		queryBuilder = queryBuilder.Set("song_name", song.Name)
	}

	if song.Group != "" {
		queryBuilder = queryBuilder.Set("group_name", song.Group)
	}

	if song.Text != "" {
		queryBuilder = queryBuilder.Set("text_song", song.Text)
	}

	if song.Genre != "" {
		queryBuilder = queryBuilder.Set("genre", song.Genre)
	}

	if song.Duration != 0 {
		queryBuilder = queryBuilder.Set("duration", song.Duration)
	}

	layout := "2006.02.01"
	defaultDate, _ := time.Parse(layout, "01.01.1970")

	if song.ReleaseData != defaultDate {
		queryBuilder = queryBuilder.Set("release_data", song.ReleaseData)
	}

	if song.Link != "" {
		queryBuilder = queryBuilder.Set("link", song.Link)
	}

	sql, args, _ := queryBuilder.Prefix("RETURNING id, song_name, group_name, text_song, genre, duration, release_date, link").ToSql()

	var newSong entity.Song

	err := s.Pool.QueryRow(ctx, sql, args...).Scan(
		&newSong.Id,
		&newSong.Name,
		&newSong.Group,
		&newSong.Text,
		&newSong.Genre,
		&newSong.Duration,
		&newSong.ReleaseData,
		&newSong.Link,
	)
	if err != nil {
		log.Errorf("pgdb - song - UpdateSong - s.Pool.QueryRow: %v", err)
		return entity.Song{}, err
	}

	return newSong, nil
}

func (s *SongRepo) AddSong(ctx context.Context, song entity.Song) (string, error) {
	sql, args, _ := s.Builder.Insert("music").
		Columns("song_name, group_name, text_song, genre, release_date, duration, link").
		Values(song.Name, song.Group, song.Text, song.Genre, song.ReleaseData, song.Duration, song.Link).
		Suffix("RETURNING id").
		ToSql()

	var id string
	err := s.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		log.Errorf("pgsb - song - UpdateSong - s.Pool.QueryRow: %v", err)
		return "", err
	}

	return id, nil
}
