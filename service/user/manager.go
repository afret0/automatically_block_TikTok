package user

import (
	"backend/utils"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"time"
)

var m *Manager
var randS *rand.Rand

func init() {
	m = new(Manager)
	m.logger = utils.Logger
	randS = rand.New(rand.NewSource(time.Now().Unix()))

}

type Manager struct {
	logger *logrus.Logger
}

func (m *Manager) newUserId() string {
	current := time.Now().UnixNano()
	return fmt.Sprintf("%d%d%d", current, randS.Intn(99), randS.Intn(99))
}

func (m *Manager) RegisterUser(ctx context.Context, WXName, phone string) error {
	//user := new(User)
	//user.Id = m.newUserId()
	//user.Phone = phone
	//user.Name = WXName
	//user.WXName = WXName
	//user.Role = role.Customer

	filter := bson.M{"phone": phone}
	upt := bson.M{"$set": bson.M{"WXName": WXName, "id": m.newUserId(), "userName": WXName, "role": role.Customer}}
	opt := new(options.UpdateOptions)
	T := true
	opt.Upsert = &T
	r, err := operator.UpdateOne(ctx, filter, upt, opt)
	m.logger.Infof("%+v", r)
	if err != nil {
		m.logger.Error(err)
		return err
	}
	return nil
}

//func GetManager() *Manager {
//	return m
//}
