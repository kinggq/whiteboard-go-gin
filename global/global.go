package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/kinggq/config"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	DB                  *gorm.DB
	REDIS               *redis.Client
	CONFIG              config.Server
	VP                  *viper.Viper
	Concurrency_Control = &singleflight.Group{}
	LOG                 zap.Logger
	BlackCache          local_cache.Cache
)
