/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:47:35
 */
package controllers

import (
	"context"
	"goredismanager/global"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type zsetController struct {
	BaseController
}

var Zsetc = zsetController{}

/**
* 删除特定值
 */
func (con *zsetController) Del(c *gin.Context) {
	key := c.PostForm("key")

	member := c.PostForm("member")

	val, _ := c.Get("username")

	ctx := context.WithValue(context.Background(), "username", val)

	res, _ := global.UseClient.Client.ZRem(ctx, key, member).Result()

	if res > 0 {
		con.Success(c, "", "删除成功")
	} else {
		con.Error(c, "删除失败")
	}

}

/**
* 新增值
**/
func (con *zsetController) Add(c *gin.Context) {
	key := c.PostForm("key")

	value := c.PostForm("value")

	score := c.PostForm("score")

	scoreInt, _ := strconv.Atoi(score)

	val, _ := c.Get("username")

	ctx := context.WithValue(context.Background(), "username", val)

	_, err := global.UseClient.Client.ZAdd(ctx, key, &redis.Z{
		Score:  float64(scoreInt),
		Member: value,
	}).Result()

	if err == nil {
		con.Success(c, "", "添加成功")
	} else {
		con.Error(c, "添加失败")
	}

}
