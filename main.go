package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wc3_game_tracker/api"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", indexHandler)

	api.RegisterTournamentEndpoints(router.Group("/api/v1/"))

	err := router.Run("0.0.0.0:3000")

	if err != nil {
		fmt.Println("Server could not be started", err)
		return
	}
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}
