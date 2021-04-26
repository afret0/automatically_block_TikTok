package utils

var resTemplateMan *ResTemplateManager

func init() {
	resTemplateMan = new(ResTemplateManager)

}

type ResTemplateManager struct {
}

type Res struct {
	Code int
	Msg  string
}

func (r *ResTemplateManager) NewRes(code int, Msg string) *Res {
	res := new(Res)
	res.Code = code
	res.Msg = Msg
	return res
}

func (r *ResTemplateManager) NewSucceedRes() *Res {
	return r.NewRes(1, "succeed")
}

func GetResTemplateManager() *ResTemplateManager {
	return resTemplateMan
}
