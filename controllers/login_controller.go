/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-07 17:20:54
 */
package controllers

import (
	"encoding/json"
	"fmt"
	"goredismanager/global"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type loginController struct {
	BaseController
}

var LoginC = loginController{}

func (con *loginController) ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login/login.html", gin.H{})
}

func (con *loginController) Login(c *gin.Context) {
	
	username := c.PostForm("account")
	password := c.PostForm("password")

	if _, ok := global.Accounts[username]; ok {
		if global.Accounts[username] == password {
			userInfo := make(map[string]interface{})
			userInfo["username"] = username
			//session 存储数据
			session := sessions.Default(c)
			userstr, _ := json.Marshal(userInfo)

			session.Set("userInfo", string(userstr))
			session.Save()

			con.Success(c, "/index", "登录成功")
		} else {
			con.Error(c, "账号密码错误")
		}
	} else {
		con.Error(c, "账号密码错误")
	}

}

func (con *loginController) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userInfo")
	session.Save()
	con.Success(c, "login", "退出成功")
}
