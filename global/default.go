/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-21 10:08:32
 */
package global

import (
	"github.com/go-redis/redis/v8"
)

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

	optionConfig := &redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}

	client := redis.NewClient(optionConfig)

	RsSlice := RedisService{
		RedisService: "本地连接",
		Config:       optionConfig,
		Client:       client,
	}

	RedisServiceStorage["本地连接"] = RsSlice

	//设置全局参数
	UseClient.ConnectName = "本地连接"
	UseClient.Db = 0
	UseClient.Client = client

}
