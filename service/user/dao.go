package user

import (
	"backend/source"
	"backend/source/tool"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
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
	dao.collection = source.DB.Collection("user")
	dao.logger = source.GetLogger()
	dao.CtxManager = tool.GetCtxManager()
}

func GetDao() *Dao {
	return dao
}

//func (d *Dao) ctx() context.Context {
//	return source.NewCtx()
//}

func (d *Dao) Find(filter interface{}, opt *options.FindOptions) ([]*User, error) {
	cur, err := d.collection.Find(d.Ctx(), filter, opt)
	if err != nil {
		return nil, err
	}
	users := make([]*User, 0)
	for cur.Next(d.Ctx()) {
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
		_ = cur.Close(d.Ctx())
	}()

	return users, err
}

func (d *Dao) UpdateOne(filter interface{}, update interface{}, opt *options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := d.collection.UpdateOne(d.Ctx(), filter, update, opt)
	return res, err
}

func (d *Dao) FindOne(filter interface{}, opt *options.FindOneOptions) (*User, error) {
	one := d.collection.FindOne(d.Ctx(), filter, opt)
	if one == nil {
		return nil, errors.New("not find")
	}
	u := new(User)
	err := one.Decode(u)
	if err != nil {
		return nil, err
	}
	u.ID = tool.ConObjectIDToString(u.ObjID)
	return u, err
}

func (d *Dao) InsertOne(name, pwd, WXName, token string, opt ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	doc := bson.M{"name": name, "password": pwd, "WXName": WXName, "tokenManager": token}
	one, err := d.collection.InsertOne(d.Ctx(), doc, opt...)
	return one, err
}
