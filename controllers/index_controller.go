/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:47:35
 */
package controllers

import (
	"context"
	"net/http"
	"strings"

	"goredismanager/comment"
	"goredismanager/global"
	"goredismanager/model"
	"goredismanager/service"

	"github.com/gin-gonic/gin"
)

type indexController struct {
	BaseController
}

var Ic = indexController{}

func (con *indexController) Show(c *gin.Context) {

	dbSlice := make([]struct{}, 16)
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"services":     global.RedisServiceStorage,
		"service_name": global.UseClient.ConnectName,
		"db":           global.UseClient.Db,
		"db_slice":     dbSlice,
		"account_num":  len(global.Accounts),
	})

}

func (con *indexController) SaveConfig(c *gin.Context) {

	var serviceReq model.ServiceConfigReq
	var err error

	err = con.FormBind(c, &serviceReq)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	err = service.AddServiceConf(serviceReq)
	if err != nil {
		con.Error(c, err.Error())
		return
	}
	con.Success(c, "", "添加连接成功")
}

func (con *indexController) SaveConfigFile(c *gin.Context) {

	len := len(global.RedisServiceStorage)
	confs := make([]model.ServiceConfigReq, len)

	i := 0
	for name, item := range global.RedisServiceStorage {

		hostPort := strings.Split(item.Config.Addr, ":")

		confs[i] = model.ServiceConfigReq{
			ServiceName: name,
			Host:        hostPort[0],
			Port:        hostPort[1],
			UseSsh:      item.UseSsh,
			Password:    item.Config.Password,
			SSHConfig:   item.SSHConfig,
		}
		i++
	}
	global.ConfigViper.Set("connections", confs)
	err := global.ConfigViper.WriteConfig()
	if err != nil {
		con.Error(c, err.Error())
		return
	}
	con.Success(c, "", "保存配置成功")
}

func (con *indexController) Switch(c *gin.Context) {

	var (
		err    error
		serReq model.ServiceSwitchReq
	)

	err = con.FormBind(c, &serReq)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	err = service.ServiceSwitch(serReq)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, "", "切换连接成功")
}

func (con *indexController) SearchKey(c *gin.Context) {

	var (
		err    error
		req    model.RedisKeyReq
		result []string
	)

	err = con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	result, err = service.SearchKeyType(req)

	if err != nil {
		con.AjaxReturn(c, AJAXFAIL, "Key值不存在")
		return
	}

	gen := comment.NewTrie()

	for _, v := range result {

		stringSlice := strings.Split(v, ":")

		gen.Insert(stringSlice, v)

	}

	resultSlice := comment.GetOne(gen.Root.Children, "")
	con.AjaxReturn(c, AJAXSUCCESS, resultSlice)
}

func (con *indexController) AddKey(c *gin.Context) {

	var (
		req model.AddKeyReq
		err error
	)

	err = con.FormBind(c, &req)

	if err != nil {
		con.Error(c, err.Error())
		return
	}

	if req.Pre != "" {
		req.Key = req.Pre + ":" + req.Key
	}

	err = service.AddKey(req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, "", "添加成功")
}

func (con *indexController) DelKey(c *gin.Context) {
	var (
		req model.DelKeyReq
		err error
	)

	err = con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	err = service.DelKey(req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.AjaxReturn(c, AJAXSUCCESS, gin.H{})
}

func (con *indexController) ShowKey(c *gin.Context) {
	key := c.Query("key")

	types, err := global.UseClient.Client.Type(context.Background(), key).Result()

	if err == nil {

		htmlString, data := service.TransView(types, key)

		c.HTML(http.StatusOK, htmlString, data)
	}

}
