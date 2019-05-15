package main

import (
	"ch_server/src/model"
	"github.com/gin-gonic/gin"
	"github.com/foolin/gin-template"
)

func main() {
	router := gin.Default()
	router.HTMLRender = gintemplate.Default()
	model.Init()
	router.POST("/register", model.RegisterHandler)
	router.GET("/list", model.ListHandler)
	router.Run(":8000")
}
