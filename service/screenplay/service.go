package screenplay

import (
	"backend/source"
	"github.com/gin-gonic/gin"
)

var svr *Service

type Service struct {
	manager *Manager
	res     *source.ResTemplateManager
}

func init() {
	svr = new(Service)
	svr.manager = GetManager()
	svr.res = source.GetResTemplateManager()
}

func (s *Service) GetOneScreenplayInfo(c *gin.Context) {
	info, err := s.manager.GetOneScreenplayInfo(c.Query("id"))
	if err != nil {
		s.res.NewResWithoutData(c, GetCode().Failed, err.Error())
		return
	}
	s.res.NewSucceedRes(c, info)
}

func GetService() *Service {
	return svr
}
