package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Tournament struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

var tournaments = []Tournament{
	{Id: uuid.New(), Name: "Creepcamp League", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 3)},
	{Id: uuid.New(), Name: "League 1", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 4)},
	{Id: uuid.New(), Name: "League 2", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 5)},
}

func allTournamentsHandler(context *gin.Context) {
	// Get all tournaments and return to client
	context.IndentedJSON(http.StatusOK, tournaments)
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", indexHandler)
	router.GET("/tournaments", allTournamentsHandler)

	err := router.Run("0.0.0.0:3000")

	if err != nil {
		fmt.Println("Server could not be started", err)
		return
	}
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}
