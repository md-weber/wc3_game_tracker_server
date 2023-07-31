package api

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (w Warcraft3ServerImpl) FindGroups(c *gin.Context) {
	groups, err := w.GroupService.GetAllGroups()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (w Warcraft3ServerImpl) AddGroup(c *gin.Context) {
	//TODO implement me
	// TODO: Extract group information from ctx
	// TODO: Create one group and store in DB and respond with the UUID
	panic("implement me")
}

func (w Warcraft3ServerImpl) FindGroup(c *gin.Context, id openapi_types.UUID) {
	//TODO implement me
	// TODO: Extract the id from the ctx
	// TODO: Get a single Groups from DB and respond accordingly
	panic("implement me")
}
