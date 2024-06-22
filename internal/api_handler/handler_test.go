package api_handler

import (
	"backend_go/graph/model"
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
	h, err := NewHandler()

	if err != nil {
		fmt.Println(err)
	}
	r.GET("/health", h.HealthCheck)
	r.GET("/artists", h.GetArtists)
	r.GET("/artists/:id", h.GetArtist)
	r.GET("/songs", h.GetSongs)
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
	var artists []model.Artist
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
	var artist model.Artist
	err := json.NewDecoder(bytes.NewReader(responseData)).Decode(&artist)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		assert.Fail(t, "Error decoding JSON:", err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetSongs(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/songs", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	// Decode the JSON bytes into a slice of Artist structs
	var songs []model.Song
	err := json.NewDecoder(bytes.NewReader(responseData)).Decode(&songs)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		assert.Fail(t, "Error decoding JSON:", err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetSong(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/songs/8253", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	// Decode the JSON bytes into a slice of Artist structs
	var song model.Song
	err := json.NewDecoder(bytes.NewReader(responseData)).Decode(&song)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		assert.Fail(t, "Error decoding JSON:", err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
}
