/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-21 14:45:08
 */
package service

import (
	"context"
	"goredismanager/global"
	"goredismanager/model"

	"github.com/go-redis/redis/v8"
)

func AddServiceConf(conf model.ServiceConfigReq) (err error) {
	optionConfig := &redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Password,
		DB:       0,
	}

	ctx := context.Background()

	client := redis.NewClient(optionConfig)

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return
	}

	RsSlice := global.RedisService{
		RedisService: conf.ServiceName,
		Config:       optionConfig,
		Client:       client,
	}

	global.RedisServiceStorage[conf.ServiceName] = RsSlice

	//设置全局参数
	global.UseClient.ConnectName = conf.ServiceName
	global.UseClient.Db = 0
	global.UseClient.Client = client

	return nil
}

func ServiceSwitch(conf model.ServiceSwitchReq) (err error) {
	config := global.RedisServiceStorage[conf.Service]
	config.Config.DB = conf.Db

	client := redis.NewClient(config.Config)
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	global.UseClient.ConnectName = conf.Service
	global.UseClient.Db = conf.Db
	global.UseClient.Client = redis.NewClient(config.Config)

	return nil
}

func SearchKeyType(conf model.RedisKeyReq) (keyType string, err error) {

	ctx := context.Background()

	keyType, err = global.UseClient.Client.Type(ctx, conf.Key).Result()
	if err != nil {
		return "", err
	}

	// switch keyType {

	// case "string":
	// 	result, err = global.UseClient.Client.Get(ctx, conf.Key).Result()
	// case "list":
	// 	len, _ := global.UseClient.Client.LLen(ctx, conf.Key).Result()
	// 	result, err = global.UseClient.Client.LRange(ctx, conf.Key, 0, len).Result()
	// case "set":
	// 	result, err = global.UseClient.Client.SMembers(ctx, conf.Key).Result()
	// case "zset":
	// 	result, err = global.UseClient.Client.ZRangeWithScores(ctx, conf.Key, 0, -1).Result()
	// }
	return
}
