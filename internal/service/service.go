package service

import "github.com/magmaheat/music-info/internal/repo"

type LibraryMusic interface {
	CreateSong()
	GetSong()
	UpdateSong()
}

type Services struct {
	LibraryMusic LibraryMusic
}

func NewServices(repo *repo.Repositories) *Services {
	return &Services{LibraryMusic: NewLibraryMusic(repo.Song)}
}
