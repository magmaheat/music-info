package repo

type MetaRepo interface {
	CreateSong()
}

type BaseRepo interface {
	CreateText() (int, error)
	GetText() error
}

type Repositories struct {
	MetaRepo
	BaseRepo
}

func New() *Repositories {
	return Repositories{
		MetaRepo: ,
	}
}
