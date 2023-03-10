package main

import (
	config2 "douyin_service/pkg/config"
	"douyin_service/services/api/config"
	"douyin_service/services/api/handlers"
	"douyin_service/services/api/middleware"
	"douyin_service/services/api/rpc"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"log"
	"net/http"
	"time"
)

func setupSetting() error {
	// 初始化配置
	cfg, err := config2.NewViperConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.ReadSection("Server", &config.ServerCfg)
	if err != nil {
		return err
	}
	config.ServerCfg.ReadTimeout *= time.Second
	config.ServerCfg.WriteTimeout *= time.Second

	err = cfg.ReadSection("App", &config.AppCfg)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	// middleware
	err = middleware.SetupLogger(config.AppCfg.LogSavePath, config.AppCfg.LogFileName)
	if err != nil {
		log.Fatalf("middleware.SetupLogger err: %v", err)
	}

	// init rpc
	rpc.NewUserClient()
	rpc.NewVideoClient()
}

func main() {
	gin.SetMode(config.ServerCfg.RunMode)
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()
	router := handlers.NewRouter()
	s := &http.Server{
		Addr:           ":" + config.ServerCfg.HttpPort,
		Handler:        router,
		ReadTimeout:    config.ServerCfg.ReadTimeout,
		WriteTimeout:   config.ServerCfg.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
