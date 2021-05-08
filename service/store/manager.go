package store

import (
	"backend/source"
	"backend/source/tool"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var m *Manager

type Manager struct {
	logger *logrus.Logger
	dao    *Dao
}

func init() {
	m = new(Manager)
	m.logger = source.GetLogger()
	m.dao = d
}

func (m *Manager) GetStoreInfoById(id string) (*Store, error) {
	filter := bson.M{"_id": tool.ConStringToObjectID(id)}
	opt := new(options.FindOneOptions)
	one, err := m.dao.FindOne(filter, opt)
	if err != nil {
		m.logger.Errorln(id, err)
	}
	return one, err
}

//func (m *Manager) GetStoreInfoByOwner(id string) (*Store, error) {
//
//}

func (m *Manager) GetStoreList(owner string) ([]*Store, error) {
	filter := bson.M{"owner": owner}
	opt := new(options.FindOptions)
	find, err := m.dao.Find(filter, opt)
	if err != nil {
		m.logger.Errorln(owner, err)
	}
	return find, err
}
