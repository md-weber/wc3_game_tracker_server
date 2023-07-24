package integration

import (
	"encoding/json"
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
