package artists

type Artist struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	NumViews int    `json:"num_views"`
}
