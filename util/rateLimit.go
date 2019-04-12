package util

import (
	gredis "github.com/go-redis/redis"
	"shorturl/middleware/redis"
	"strconv"
	"time"
)

func (u *Util) Rate(key string, num, gap int) bool {
	rateLimitKey := "rate_limit_" + key
	expires := gap * 2
	nTime := time.Now()
	now := nTime.Unix()
	pipe := redis.Redis().Pipeline()
	pipe.LPush(rateLimitKey, now)
	pipe.Expire(rateLimitKey, time.Duration(expires)*time.Second)
	pipes, err := pipe.Exec()
	if err != nil {
		return false
	}
	len := pipes[0].(*gredis.IntCmd).Val()
	if len > int64(num) {
		redisNow := redis.Redis().LIndex(rateLimitKey, int64(num-1))
		redisIntNow, _ := strconv.Atoi(redisNow.Val())
		if (now - int64(redisIntNow)) < int64(gap) {
			redis.Redis().LTrim(rateLimitKey, int64(-num), int64(-1))
			return true
		} else {
			redis.Redis().LTrim(rateLimitKey, 0, int64(num-1))
			return false
		}

	}
	return false
}
