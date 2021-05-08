package screenplay

import (
	"backend/source"
	"backend/source/tool"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dao *Dao

type Dao struct {
	collection *mongo.Collection
	logger     *logrus.Logger
	*tool.CtxManager
}

func init() {
	dao = new(Dao)
	dao.logger = source.GetLogger()
	dao.collection = source.DB.Collection("screenplay")
	dao.CtxManager = tool.GetCtxManager()
}

func (d *Dao) FindOne(filter interface{}, opt *options.FindOneOptions) (*Screenplay, error) {
	one := d.collection.FindOne(d.Ctx(), filter, opt)
	if one == nil {
		return nil, errors.New("not find")
	}
	s := new(Screenplay)
	err := one.Decode(s)
	s.Id = tool.ConObjectIDToString(s.ObjId)
	return s, err
}

func GetDao() *Dao {
	return dao
}
