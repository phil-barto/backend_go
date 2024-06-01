package json_database

import (
	"backend_go/internal/artists"
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"runtime"

	"github.com/chigopher/pathlib"
)

func (JsonDB) getJsonData() ([]artists.Artist, error) {
	var artists []artists.Artist

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("Error getting current file path")
	}

	// Create a new Path object with the current file path
	path := pathlib.NewPath(filename).Parent().Parent().Parent().Join("data", "artists.json")

	data, err := os.ReadFile(path.String())
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
		return artists, &fs.PathError{}
	}

	if err := json.Unmarshal(data, &artists); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
		return artists, &json.UnmarshalTypeError{}
	}

	return artists, nil
}

func (j JsonDB) GetArtist(uuid string) (artists.Artist, error) {
	var a artists.Artist
	artists, err := j.getJsonData()
	if err != nil {
		return a, &json.InvalidUnmarshalError{}
	}

	for _, artist := range artists {
		if artist.ID == uuid {
			return artist, nil
		}
	}

	return a, os.ErrNotExist
}

func (j JsonDB) GetArtists() ([]artists.Artist, error) {
	artists, err := j.getJsonData()
	return artists, err
}

func (j *JsonDB) PostArtist(id string) (artists.Artist, error) {
	return artists.Artist{
		ID:       "1234",
		Name:     "Johnny Depp",
		Email:    "johnny.depp@gmail.com",
		NumViews: 123,
	}, nil
}
