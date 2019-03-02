package service

import (
	"fmt"
	gredis "github.com/go-redis/redis"
	"github.com/spaolacci/murmur3"
	"log"
	"math/rand"
	"regexp"
	"shorturl/conf"
	"shorturl/middleware/redis"
	"shorturl/util"
	"strconv"
	"strings"
	"time"
)

var (
	ut *util.Util
)

var urlRegexp = regexp.MustCompile(`\[!@[A-Z]+@!\]`)

const redisKeyPrefix string = "shorturl:%s"
const domainSchema string = "http://"
const leftPrefix string = "[!@"
const rightPrefix string = "@!]"

func getRandStr() string {
	strMap := map[string]string{
		"0": "A",
		"1": "B",
		"2": "C",
		"3": "D",
		"4": "E",
		"5": "F",
		"6": "G",
		"7": "H",
		"8": "I",
		"9": "J",
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numStr := strconv.Itoa(r.Intn(1000000))
	retStr := ""
	for _, value := range strings.Split(numStr, "") {
		retStr = retStr + strMap[value]
	}
	return retStr
}

func getHash(url string) string {
	hasher := murmur3.New32()
	hasher.Write([]byte(url))
	uint32Num := hasher.Sum32()
	hasher.Reset()
	shortHashstr := ut.DecimalTo62(uint32Num)
	return shortHashstr
}

func (s *Service) GenShortUrl(originUrl string) string {
	shorturlDomain := domainSchema + conf.Conf.Domain.Name + "/"
	shortHashstr := getHash(originUrl)
	redisKey := fmt.Sprintf(redisKeyPrefix, shortHashstr)
	redisLongUrl, err := redis.Redis().Get(redisKey).Result()
	if err == gredis.Nil {
		//nothing
	} else if err != nil {
		log.Println("redis get error", err)
		return ""
	}
	//hash 值相同 原地址不同
	if redisLongUrl == "" {
		err := redis.Redis().Set(redisKey, originUrl, 62073600*time.Second).Err()
		if err != nil {
			log.Println("redis set error", err)
			return ""
		}
		return shorturlDomain + shortHashstr
	} else {
		if redisLongUrl != originUrl {
			newOriginUrl := originUrl + leftPrefix + getRandStr() + rightPrefix
			newShortHashstr := getHash(newOriginUrl)
			redisKey = fmt.Sprintf(redisKeyPrefix, newShortHashstr)
			err := redis.Redis().Set(redisKey, newOriginUrl, 62073600*time.Second).Err()
			if err != nil {
				log.Println("redis set error", err)
				return ""
			}
			return shorturlDomain + newShortHashstr
		}
	}
	return shorturlDomain + shortHashstr
}

func (s *Service) GetOriginUrl(shortHashstr string) string {
	redisKey := fmt.Sprintf(redisKeyPrefix, shortHashstr)
	redisLongUrl, err := redis.Redis().Get(redisKey).Result()
	if err == gredis.Nil {
		//nothing
	} else if err != nil {
		log.Println("redis get error", err)
		return ""
	}

	matchRet := urlRegexp.MatchString(redisLongUrl)
	if !matchRet {
		return redisLongUrl
	}
	return urlRegexp.ReplaceAllString(redisLongUrl, "")
}
