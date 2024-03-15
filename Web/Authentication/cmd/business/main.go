package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	cfg "course.project/authentication/config"
	cache "course.project/authentication/internal/common/cache"
	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {

}

func initCfg() {
	err := cfg.InitConfig()
	if err != nil {
		panic("init config error: " + err.Error())
	}
}

func initCache() {
	err := cache.Ca.InitCache()
	if err != nil {
		panic("init cache error: " + err.Error())
	}
}

func main() {
	initCfg()
	initCache()
	defer cache.Ca.CloseCache()

	router := gin.Default()
	register(router)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
