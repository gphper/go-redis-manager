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

	"goredismanager/common"
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

	if len == 0 {
		con.Error(c, "当前没有有效的连接配置信息")
		return
	}

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

	type ResultStruct struct {
		Course      uint64
		ResultSlice []common.Node
		Key         string
	}

	var (
		err    error
		req    model.RedisKeyReq
		result []string
		rs     ResultStruct
		course uint64
	)

	err = con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	result, course, err = service.SearchKeyType(req, c)

	if err != nil {
		con.AjaxReturn(c, AJAXFAIL, "Key值不存在")
		return
	}

	gen := common.NewTrie()

	for _, v := range result {

		stringSlice := strings.Split(v, ":")

		gen.Insert(stringSlice, v)

	}

	rs.ResultSlice = common.GetOne(gen.Root.Children, "")
	rs.Course = course
	rs.Key = req.SearchKey

	con.AjaxReturn(c, AJAXSUCCESS, rs)
}

func (con *indexController) AddKey(c *gin.Context) {

	var (
		req model.AddKeyReq
		err error
		val interface{}
		ctx context.Context
	)

	err = con.FormBind(c, &req)

	if err != nil {
		con.Error(c, err.Error())
		return
	}

	if req.Pre != "" {
		req.Key = req.Pre + ":" + req.Key
	}

	val, _ = c.Get("username")

	ctx = context.WithValue(context.Background(), "username", val)

	err = service.AddKey(req, ctx)

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
		val interface{}
		ctx context.Context
	)

	err = con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	val, _ = c.Get("username")

	ctx = context.WithValue(context.Background(), "username", val)

	err = service.DelKey(req, ctx)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.AjaxReturn(c, AJAXSUCCESS, gin.H{})
}

func (con *indexController) ShowKey(c *gin.Context) {
	key := c.Query("key")

	val, _ := c.Get("username")

	ctx := context.WithValue(context.Background(), "username", val)

	types, err := global.UseClient.Client.Type(ctx, key).Result()

	if err == nil {

		htmlString, data := service.TransView(types, key, ctx)

		c.HTML(http.StatusOK, htmlString, data)
	}

}
