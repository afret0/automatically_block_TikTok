package user

import (
	"backend/utils"
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var operator *Operator

type Operator struct {
	Collection *mongo.Collection
	L          *logrus.Logger
}

func init() {
	operator = new(Operator)
	operator.Collection = utils.DB.Collection("user")
	operator.L = utils.Logger
}

func (o *Operator) FindValue(ctx context.Context, filter interface{}, opt *options.FindOptions) ([]*User, error) {
	cur, err := o.Collection.Find(ctx, filter, opt)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	users := make([]*User, 0)
	for cur.Next(ctx) {
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
		_ = cur.Close(ctx)
	}()

	return users, err
}

func (o *Operator) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opt ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := o.Collection.UpdateOne(ctx, filter, update, opt...)
	return res, err
}
