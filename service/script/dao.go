package script

import (
	"backend/source"
	"backend/source/tool"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dao *Dao

type Dao struct {
	collection *mongo.Collection
	logger     *logrus.Logger
	ctx        context.Context
}

func init() {
	dao = new(Dao)
	dao.ctx = source.GetCtx()
	dao.logger = source.GetLogger()
	dao.collection = source.DB.Collection("script")
}

func (d *Dao) FindOne(filter interface{}, opt *options.FindOneOptions) (*Script, error) {
	one := d.collection.FindOne(d.ctx, filter, opt)
	if one == nil {
		return nil, errors.New("not find")
	}
	s := new(Script)
	err := one.Decode(s)
	s.Id = tool.ConObjectIDToString(s.ObjId)
	return s, err
}

func GetDao() *Dao {
	return dao
}
