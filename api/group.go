package api

import (
	"encoding/json"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strings"
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
	group, err := w.GroupService.GetGroup(id)

	if err != nil && strings.Contains(err.Error(), "no rows in result set") {
		c.JSON(http.StatusNotFound, err)
	}

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, group)
}
