package v1

import (
	"context"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"douyin_service/pkg/upload"
	"douyin_service/pkg/util"
	"douyin_service/services/api/config"
	"douyin_service/services/api/rpc"
	"douyin_service/services/api/service"
	"douyin_service/services/kitex_gen/user"
	"douyin_service/services/kitex_gen/video"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

type Publish struct {
}

func NewPublish() Publish {
	return Publish{}
}

func (p Publish) List(c *gin.Context) {
	param := service.PublishListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	valid, tokenErr := rpc.ValidToken(context.Background(), &user.TokenRequest{
		Token:     param.Token,
		SkipCheck: errcode.SkipCheckUserID,
	})
	if !valid {
		response.ToErrorResponse(errcode.UnauthorizedTokenError.WithDetails(tokenErr.Error()))
		return
	}
	// 获取视频发布列表
	resp, err := rpc.PublishList(context.Background(), &video.PublishListRequest{
		UserId: int64(param.UserId),
	})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorListPublishFail.WithDetails(err.Error()))
		return
	}

	var res service.PublishListResponse
	res.StatusCode = 0
	res.StatusMsg = "获取视频列表成功"
	// TODO: 需要优化
	res.VideoList = make([]service.VideoInfo, len(resp.VideoList))
	for i := range resp.VideoList {
		res.VideoList[i] = service.VideoInfo{
			Id: resp.VideoList[i].Id,
			Author: service.UserInfo{
				ID:   resp.VideoList[i].Author.Id,
				Name: resp.VideoList[i].Author.Name,
				// FollowCount:   *resp.VideoList[i].Author.FollowCount,
				// FollowerCount: *resp.VideoList[i].Author.FollowerCount,
				IsFollow: resp.VideoList[i].Author.IsFollow,
			},
			PlayUrl:       resp.VideoList[i].PlayUrl,
			CoverUrl:      resp.VideoList[i].CoverUrl,
			FavoriteCount: resp.VideoList[i].FavoriteCount,
			CommentCount:  resp.VideoList[i].CommentCount,
			IsFavorite:    resp.VideoList[i].IsFavorite,
			Title:         resp.VideoList[i].Title,
		}
	}
	response.ToResponse(res)
}

func (p Publish) Action(c *gin.Context) {
	data, _ := c.FormFile("data")
	param := service.PublishActionRequest{
		Data:  data,
		Token: c.PostForm("token"),
		Title: c.PostForm("title"),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 保存视频，不能通过protobuf传递*multipart.FileHeader
	respUser, err := rpc.GetUserByToken(context.Background(), &user.UserTokenRequest{Token: param.Token})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorActionPublishFail.WithDetails(err.Error()))
		return
	}
	userId := respUser.User.Id
	playUrl, coverUrl, err := uploadFile(data, userId)

	// 发布视频
	var res service.ResponseCommon
	_, err = rpc.PublishAction(context.Background(), &video.PublishActionRequest{
		Title:    param.Title,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		UserId:   userId,
	})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorActionPublishFail.WithDetails(err.Error()))
		return
	}

	res.StatusCode = 0
	res.StatusMsg = "视频发布成功"
	response.ToResponse(res)
}

func uploadFile(data *multipart.FileHeader, userId int64) (playUrl string, coverUrl string, err error) {
	// 上传校验
	fileName := upload.GetFileNameWithTime(data.Filename) // 防止重名，文件名+时间戳+MD5
	if !upload.CheckContainExt(upload.TypeVideo, fileName, config.AppCfg.UploadVideoAllowExts) {
		err = fmt.Errorf("文件格式不支持，仅支持格式: %v", config.AppCfg.UploadVideoAllowExts)
		return
	}
	if upload.CheckMaxSizeByHeader(upload.TypeVideo, int(data.Size), config.AppCfg.UploadVideoMaxSize) {
		err = fmt.Errorf("请上传文件大小不超过%v的视频", config.AppCfg.UploadVideoMaxSize)
		return
	}
	uploadSavePath := path.Join(config.AppCfg.UploadSavePath, strconv.Itoa(int(userId)))
	if upload.CheckSavePath(uploadSavePath) {
		if err = upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			err = errors.New("无法创建保存文件夹")
			return
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		err = errors.New("保存路径权限不够")
		return
	}

	// 上传视频
	dst := path.Join(uploadSavePath, fileName)
	if err = upload.SaveFile(data, dst); err != nil {
		return
	}
	playUrl = util.UrlJoin(config.AppCfg.UploadServerUrl, strconv.Itoa(int(userId)), fileName)

	// 获取视频封面并上传
	coverName := fmt.Sprintf("%s.png", upload.GetFilenameWithoutExt(fileName))
	cdst := path.Join(uploadSavePath, coverName)
	if err = upload.ExactCoverFromVideo(dst, cdst); err != nil {
		// 提取封面失败
		if strings.HasSuffix(err.Error(), exec.ErrNotFound.Error()) {
			coverUrl = "https://c-ssl.dtstatic.com/uploads/item/201803/13/20180313083933_olurq.thumb.1000_0.jpg"
		} else {
			return
		}
	} else {
		coverUrl = util.UrlJoin(config.AppCfg.UploadServerUrl, strconv.Itoa(int(userId)), coverName)
	}

	return
}
