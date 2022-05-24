package util

import (
	"strconv"
)

//------------ reds-favorite的一些工具函数 ------------

const (
	USERPRIFIX          = "U"
	VIDEOPRIFX          = "V"
	USERFAVORCNTPRIFIX  = "UC"
	VIDEOFAVORCNTPRIFIX = "VC"
	VIDEO               = "VI"
)

// UserFavorKey 拼接userId点赞的key
func UserFavorKey(userId int64) string {
	return USERPRIFIX + strconv.FormatInt(userId, 10)
}

// VideoFavoredKey 拼接点赞videoId的key
func VideoFavoredKey(videoId int64) string {
	return VIDEOPRIFX + strconv.FormatInt(videoId, 10)
}

// VideoFavorCntKey 拼接videoId点赞的数量
func VideoFavorCntKey(videoId int64) string {
	return VIDEOFAVORCNTPRIFIX + strconv.FormatInt(videoId, 10)
}

// VideoKey 拼接video的Key
func VideoKey(videoId int64) string {
	return VIDEO + strconv.FormatInt(videoId, 10)
}

// GetListFromBitMap 从bitmap中获取值为1的位数
func GetListFromBitMap(bitmap []byte) []int64 {
	res := make([]int64, 0)
	//byte: 0000_0000
	for i, v := range bitmap {
		for mask, j := uint8(128), 0; j < 8; j++ {
			if v&mask != 0 {
				res = append(res, int64((i<<3)+j))
			}
			mask >>= 1
		}
	}
	return res
}

func FollowCountKey(id int64) string {
	return strconv.Itoa(int(id)) + "_FOLLOW_COUNT"
}

func FanCountKey(id int64) string {
	return strconv.Itoa(int(id)) + "_FAN_COUNT"
}