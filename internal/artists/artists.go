package artists

type Artist struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	NumViews int    `json:"num_views"`
}

type ArtistStore interface {
	GetArtist(string) (Artist, error)
	GetArtists() ([]Artist, error)
}

type ArtistService struct {
	Store ArtistStore
}

func NewArtistService(store ArtistStore) *ArtistService {
	return &ArtistService{
		Store: store,
	}
}
