package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/kinggq/api/v1"
)

type BaseRouter struct{}

func (self *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
		baseRouter.POST("/login", baseApi.Login)
		baseRouter.POST("/register", baseApi.Register)
	}
	return baseRouter
}
