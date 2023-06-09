package global

import (
	"github.com/kinggq/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG config.Server
	VP     *viper.Viper
)
