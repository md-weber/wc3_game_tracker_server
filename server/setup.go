package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"wc3_game_tracker/api"
	"wc3_game_tracker/repo"
	"wc3_game_tracker/services"
)

const ApiPrefix = "/api/v1/"

func SetupServer() *gin.Engine {
	router := gin.Default()

	registerLogger(router)
	env := os.Getenv("ENVIRONMENT")

	if env == "prod" || env == "dev" {
		router.LoadHTMLGlob("templates/*")
	} else {
		router.LoadHTMLGlob("../templates/*")
	}

	groupApi := registerGroupApi()

	api.RegisterWebEndpoints(router.Group(""))
	api.RegisterGroupEndpoints(router.Group(ApiPrefix), groupApi)
	api.RegisterLeagueEndpoints(router.Group(ApiPrefix))
	return router
}

func registerGroupApi() *api.GroupApi {
	groupApi := new(api.GroupApi)
	groupService := new(services.GroupService)
	groupRepo := new(repo.GroupRepository)

	groupService.GroupRepo = groupRepo
	groupApi.GroupService = groupService
	return groupApi
}

func Setup() *gin.Engine {
	router := SetupServer()
	err := router.Run("0.0.0.0:3000")

	if err != nil {
		fmt.Println("Server could not be started", err)
		return nil
	}
	return router

}
