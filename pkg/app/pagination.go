package app

import (
	"douyin_service/pkg/convert"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context, maxPageSize int) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return 10
	}
	if pageSize > maxPageSize {
		return maxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	ret := 0
	if page > 0 {
		ret = (page - 1) * pageSize
	}
	return ret
}
