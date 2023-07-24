package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wc3_game_tracker/services"
)

type GroupApiI interface {
	allGroupsHandler(ctx *gin.Context)
	oneGroupHandler(ctx *gin.Context)
	createNewGroupHandler(ctx *gin.Context)
}

type GroupApi struct {
	GroupService services.GroupServiceI
}

func (g GroupApi) allGroupsHandler(ctx *gin.Context) {
	groups, err := g.GroupService.GetAllGroups()
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, groups)
}

func (g GroupApi) oneGroupHandler(ctx *gin.Context) {
	// TODO: Extract the id from the ctx
	// TODO: Get a single Groups from DB and respond accordingly
}

func (g GroupApi) createNewGroupHandler(ctx *gin.Context) {
	// TODO: Extract group information from ctx
	// TODO: Create one group and store in DB and respond with the UUID
}

func RegisterGroupEndpoints(router *gin.RouterGroup, service GroupApiI) {
	router.GET("/groups", service.allGroupsHandler)
	router.GET("/group", service.oneGroupHandler)
	router.PUT("/group", service.createNewGroupHandler)
}
