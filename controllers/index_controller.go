/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:47:35
 */
package controllers

import (
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
	})

}

func (con *indexController) SaveConfig(c *gin.Context) {

	var serviceReq model.ServiceConfigReq
	var err error

	err = con.FormBind(c, &serviceReq)
	if err != nil {
		con.Error(c, err.Error())
	}

	err = service.AddServiceConf(serviceReq)
	if err != nil {
		con.Error(c, err.Error())
		return
	}
	con.Success(c, "", "添加连接成功")
}

func (con *indexController) Switch(c *gin.Context) {

	var (
		err    error
		serReq model.ServiceSwitchReq
	)

	err = con.FormBind(c, &serReq)
	if err != nil {
		con.Error(c, err.Error())
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
		result string
	)

	err = con.FormBind(c, &req)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	result, err = service.SearchKeyType(req)
	if err != nil || result == "none" {
		con.AjaxReturn(c, AJAXFAIL, "当前key值不存在")
		return
	}

	req.Key += "_" + result

	stringSlice := strings.Split(req.Key, ":")

	gen := comment.NewTrie()
	gen.Insert(stringSlice)

	// g := comment.NewTrie()
	// g.Insert([]string{"hello", "hello1", "hello2", "hello3_string"})
	// g.Insert([]string{"hello", "hello1", "hello2", "hello33_string"})
	// g.Insert([]string{"hello", "hello1", "hello22", "hello333_string"})
	// g.Insert([]string{"hello", "hello1", "hello22_string"})

	// json.NewEncoder(os.Stdout).Encode()
	// fmt.Println(comment.GetOne(g.Root, ""))

	resultSlice := comment.GetOne(gen.Root, "")
	// resultSlice := comment.GetOne(g.Root, "")

	con.AjaxReturn(c, AJAXSUCCESS, resultSlice)
}

func (con *indexController) StringShow(c *gin.Context) {

}

func (con *indexController) ListShow(c *gin.Context) {

}

func (con *indexController) SetShow(c *gin.Context) {

}
