package system

import (
	"errors"

	"github.com/kinggq/global"
	"github.com/kinggq/model/system"
	"gorm.io/gorm"
)

type TaskService struct{}

func (s *TaskService) Create(params system.ProjTasks) (err error) {
	var task system.ProjTaskLists
	if errors.Is(global.DB.Where("list_id = ?", params.ListId).First(&task).Error, gorm.ErrRecordNotFound) {
		return errors.New("列表不存在")
	}
	err = global.DB.Create(&params).Error
	return err
}
