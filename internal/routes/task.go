package routes

import (
	"github.com/gin-gonic/gin"
	"todoApp/internal/controllers"
)

func task(r *gin.Engine, c controllers.TaskController) {
	r.GET("/", c.Index)
	r.POST("/create", c.Create)
	r.DELETE("/remove", c.Remove)
	r.GET("/:id", c.GetSingle)
}
