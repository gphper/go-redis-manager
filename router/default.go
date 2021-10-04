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

	string_admin := admin.Group("/string")
	{
		string_admin.GET("/show", controllers.Sc.Show)
		string_admin.POST("/ttl", controllers.Sc.Ttl)
	}

	list_admin := admin.Group("/list")
	{
		list_admin.GET("/show", controllers.Lc.Show)
		list_admin.POST("/del", controllers.Lc.Del)
	}

	set_admin := admin.Group("/set")
	{
		set_admin.GET("/show", controllers.Setc.Show)
		set_admin.POST("/del", controllers.Setc.Del)
	}

	zset_admin := admin.Group("/zset")
	{
		zset_admin.GET("/show", controllers.Zsetc.Show)
		zset_admin.POST("/del", controllers.Zsetc.Del)
	}

	hash_admin := admin.Group("/hash")
	{
		hash_admin.GET("/show", controllers.Hashc.Show)
		hash_admin.POST("/del", controllers.Hashc.Del)
	}

	return router
}
