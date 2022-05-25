package user

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"service/smtp"
	"service/source"
	"service/source/tool"
	"service/user/verificationCode"
	"time"
)

var man *Manager

type Manager struct {
	logger   *logrus.Logger
	ctx      context.Context
	verifier *verificationCode.Verifier
	smtp     *smtp.Manager
	dao      *Dao
	tool     *tool.Tool
	locker   *source.Locker
}

func init() {
	man = new(Manager)
	man.logger = source.GetLogger()
	man.verifier = verificationCode.GetVerifier()
	man.dao = getDao()
	man.tool = tool.GetTool()
	man.smtp = smtp.GetManager()
	man.locker = source.GetLocker()
}

func (m *Manager) RegisterUser(ctx context.Context, Name, email, password, verificationCode string) (string, error) {
	check := m.verifier.CheckVCode(email, verificationCode)
	if check != true {
		return "", CheckVerificationCodeError
	}
	filter := bson.M{"email": email}
	now := time.Now().Unix()
	upt := bson.M{"name": Name, "password": password, "registerTime": now, "updateTime": now}
	opt := new(options.UpdateOptions)
	upsert := true
	opt.Upsert = &upsert
	lock := m.locker.Lock(email)
	if !lock {
		return "", m.locker.LockFailed
	}
	defer m.locker.Unlock(email)
	_, err := m.dao.UpdateOne(ctx, filter, upt, opt)
	if err != nil {
		m.logger.Errorln(email, err)
		return "", err
	}
	return "", nil
}

//func (m *Manager) Login(ctx context.Context, phone, verificationCode string) (string, error) {
//	if !m.verifier.CheckVCode(phone, verificationCode) {
//		return "", errors.New("verificationCode error")
//	}
//	pjt := bson.M{"_id": 1, "name": 1, "token": 1}
//	user, err := m.getUserInfoByPhone(ctx, phone, pjt)
//	if err != nil {
//		m.logger.Errorln(phone, err)
//		return "", err
//	}
//
//	t, err := GetJWT().GenerateToken(user.ID, user.Name, user.Email)
//	if err != nil {
//		return "", err
//	}
//	if t == user.Token {
//		return t, nil
//	}
//	filter := bson.M{"phone": phone}
//	upt := bson.M{"$set": bson.M{"token": t}}
//	opt := new(options.UpdateOptions)
//	_, err = m.dao.UpdateOne(ctx, filter, upt, opt)
//	if err != nil {
//		m.logger.Errorln(phone, err)
//	}
//	return t, err
//}

func (m *Manager) SendVerificationCode(email string) error {
	vCode := m.verifier.GenVerifyCode()
	m.verifier.SetVerifyCode(email, vCode, 10)
	subject := "pancake 验证码"
	text := fmt.Sprintf("霓为衣兮风为马，云之君兮纷纷而来下 ~ \n\n虎鼓瑟兮鸾回车，仙之人兮列如麻 ~ \n\n\n您的验证码为: %s", vCode)
	err := m.smtp.SendEmail(email, subject, text)
	return err
}

//func (m *Manager) getUserInfoByPhone(ctx context.Context, phone string, pjt primitive.M) (*User, error) {
//	//phoneRev := tool.ReverseString(phone)
//	filter := bson.M{"phone": phone}
//	opt := new(options.FindOneOptions)
//	opt.Projection = pjt
//	user, err := m.dao.FindOne(ctx, filter, opt)
//	if err != nil {
//		m.logger.Errorln(phone, err)
//	}
//	return user, err
//}

//func (m *Manager) GetUserInfoByID(ctx context.Context, id string) (*User, error) {
//	//TODO 缓存
//	filter := bson.M{"_id": m.tool.ConStringToObjectID(id)}
//	opt := new(options.FindOneOptions)
//	pjt := bson.M{"name": 1, "avatar": 1, "dm": 1}
//	opt.Projection = pjt
//	one, err := m.dao.FindOne(ctx, filter, opt)
//	if err != nil {
//		m.logger.Errorln(id, err)
//	}
//	return one, err
//}

func GetManager() *Manager {
	return man
}
