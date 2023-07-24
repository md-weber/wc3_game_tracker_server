package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wc3_game_tracker/repo"
)

func allGroupsHandler(ctx *gin.Context) {
	err, groups := repo.GetAllGroups()
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, groups)
}

func oneGroupHandler(ctx *gin.Context) {
	// TODO: Extract the id from the ctx
	// TODO: Get a single Groups from DB and respond accordingly
}

func createNewGroupHandler(ctx *gin.Context) {
	// TODO: Extract group information from ctx
	// TODO: Create one group and store in DB and respond with the UUID
}

func RegisterGroupEndpoints(router *gin.RouterGroup) {
	router.GET("/groups", allGroupsHandler)
	router.GET("/group", oneGroupHandler)
	router.PUT("/group", createNewGroupHandler)
}
