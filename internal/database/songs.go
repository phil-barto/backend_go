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

func (j JsonDB) getSongData() ([]model.Song, error) {

	var songs []model.Song

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("Error getting current file path")
	}

	// Create a new Path object with the current file path
	path := pathlib.NewPath(filename).Parent().Parent().Parent().Join("data", "songs.json")

	data, err := os.ReadFile(path.String())
	fmt.Println("data: %s", data)
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
	fmt.Println("getting songs beginning")
	songs, err := j.getSongData()
	fmt.Println("getting songs end")
	return songs, err
}
