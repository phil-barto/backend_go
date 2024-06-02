package json_database

import (
	"backend_go/internal/songs"
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"runtime"

	"github.com/chigopher/pathlib"
)

func (j JsonDB) getSongData() ([]songs.Song, error) {

	var songs []songs.Song

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

func (j JsonDB) GetSong(uuid string) (songs.Song, error) {
	var s songs.Song

	songs, err := j.getSongData()
	if err != nil {
		return s, &json.InvalidUnmarshalError{}
	}

	for _, song := range songs {
		if song.ID == uuid {
			return song, nil
		}
	}

	return s, os.ErrNotExist
}

func (j JsonDB) GetSongs() ([]songs.Song, error) {
	songs, err := j.getSongData()
	return songs, err
}
