package initialize


import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
	_"fairy/docs"
	"fairy/router"
	"fairy/middleware"

)

func RunServer() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 初始化路由实例
	var routerGroup = new(router.RouterGroup)

	routerGroup.InitRouter(r)
	// 实例化路由
	r.Run(":5001")
}