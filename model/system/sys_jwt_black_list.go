package system

import "github.com/kinggq/global"

type JwtBlackList struct {
	global.MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
