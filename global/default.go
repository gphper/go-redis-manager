/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-21 10:08:32
 */
package global

import "github.com/go-redis/redis/v8"

var RedisServiceStorage map[string]RedisService

var UseClient GlobalClient

type GlobalClient struct {
	ConnectName string
	Db          int
	Client      *redis.Client
}

type RedisService struct {
	RedisService string
	Config       *redis.Options
	Client       *redis.Client
}

func init() {
	RedisServiceStorage = make(map[string]RedisService)
}
