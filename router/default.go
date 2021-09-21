/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:17:35
 */
package router

import (
	"goredismanager/controllers"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/api/", "/metrics"})))

	admin := router.Group("/admin")
	admin.GET("/index", controllers.Ic.Show)

	return router
}
