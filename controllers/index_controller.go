/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:47:35
 */
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type indexController struct {
	BaseController
}

var Ic = indexController{}

func (con *indexController) Show(c *gin.Context) {

	c.HTML(http.StatusOK, "index/index.html", gin.H{})

}
