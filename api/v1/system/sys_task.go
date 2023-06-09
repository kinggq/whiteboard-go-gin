package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kinggq/model/common/response"
	"github.com/kinggq/model/system"
	"github.com/kinggq/utils"
)

type TaskApi struct{}

func (t *TaskApi) Create(ctx *gin.Context) {
	var task system.ProjTasks
	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = utils.Verify(task, utils.TaskCreateVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	if task.PriorityId == nil {
		task.PriorityId = nil
	}

	taskParams := &system.ProjTasks{
		TaskTitle:  task.TaskTitle,
		CreatorId:  task.CreatorId,
		PriorityId: task.PriorityId,
		ListId:     task.ListId,
	}
	err = taskService.Create(*taskParams)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithDetailed("", "创建成功", ctx)
}
