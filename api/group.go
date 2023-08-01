package api

import (
	"encoding/json"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"wc3_game_tracker/api/models"
)

func (w Warcraft3ServerImpl) FindGroups(c *gin.Context) {
	groups, err := w.GroupService.GetAllGroups()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (w Warcraft3ServerImpl) AddGroup(c *gin.Context) {
	var group *models.Group

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, "The request body did not have a group")
		log.Println("We cannot read the request body")
	}
	err1 := json.Unmarshal(jsonData, &group)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, "The group was not in the right format")
		log.Fatal("We cannot unmarshall the jsonData", err1)
	}

	id, err := w.GroupService.SaveGroup(group)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, id)
}

func (w Warcraft3ServerImpl) FindGroup(c *gin.Context, id openapi_types.UUID) {
	//TODO implement me
	// TODO: Extract the id from the ctx
	// TODO: Get a single Groups from DB and respond accordingly
	panic("implement me")
}
