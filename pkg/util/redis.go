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
)

// UserFavorKey 拼接userId点赞的key
func UserFavorKey(userId int64) string {
	return USERPRIFIX + strconv.FormatInt(userId, 10)
}

// VideoFavoredKey 拼接点赞videoId的key
func VideoFavoredKey(videoId int64) string {
	return VIDEOPRIFX + strconv.FormatInt(videoId, 10)
}

// UserFavorCntKey 拼接userId点赞数量的key
func UserFavorCntKey(userId int64) string {
	return USERFAVORCNTPRIFIX + strconv.FormatInt(userId, 10)
}

// VideoFavoredCntKey 拼接videoId的点赞数量的key
func VideoFavoredCntKey(videoId int64) string {
	return VIDEOFAVORCNTPRIFIX + strconv.FormatInt(videoId, 10)
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
