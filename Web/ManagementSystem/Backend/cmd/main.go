package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	cfg "course.project/management_system/config"
	DB "course.project/management_system/internal/common/db"
	"course.project/management_system/internal/common/logger"
	"course.project/management_system/internal/phase"
	"course.project/management_system/internal/semester"
)

func register(r *gin.Engine) {
	// phase
	r.POST("/api/v1/start_course_selection", phase.StartCourseSelectionPhase)
	r.POST("/api/v1/end_course_selection", phase.EndCourseSelection)
	r.GET("/api/v1/get_phase", phase.GetPhase)
	// semester
	r.POST("/api/v1/create_semester", semester.CreateSemester)
}

func initDB() {
	err := DB.Db.InitDB()
	if err != nil {
		panic("init database error: " + err.Error())
	}
}

func initCfg() {
	err := cfg.InitConfig()
	if err != nil {
		panic("init config error: " + err.Error())
	}
}

func main() {
	initCfg()
	initDB()
	defer DB.Db.CloseDB()

	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	router := gin.Default()
	// the middleware setting should be before the API register
	router.Use(logger.Logger)
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
