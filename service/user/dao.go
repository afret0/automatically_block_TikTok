package user

import (
	"backend/source"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var operator *Operator

type Operator struct {
	collection *mongo.Collection
	logger     *logrus.Logger
	ctx        context.Context
}

func init() {
	operator = new(Operator)
	operator.collection = source.DB.Collection("user")
	operator.logger = source.Logger
}

func (o *Operator) Find(filter interface{}, opt *options.FindOptions) ([]*User, error) {
	cur, err := o.collection.Find(o.ctx, filter, opt)
	if err != nil {
		source.Logger.Error(err)
		return nil, err
	}
	users := make([]*User, 0)
	for cur.Next(o.ctx) {
		item := new(User)
		err = cur.Decode(item)
		if err != nil {
			return nil, err
		}
		users = append(users, item)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}

	defer func() {
		_ = cur.Close(o.ctx)
	}()

	return users, err
}

func (o *Operator) UpdateOne(filter interface{}, update interface{}, opt ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := o.collection.UpdateOne(o.ctx, filter, update, opt...)
	return res, err
}

func (o *Operator) FindOne(filter interface{}, opt *options.FindOneOptions) (*User, error) {
	one := o.collection.FindOne(o.ctx, filter, opt)
	if one == nil {
		return nil, errors.New("not find")
	}
	user := new(User)
	err := one.Decode(user)
	user.ID = source.ConObjectIDToString(user.ObjID)
	return user, err
}

func (o *Operator) InsertOne(name, pwd, WXName, token string, opt ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	doc := bson.M{"name": name, "password": pwd, "WXName": WXName, "role": role.Customer, "token": token}
	one, err := o.collection.InsertOne(o.ctx, doc, opt...)
	return one, err
}
