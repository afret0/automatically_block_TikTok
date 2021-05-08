package user

import (
	"backend/source"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var svr *Service

func init() {
	svr = new(Service)
	svr.logger = source.GetLogger()
	svr.res = source.GetResTemplateManager()
}

type Service struct {
	logger *logrus.Logger
	res    *source.ResTemplateManager
}

func (s *Service) RegisterUser(ctx *gin.Context) {
	_ = man.RegisterUser(ctx.GetString("name"), ctx.GetString("phone"))
}

func (s *Service) Login(ctx *gin.Context) {
	token, err := man.Login(ctx.Query("phone"), ctx.Query("vCode"))
	if err != nil {
		s.res.NewResWithoutData(ctx, code.Failed, err.Error())
		return
	}
	s.res.NewSucceedRes(ctx, map[string]string{"token": token})
	return
}

func (s *Service) SendVerificationCode(ctx *gin.Context) {
	err := man.SendVerificationCode(ctx.Query("sender"), ctx.Query("phone"))
	if err != nil {
		s.res.NewResWithoutData(ctx, code.Failed, "send verification code failed")
		return
	}
	s.res.NewSucceedResWithoutData(ctx)
}

func (s *Service) GetUserInfo(ctx *gin.Context) {
	user, err := man.GetUserInfoByPhone(ctx.Query("phone"))
	if err != nil {
		s.res.NewRes(ctx, code.Failed, "get user info failed", user)
		return
	}
	s.res.NewRes(ctx, 1, "ok", user)
}

func GetService() *Service {
	return svr
}
