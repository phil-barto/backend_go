package database

import (
	"backend_go/graph/model"
	"errors"
)

const (
	JSON     = "jsondb"
	POSTGRES = "postgres"
)

type Database interface {
	GetSong(string) (model.Song, error)
	GetSongs() ([]model.Song, error)
	GetArtist(string) (model.Artist, error)
	GetArtists() ([]model.Artist, error)
}

func NewDB(dbType string) (Database, error) {
	if dbType == JSON {
		return newJsonDB()
	} else {
		return nil, errors.ErrUnsupported
	}
}
