package lockUtil

import (
	"pz-infra/log"
	"pz-infra/redisUtil"
	"time"

	"pz-infra/redsync"

	"github.com/garyburd/redigo/redis"
)

func GetLockerAndLock(name string, expiry ...time.Duration) (redsync.Locker, error) {
	mutex, err := redsync.NewMutexWithPool(name, []*redis.Pool{redisUtil.GetPool()})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if len(expiry) > 0 {
		mutex.Expiry = expiry[0]
	}
	if err := mutex.Lock(); err != nil {
		log.Error(err)
		return nil, err
	}
	return mutex, nil
}
