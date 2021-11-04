/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-21 14:45:08
 */
package service

import (
	"context"
	"encoding/json"
	"goredismanager/global"
	"goredismanager/model"
	"time"

	"github.com/gin-gonic/gin"
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

func SearchKeyType(conf model.RedisKeyReq) (keys []string, err error) {

	ctx := context.Background()

	if conf.SearchType == 1 {
		//查询指定值
		keys = append(keys, conf.SearchKey)

		_, err = global.UseClient.Client.Exists(ctx, conf.SearchKey).Result()
		return
	} else {
		//模糊匹配
		var cursor uint64
		var tmpKeys []string

		for {

			tmpKeys, cursor, err = global.UseClient.Client.Scan(ctx, cursor, conf.SearchKey, 100).Result()

			keys = append(keys, tmpKeys...)

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
	case "stream":
		obj := make(map[string]interface{})

		err = json.Unmarshal([]byte(req.Value), &obj)
		if err != nil {
			return
		}

		_, err = global.UseClient.Client.XAdd(ctx, &redis.XAddArgs{
			Stream: req.Key,
			ID:     req.Id,
			Values: obj,
		}).Result()
	}

	return
}

func DelKey(req model.DelKeyReq) (err error) {
	ctx := context.Background()

	if req.Type == "none" {

		var cursor uint64
		var tmpKeys []string

		for {

			tmpKeys, cursor, err = global.UseClient.Client.Scan(ctx, cursor, req.Key+":*", 1000).Result()
			if err != nil {
				return
			}

			_, err = global.UseClient.Client.Del(ctx, tmpKeys...).Result()
			if err != nil {
				return
			}

			if cursor == 0 {
				break
			}

		}

	} else {
		_, err = global.UseClient.Client.Del(ctx, req.Key).Result()
		if err != nil {
			return
		}
	}
	return
}

func TransView(keyType string, keys string) (htmlString string, data gin.H) {

	ctx := context.Background()

	time, _ := global.UseClient.Client.TTL(ctx, keys).Result()

	switch keyType {
	case "string":

		value, _ := global.UseClient.Client.Get(ctx, keys).Result()

		htmlString = "show/string.html"
		data = gin.H{
			"key":   keys,
			"value": value,
			"time":  time.Seconds(),
		}
	case "list":

		value, _ := global.UseClient.Client.LRange(ctx, keys, 0, -1).Result()

		htmlString = "show/list.html"
		data = gin.H{
			"key":   keys,
			"value": value,
			"time":  time.Seconds(),
		}
	case "zset":

		zvalue, _ := global.UseClient.Client.ZRangeWithScores(ctx, keys, 0, -1).Result()

		htmlString = "show/zset.html"
		data = gin.H{
			"key":    keys,
			"zvalue": zvalue,
			"time":   time.Seconds(),
		}
	case "set":

		value, _ := global.UseClient.Client.SMembers(ctx, keys).Result()

		htmlString = "show/set.html"
		data = gin.H{
			"key":   keys,
			"value": value,
			"time":  time.Seconds(),
		}
	case "hash":

		res, _ := global.UseClient.Client.HGetAll(ctx, keys).Result()

		htmlString = "show/hash.html"
		data = gin.H{
			"key":    keys,
			"result": res,
			"time":   time.Seconds(),
		}
	}
	return
}
