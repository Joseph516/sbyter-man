package main

import (
	config2 "douyin_service/pkg/config"
	user "douyin_service/services/kitex_gen/user/userservice"
	"douyin_service/services/user/config"
	"douyin_service/services/user/repository"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func init() {
	// 初始化配置
	cfg, err := config2.NewViperConfig("config/config.yaml")
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
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	svr := user.NewServer(
		new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user-service"}),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
