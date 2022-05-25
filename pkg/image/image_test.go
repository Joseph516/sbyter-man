package image

import (
	"fmt"
	"testing"
)

// 如何在接口中使用
//svc := service.New(c)
//url := "https://tiia-demo-default-1254418846.cos.ap-guangzhou.myqcloud.com/DetectLabel1.jpg"
//tags, _ := image.DetectTag(url) // 获取所有标签
//ids, err := svc.AddTagList(service.TagListRequest{Tags: tags}) // 获取标签对应的id

func TestDetectTag(t *testing.T) {
	url := "https://tiia-demo-default-1254418846.cos.ap-guangzhou.myqcloud.com/DetectLabel1.jpg"
	tags, err := DetectTag(url)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tags)
}
