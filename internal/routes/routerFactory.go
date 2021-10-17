package routes

import (
	"github.com/gin-gonic/gin"
	"todoApp/internal/app"
	"todoApp/internal/controllers"
)

func BuildHttpRouter(app app.Application) *gin.Engine {
	r := gin.Default()
	c := controllers.NewTaskController(app)
	task(r, *c)

	return r
}
