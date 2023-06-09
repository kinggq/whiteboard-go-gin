package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ERROR_CODE   = -5
	SUCCESS_CODE = 0
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR_CODE, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR_CODE, data, message, c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS_CODE, data, message, c)
}
