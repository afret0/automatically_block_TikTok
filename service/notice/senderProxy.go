package notice

import (
	"backend/source"
	"github.com/sirupsen/logrus"
)

var senderProxy *SenderProxy

func init() {
	senderProxy = new(SenderProxy)
	senderProxy.logger = source.GetLogger()
}

type Sender interface {
	Send() error
	SendVerificationCode(phone string, vCode string) error
}

type SenderProxy struct {
	logger    *logrus.Logger
	smsSender *SMSSender
}

func GetSender(sender string) *SenderProxy {
	return senderProxy
}

func (s *SenderProxy) Send() error {
	return nil
}

func (s *SenderProxy) SendVerificationCode(phone string, vCode string) error {
	s.logger.Infoln(phone, vCode)
	return nil
}
