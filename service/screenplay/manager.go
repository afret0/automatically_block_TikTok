package screenplay

import (
	"backend/source"
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
	m.dao = GetDao()
}

func (m *Manager) GetOneScriptInfo(id string) (*Screenplay, error) {
	filter := bson.M{"id": id}
	opt := new(options.FindOneOptions)
	one, err := m.dao.FindOne(filter, opt)
	if err != nil {
		m.logger.Errorln(id, err)
	}
	return one, err
}

func (m *Manager) InsertScript() (string, error) {
	return "", nil
}

func GetManager() *Manager {
	return m
}
