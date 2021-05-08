package notice

import (
	"backend/source"
	"github.com/sirupsen/logrus"
)

var smsSender *SMSSender

func init() {
	smsSender = new(SMSSender)
	smsSender.logger = source.GetLogger()
}

type SMSSender struct {
	logger *logrus.Logger
}

func (s *SMSSender) Send() error {
	return nil
}

func (s *SMSSender) SendVerificationCode(phone string, vCode string) error {
	//#TODO 限流
	return nil
}

func GetSmsSender() *SMSSender {
	return smsSender
}
