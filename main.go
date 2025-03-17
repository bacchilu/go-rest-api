package main

import (
	"net/http"

	"github.com/bacchilu/rest-api/db"
	"github.com/bacchilu/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", func(context *gin.Context) {
		res, err := models.GetAllEvents()
		if err != nil {
			context.JSON(http.StatusInternalServerError, map[string]string{"msg": "error"})
		}
		context.JSON(http.StatusOK, res)
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
