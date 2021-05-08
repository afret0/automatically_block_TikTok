package dao

import (
	"backend/source"
	"backend/source/tool"
	"backend/store"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var d *Dao

type Dao struct {
	collection *mongo.Collection
	logger     *logrus.Logger
	ctx        context.Context
}

func init() {
	d = new(Dao)
	d.collection = source.DB.Collection("store")
	d.logger = source.GetLogger()
	d.ctx = source.GetCtx()
}

func (d *Dao) FindOne(filter interface{}, opt *options.FindOneOptions) (*store.Store, error) {
	one := d.collection.FindOne(d.ctx, filter, opt)
	if one == nil {
		return nil, errors.New("not find")
	}
	s := new(store.Store)
	err := one.Decode(s)
	if err != nil {
		return nil, err
	}
	s.ID = tool.ConObjectIDToString(s.ObjID)
	return s, nil
}

func (d *Dao) Find(filter interface{}, opt *options.FindOptions) ([]*store.Store, error) {
	cur, err := d.collection.Find(d.ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	stores := make([]*store.Store, 0)
	for cur.Next(d.ctx) {
		item := new(store.Store)
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
		_ = cur.Close(d.ctx)
	}()

	return stores, nil
}
