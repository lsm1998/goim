package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct {
	Context *gin.Context
	Errors  error
}

func (e *Api) AddError(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else if err != nil {
		e.Errors = fmt.Errorf("%v; %w", e.Errors, err)
	}
}

// Error 通常错误数据处理
func (e Api) Error(code int, err error, msg string) {
	Error(e.Context, code, msg)
}

// OK 通常成功数据处理
func (e Api) OK(data interface{}, msg string) {
	OK(e.Context, data, msg)
}

// Error 失败数据处理
func Error(c *gin.Context, code int, msg string) {
	result := map[string]interface{}{}
	result["code"] = code
	result["msg"] = msg
	c.AbortWithStatusJSON(http.StatusOK, result)
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	result := map[string]interface{}{}
	result["code"] = http.StatusOK
	result["msg"] = msg
	result["data"] = data
	c.AbortWithStatusJSON(http.StatusOK, result)
}
