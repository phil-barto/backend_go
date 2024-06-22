package database

import (
	"backend_go/graph/model"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"runtime"

	"github.com/chigopher/pathlib"
)

type JsonDB struct {
}

func newJsonDB() (JsonDB, error) {
	return JsonDB{}, nil
}

func (j JsonDB) getSongData() ([]model.Song, error) {

	var songs []model.Song

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("Error getting current file path")
	}

	// Create a new Path object with the current file path
	path := pathlib.NewPath(filename).Parent().Parent().Parent().Join("data", "songs.json")

	data, err := os.ReadFile(path.String())
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
		return songs, &fs.PathError{}
	}

	if err := json.Unmarshal(data, &songs); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
		return songs, &json.UnmarshalTypeError{}
	}

	return songs, nil
}

func (j JsonDB) GetSong(uuid string) (model.Song, error) {
	var s model.Song

	songs, err := j.getSongData()
	if err != nil {
		fmt.Println(err)
		return s, &json.InvalidUnmarshalError{}
	}

	for _, song := range songs {
		if song.ID == uuid {
			return song, nil
		}
	}

	return s, os.ErrNotExist
}

func (j JsonDB) GetSongs() ([]model.Song, error) {
	songs, err := j.getSongData()
	return songs, err
}

func (JsonDB) getJsonData() ([]model.Artist, error) {
	var artists []model.Artist

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

func (j JsonDB) GetArtist(uuid string) (model.Artist, error) {
	var a model.Artist
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

func (j JsonDB) GetArtists() ([]model.Artist, error) {
	artists, err := j.getJsonData()
	return artists, err
}
