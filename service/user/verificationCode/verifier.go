package verificationCode

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"service/source"
	"time"
)
import "github.com/go-redis/redis"

var v *Verifier

func init() {
	v = new(Verifier)
	v.logger = source.GetLogger()
	v.redisClient = source.GetRedisClient()
}

type Verifier struct {
	logger      *logrus.Logger
	redisClient *redis.Client
}

func (v *Verifier) GenVerifyCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vCode
}

func (v *Verifier) GetVerifyKey(email string) string {
	key := fmt.Sprintf("verifyCode:%s", email)
	return key
}

func (v *Verifier) SetVerifyCode(email, vCode string, expiration int) {
	key := v.GetVerifyKey(email)
	v.redisClient.Set(key, vCode, time.Duration(expiration)*time.Minute)
}

func (v *Verifier) CheckVCode(email, vCodeForCheck string) bool {
	key := v.GetVerifyKey(email)
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
