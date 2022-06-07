package app

import (
	"douyin_service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalRows int64 `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{})  {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int64)  {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"code":http.StatusOK,
		"message": "获取列表成功",
		"list": list,
		"pager": Pager{
			Page: GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error)  {
	//response := gin.H{"status_code": err.Code(), "status_msg": err.Msg()}
	response := gin.H{"status_code": 0, "status_msg": err.Msg()} // 由于前端设计，这里只能用0
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
