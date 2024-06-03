package main

import (
	"backend_go/internal/api_handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	h, err := api_handler.NewHandler()
	if err != nil {
		fmt.Println("error instantiating handler")
	}

	r.GET("/health", h.HealthCheck)
	r.GET("/artists", h.GetArtists)
	r.GET("/artists/:id", h.GetArtist)
	r.GET("/songs", h.GetSongs)
	r.GET("/songs/:id", h.GetSong)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
