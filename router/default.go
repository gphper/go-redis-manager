/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:17:35
 */
package router

import (
	"goredismanager/controllers"
	"goredismanager/global"
	"goredismanager/middleware"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	store := cookie.NewStore([]byte("goredismanagerphper"))
	router.Use(gzip.Gzip(gzip.DefaultCompression), sessions.Sessions("goredismanager", store))

	router.GET("/login", controllers.LoginC.ShowLogin)
	router.POST("/login", controllers.LoginC.Login)
	router.POST("/login_out", controllers.LoginC.LoginOut)

	app := router.Group("/")

	if len(global.Accounts) > 0 {
		//用户验证
		app.Use(middleware.UserAuth())
	}

	app.GET("/index", controllers.Ic.Show)
	app.POST("/save_config", controllers.Ic.SaveConfig)
	app.POST("/save_configfile", controllers.Ic.SaveConfigFile)
	app.POST("/switch", controllers.Ic.Switch)
	app.POST("/search_key", controllers.Ic.SearchKey)
	app.POST("/add_key", controllers.Ic.AddKey)
	app.POST("/del_key", controllers.Ic.DelKey)
	app.GET("/show_key", controllers.Ic.ShowKey)

	string_admin := app.Group("/string")
	{
		string_admin.POST("/ttl", controllers.Sc.Ttl)
		string_admin.POST("/save", controllers.Sc.Save)
	}

	list_admin := app.Group("/list")
	{
		list_admin.POST("/del", controllers.Lc.Del)
		list_admin.POST("/add", controllers.Lc.Add)
	}

	set_admin := app.Group("/set")
	{
		set_admin.POST("/del", controllers.Setc.Del)
		set_admin.POST("/add", controllers.Setc.Add)
	}

	zset_admin := app.Group("/zset")
	{
		zset_admin.POST("/del", controllers.Zsetc.Del)
		zset_admin.POST("/add", controllers.Zsetc.Add)
	}

	hash_admin := app.Group("/hash")
	{
		hash_admin.POST("/del", controllers.Hashc.Del)
		hash_admin.POST("/add", controllers.Hashc.Add)
	}

	stream_admin := app.Group("/stream")
	{
		stream_admin.POST("/del", controllers.Streamc.Del)
		stream_admin.POST("/add", controllers.Streamc.Add)
	}

	return router
}
