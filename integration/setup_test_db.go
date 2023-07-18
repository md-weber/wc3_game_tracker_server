package integration

import (
	"github.com/gin-gonic/gin"
	"os"
	"time"
	"wc3_game_tracker/api/models"
	"wc3_game_tracker/repo"
	"wc3_game_tracker/server"
)

func initTesting() *gin.Engine {
	_ = os.Setenv("DATABASE_URL", "postgres://local-dev@localhost:5433/api?sslmode=disable")
	r := server.SetupServer()
	initDatabase()
	return r
}

// Do not use this method will clean the Database, only for testing purposes

//goland:noinspection ALL - This is a test only function and should never be called outside
func cleanDatabase() {
	db := repo.GetOpenConnection()
	_, err := db.Exec("DELETE FROM league")
	if err != nil {
		return
	}
}

func initDatabase() {
	cleanDatabase()
	var testLeague = new(models.League)

	testLeague.Name = "Test League"
	testLeague.StartDate = models.CustomTime{Time: time.Now()}
	testLeague.EndDate = models.CustomTime{Time: time.Now()}
	testLeague.Website = "https://creepcamp.de/"

	repo.SaveLeague(testLeague)
}