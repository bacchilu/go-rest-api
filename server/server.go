package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bacchilu/rest-api/interactor"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	app    interactor.Application
}

func NewServer(app interactor.Application) Server {
	engine := gin.Default()

	engine.GET("/events", func(context *gin.Context) {
		res, err := app.ListEvents()
		if err != nil {
			context.JSON(http.StatusInternalServerError, map[string]string{"msg": "error"})
			return
		}
		context.JSON(http.StatusOK, res)
	})

	engine.GET("/events/:id", func(context *gin.Context) {
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, map[string]string{"msg": "error"})
			return
		}

		res, err := app.GetEvent(id)
		if err != nil {
			context.JSON(http.StatusInternalServerError, map[string]string{"msg": "error"})
			return
		}
		context.JSON(http.StatusOK, res)
	})

	engine.POST("/events", func(context *gin.Context) {
		event := interactor.Event{}
		err := context.ShouldBindJSON(&event)
		if err != nil {
			context.JSON(http.StatusBadRequest, map[string]string{"msg": "missing data"})
			return
		}
		event, err = app.CreateEvent(event)
		if err != nil {
			context.JSON(http.StatusBadRequest, map[string]string{"msg": "save error"})
			return
		}
		context.JSON(http.StatusCreated, event)
	})

	engine.PUT("/events/:id", func(context *gin.Context) {
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, map[string]string{"msg": "error"})
			return
		}

		event := interactor.Event{}
		err = context.ShouldBindJSON(&event)
		if err != nil {
			context.JSON(http.StatusBadRequest, map[string]string{"msg": "missing data"})
			return
		}
		event.ID = id
		err = app.UpdateEvent(event)
		if err != nil {
			context.JSON(http.StatusBadRequest, map[string]string{"msg": fmt.Sprintf("save error - %v", err)})
			return
		}
		context.JSON(http.StatusCreated, event)
	})

	return Server{engine, app}
}

func (s Server) Run() {
	s.engine.Run()
}
