package app

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	val "github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []ValidError

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		fieldErrMsg := fmt.Sprintf("parameter %v is invalid", err.Key)
		errs = append(errs, fieldErrMsg)
	}
	return errs
}

// BindAndValid 参数绑定和校验
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	// 参数绑定和入参校验
	err := c.ShouldBind(v)
	if err != nil {
		// 获取错误消息体
		var ve val.ValidationErrors
		if errors.As(err, &ve) {
			errs := make(ValidErrors, len(ve))
			for i, fe := range ve {
				errs[i] = ValidError{fe.Field(), fe.Error()}
			}
			return false, errs
		}
		return false, nil
	}
	return true, nil
}
