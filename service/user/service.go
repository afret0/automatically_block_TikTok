package user

import (
	"context"
	"github.com/gin-gonic/gin"
)

var svr *Service

func init() {
	svr = new(Service)
}

type Service struct {
}

func (s *Service) RegisterUser(ctx gin.Context, WXName, phone string) {
	err := m.RegisterUser(ctx, "test", "18435155427")
	if err != nil {

	}

}
