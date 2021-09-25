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
	admin.POST("/save_config", controllers.Ic.SaveConfig)
	admin.POST("/switch", controllers.Ic.Switch)
	admin.POST("/search_key", controllers.Ic.SearchKey)

	admin.GET("/show/string", controllers.Ic.StringShow)
	admin.GET("/show/list", controllers.Ic.ListShow)
	admin.GET("/show/set", controllers.Ic.ListShow)

	return router
}
