package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// RainbowFart 彩虹屁结构体
type RainbowFart struct {
	Data `json:"data"`
}

type Data struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type BackgroundImg struct {
	Error  int    `json:"error"`
	Result int    `json:"result"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Img    string `json:"img"`
}

// RandomAvatar 获取一些随机头像
func RandomAvatar(id string) string {
	url := fmt.Sprintf("https://api.multiavatar.com/%s.png?apikey=A5wbsoJPETy1uk", id)
	return url
}

// RandomSign 获取一些随机个签
func RandomSign() (string, error) {
	url := "https://api.shadiao.app/chp"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(make([]byte, 0, 512))
	_, err = buf.ReadFrom(resp.Body)

	if err != nil {
		return "", err
	}

	rainbowFart := RainbowFart{}
	err = json.Unmarshal(buf.Bytes(), &rainbowFart)
	if err != nil {
		return "", err
	}

	return rainbowFart.Text, nil
}

//RandomBackground 获得一张随机背景图
func RandomBackground() (string, error) {
	url := "https://tuapi.eees.cc/api.php?category={fengjing}&type=json"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(make([]byte, 0, 512))
	_, err = buf.ReadFrom(resp.Body)

	if err != nil {
		return "", err
	}

	backgroundImg := BackgroundImg{}
	err = json.Unmarshal(buf.Bytes(), &backgroundImg)
	if err != nil {
		return "", err
	}

	return backgroundImg.Img, nil
}