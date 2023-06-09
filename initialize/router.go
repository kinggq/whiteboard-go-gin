package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/kinggq/middleware"
	"github.com/kinggq/router"
)

func Routes() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System

	// 放行全部跨域请求
	Router.Use(middleware.Cors())

	PublicGroup := Router.Group("")
	{
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitTaskRouter(PublicGroup)
	}
	return Router
}
