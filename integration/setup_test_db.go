package integration

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	testLeague.Id = uuid.MustParse("5ed90ba3-1ff9-4f40-ac29-c6f5a0ab230b")
	testLeague.Name = "Test League"
	testLeague.StartDate = models.CustomTime{Time: time.Now()}
	testLeague.EndDate = models.CustomTime{Time: time.Now()}
	testLeague.Website = "https://creepcamp.de/"

	var testGroup = new(models.Group)
	testGroup.Id = uuid.MustParse("6c75ad7d-6dfa-4f14-8d71-df41fac6fead")
	testGroup.Name = "Test Group"
	testGroup.LeagueId = uuid.MustParse("5ed90ba3-1ff9-4f40-ac29-c6f5a0ab230b")
	testGroup.Admin = "Legolas"
	testGroup.Vetos = 3
	testGroup.MapsPerMatch = 3
	testGroup.Status = "In Progress"

	groupRepo := new(repo.GroupRepository)
	_, _ = groupRepo.SaveGroup(testGroup)
	repo.SaveLeague(testLeague)
}
