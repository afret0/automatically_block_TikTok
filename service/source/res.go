package source

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var resTemplateMan *ResTemplateManager

func init() {
	resTemplateMan = new(ResTemplateManager)

}

type ResTemplateManager struct {
}

func (r *ResTemplateManager) NewRes(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{"code": code, "msg": msg, "data": data})
	return
}

func (r *ResTemplateManager) NewResWithoutData(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, map[string]interface{}{"code": code, "msg": msg, "data": nil})
	return
}

func (r *ResTemplateManager) NewSucceedRes(c *gin.Context, data interface{}) {
	r.NewRes(c, 1, "succeed", data)
	return
}

func (r *ResTemplateManager) NewSucceedResWithoutData(c *gin.Context) {
	r.NewRes(c, 1, "succeed", map[string]string{})
	return
}

func GetResTemplateManager() *ResTemplateManager {
	return resTemplateMan
}
