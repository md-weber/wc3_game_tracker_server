package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wc3_game_tracker/api"
)

func Setup() {
	router := gin.Default()

	registerLogger(router)
	router.LoadHTMLGlob("templates/*")

	api.RegisterWebEndpoints(router.Group(""))
	api.RegisterTournamentEndpoints(router.Group("/api/v1/"))

	err := router.Run("0.0.0.0:3000")

	if err != nil {
		fmt.Println("Server could not be started", err)
		return
	}
}
