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
	"strconv"

	"github.com/gin-gonic/gin"
)

type listController struct {
	BaseController
}

var Lc = listController{}

/**
* list 列表展示
**/
func (con *listController) Show(c *gin.Context) {
	key := c.Query("key")

	ctx := context.Background()
	value, _ := global.UseClient.Client.LRange(ctx, key, 0, -1).Result()

	time, _ := global.UseClient.Client.TTL(ctx, key).Result()

	c.HTML(http.StatusOK, "show/list.html", gin.H{
		"key":   key,
		"value": value,
		"time":  time.Seconds(),
	})
}

/**
* 删除特定值
**/
func (con *listController) Del(c *gin.Context) {
	key := c.PostForm("key")
	index := c.PostForm("index")
	indexs, _ := strconv.Atoi(index)

	ctx := context.Background()

	_, err := global.UseClient.Client.LSet(ctx, key, int64(indexs), "listdel").Result()

	_, err1 := global.UseClient.Client.LRem(ctx, key, 1, "listdel").Result()

	if err == nil && err1 == nil {
		con.Success(c, "", "删除成功")
	} else {
		con.Error(c, "删除失败")
	}

}

/**
* 新增值
**/
func (con *listController) Add(c *gin.Context) {
	key := c.PostForm("key")
	value := c.PostForm("value")

	ctx := context.Background()

	_, err := global.UseClient.Client.LPush(ctx, key, value).Result()

	if err == nil {
		con.Success(c, "", "添加成功")
	} else {
		con.Error(c, "添加失败")
	}

}
