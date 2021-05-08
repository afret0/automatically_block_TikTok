package tool

import (
	"context"
	"time"
)

var cm *CtxManager

func init() {
	cm = new(CtxManager)
}

type CtxManager struct {
}

func (c CtxManager) Ctx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func GetCtxManager() *CtxManager {
	return cm
}
