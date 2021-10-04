/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:47:35
 */
package controllers

import (
	"context"
	"fmt"
	"goredismanager/global"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type zsetController struct {
	BaseController
}

var Zsetc = zsetController{}

/**
* 展示列表
 */
func (con *zsetController) Show(c *gin.Context) {
	key := c.Query("key")

	ctx := context.Background()

	time, _ := global.UseClient.Client.TTL(ctx, key).Result()

	zvalue, _ := global.UseClient.Client.ZRangeWithScores(ctx, key, 0, -1).Result()
	fmt.Printf("%+v", zvalue)

	c.HTML(http.StatusOK, "show/zset.html", gin.H{
		"key":    key,
		"zvalue": zvalue,
		"time":   time.Seconds(),
	})
}

/**
* 删除特定值
 */
func (con *zsetController) Del(c *gin.Context) {
	key := c.PostForm("key")
	member := c.PostForm("member")
	ctx := context.Background()

	res, _ := global.UseClient.Client.ZRem(ctx, key, member).Result()
	fmt.Println(res)
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

	ctx := context.Background()

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
