package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPastes(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/pastes/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetPastesById(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/pastes/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPostPastes(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	blob := []byte(`{"blob": "bar"}`)
	req, _ := http.NewRequest("POST", "/api/v1/pastes/", bytes.NewBuffer(blob))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestDeletePastes(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/pastes/123", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
