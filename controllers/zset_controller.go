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

	"github.com/gin-gonic/gin"
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
