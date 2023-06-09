package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kinggq/model/common/response"
	"github.com/kinggq/model/system"
	"github.com/kinggq/utils"
)

type TaskListApi struct{}

func (t TaskListApi) Create(ctx *gin.Context) {
	var tl system.ProjTaskLists
	err := ctx.ShouldBindJSON(&tl)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = utils.Verify(tl, utils.TaskListVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	l := &system.ProjTaskLists{
		ListTitle: tl.ListTitle,
		ProjectId: tl.ProjectId,
	}
	err = taskListService.Create(*l)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithDetailed("", "创建成功", ctx)
}
