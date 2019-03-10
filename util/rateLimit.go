package util

import (
	"fmt"
	"shorturl/middleware/redis"
	"strconv"
	"time"
)

func (u *Util) Rate(key string, num, gap int) bool {
	rateLimitKey := "rate_limit_prefix_" + key
	expires := gap * 2
	lens, _ := redis.Redis().LLen(rateLimitKey).Result()
	nTime := time.Now()
	now := nTime.Unix()
	if lens < int64(num) {
		pipe := redis.Redis().Pipeline()
		redis.Redis().LPush(rateLimitKey, now)
		pipe.Expire(rateLimitKey, time.Duration(expires)*time.Second)
		pipe.Exec()
		return false
	} else {
		redisNow := redis.Redis().LIndex(rateLimitKey, -1)
		redisIntNow, _ := strconv.Atoi(redisNow.Val())
		if (now - int64(redisIntNow)) < int64(gap) {
			fmt.Print("rate error")
			return true
		} else {
			pipe := redis.Redis().Pipeline()
			redis.Redis().LPush(rateLimitKey, now)
			redis.Redis().LTrim(rateLimitKey, 0, int64(num-1))
			pipe.Expire(rateLimitKey, time.Duration(expires)*time.Second)
			pipe.Exec()
			return false
		}
	}
}
