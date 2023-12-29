package router

import (
	"github.com/gin-gonic/gin"
	"fairy/router/sysRouter"
)

type RouterGroup struct {}

func (router RouterGroup) InitRouter(r *gin.Engine) {
	sysRouter.SysRouter(r)
}