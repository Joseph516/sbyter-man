package service

import (
	"douyin_service/global"
	"douyin_service/internal/model"
	"douyin_service/pkg/tag"
	"douyin_service/pkg/upload"
	"douyin_service/pkg/util"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

type PublishListRequest struct {
	UserId uint   `form:"user_id"  binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type VideoInfo struct {
	Id            uint     `json:"id"`
	Author        UserInfo `json:"author"`
	PlayUrl       string   `json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	CommentCount  int64    `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   `json:"title"`
}

type PublishListResponse struct {
	ResponseCommon
	VideoList []VideoInfo `json:"video_list"`
}

type PublishActionRequest struct {
	Data  *multipart.FileHeader `form:"data" binding:"required"`
	Token string                `form:"token" binding:"required"`
	Title string                `form:"title" binding:"required"`
}

func (svc *Service) PublishList(userId uint) (pubResp PublishListResponse, err error) {
	// 根据用户id获取发布视频信息
	video, err := svc.dao.ListVideoByUserId(userId)
	if err != nil {
		return
	}

	// 根据用户id获取用户自身信息
	user, err := svc.dao.GetUserById(userId)
	if err != nil {
		return
	}

	// 遍历赋值
	pubResp.VideoList = make([]VideoInfo, len(video))
	for i := range video {
		isFavorite, _ := svc.IsFavor(userId, video[i].ID)
		favoriteCnt, _ := svc.QueryFavorCnt(video[i].ID)
		pubResp.VideoList[i] = VideoInfo{
			Id: video[i].ID,
			Author: UserInfo{
				ID:            user.ID,
				Name:          user.UserName,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      false,
			},
			PlayUrl:       video[i].PlayUrl,
			CoverUrl:      video[i].CoverUrl,
			FavoriteCount: favoriteCnt,
			CommentCount:  video[i].CommentCount,
			IsFavorite:    isFavorite,
			Title:         video[i].Title,
		}
	}
	return
}

func (svc *Service) PublishAction(data *multipart.FileHeader, token, title string, userId uint) error {
	// 上传校验
	fileName := upload.GetFileNameWithTime(data.Filename) // 防止重名，文件名+时间戳+MD5
	if !upload.CheckContainExt(upload.TypeVideo, fileName) {
		return fmt.Errorf("文件格式不支持，仅支持格式: %v", global.AppSetting.UploadVideoAllowExts)
	}
	if upload.CheckMaxSizeByHeader(upload.TypeVideo, int(data.Size)) {
		return fmt.Errorf("请上传文件大小不超过%v的视频", global.AppSetting.UploadVideoMaxSize)
	}
	uploadSavePath := path.Join(upload.GetSavePath(), strconv.Itoa(int(userId)))
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return errors.New("无法创建保存文件夹")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return errors.New("保存路径权限不够")
	}

	// 上传视频
	dst := path.Join(uploadSavePath, fileName)
	if err := upload.SaveFile(data, dst); err != nil {
		return err
	}
	playUrl := util.UrlJoin(global.AppSetting.UploadServerUrl, strconv.Itoa(int(userId)), fileName)

	// 获取视频封面并上传
	coverName := fmt.Sprintf("%s.png", upload.GetFilenameWithoutExt(fileName))
	cdst := path.Join(uploadSavePath, coverName)
	var coverUrl string
	if err := upload.ExactCoverFromVideo(dst, cdst); err != nil {
		// 提取封面失败
		if strings.HasSuffix(err.Error(), exec.ErrNotFound.Error()) {
			coverUrl = "https://c-ssl.dtstatic.com/uploads/item/201803/13/20180313083933_olurq.thumb.1000_0.jpg"
		} else {
			return err
		}
	} else {
		coverUrl = util.UrlJoin(global.AppSetting.UploadServerUrl, strconv.Itoa(int(userId)), coverName)
	}

	// 获取视频标签信息
	tags, _ := tag.ExtractTagFromText(title)

	// 更新数据库
	err := svc.dao.PublishVideoWithTag(userId, playUrl, coverUrl, title, tags)

	return err
}

func (svc *Service) QueryBatchVdieoById(favorList []uint) ([]model.Video, error) {
	return svc.dao.QueryBatchVideoById(favorList)
}
