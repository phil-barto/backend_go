package songs

import "github.com/gin-gonic/gin"

type Song struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ArtistId string `json:"artist_id"`
	Length   int    `json:"length"`
}

type SongStore interface {
	GetSong(gin.Context, string) (Song, error)
	GetSongs(gin.Context) ([]Song, error)
	PostSong(gin.Context, string) (Song, error)
}

type SongService struct {
	store SongStore
}

func NewSongService(store SongStore) *SongService {
	return &SongService{
		store: store,
	}
}
