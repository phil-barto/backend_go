package database

import (
	"backend_go/internal/artists"
	"backend_go/internal/songs"
	"errors"
)

const (
	JSON     = "jsondb"
	POSTGRES = "postgres"
)

type Database interface {
	GetSong(string) (songs.Song, error)
	GetSongs() ([]songs.Song, error)
	GetArtist(string) (artists.Artist, error)
	GetArtists() ([]artists.Artist, error)
}

type JsonDB struct {
}

func newJsonDB() (JsonDB, error) {
	return JsonDB{}, nil
}

func NewDB(dbType string) (Database, error) {
	if dbType == JSON {
		return newJsonDB()
	} else {
		return nil, errors.ErrUnsupported
	}
}
