package main

import (
	"net/http"

	"github.com/bacchilu/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", func(context *gin.Context) {
		context.JSON(http.StatusOK, models.GetEvents())
	})

	server.Run()
}
