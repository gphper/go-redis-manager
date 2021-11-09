/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:47:35
 */
package controllers

import (
	"context"
	"strconv"
	"time"

	"goredismanager/global"

	"github.com/gin-gonic/gin"
)

type stringController struct {
	BaseController
}

var Sc = stringController{}

func (con *stringController) Ttl(c *gin.Context) {

	ttl := c.PostForm("ttl")

	key := c.PostForm("key")

	val, _ := c.Get("username")

	ctx := context.WithValue(context.Background(), "username", val)

	ttlInt, _ := strconv.Atoi(ttl)

	time := time.Duration(ttlInt) * time.Second

	if ttlInt < 0 {
		global.UseClient.Client.Persist(ctx, key)
	} else {
		global.UseClient.Client.Expire(ctx, key, time).Result()
	}

	con.Success(c, "", "设置成功")
}

/**
* 编辑内容
**/
func (con *stringController) Save(c *gin.Context) {

	key := c.PostForm("key")

	content := c.PostForm("content")

	val, _ := c.Get("username")

	ctx := context.WithValue(context.Background(), "username", val)

	_, err := global.UseClient.Client.Set(ctx, key, content, 0*time.Second).Result()

	if err != nil {
		con.Error(c, "设置失败")
	} else {
		con.Success(c, "", "设置成功")
	}

}
