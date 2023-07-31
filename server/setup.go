package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"wc3_game_tracker/api"
	"wc3_game_tracker/controller"
	"wc3_game_tracker/repo"
	"wc3_game_tracker/services"
)

func SetupHandler(r *gin.Engine) {
	groupRepo := new(repo.GroupRepository)
	groupService := services.GroupService{GroupRepo: groupRepo}
	myApi := api.Warcraft3ServerImpl{GroupService: groupService}

	controller.RegisterHandlers(r.Group("/api/v1"), myApi)
}

func SetupServer() *gin.Engine {
	router := gin.Default()

	registerLogger(router)
	env := os.Getenv("ENVIRONMENT")

	if env == "prod" || env == "dev" {
		router.LoadHTMLGlob("templates/*")
	} else {
		router.LoadHTMLGlob("../templates/*")
	}

	api.RegisterWebEndpoints(router.Group(""))
	SetupHandler(router)
	return router
}

func Setup() *gin.Engine {
	router := SetupServer()
	err := router.Run(":3000")

	if err != nil {
		fmt.Println("Server could not be started", err)
		return nil
	}
	return router

}
