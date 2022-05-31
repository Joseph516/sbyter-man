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
	FOLLOWCOUNTFREFIX   = "FOLLOW_COUNT:"
	FANCOUNTFREFIX      = "FAN_COUNT:"
)

// UserFavorKey 拼接userId点赞的key
func UserFavorKey(userId uint) string {
	return USERPRIFIX + strconv.FormatInt(int64(userId), 10)
}

// VideoFavoredKey 拼接点赞videoId的key
func VideoFavoredKey(videoId uint) string {
	return VIDEOPRIFX + strconv.FormatInt(int64(videoId), 10)
}

// VideoFavorCntKey 拼接videoId点赞的数量
func VideoFavorCntKey(videoId uint) string {
	return VIDEOFAVORCNTPRIFIX + strconv.FormatInt(int64(videoId), 10)
}

// VideoKey 拼接video的Key
func VideoKey(videoId uint) string {
	return VIDEO + strconv.FormatInt(int64(videoId), 10)
}

// GetListFromBitMap 从bitmap中获取值为1的位数
func GetListFromBitMap(bitmap []byte) []uint {
	res := make([]uint, 0)
	//byte: 0000_0000
	for i, v := range bitmap {
		for mask, j := uint8(128), 0; j < 8; j++ {
			if v&mask != 0 {
				res = append(res, uint((i<<3)+j))
			}
			mask >>= 1
		}
	}
	return res
}

func FollowCountKey(id uint) string {
	return   FOLLOWCOUNTFREFIX + strconv.Itoa(int(id))
}

func FanCountKey(id uint) string {
	return FANCOUNTFREFIX +  strconv.Itoa(int(id))
}
