package main

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/internal/model"
	"go-blog/internal/routers"
	"go-blog/pkg/logger"
	"go-blog/pkg/setting"
	"go-blog/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 编程之旅：一起用Go做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":" + global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	global.AppSetting.DefaultContextTimeout *= time.Second

	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	global.JWTSetting.Expire *= time.Second

	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize: 600,
		MaxAge: 10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "localhost:6831")
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}