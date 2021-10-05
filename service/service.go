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
	"time"

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

func SearchKeyType(conf model.RedisKeyReq) (keys []model.RedisKey, err error) {

	ctx := context.Background()

	if conf.SearchType == 1 {
		//查询指定值
		var keyType string

		keyType, err = global.UseClient.Client.Type(ctx, conf.SearchKey).Result()
		keys = append(keys, model.RedisKey{
			Key:  conf.SearchKey,
			Type: keyType,
		})

	} else {
		//模糊匹配
		var cursor uint64
		var tmpKeys []string

		for {

			tmpKeys, cursor, err = global.UseClient.Client.Scan(ctx, cursor, conf.SearchKey, 100).Result()

			for _, tmpKey := range tmpKeys {
				kryType, _ := global.UseClient.Client.Type(ctx, tmpKey).Result()
				keys = append(keys, model.RedisKey{
					Key:  tmpKey,
					Type: kryType,
				})
			}

			if cursor == 0 {
				break
			}

		}

	}
	return
}

func AddKey(req model.AddKeyReq) (err error) {

	ctx := context.Background()

	switch req.Type {
	case "string":
		_, err = global.UseClient.Client.Set(ctx, req.Key, req.Value, 0*time.Second).Result()
	case "list":
		_, err = global.UseClient.Client.LPush(ctx, req.Key, req.Value).Result()
	case "set":
		_, err = global.UseClient.Client.SAdd(ctx, req.Key, req.Value).Result()
	case "zset":
		_, err = global.UseClient.Client.ZAdd(ctx, req.Key, &redis.Z{
			Score:  float64(req.Score),
			Member: req.Value,
		}).Result()
	case "hash":
		_, err = global.UseClient.Client.HSet(ctx, req.Key, req.Field, req.Value).Result()
	}

	return
}
