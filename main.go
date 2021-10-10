/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 19:57:42
 */
package main

import (
	"context"
	"goredismanager/global"
	"goredismanager/router"
	"goredismanager/web"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

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
