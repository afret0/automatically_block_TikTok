package user

import (
	"backend/source"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var svr *Service
var resTM *source.ResTemplateManager

func init() {
	svr = new(Service)
	svr.logger = source.Logger
	resTM = source.GetResTemplateManager()
}

type Service struct {
	logger *logrus.Logger
}

func (s *Service) RegisterUser(ctx *gin.Context) {
	_ = man.RegisterUser(ctx.GetString("name"), ctx.GetString("phone"))
}

func (s *Service) Login(ctx *gin.Context) {
	token, err := man.Login(ctx.Query("phone"), ctx.Query("vCode"))
	if err != nil {
		resTM.NewResWithoutData(ctx, 100101, "login failed")
		return
	}
	resTM.NewSucceedRes(ctx, map[string]string{"token": token})
	return
}

func (s *Service) SendVerificationCode(ctx *gin.Context) {
	err := man.SendVerificationCode(ctx.Query("sender"), ctx.Query("phone"))
	if err != nil {
		resTM.NewResWithoutData(ctx, 100101, "send failed")
		return
	}
	resTM.NewSucceedResWithoutData(ctx)
}

func GetService() *Service {
	return svr
}
