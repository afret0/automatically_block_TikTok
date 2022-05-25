package smtp

import (
	"errors"
	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"service/source"
)

type Manager struct {
	logger *logrus.Logger
	cfg    *viper.Viper
}

type SendEmailResponse struct {
	source.RequestResponse
}

var manager *Manager
var SendEmailError error

func init() {
	manager = new(Manager)
	manager.logger = source.GetLogger()
	manager.cfg = source.GetConfig()

	SendEmailError = errors.New("send smtp error")
}

func GetManager() *Manager {
	return manager
}

func (m *Manager) SendEmail(to, subject, text string) error {
	response := new(SendEmailResponse)
	smtpAddr := m.cfg.GetString("smtp")
	request := gorequest.New()
	_, _, err := request.Post(smtpAddr).SendMap(map[string]string{"to": to, "subject": subject, "text": text}).EndStruct(response)
	if err != nil {
		m.logger.Errorln(to, subject, err)
		return SendEmailError
	}
	return nil
}
