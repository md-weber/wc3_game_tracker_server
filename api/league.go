package api

import (
	"encoding/json"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strings"
	"wc3_game_tracker/api/models"
	"wc3_game_tracker/repo"
)

func (w Warcraft3ServerImpl) FindLeagues(c *gin.Context) {
	leagues, err := repo.GetAllLeagues()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, leagues)
}

func (w Warcraft3ServerImpl) AddLeague(c *gin.Context) {
	var league *models.League

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Fatal("We cannot read the request body")
	}
	err1 := json.Unmarshal(jsonData, &league)
	if err1 != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Fatal("We cannot unmarshall the jsonData", err1)
	}

	storedLeague := repo.SaveLeague(league)
	if storedLeague == nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, storedLeague.Id)
}

func (w Warcraft3ServerImpl) FindLeague(c *gin.Context, id openapi_types.UUID) {
	league, err := repo.GetLeague(id)

	if err != nil && strings.Contains(err.Error(), "no rows in result set") {
		c.JSON(http.StatusNotFound, err)
	}

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, league)
}
