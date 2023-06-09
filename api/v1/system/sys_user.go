package system

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kinggq/global"
	"github.com/kinggq/model/common/response"
	"github.com/kinggq/model/system"
	req "github.com/kinggq/model/system/request"
	res "github.com/kinggq/model/system/response"
	"github.com/kinggq/utils"
)

type BaseApi struct{}

func (b *BaseApi) Login(ctx *gin.Context) {
	var u system.SysUser
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = utils.Verify(u, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	uu := system.SysUser{Username: u.Username, Password: u.Password}
	user, err := userService.Login(uu)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	b.TokenNext(ctx, *user)
}

func (b *BaseApi) TokenNext(ctx *gin.Context, u system.SysUser) {
	j := utils.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)}
	claims := j.CreateClaims(req.BaseClaims{
		Username: u.Username,
		ID:       u.ID,
		NickName: u.NickName,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage("获取Token失败", ctx)
		return
	}

	user := res.SysUserResponse{
		Username: u.Username,
		Id:       u.Id,
		Phone:    u.Phone,
		Avatar:   u.Avatar,
		NickName: u.NickName,
	}

	response.OkWithDetailed(res.LoginResponse{
		Token:     token,
		User:      user,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登录成功", ctx)
}

func (b *BaseApi) Register(ctx *gin.Context) {
	var r req.Register
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	user := &system.SysUser{
		Username: r.Phone,
		Phone:    r.Phone,
		Password: r.Password,
	}
	userReturn, err := userService.Register(*user)

	u := res.SysUserResponse{
		Username: userReturn.Username,
		Id:       userReturn.Id,
		Phone:    userReturn.Phone,
		Avatar:   userReturn.Avatar,
		NickName: userReturn.NickName,
	}
	if err != nil {
		fmt.Println("user created is error: ", err)
		response.FailWithDetailed(u, err.Error(), ctx)
		return
	}
	response.OkWithDetailed(u, "注册成功", ctx)
}
