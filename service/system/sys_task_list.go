package system

import (
	"errors"

	"github.com/kinggq/global"
	"github.com/kinggq/model/system"
	"gorm.io/gorm"
)

type TaskListService struct{}

func (s *TaskListService) Create(id uint, list system.ProjTaskLists) (err error) {
	assignErr := global.DB.Where("user_id = ? AND project_id = ?", id, list.ProjectId).First(&system.ProjUserProjects{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("无权限，请联系管理员")
	}
	var l system.ProjProjects
	if errors.Is(global.DB.Where("project_id = ?", list.ProjectId).First(&l).Error, gorm.ErrRecordNotFound) {
		return errors.New("项目不存在")
	}
	err = global.DB.Create(&list).Error
	return err
}
