package api_handler

import (
	"backend_go/internal/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	log log.Logger
	db  database.Database
}

func (h Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "health check okay",
	})
}

func (h Handler) GetArtist(c *gin.Context) {
	id := c.Param("id")
	artist, err := h.db.GetArtist(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "artist not found"})
	} else {
		c.JSON(http.StatusOK, artist)
	}
}

func (h Handler) GetArtists(c *gin.Context) {
	artists, err := h.db.GetArtists()
	if err != nil {
		fmt.Println("artists failed")
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	} else {
		c.JSON(http.StatusOK, artists)
	}
}

func (h Handler) GetSongs(c *gin.Context) {
	songs, err := h.db.GetSongs()
	if err != nil {
		fmt.Println("songs failed")
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	} else {
		c.JSON(http.StatusOK, songs)
	}
}

func (h Handler) GetSong(c *gin.Context) {
	id := c.Param("id")
	song, err := h.db.GetSong(id)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	} else {
		c.JSON(http.StatusOK, song)
	}
}

func NewHandler() (Handler, error) {
	db, err := database.NewDB(database.JSON)
	if err != nil {
		fmt.Println("formatting error")
	}

	return Handler{
		db: db,
	}, nil
}
