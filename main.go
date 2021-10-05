/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 19:57:42
 */
package main

import (
	"context"
	"flag"
	"goredismanager/router"
	"goredismanager/web"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	var host string
	var port string

	flag.StringVar(&host, "h", "127.0.0.1", "ip地址默认：127.0.0.1")
	flag.StringVar(&port, "p", "8000", "端口地址默认：8000")
	flag.Parse()

	router := router.Init()

	router.StaticFS("/statics", web.StaticsFs)
	router.HTMLRender = web.LoadTemplates()

	srv := &http.Server{
		Addr:    host + ":" + port,
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
