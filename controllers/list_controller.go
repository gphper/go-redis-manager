/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:47:35
 */
package controllers

import (
	"github.com/gin-gonic/gin"
)

type listController struct {
	BaseController
}

var Lc = listController{}

func (con *listController) Show(c *gin.Context) {

}
