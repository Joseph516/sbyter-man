package main

import (
	"context"
	"douyin_service/services/user/kitex_gen/user"
	"douyin_service/services/user/kitex_gen/user/userservice"
	client2 "github.com/cloudwego/kitex/client"
	"log"
)

func main() {
	client, err := userservice.NewClient("userservice", client2.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &user.UserRegisterRequest{
		Username: "hyh",
		Password: "hyh123456",
	}
	resp, err := client.Register(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
