// Code generated by Kitex v0.3.2. DO NOT EDIT.

package videoservice

import (
	"context"
	"douyin_service/services/kitex_gen/video"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"PublishAction":  kitex.NewMethodInfo(publishActionHandler, newPublishActionArgs, newPublishActionResult, false),
		"PublishList":    kitex.NewMethodInfo(publishListHandler, newPublishListArgs, newPublishListResult, false),
		"Feed":           kitex.NewMethodInfo(feedHandler, newFeedArgs, newFeedResult, false),
		"FavoriteAction": kitex.NewMethodInfo(favoriteActionHandler, newFavoriteActionArgs, newFavoriteActionResult, false),
		"FavoriteList":   kitex.NewMethodInfo(favoriteListHandler, newFavoriteListArgs, newFavoriteListResult, false),
		"CommentAction":  kitex.NewMethodInfo(commentActionHandler, newCommentActionArgs, newCommentActionResult, false),
		"CommentList":    kitex.NewMethodInfo(commentListHandler, newCommentListArgs, newCommentListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.3.2",
		Extra:           extra,
	}
	return svcInfo
}

func publishActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.PublishActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).PublishAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishActionArgs:
		success, err := handler.(video.VideoService).PublishAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishActionResult)
		realResult.Success = success
	}
	return nil
}
func newPublishActionArgs() interface{} {
	return &PublishActionArgs{}
}

func newPublishActionResult() interface{} {
	return &PublishActionResult{}
}

type PublishActionArgs struct {
	Req *video.PublishActionRequest
}

