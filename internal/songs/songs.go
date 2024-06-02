package songs

type Song struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ArtistId string `json:"artist_id"`
	Length   int    `json:"length"`
}

type SongStore interface {
	GetSong(string) (Song, error)
	GetSongs() ([]Song, error)
}

type SongService struct {
	store SongStore
}

func NewSongService(store SongStore) *SongService {
	return &SongService{
		store: store,
	}
}
