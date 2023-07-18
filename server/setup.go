package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"wc3_game_tracker/api"
)

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
	api.RegisterLeagueEndpoints(router.Group("/api/v1/"))
	return router
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
