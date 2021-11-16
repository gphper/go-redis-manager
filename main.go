/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 19:57:42
 */
package main

import (
	"context"
	"fmt"
	"goredismanager/global"
	"goredismanager/router"
	"goredismanager/web"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	release bool = true
)

func main() {

	if release {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	router := router.Init()

	router.StaticFS("/statics", web.StaticsFs)
	router.HTMLRender = web.LoadTemplates()

	srv := &http.Server{
		Addr:    global.HostName + ":" + global.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}

	}()

	go func() {
		//windows 自动打开访问地址
		if runtime.GOOS == "windows" {
			url := "http://" + srv.Addr + "/index"
			fmt.Println("访问地址：" + url)
			cmd := exec.Command("cmd", "/c start "+url)
			cmd.Start()
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
