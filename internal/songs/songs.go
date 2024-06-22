package songs

type Song struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ArtistId string `json:"artist_id"`
	Length   int    `json:"length"`
}
