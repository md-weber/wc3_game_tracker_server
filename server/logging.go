package server

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"strconv"
	"time"
)

func RegisterLogger(router *gin.Engine) {
	prod, err := strconv.ParseBool(os.Getenv("PRODUCTION"))
	if err != nil {
		prod = false
	}

	var logger *zap.Logger

	if prod == true {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(logger, true))
}
