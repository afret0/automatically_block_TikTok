package user

import (
	"backend/notice"
	"backend/source"
	"backend/source/tool"
	verificationCode2 "backend/user/verificationCode"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var man *Manager

func init() {
	man = new(Manager)
	man.logger = source.GetLogger()
	man.verifier = verificationCode2.GetVerifier()
	man.dao = GetDao()
}

type Manager struct {
	logger   *logrus.Logger
	ctx      context.Context
	verifier *verificationCode2.Verifier
	dao      *Dao
}

func (m *Manager) RegisterUser(WXName, phone string) error {
	//user := new(User)
	//user.ID = man.newUserId()
	//user.Phone = phone
	//user.Name = WXName
	//user.WXName = WXName
	//user.Role = role.Customer

	//filter := bson.M{"phone": phone}
	//upt := bson.M{"$set": bson.M{"WXName": WXName, "id": m.newUserId(), "userName": WXName, "role": role.Customer}}
	//opt := new(options.UpdateOptions)
	//T := true
	//opt.Upsert = &T
	//r, code := dao.UpdateOne(m.ctx, filter, upt, opt)
	//if code != nil {
	//	m.logger.Errorln(phone, code)
	//	return code
	//}
	return nil
}

func (m *Manager) Login(phone, verificationCode string) (string, error) {
	if !m.verifier.CheckVCode(phone, verificationCode) {
		return "", errors.New("verificationCode error")
	}
	pjt := bson.M{"_id": 1, "name": 1, "token": 1}
	user, err := m.getUserInfoByPhone(phone, pjt)
	if err != nil {
		m.logger.Errorln(phone, err)
		return "", err
	}

	t, err := source.GetJWT().GenerateToken(user.ID, user.Name, user.Phone)
	if err != nil {
		return "", err
	}
	if t == user.Token {
		return t, nil
	}
	filter := bson.M{"phone": phone}
	upt := bson.M{"$set": bson.M{"token": t}}
	opt := new(options.UpdateOptions)
	_, err = m.dao.UpdateOne(filter, upt, opt)
	if err != nil {
		m.logger.Errorln(phone, err)
	}
	return t, err
}

func (m *Manager) SendVerificationCode(senderName, phone string) error {
	//TODO rate limit
	vCode := m.verifier.GenVerifyCode()
	m.verifier.SetVerifyCode(phone, vCode, 10)
	sender := notice.GetSender(senderName)
	return sender.SendVerificationCode(phone, vCode)
}

func (m *Manager) getUserInfoByPhone(phone string, pjt primitive.M) (*User, error) {
	//phoneRev := tool.ReverseString(phone)
	filter := bson.M{"phone": phone}
	opt := new(options.FindOneOptions)
	opt.Projection = pjt
	user, err := m.dao.FindOne(filter, opt)
	if err != nil {
		m.logger.Errorln(phone, err)
	}
	return user, err
}

func (m *Manager) GetUserInfoByPhone(phone string) (*User, error) {
	pjt := bson.M{"_id": 1, "name": 1, "avatar": 1, "dm": 1}
	return m.getUserInfoByPhone(phone, pjt)
}

func (m *Manager) GetUserInfoByID(id string) (*User, error) {
	filter := bson.M{"_id": tool.ConStringToObjectID(id)}
	opt := new(options.FindOneOptions)
	pjt := bson.M{"name": 1, "avatar": 1, "dm": 1}
	opt.Projection = pjt
	one, err := m.dao.FindOne(filter, opt)
	if err != nil {
		m.logger.Errorln(id, err)
	}
	return one, err
}
