/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:47:35
 */
package controllers

import (
	"context"
	"goredismanager/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

type setController struct {
	BaseController
}

var Setc = setController{}

func (con *setController) Show(c *gin.Context) {
	key := c.Query("key")

	ctx := context.Background()
	value, _ := global.UseClient.Client.SMembers(ctx, key).Result()

	time, _ := global.UseClient.Client.TTL(ctx, key).Result()

	c.HTML(http.StatusOK, "show/set.html", gin.H{
		"key":   key,
		"value": value,
		"time":  time.Seconds(),
	})
}

/**
* 删除特定值
**/
func (con *setController) Del(c *gin.Context) {
	key := c.PostForm("key")
	member := c.PostForm("member")
	ctx := context.Background()

	res, _ := global.UseClient.Client.SRem(ctx, key, member).Result()

	if res > 0 {
		con.Success(c, "", "删除成功")
	} else {
		con.Error(c, "删除失败")
	}

}
