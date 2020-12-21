package server

import (
"tokenAuthGolang/server/controller"
"github.com/gin-gonic/gin"
)

func LoadRoutes() {
	r := gin.Default()
	r.GET("/", controller.UpAndRunningResponse)
	r.POST("/login", controller.ValidateFormLoginInput, controller.ProcessLogin)
	r.Run()
}