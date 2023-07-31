package integration

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"wc3_game_tracker/api/models"
)

func TestGetAllLeagues(t *testing.T) {
	r := initTesting()

	req, _ := http.NewRequest("GET", "/api/v1/leagues", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var result []*models.League

	responseData, _ := io.ReadAll(w.Body)
	_ = json.Unmarshal(responseData, &result)

	assert.Equal(t, "Test League", result[0].Name)
	assert.Equal(t, "https://creepcamp.de/", result[0].Website)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestInsertLeagueInDB(t *testing.T) {
	r := initTesting()
	m, b := map[string]any{
		"name":      "Test League",
		"startDate": "2023-07-17T07:45:42.914Z",
		"endDate":   "2023-07-25T07:45:42.914Z",
		"website":   "https://creepcamp.de/",
	}, new(bytes.Buffer)

	_ = json.NewEncoder(b).Encode(m)

	req, _ := http.NewRequest("POST", "/api/v1/leagues", b)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	_, err := uuid.Parse(w.Body.String())
	if err != nil {
		t.Error("The given String is not a UUID.", err)
	}

}
