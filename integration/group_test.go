package integration

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"wc3_game_tracker/api/models"
)

func TestGetAllGroups(t *testing.T) {
	r := initTesting()

	req, _ := http.NewRequest("GET", "/api/v1/groups", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var result []*models.Group

	responseData, _ := io.ReadAll(w.Body)
	_ = json.Unmarshal(responseData, &result)

	assert.Equal(t, "Test Group", result[0].Name)
	assert.Equal(t, 3, result[0].Vetos)
	assert.Equal(t, 3, result[0].MapsPerMatch)
	assert.Equal(t, "In Progress", result[0].Status)
	assert.Equal(t, "Legolas", result[0].Admin)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestStoreAGroup(t *testing.T) {
	r := initTesting()
	m, b := map[string]any{
		"name":           "Test League",
		"league_id":      "5ed90ba3-1ff9-4f40-ac29-c6f5a0ab230b",
		"admin":          "Max Fosh",
		"vetos":          3,
		"maps_per_match": 3,
		"status":         "In Progress",
	}, new(bytes.Buffer)

	_ = json.NewEncoder(b).Encode(m)

	req, _ := http.NewRequest("POST", "/api/v1/groups", b)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	_, err := uuid.Parse(w.Body.String())
	if err != nil {
		t.Error("The given String is not a UUID.", err)
	}
}
