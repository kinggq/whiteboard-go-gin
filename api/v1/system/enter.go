package system

import "github.com/kinggq/service"

type ApiGroup struct {
	BaseApi
	TaskApi
	TaskListApi
	ProjectApi
}

var (
	userService     = service.ServiceGroupApp.SystemServiceGroup.UserService
	taskService     = service.ServiceGroupApp.SystemServiceGroup.TaskService
	taskListService = service.ServiceGroupApp.SystemServiceGroup.TaskListService
	projectService  = service.ServiceGroupApp.SystemServiceGroup.ProjectService
)
