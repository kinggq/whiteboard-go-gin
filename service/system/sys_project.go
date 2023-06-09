package system

import (
	"github.com/kinggq/global"
	"github.com/kinggq/model/system"
)

type ProjectService struct{}

func (s *ProjectService) Create(p system.ProjProjects) (err error) {
	err = global.DB.Create(&p).Error
	return err
}
