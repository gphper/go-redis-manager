/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:47:35
 */
package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"goredismanager/global"

	"github.com/gin-gonic/gin"
)

type stringController struct {
	BaseController
}

var Sc = stringController{}

func (con *stringController) Show(c *gin.Context) {
	key := c.Query("key")

	ctx := context.Background()
	value, _ := global.UseClient.Client.Get(ctx, key).Result()

	time, _ := global.UseClient.Client.TTL(ctx, key).Result()
	fmt.Println(time.Seconds())
	c.HTML(http.StatusOK, "show/string.html", gin.H{
		"key":   key,
		"value": value,
		"time":  time.Seconds(),
	})
}

func (con *stringController) Ttl(c *gin.Context) {

	ttl := c.PostForm("ttl")
	key := c.PostForm("key")
	ctx := context.Background()

	ttlInt, _ := strconv.Atoi(ttl)
	time := time.Duration(ttlInt) * time.Second

	if ttlInt < 0 {
		global.UseClient.Client.Persist(ctx, key)
	} else {
		global.UseClient.Client.Expire(ctx, key, time).Result()
	}

	con.Success(c, "", "设置成功")
}
