package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"wc3_game_tracker/api/models"
	"wc3_game_tracker/repo"
)

func allLeaguesHandler(context *gin.Context) {
	leagues, err := repo.GetAllLeagues()
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
	}
	context.JSON(http.StatusOK, leagues)
}

func createNewLeagueHandler(context *gin.Context) {
	var league *models.League

	jsonData, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		log.Fatal("We cannot read the request body")
	}
	err1 := json.Unmarshal(jsonData, &league)
	if err1 != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		log.Fatal("We cannot unmarshall the jsonData", err1)
	}

	storedLeague := repo.SaveLeague(league)
	if storedLeague == nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	}

	context.JSON(http.StatusOK, storedLeague.Id)
}

func RegisterLeagueEndpoints(router *gin.RouterGroup) {
	router.GET("/leagues", allLeaguesHandler)
	router.PUT("/league", createNewLeagueHandler)
}
