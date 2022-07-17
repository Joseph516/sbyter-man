package main

import (
	"douyin_service/services/user/config"
	user "douyin_service/services/user/kitex_gen/user/userservice"
	"douyin_service/services/user/repository"
	"log"
)

func init() {
	// 初始化配置
	cfg, err := config.NewViperConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.ReadSection("JWT", &config.JwtCfg)
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.ReadSection("Database", &config.DatabaseCfg)
	if err != nil {
		log.Fatal(err)
	}

	// 初始化数据库
	if err := repository.Init(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
