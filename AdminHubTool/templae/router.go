package templae

const TemplateRouterStr = `package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/exec/router"
)

func writeRouter() *gin.Engine {
	r := router.CreateRouter()
	
	r.Write()
	return r.R
}
`
