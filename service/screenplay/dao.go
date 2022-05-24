package screenplay

import (
	"backend/source"
	"backend/source/tool"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dao *Dao

type Dao struct {
	collection *mongo.Collection
	logger     *logrus.Logger
	*tool.CtxManager
	err *source.Err
}

func init() {
	dao = new(Dao)
	dao.logger = source.GetLogger()
	dao.collection = source.DB.Collection("screenplay")
	dao.CtxManager = tool.GetCtxManager()
	dao.err = source.GetErr()
}

func (d *Dao) FindOne(filter interface{}, opt *options.FindOneOptions) (*Screenplay, error) {
	one := d.collection.FindOne(d.Ctx(), filter, opt)
	if one == nil {
		return nil, d.err.NoFind
	}
	s := new(Screenplay)
	err := one.Decode(s)
	s.Id = tool.ConObjectIDToString(s.ObjId)
	return s, err
}

func (d *Dao) InsertOne(doc interface{}, opt *options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	one, err := d.collection.InsertOne(d.Ctx(), doc, opt)
	return one, err
}

func (d *Dao) UpdateOne(filter interface{}, upt interface{}, opt *options.UpdateOptions) (*mongo.UpdateResult, error) {
	one, err := d.collection.UpdateOne(d.Ctx(), filter, upt, opt)
	return one, err
}

func GetDao() *Dao {
	return dao
}
