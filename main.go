package main

import (
	"github.com/kinggq/core"
	"github.com/kinggq/global"
	"github.com/kinggq/initialize"
)

type SysUser struct {
	Id       int    `json:"id" gorm:"index;comment:用户唯一ID"`
	Username string `json:"username" gorm:"index;comment:用户名"`
	Password string `json:"password" gorm:"index;comment:用户密码"`
}

func main() {
	global.VP = core.Viper()
	global.DB = initialize.Gorm()

	core.RunServer()
}
