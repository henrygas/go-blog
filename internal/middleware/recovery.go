package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/pkg/app"
	"go-blog/pkg/email"
	"go-blog/pkg/errcode"
	"time"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewMail(&email.SMTPInfo{
		Host: global.EmailSetting.Host,
		Port: global.EmailSetting.Port,
		IsSSL: global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From: global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().WithContext(c).Errorf(c, "panic recover err: %v", err)
				err := defaultMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					global.Logger.WithContext(c).Panicf(c, "mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
			}
		}()
		c.Next()
	}
}
