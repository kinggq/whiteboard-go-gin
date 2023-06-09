package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kinggq/model/common/response"
	"github.com/kinggq/model/system"
	"github.com/kinggq/utils"
)

type ProjectApi struct{}

func (p *ProjectApi) Create(ctx *gin.Context) {
	var proj system.ProjProjects
	err := ctx.ShouldBindJSON(&proj)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = utils.Verify(proj, utils.ProjectVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	project := &system.ProjProjects{
		ProjectName: proj.ProjectName,
	}
	err = projectService.Create(*project)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithDetailed("", "创建成功", ctx)
}
