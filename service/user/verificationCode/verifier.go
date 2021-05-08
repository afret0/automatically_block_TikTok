package verificationCode

import (
	"backend/notice"
	"backend/source"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)
import "github.com/go-redis/redis"

var v *Verifier

func init() {
	v = new(Verifier)
	v.logger = source.GetLogger()
	v.redisClient = source.GetRedisClient()
	v.sender = notice.GetSmsSender()
}

type Verifier struct {
	logger      *logrus.Logger
	redisClient *redis.Client
	sender      notice.Sender
}

func (v *Verifier) GenVerifyCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vCode
}

func (v *Verifier) GetVerifyKey(phone string) string {
	key := fmt.Sprintf("verifyCode:%s", phone)
	return key
}

func (v *Verifier) SetVerifyCode(phone, vCode string, expiration int) {
	key := v.GetVerifyKey(phone)
	v.redisClient.Set(key, vCode, time.Duration(expiration)*time.Minute)
}

func (v *Verifier) CheckVCode(phone string, vCodeForCheck string) bool {
	key := v.GetVerifyKey(phone)
	get := v.redisClient.Get(key)
	vCode, _ := get.Result()
	if vCode == vCodeForCheck {
		return true
	}
	return false
}

func GetVerifier() *Verifier {
	return v
}
