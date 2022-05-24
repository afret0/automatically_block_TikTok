package screenplay

import (
	"backend/source"
	"backend/source/tool"
	"backend/user"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var m *Manager

type Manager struct {
	logger *logrus.Logger
	dao    *Dao
	err    *source.Err
}

func init() {
	m = new(Manager)
	m.logger = source.GetLogger()
	m.dao = GetDao()
	m.err = source.GetErr()
}

func (m *Manager) GetOneScreenplayInfo(id string) (*Screenplay, error) {
	filter := bson.M{"_id": tool.ConStringToObjectID(id)}
	opt := new(options.FindOneOptions)
	one, err := m.dao.FindOne(filter, opt)
	if err != nil {
		m.logger.Errorln(id, err)
	}
	return one, err
}

func (m *Manager) NewScreenplay(s *Screenplay, userID string) (interface{}, error) {
	userInfo, err := user.GetManager().GetUserInfoByID(userID)
	if err != nil {
		return nil, err
	}
	if !userInfo.Boss {
		m.logger.Infoln(userID, s.Id, "no authority")
		return nil, m.err.NoAuthority
	}
	doc := bson.M{"name": s.Name, "cover": s.Cover, "instruction": s.Instruction, "store": s.Store, "note": s.Note, "noun": s.Noun, "tag": s.Tag, "style": s.Style, "registerTime": time.Now().Unix(), "updateTime": time.Now().Unix()}
	opt := new(options.InsertOneOptions)
	one, err := m.dao.InsertOne(doc, opt)
	if err != nil {
		m.logger.Errorln(s.Name, err)
	}
	return one.InsertedID, err
}

func (m *Manager) UpdateScreenplay(s *Screenplay, userID string) error {
	userMan := user.GetManager()
	userInfo, err := userMan.GetUserInfoByID(userID)
	if err != nil {
		return err
	}
	unUpdatedScreenplay, err := m.GetOneScreenplayInfo(s.Id)
	pm := GetPermissionsManager()
	permissions, err := pm.UpdateScreenplay(userInfo, unUpdatedScreenplay)
	if err != nil {
		return err
	}
	if !permissions {
		return m.err.NoAuthority
	}

	filter := bson.M{"_id": tool.ConStringToObjectID(s.Id)}
	opt := new(options.UpdateOptions)
	upt := bson.M{"updateTime": time.Now().Unix()}
	if len(s.Cover) > 1 {
		upt["cover"] = s.Cover
	}
	if len(s.Instruction) > 1 {
		upt["instruction"] = s.Instruction
	}
	if len(s.Note) > 1 {
		upt["note"] = s.Note
	}
	if s.Noun != nil {
		upt["noun"] = s.Noun
	}
	_, err = m.dao.UpdateOne(filter, upt, opt)
	if err != nil {
		m.logger.Errorln(err)
	}
	return err
}

func GetManager() *Manager {
	return m
}
