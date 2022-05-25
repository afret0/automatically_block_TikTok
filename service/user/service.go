package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"service/source"
	"service/source/tool"
)

var svr *Service

type Service struct {
	logger *logrus.Logger
	tool   *tool.Tool
}

func init() {
	svr = new(Service)
	svr.logger = source.GetLogger()
	svr.tool = tool.GetTool()
}

func (s *Service) RegisterUser(ctx *gin.Context) {
	registerInformation := new(RegisterInformation)
	err := ctx.Bind(registerInformation)
	if err != nil {
		s.logger.Errorln(err)
		return
	}
	token, err := man.RegisterUser(ctx, registerInformation.Name, registerInformation.Email, registerInformation.Password, registerInformation.VerificationCode)
	if err != nil {
		//s.logger.Errorln(registerInformation.Email, err)
		responseManager.ReturnFailedResponse(ctx, s.tool.SprintfErr(err))
		return
	}
	responseManager.ReturnSucceedResponse(ctx, map[string]string{"token": token})
}

//func (s *Service) Login(ctx *gin.Context) {
//	token, err := man.Login(ctx.Query("phone"), ctx.Query("vCode"))
//	if err != nil {
//		s.res.NewResWithoutData(ctx, code.Failed, err.Error())
//		return
//	}
//	s.res.NewSucceedRes(ctx, map[string]string{"token": token})
//}

func (s *Service) SendVerificationCode(ctx *gin.Context) {
	email := ctx.Query("email")
	err := man.SendVerificationCode(email)
	if err != nil {
		s.logger.Errorln(email, err)
		responseManager.ReturnFailedResponse(ctx, s.tool.SprintfErr(err))
		return
	}
	responseManager.ReturnSucceedResponseWithoutData(ctx, nil)
}

//
//func (s *Service) GetUserInfo(ctx *gin.Context) {
//	phone := ctx.Query("phone")
//	id := ctx.Query("id")
//	var user *User
//	var err error
//	if len(phone) > 0 {
//		user, err = man.GetUserInfoByPhone(phone)
//	}
//	if len(id) > 0 {
//		user, err = man.GetUserInfoByID(id)
//	}
//	if err != nil {
//		s.res.NewRes(ctx, code.Failed, "get user info failed", user)
//		return
//	}
//	s.res.NewRes(ctx, 1, "ok", user)
//}

func GetService() *Service {
	return svr
}
