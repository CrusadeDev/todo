package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"todoApp/internal/app"
	"todoApp/internal/controllers"
)

func BuildHttpRouter(app app.Application, log *logrus.Logger) *gin.Engine {
	r := gin.Default()
	c := controllers.NewTaskController(app, log)
	task(r, *c)

	return r
}
