package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
	"todoApp/internal/app"
	"todoApp/internal/app/command"
)

type TaskController struct {
	app app.Application
	log *logrus.Logger
}

func NewTaskController(app app.Application) *TaskController {
	return &TaskController{app: app}
}

func (c TaskController) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": c.app.Queries.ShowAll.Handle(),
	})
}

func (c TaskController) Create(ctx *gin.Context) {
	req := struct {
		Id   int    `json:"id"`
		Item string `json:"item"`
	}{}
	err := ctx.BindJSON(&req)

	if err != nil {
		c.log.Error(err)
		ctx.JSON(500, gin.H{"data": err.Error()})
	}

	c.app.Commands.Create.Handle(command.CreateCommand{Id: req.Id, Message: req.Item})
	ctx.JSON(200, gin.H{})
}

func (c TaskController) Remove(ctx *gin.Context) {
	req := struct {
		Id int `json:"id"`
	}{}
	err := ctx.BindJSON(&req)

	if err != nil {
		c.log.Error(err)
		ctx.JSON(500, gin.H{"data": err.Error()})
	}

	c.app.Commands.Remove.Handle(command.RemoveCommand{Id: req.Id})
	ctx.JSON(200, gin.H{})
}

func (c TaskController) GetSingle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		c.log.Error(err)
		ctx.JSON(400, gin.H{"data": "id param needs to be numeric"})
	}

	ctx.JSON(200, gin.H{"data": c.app.Queries.GetSingle.Handle(id)})
}
