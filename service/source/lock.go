package source

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
)

var l *Locker

func init() {
	l = new(Locker)
	l.redis = GetRedisClient()
	l.logger = GetLogger()
}

type Locker struct {
	redis  *redis.Client
	logger *logrus.Logger
}

type Lock struct {
	Key string
}

//TODO 锁
func (l *Locker) Lock(key string) {
	key = fmt.Sprintf("lock_%s", key)
	nx := l.redis.SetNX(key, 1, 5*time.Second)
	lockSuccess, err := nx.Result()
	if err != nil || !lockSuccess {
		l.logger.Errorln(err, "lock result: ", lockSuccess)
		return
	}
}

func (l *Lock) Unlock() {

}

func GetLock() *Locker {
	return l
}
