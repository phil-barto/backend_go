package api_handler

import (
	"backend_go/internal/artists"
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
	router := gin.Default()
	return router
}

func TestHealthCheck(t *testing.T) {
	h := Handler{}
	mockResponse := `{"message":"health check okay"}`

	r := setupRouter()
	r.GET("/health", h.HealthCheck)

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetArtist(t *testing.T) {
	h, err := NewHandler()
	if err != nil {
		fmt.Println(err)
	}

	r := setupRouter()
	r.GET("/artists", h.GetArtists)

	req, _ := http.NewRequest("GET", "/artists", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	// Decode the JSON bytes into a slice of Artist structs
	var artists []artists.Artist
	err = json.NewDecoder(bytes.NewReader(responseData)).Decode(&artists)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestGetArtistId(t *testing.T) {
	h, err := NewHandler()
	if err != nil {
		fmt.Println(err)
	}

	r := setupRouter()
	r.GET("/artists/:id", h.GetArtist)

	req, _ := http.NewRequest("GET", "/artists/1432", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	fmt.Println(responseData)

	// Decode the JSON bytes into a slice of Artist structs
	var artist artists.Artist
	err = json.NewDecoder(bytes.NewReader(responseData)).Decode(&artist)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		assert.Fail(t, "Error decoding JSON:", err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
}
