package api_handler

import (
	"backend_go/internal/artists"
	"backend_go/internal/songs"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := Handler{}
	r.GET("/health", h.HealthCheck)
	r.GET("/artists", h.GetArtists)
	r.GET("/artists/:id", h.GetArtist)
	r.GET("/songs/", h.GetSongs)
	r.GET("/songs/:id", h.GetSong)
	return r
}

func TestHealthCheck(t *testing.T) {
	mockResponse := `{"message":"health check okay"}`
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetArtist(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/artists", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	// Decode the JSON bytes into a slice of Artist structs
	var artists []artists.Artist
	err := json.NewDecoder(bytes.NewReader(responseData)).Decode(&artists)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestGetArtistId(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/artists/1432", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	fmt.Println(responseData)

	// Decode the JSON bytes into a slice of Artist structs
	var artist artists.Artist
	err := json.NewDecoder(bytes.NewReader(responseData)).Decode(&artist)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		assert.Fail(t, "Error decoding JSON:", err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetSongId(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/songs/8253", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	// Decode the JSON bytes into a slice of Artist structs
	var song songs.Song
	err := json.NewDecoder(bytes.NewReader(responseData)).Decode(&song)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		assert.Fail(t, "Error decoding JSON:", err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
}
