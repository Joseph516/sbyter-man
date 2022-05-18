package service

import (
	"douyin_service/global"
	"strconv"
)

const (
	USERPRIFIX          = "U"
	VIDEOPRIFX          = "V"
	USERFAVORCNTPRIFIX  = "UC"
	VIDEOFAVORCNTPRIFIX = "CV"
)

type ActionRequest struct {
	UserId     uint   `json:"user_id" form:"user_id" binding:"required"`
	Token      string `json:"token" form:"token" binding:"required"`
	VideoId    uint   `json:"video_id" form:"video_id" binding:"required"`
	ActionType int    `json:"action_type" form:"action_type" binding:"required"`
}

type ActionResponse struct {
	ResponseCommon
}

type FavoriteListRequest struct {
	UserId uint   `json:"user_id" form:"user_id" binding:"required"`
	Token  string `json:"token" form:"token" binding:"required"`
}

func (svc *Service) Action(param *ActionRequest) error {
	userId := param.UserId
	videoId := param.VideoId
	actionType := param.ActionType

	if actionType == 1 {
		return favor(userId, videoId)
	}

	return cancelFavor(userId, videoId)
}

// FavoriteList 返回(用户点赞列表，err)
func (svc *Service) FavoriteList(param *FavoriteListRequest) ([]uint, error) {
	//查询该用户点赞作品id
	userId := USERPRIFIX + strconv.Itoa(int(param.UserId))
	//返回点赞的二进制，下标从左到右！
	favorStr, err := global.Rd.Get(userId).Result()
	if err != nil {
		return nil, err
	}
	favorByte := []byte(favorStr)
	favorList := getIndexFromByte(favorByte)
	return favorList, nil
}

// userId对videoId点赞
func favor(userId uint, videoId uint) error {
	//已经点赞了直接返回
	ok, err := isFavor(userId, videoId)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}

	user := USERPRIFIX + strconv.Itoa(int(userId))
	video := VIDEOPRIFX + strconv.Itoa(int(videoId))
	usercnt := USERFAVORCNTPRIFIX + strconv.Itoa(int(userId))
	videocnt := VIDEOFAVORCNTPRIFIX + strconv.Itoa(int(videoId))

	err = global.Rd.SetBit(user, int64(videoId), 1).Err()
	if err != nil {
		return err
	}

	err = global.Rd.Incr(usercnt).Err()
	if err != nil {
		return err
	}

	err = global.Rd.SetBit(video, int64(userId), 1).Err()
	if err != nil {
		return err
	}

	err = global.Rd.Incr(videocnt).Err()
	if err != nil {
		return err
	}
	return nil
}

// userId取消点赞videoId
func cancelFavor(userId uint, videoId uint) error {
	//未点赞直接返回
	ok, err := isFavor(userId, videoId)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	user := USERPRIFIX + strconv.Itoa(int(userId))
	video := VIDEOPRIFX + strconv.Itoa(int(videoId))
	usercnt := USERFAVORCNTPRIFIX + strconv.Itoa(int(userId))
	videocnt := VIDEOFAVORCNTPRIFIX + strconv.Itoa(int(videoId))

	err = global.Rd.SetBit(user, int64(videoId), 0).Err()
	if err != nil {
		return err
	}

	err = global.Rd.Decr(usercnt).Err()
	if err != nil {
		return err
	}

	err = global.Rd.SetBit(video, int64(userId), 0).Err()
	if err != nil {
		return err
	}

	err = global.Rd.Decr(videocnt).Err()
	if err != nil {
		return err
	}
	return nil
}

// 根据给定的byte数组找到为1的下标
func getIndexFromByte(favor []byte) []uint {
	res := make([]uint, 0)
	//byte: 0000_0000
	for i, v := range favor {
		for mask, j := uint8(128), 0; j < 8; j++ {
			if v&mask != 0 {
				res = append(res, uint((i<<3)+j))
			}
			mask >>= 1
		}
	}
	return res
}

// userId是否点赞了videoId
func isFavor(userId uint, videoId uint) (bool, error) {
	user := USERPRIFIX + strconv.Itoa(int(userId))
	result, err := global.Rd.GetBit(user, int64(videoId)).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

// videoId是否被userId点赞
func isFavored(videoId uint, userId uint) (bool, error) {
	video := VIDEOPRIFX + strconv.Itoa(int(videoId))
	result, err := global.Rd.GetBit(video, int64(userId)).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

// 返回某个用户的点赞总数
func getUserFavorCnt(userId uint) (uint, error) {
	user := USERFAVORCNTPRIFIX + strconv.Itoa(int(userId))
	result, err := global.Rd.Get(user).Result()
	if err != nil {
		return 0, err
	}
	cnt, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}
	return uint(cnt), err
}

// 返回某个视频的被赞总数
func getVideoFavoredCnt(videoId uint) (uint, error) {
	video := VIDEOFAVORCNTPRIFIX + strconv.Itoa(int(videoId))
	result, err := global.Rd.Get(video).Result()
	if err != nil {
		return 0, err
	}
	cnt, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}
	return uint(cnt), err
}
