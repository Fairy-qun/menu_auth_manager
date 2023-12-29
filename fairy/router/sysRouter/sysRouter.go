package sysRouter

import (
	"github.com/gin-gonic/gin"
	"fairy/api"
)

// 系统相关路由
func SysRouter(r *gin.Engine) {
	sysAPI := api.SysApi{}
	r.POST("/register",sysAPI.Register)
	r.POST("/login",sysAPI.Login)
}

