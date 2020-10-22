package api

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/internal/service"
	"go-blog/pkg/app"
	"go-blog/pkg/errcode"
)

// 从请求中获取
func GetAuth(c *gin.Context) {
	// 从请求中解析auth相关参数
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid == false {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		panic("发生鉴权错误啦，快来看看吧")
		return
	}

	// 鉴定auth相关参数的合法性
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	// 生成token
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	// 将生成的token返回给客户端
	response.ToResponse(gin.H{
		"token": token,
	})
}
