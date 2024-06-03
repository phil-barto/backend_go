package api_handler

import (
	"backend_go/internal/artists"
	"backend_go/internal/json_database"
	"backend_go/internal/songs"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	log       log.Logger
	a_service artists.ArtistService
	s_service songs.SongService
}

func (h Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "health check okay",
	})
}

func (h Handler) GetArtist(c *gin.Context) {
	id := c.Param("id")
	artist, err := h.a_service.Store.GetArtist(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "artist not found"})
	} else {
		c.JSON(http.StatusOK, artist)
	}
}

func (h Handler) GetArtists(c *gin.Context) {
	artists, err := h.a_service.Store.GetArtists()
	if err != nil {
		fmt.Println("artists failed")
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	} else {
		c.JSON(http.StatusOK, artists)
	}
}

func (h Handler) GetSongs(c *gin.Context) {
	songs, err := h.s_service.Store.GetSongs()
	if err != nil {
		fmt.Println("songs failed")
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	} else {
		c.JSON(http.StatusOK, songs)
	}
}

func (h Handler) GetSong(c *gin.Context) {
	id := c.Param("id")
	song, err := h.s_service.Store.GetSong(id)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	} else {
		c.JSON(http.StatusOK, song)
	}
}

func NewHandler() (Handler, error) {
	db, err := json_database.NewJsonDB()
	if err != nil {
		fmt.Println("formatting error")
	}
	a_service := artists.NewArtistService(db)
	s_service := songs.NewSongService(db)

	return Handler{
		a_service: *a_service,
		s_service: *s_service,
	}, nil
}
