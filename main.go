package main

import (
	"net/http"

	"github.com/bacchilu/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", func(context *gin.Context) {
		context.JSON(http.StatusOK, models.GetAllEvents())
	})
	server.POST("/events", func(context *gin.Context) {
		event := models.Event{}
		err := context.ShouldBindJSON(&event)
		if err != nil {
			context.JSON(http.StatusBadRequest, map[string]string{"msg": "missing data"})
			return
		}
		event.Save()
		context.JSON(http.StatusCreated, event)
	})

	server.Run()
}
