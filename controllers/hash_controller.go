/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:47:35
 */
package controllers

import (
	"context"
	"goredismanager/global"

	"github.com/gin-gonic/gin"
)

type hashController struct {
	BaseController
}

var Hashc = hashController{}

/**
* 删除特定值
**/
func (con *hashController) Del(c *gin.Context) {
	key := c.PostForm("key")
	member := c.PostForm("member")
	ctx := context.Background()

	res, _ := global.UseClient.Client.HDel(ctx, key, member).Result()

	if res > 0 {
		con.Success(c, "", "删除成功")
	} else {
		con.Error(c, "删除失败")
	}

}

/**
* 新增值
**/
func (con *hashController) Add(c *gin.Context) {
	key := c.PostForm("key")
	field := c.PostForm("field")
	value := c.PostForm("value")

	ctx := context.Background()

	_, err := global.UseClient.Client.HSet(ctx, key, field, value).Result()

	if err == nil {
		con.Success(c, "", "添加成功")
	} else {
		con.Error(c, "添加失败")
	}

}
