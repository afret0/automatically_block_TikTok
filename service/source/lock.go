package source

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
)

var l *Locker

type Locker struct {
	redis      *redis.Client
	logger     *logrus.Logger
	LockFailed error
}

type Lock struct {
	Key string
}

func (l *Locker) getLockKey(key string) string {
	return fmt.Sprintf("lock_%s", key)
}

func (l *Locker) Lock(key string) bool {
	nx := l.redis.SetNX(l.getLockKey(key), 1, 5*time.Second)
	lockSuccess, err := nx.Result()
	if err != nil {
		l.logger.Errorln(err)
	}
	return lockSuccess
}

func (l *Locker) Unlock(key string) {
	nx := l.redis.Del(l.getLockKey(key))
	unlockSuccess, err := nx.Result()
	if err == nil && unlockSuccess > 0 {
		l.logger.Errorln(err, "unlock result: ", unlockSuccess)
	}
}

func GetLocker() *Locker {
	if l == nil {
		l = new(Locker)
		l.redis = GetRedisClient()
		l.logger = GetLogger()
		l.LockFailed = errors.New("lock failed")
	}
	return l
}