func (p *PublishActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishActionArgs) Unmarshal(in []byte) error {
	msg := new(video.PublishActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishActionArgs_Req_DEFAULT *video.PublishActionRequest

func (p *PublishActionArgs) GetReq() *video.PublishActionRequest {
	if !p.IsSetReq() {
		return PublishActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type PublishActionResult struct {
	Success *video.PublishActionResponse
}

var PublishActionResult_Success_DEFAULT *video.PublishActionResponse

func (p *PublishActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishActionResult) Unmarshal(in []byte) error {
	msg := new(video.PublishActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishActionResult) GetSuccess() *video.PublishActionResponse {
	if !p.IsSetSuccess() {
		return PublishActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.PublishActionResponse)
}

func (p *PublishActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func publishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.PublishListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).PublishList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishListArgs:
		success, err := handler.(video.VideoService).PublishList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishListResult)
		realResult.Success = success
	}
	return nil
}
func newPublishListArgs() interface{} {
	return &PublishListArgs{}
}

func newPublishListResult() interface{} {
	return &PublishListResult{}
}

type PublishListArgs struct {
	Req *video.PublishListRequest
}

func (p *PublishListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishListArgs) Unmarshal(in []byte) error {
	msg := new(video.PublishListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishListArgs_Req_DEFAULT *video.PublishListRequest

func (p *PublishListArgs) GetReq() *video.PublishListRequest {
	if !p.IsSetReq() {
		return PublishListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishListArgs) IsSetReq() bool {
	return p.Req != nil
}

type PublishListResult struct {
	Success *video.PublishListResponse
}

var PublishListResult_Success_DEFAULT *video.PublishListResponse

func (p *PublishListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishListResult) Unmarshal(in []byte) error {
	msg := new(video.PublishListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishListResult) GetSuccess() *video.PublishListResponse {
	if !p.IsSetSuccess() {
		return PublishListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishListResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.PublishListResponse)
}

func (p *PublishListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.FeedRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).Feed(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FeedArgs:
		success, err := handler.(video.VideoService).Feed(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FeedResult)
		realResult.Success = success
	}
	return nil
}
func newFeedArgs() interface{} {
	return &FeedArgs{}
}

func newFeedResult() interface{} {
	return &FeedResult{}
}

type FeedArgs struct {
	Req *video.FeedRequest
}

func (p *FeedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FeedArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FeedArgs) Unmarshal(in []byte) error {
	msg := new(video.FeedRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FeedArgs_Req_DEFAULT *video.FeedRequest

func (p *FeedArgs) GetReq() *video.FeedRequest {
	if !p.IsSetReq() {
		return FeedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FeedArgs) IsSetReq() bool {
	return p.Req != nil
}

type FeedResult struct {
	Success *video.FeedResponse
}

var FeedResult_Success_DEFAULT *video.FeedResponse

func (p *FeedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FeedResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FeedResult) Unmarshal(in []byte) error {
	msg := new(video.FeedResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FeedResult) GetSuccess() *video.FeedResponse {
	if !p.IsSetSuccess() {
		return FeedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FeedResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.FeedResponse)
}

func (p *FeedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.FavoriteActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).FavoriteAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteActionArgs:
		success, err := handler.(video.VideoService).FavoriteAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteActionResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteActionArgs() interface{} {
	return &FavoriteActionArgs{}
}

func newFavoriteActionResult() interface{} {
	return &FavoriteActionResult{}
}

type FavoriteActionArgs struct {
	Req *video.FavoriteActionRequest
}

func (p *FavoriteActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteActionArgs) Unmarshal(in []byte) error {
	msg := new(video.FavoriteActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteActionArgs_Req_DEFAULT *video.FavoriteActionRequest

func (p *FavoriteActionArgs) GetReq() *video.FavoriteActionRequest {
	if !p.IsSetReq() {
		return FavoriteActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type FavoriteActionResult struct {
	Success *video.FavoriteActionResponse
}

var FavoriteActionResult_Success_DEFAULT *video.FavoriteActionResponse

func (p *FavoriteActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteActionResult) Unmarshal(in []byte) error {
	msg := new(video.FavoriteActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteActionResult) GetSuccess() *video.FavoriteActionResponse {
	if !p.IsSetSuccess() {
		return FavoriteActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.FavoriteActionResponse)
}

func (p *FavoriteActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.FavoriteListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).FavoriteList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteListArgs:
		success, err := handler.(video.VideoService).FavoriteList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteListResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteListArgs() interface{} {
	return &FavoriteListArgs{}
}

func newFavoriteListResult() interface{} {
	return &FavoriteListResult{}
}

type FavoriteListArgs struct {
	Req *video.FavoriteListRequest
}

func (p *FavoriteListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteListArgs) Unmarshal(in []byte) error {
	msg := new(video.FavoriteListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteListArgs_Req_DEFAULT *video.FavoriteListRequest

func (p *FavoriteListArgs) GetReq() *video.FavoriteListRequest {
	if !p.IsSetReq() {
		return FavoriteListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FavoriteListResult struct {
	Success *video.FavoriteListResponse
}

var FavoriteListResult_Success_DEFAULT *video.FavoriteListResponse

func (p *FavoriteListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteListResult) Unmarshal(in []byte) error {
	msg := new(video.FavoriteListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteListResult) GetSuccess() *video.FavoriteListResponse {
	if !p.IsSetSuccess() {
		return FavoriteListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteListResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.FavoriteListResponse)
}

func (p *FavoriteListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.CommentActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).CommentAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CommentActionArgs:
		success, err := handler.(video.VideoService).CommentAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CommentActionResult)
		realResult.Success = success
	}
	return nil
}
func newCommentActionArgs() interface{} {
	return &CommentActionArgs{}
}

func newCommentActionResult() interface{} {
	return &CommentActionResult{}
}

type CommentActionArgs struct {
	Req *video.CommentActionRequest
}

func (p *CommentActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CommentActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CommentActionArgs) Unmarshal(in []byte) error {
	msg := new(video.CommentActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CommentActionArgs_Req_DEFAULT *video.CommentActionRequest

func (p *CommentActionArgs) GetReq() *video.CommentActionRequest {
	if !p.IsSetReq() {
		return CommentActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CommentActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type CommentActionResult struct {
	Success *video.CommentActionResponse
}

var CommentActionResult_Success_DEFAULT *video.CommentActionResponse

func (p *CommentActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CommentActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CommentActionResult) Unmarshal(in []byte) error {
	msg := new(video.CommentActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CommentActionResult) GetSuccess() *video.CommentActionResponse {
	if !p.IsSetSuccess() {
		return CommentActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CommentActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.CommentActionResponse)
}

func (p *CommentActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func commentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.CommentListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).CommentList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CommentListArgs:
		success, err := handler.(video.VideoService).CommentList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CommentListResult)
		realResult.Success = success
	}
	return nil
}
func newCommentListArgs() interface{} {
	return &CommentListArgs{}
}

func newCommentListResult() interface{} {
	return &CommentListResult{}
}

type CommentListArgs struct {
	Req *video.CommentListRequest
}

func (p *CommentListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CommentListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CommentListArgs) Unmarshal(in []byte) error {
	msg := new(video.CommentListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CommentListArgs_Req_DEFAULT *video.CommentListRequest

func (p *CommentListArgs) GetReq() *video.CommentListRequest {
	if !p.IsSetReq() {
		return CommentListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CommentListArgs) IsSetReq() bool {
	return p.Req != nil
}

type CommentListResult struct {
	Success *video.CommentListResponse
}

var CommentListResult_Success_DEFAULT *video.CommentListResponse

func (p *CommentListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CommentListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CommentListResult) Unmarshal(in []byte) error {
	msg := new(video.CommentListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CommentListResult) GetSuccess() *video.CommentListResponse {
	if !p.IsSetSuccess() {
		return CommentListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CommentListResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.CommentListResponse)
}

func (p *CommentListResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PublishAction(ctx context.Context, Req *video.PublishActionRequest) (r *video.PublishActionResponse, err error) {
	var _args PublishActionArgs
	_args.Req = Req
	var _result PublishActionResult
	if err = p.c.Call(ctx, "PublishAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishList(ctx context.Context, Req *video.PublishListRequest) (r *video.PublishListResponse, err error) {
	var _args PublishListArgs
	_args.Req = Req
	var _result PublishListResult
	if err = p.c.Call(ctx, "PublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Feed(ctx context.Context, Req *video.FeedRequest) (r *video.FeedResponse, err error) {
	var _args FeedArgs
	_args.Req = Req
	var _result FeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteAction(ctx context.Context, Req *video.FavoriteActionRequest) (r *video.FavoriteActionResponse, err error) {
	var _args FavoriteActionArgs
	_args.Req = Req
	var _result FavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteList(ctx context.Context, Req *video.FavoriteListRequest) (r *video.FavoriteListResponse, err error) {
	var _args FavoriteListArgs
	_args.Req = Req
	var _result FavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentAction(ctx context.Context, Req *video.CommentActionRequest) (r *video.CommentActionResponse, err error) {
	var _args CommentActionArgs
	_args.Req = Req
	var _result CommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentList(ctx context.Context, Req *video.CommentListRequest) (r *video.CommentListResponse, err error) {
	var _args CommentListArgs
	_args.Req = Req
	var _result CommentListResult
	if err = p.c.Call(ctx, "CommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
