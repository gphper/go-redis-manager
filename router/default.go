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
	admin.POST("/save_configfile", controllers.Ic.SaveConfigFile)
	admin.POST("/switch", controllers.Ic.Switch)
	admin.POST("/search_key", controllers.Ic.SearchKey)
	admin.POST("/add_key", controllers.Ic.AddKey)
	admin.POST("/del_key", controllers.Ic.DelKey)
	admin.GET("/show_key", controllers.Ic.ShowKey)

	string_admin := admin.Group("/string")
	{
		string_admin.POST("/ttl", controllers.Sc.Ttl)
		string_admin.POST("/save", controllers.Sc.Save)
	}

	list_admin := admin.Group("/list")
	{
		list_admin.POST("/del", controllers.Lc.Del)
		list_admin.POST("/add", controllers.Lc.Add)
	}

	set_admin := admin.Group("/set")
	{
		set_admin.POST("/del", controllers.Setc.Del)
		set_admin.POST("/add", controllers.Setc.Add)
	}

	zset_admin := admin.Group("/zset")
	{
		zset_admin.POST("/del", controllers.Zsetc.Del)
		zset_admin.POST("/add", controllers.Zsetc.Add)
	}

	hash_admin := admin.Group("/hash")
	{
		hash_admin.POST("/del", controllers.Hashc.Del)
		hash_admin.POST("/add", controllers.Hashc.Add)
	}

	stream_admin := admin.Group("/stream")
	{
		stream_admin.POST("/del", controllers.Streamc.Del)
		stream_admin.POST("/add", controllers.Streamc.Add)
	}

	return router
}
