package system

import (
	"context"

	"github.com/kinggq/global"
	"github.com/kinggq/model/system"
	"github.com/kinggq/utils"
)

type JwtService struct{}

func (j *JwtService) JsonInBlackList(jwtList system.JwtBlackList) (err error) {
	err = global.DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

func (j *JwtService) GetRedisJWT(username string) (redisJWT string, err error) {
	redisJWT, err = global.REDIS.Get(context.Background(), username).Result()
	return redisJWT, err
}

func (j *JwtService) SetRedisJWT(jwt string, username string) (err error) {
	dr, err := utils.ParseDuration(global.CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.REDIS.Set(context.Background(), username, jwt, timer).Err()
	return err
}
