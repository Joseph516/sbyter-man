package util

import "strconv"

//------------ reds-favorite的一些工具函数 ------------

const (
	USERPRIFIX          = "U"
	VIDEOPRIFX          = "V"
	USERFAVORCNTPRIFIX  = "UC"
	VIDEOFAVORCNTPRIFIX = "VC"
)

// UserFavorKey 拼接userId点赞的key
func UserFavorKey(userId uint) string {
	return USERPRIFIX + strconv.Itoa(int(userId))
}

// VideoFavoredKey 拼接点赞videoId的key
func VideoFavoredKey(videoId uint) string {
	return VIDEOPRIFX + strconv.Itoa(int(videoId))
}

// UserFavorCntKey 拼接userId点赞数量的key
func UserFavorCntKey(userId uint) string {
	return USERFAVORCNTPRIFIX + strconv.Itoa(int(userId))
}

// VideoFavoredCntKey 拼接videoId的点赞数量的key
func VideoFavoredCntKey(videoId uint) string {
	return VIDEOFAVORCNTPRIFIX + strconv.Itoa(int(videoId))
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
