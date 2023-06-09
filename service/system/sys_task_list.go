package system

import (
	"errors"

	"github.com/kinggq/global"
	"github.com/kinggq/model/system"
	"gorm.io/gorm"
)

type TaskListService struct{}

func (s *TaskListService) Create(list system.ProjTaskLists) (err error) {
	var l system.ProjProjects
	if errors.Is(global.DB.Where("project_id = ?", list.ProjectId).First(&l).Error, gorm.ErrRecordNotFound) {
		return errors.New("项目不存在")
	}
	err = global.DB.Create(&list).Error
	return err
}
