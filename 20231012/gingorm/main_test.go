package main

import (
	"net/http"
	"bytes"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetAirportRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/airport/8", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"airport-name\":\"Updated Airport\",\"airport-place\":\"Updated Place\",\"status\":\"ok\"}", w.Body.String())
}

func TestPostAirportRoute(t *testing.T) {
	router := setupRouter()

	requestBody := []byte(`{"name": "post_test_airport", "place": "post_test_place"}`)

	req, _ := http.NewRequest("POST", "/airport/create", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"airport-name\":\"post_test_airport\",\"airport-place\":\"post_test_place\",\"status\":\"ok\"}", w.Body.String())
}

func TestPostAirportRouteWithEmptyBody(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("POST", "/airport/create", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"status\":\"bad request\"}", w.Body.String())
}

func TestPostAirportRouteWithLuckOfName(t *testing.T) {
	router := setupRouter()

	requestBody := []byte(`{"place": "post_test_place"}`)

	req, _ := http.NewRequest("POST", "/airport/create", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"status\":\"bad request\"}", w.Body.String())
}

// func TestPostAirportRouteFailToCreate(t *testing.T) {
// 	router := setupRouter()

// 	requestBody := []byte(`{"name": "post_test_airport", "place": "post_test_place"}`)

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	mock.ExpectExec("INSERT INTO `airports`").WillReturnError(fmt.Errorf("some error"))

// 	req, _ := http.NewRequest("POST", "/airport/create", bytes.NewBuffer(requestBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 500, w.Code)
// 	assert.Equal(t, "{\"status\":\"Something went wrong\"}", w.Body.String())
// }

func TestPutAirportRoute(t *testing.T) {
	router := setupRouter()

	requestBody := []byte(`{"id": 10, "name": "Updated Airport", "place": "Updated Place"}`)

	req, _ := http.NewRequest("PUT", "/airport/update", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"airport-name\":\"Updated Airport\",\"airport-place\":\"Updated Place\",\"status\":\"ok\"}", w.Body.String())
}

func TestPutAirportRouteWithEmptyBody(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("PUT", "/airport/update", nil)
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"status\":\"bad request\"}", w.Body.String())
}

func TestPutAirportRouteNotFound(t *testing.T){
	router := setupRouter()

	// 存在しないIDを指定
	requestBody := []byte(`{"id": 1000, "name": "Updated Airport", "place": "Updated Place"}`)

	req, _ := http.NewRequest("PUT", "/airport/update", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "{\"status\":\"not found\"}", w.Body.String())
}

func TestPutAirportRouteUpdateOnlyName(t *testing.T){
	router := setupRouter()

	requestBody := []byte(`{"id": 10, "name": "Test Updated Airport"}`)

	req, _ := http.NewRequest("PUT", "/airport/update", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"airport-name\":\"Test Updated Airport\",\"airport-place\":\"Updated Place\",\"status\":\"ok\"}", w.Body.String())
}

func TestPutAirportRouteUpdateOnlyPlace(t *testing.T){
	router := setupRouter()

	requestBody := []byte(`{"id": 10, "place": "Test Updated Place"}`)

	req, _ := http.NewRequest("PUT", "/airport/update", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"airport-name\":\"Test Updated Airport\",\"airport-place\":\"Test Updated Place\",\"status\":\"ok\"}", w.Body.String())
}

func TestDeleteAirportRoute(t *testing.T){
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/airport/delete/11", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"status\":\"ok\"}", w.Body.String())
}

func TestDeleteAirportRouteNotFound(t *testing.T){
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/airport/delete/1000", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "{\"status\":\"not found\"}", w.Body.String())
}

