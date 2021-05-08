package store

import (
	"backend/source"
	"backend/source/tool"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var d *Dao

type Dao struct {
	collection *mongo.Collection
	logger     *logrus.Logger
	*tool.CtxManager
}

func init() {
	d = new(Dao)
	d.collection = source.DB.Collection("store")
	d.logger = source.GetLogger()
	d.CtxManager = tool.GetCtxManager()
}

func (d *Dao) FindOne(filter interface{}, opt *options.FindOneOptions) (*Store, error) {
	one := d.collection.FindOne(d.Ctx(), filter, opt)
	if one == nil {
		return nil, errors.New("not find")
	}
	s := new(Store)
	err := one.Decode(s)
	if err != nil {
		return nil, err
	}
	s.ID = tool.ConObjectIDToString(s.ObjID)
	return s, nil
}

func (d *Dao) Find(filter interface{}, opt *options.FindOptions) ([]*Store, error) {
	cur, err := d.collection.Find(d.Ctx(), filter, opt)
	if err != nil {
		return nil, err
	}
	stores := make([]*Store, 0)
	for cur.Next(d.Ctx()) {
		item := new(Store)
		err = cur.Decode(item)
		if err != nil {
			return nil, err
		}
		stores = append(stores, item)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}

	defer func() {
		_ = cur.Close(d.Ctx())
	}()

	return stores, nil
}
