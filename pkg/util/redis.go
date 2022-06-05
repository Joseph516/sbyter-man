package util

import (
	"strconv"
)

//------------ reds-favorite的一些工具函数 ------------

const (
	USERPRIFIX            = "U"             //用户点赞列表key前缀
	VIDEOPRIFX            = "V"             //点赞视频的用户列表key前缀
	VIDEOFAVORCNTPRIFIX   = "VC"            //视频被点赞数量key前缀
	VIDEO                 = "VI"            //视频key前缀
	USERTOTALFAVORITEDCNT = "UTDC"          //用户被赞的数量key前缀
	USERTOTALFAVORITECNT  = "UTC"           //用户点赞数量key前缀
	VIDEO_AUTHOR          = "VIDEO_AUTHOR:" //视频对应的作者id的key前缀
	FOLLOWCOUNTFREFIX     = "FOLLOW_COUNT:"
	FANCOUNTFREFIX        = "FAN_COUNT:"
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

//UserFavoritedCntKey 拼接用户被赞数量key
func UserFavoritedCntKey(userId uint) string {
	return USERTOTALFAVORITEDCNT + strconv.FormatInt(int64(userId), 10)
}

// UserFavoriteCntKey 拼接用户点赞数量key
func UserFavoriteCntKey(userId uint) string {
	return USERTOTALFAVORITECNT + strconv.FormatInt(int64(userId), 10)
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
	return FOLLOWCOUNTFREFIX + strconv.Itoa(int(id))
}

func FanCountKey(id uint) string {
	return FANCOUNTFREFIX + strconv.Itoa(int(id))
}

func VideoAuthorKey(videoId uint) string {
	return VIDEO_AUTHOR + strconv.Itoa(int(videoId))
}
