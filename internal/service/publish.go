package service

import (
	"douyin_service/global"
	"douyin_service/internal/model"
	"douyin_service/pkg/oss"
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
	"sync"
	"time"
)

type PublishListRequest struct {
	UserId uint   `form:"user_id"` //考虑到客户端在未登录下仍会尝试调用publish/list,因此将UserId和Token的required标签取消了
	Token  string `form:"token"`
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

var video_lock sync.Mutex

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
		isFollow, _ := svc.dao.IsFollow(userId, video[i].AuthorId)
		pubResp.VideoList[i] = VideoInfo{
			Id: video[i].ID,
			Author: UserInfo{
				ID:            user.ID,
				Name:          user.UserName,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      isFollow,
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
		coverUrl = util.UrlJoin(global.AppSetting.UploadServerUrl, global.AppSetting.UploadSavePath, strconv.Itoa(int(userId)), coverName)
	}

	// 更新数据库, 检查一下底下两个路径是否正确
	imgPath := util.UrlJoin(global.AppSetting.UploadSavePath, strconv.Itoa(int(userId)), coverName)
	videoPath := util.UrlJoin(global.AppSetting.UploadSavePath, strconv.Itoa(int(userId)), fileName)
	pre := "https://sbyterman.oss-cn-hangzhou.aliyuncs.com/" // OSS地址前缀
	err := oss.UploadOSS(videoPath, "video/" + fileName) // 上传视频
	if err == nil {
		playUrl = pre + "video/" + fileName

	}
	err = oss.UploadOSS(imgPath, "img/" + coverName)  //上传封面
	if err == nil {
		coverUrl = pre + "img/"+ coverName
	}

	// 下面的playUrl和coverUrl换成OSS地址
	err = svc.dao.PublishVideo(userId, playUrl, coverUrl, title)

	return err
}

func (svc *Service) QueryBatchVdieoById(favorList []uint) ([]model.Video, error) {
	return svc.dao.QueryBatchVideoById(favorList)
}

// QueryAuthorIdByVideoId 根据videoId查询authorId的缓存查找方法
func (svc *Service) QueryAuthorIdByVideoId(videoId uint) (uint, error) {
	exist, authorId, err := svc.redis.QueryAuthorIdByVideoId(videoId)
	if err != nil {
		return 0, err
	}
	if exist {
		return authorId, nil
	}
	key := util.VideoAuthorKey(videoId)
	video_lock.Lock()
	defer video_lock.Unlock()
	exist, authorId, err = svc.redis.QueryAuthorIdByVideoId(videoId)
	if err != nil {
		return 0, err
	}
	if exist {
		return authorId, nil
	} else {
		video, err := svc.dao.QueryVideoById(videoId)
		if err != nil {
			return 0, err
		}
		authorId = video.AuthorId
		err = svc.redis.Set(key, authorId, time.Hour*12)
		if err != nil {
			return 0, err
		}
		return authorId, nil
	}
}
