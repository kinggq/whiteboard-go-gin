package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/kinggq/api/v1"
)

type TaskRouter struct{}

func (r *TaskRouter) InitTaskRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	taskRouter := Router.Group("task")
	{
		taskApi := v1.ApiGroupApp.SystemApiGroup.TaskApi
		taskListApi := v1.ApiGroupApp.SystemApiGroup.TaskListApi
		projectApi := v1.ApiGroupApp.SystemApiGroup.ProjectApi

		taskRouter.POST("/create", taskApi.Create)
		taskRouter.POST("/list/create", taskListApi.Create)
		taskRouter.POST("/project/create", projectApi.Create)
	}

	return taskRouter
}
