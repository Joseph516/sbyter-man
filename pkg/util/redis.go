package util

import (
	"strconv"
)

//------------ reds-favorite的一些工具函数 ------------

const (
	_USERPRIFIX          = "U"
	_VIDEOPRIFX          = "V"
	_USERFAVORCNTPRIFIX  = "UC"
	_VIDEOFAVORCNTPRIFIX = "VC"
	_VIDEO               = "VI"
)

// UserFavorKey 拼接userId点赞的key
func UserFavorKey(userId int64) string {
	return _USERPRIFIX + strconv.FormatInt(userId, 10)
}

// VideoFavoredKey 拼接点赞videoId的key
func VideoFavoredKey(videoId int64) string {
	return _VIDEOPRIFX + strconv.FormatInt(videoId, 10)
}

// VideoFavorCntKey 拼接videoId点赞的数量
func VideoFavorCntKey(videoId int64) string {
	return _VIDEOFAVORCNTPRIFIX + strconv.FormatInt(videoId, 10)
}

// VideoKey 拼接video的Key
func VideoKey(videoId int64) string {
	return _VIDEO + strconv.FormatInt(videoId, 10)
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
