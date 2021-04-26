package notice

import (
	"backend/source"
	"github.com/sirupsen/logrus"
)

var senderProxy *SenderProxy

func init() {
	senderProxy = new(SenderProxy)
	senderProxy.logger = source.Logger
}

type Sender interface {
	Send() error
	SendVerificationCode(phone string, vCode int) error
}

type SenderProxy struct {
	logger    *logrus.Logger
	smsSender *SMSSender
}

func GetSender() *SenderProxy {
	return senderProxy
}

func (s *SenderProxy) Send() error {
	return nil
}
