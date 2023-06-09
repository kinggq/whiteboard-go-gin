package system

import "github.com/kinggq/global"

type SysUser struct {
	global.MODEL
	Id       int    `json:"id" gorm:"index;comment:用户唯一ID"`
	Username string `json:"username" gorm:"index;comment:用户名"`
	Password string `json:"password" gorm:"index;comment:用户密码"`
	Phone    string `json:"phone" gorm:"index;comment:手机号"`
	NickName string `json:"nike_name" gorm:"index;comment:名称"`
	Avatar   string `json:"avatar" gorm:"index;comment:头像"`
}
