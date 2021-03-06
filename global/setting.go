package global

import (
	"go-blog/pkg/logger"
	"go-blog/pkg/setting"
)

var (
	ServerSetting *setting.ServerSetting
	AppSetting *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	JWTSetting *setting.JWTSetting
	EmailSetting *setting.EmailSetting
	Logger *logger.Logger
)
