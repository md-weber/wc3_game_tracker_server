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

func allTournamentsHandler(context *gin.Context) {
	tournaments, err := repo.GetAllTournaments()
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
	}
	context.JSON(http.StatusOK, tournaments)
}

func createNewTournamentHandler(context *gin.Context) {
	var testTournament *models.Tournament

	jsonData, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		log.Fatal("We cannot read the request body")
	}
	err1 := json.Unmarshal(jsonData, &testTournament)
	if err1 != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		log.Fatal("We cannot unmarshall the jsonData", err1)
	}

	storedTournament := repo.SaveTournament(testTournament)
	if storedTournament == nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	}

	context.JSON(http.StatusOK, storedTournament.Id)
}

func RegisterTournamentEndpoints(router *gin.RouterGroup) {
	router.GET("/tournaments", allTournamentsHandler)
	router.PUT("/tournament", createNewTournamentHandler)
}
