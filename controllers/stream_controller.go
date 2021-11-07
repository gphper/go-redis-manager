/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:47:35
 */
package controllers

import (
	"context"
	"encoding/json"
	"goredismanager/global"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type streamController struct {
	BaseController
}

var Streamc = streamController{}

/**
* 添加
**/
func (con *streamController) Add(c *gin.Context) {

	var err error

	key := c.PostForm("key")

	id := c.PostForm("id")

	value := c.PostForm("value")

	obj := make(map[string]interface{})

	err = json.Unmarshal([]byte(value), &obj)

	if err != nil {
		con.Error(c, "json格式错误")
		return
	}

	ctx := context.Background()

	_, err = global.UseClient.Client.XAdd(ctx, &redis.XAddArgs{
		Stream: key,
		ID:     id,
		Values: obj,
	}).Result()

	if err == nil {
		con.Success(c, "", "添加成功")
	} else {
		con.Error(c, "添加失败")
	}
}

/**
* 删除
**/
func (con *streamController) Del(c *gin.Context) {
	key := c.PostForm("key")
	id := c.PostForm("id")

	ctx := context.Background()

	_, err := global.UseClient.Client.XDel(ctx, key, id).Result()

	if err == nil {
		con.Success(c, "", "删除成功")
	} else {
		con.Error(c, "删除失败")
	}
}
