package client

import (
	"douyin_service/services/video/kitex_gen/video/videoservice"
	"fmt"
	client2 "github.com/cloudwego/kitex/client"
	"log"
)

func main() {
	client, err := videoservice.NewClient("videoservice", client2.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client)
	// TODO: do something with client
}
