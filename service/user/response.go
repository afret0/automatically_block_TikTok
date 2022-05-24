package user

import (
	"github.com/gin-gonic/gin"
)

type ResponseManager struct {
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var responseManager *ResponseManager

func init() {
	responseManager = new(ResponseManager)
}

func (r ResponseManager) ReturnSucceedResponse(ctx *gin.Context, data interface{}) {
	succeedResponse := new(Response)
	succeedResponse.Code = 1
	succeedResponse.Message = "succeed"
	succeedResponse.Data = data
	ctx.JSON(200, succeedResponse)
}

func (r ResponseManager) ReturnFailedResponse(ctx *gin.Context, msg string) {
	s := new(Response)
	s.Code = 0
	s.Message = msg
	ctx.JSON(200, s)
}
