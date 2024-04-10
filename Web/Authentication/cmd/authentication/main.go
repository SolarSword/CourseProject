package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	cfg "course.project/authentication/config"
	cache "course.project/authentication/internal/common/cache"
	DB "course.project/authentication/internal/common/db"
	"course.project/authentication/internal/processor/login"
	pb "course.project/authentication/proto/authentication"
)

const (
	port = 6699
)

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

func initCache() {
	err := cache.Ca.InitCache()
	if err != nil {
		panic("init cache error: " + err.Error())
	}
}

func main() {
	initCfg()
	initDB()
	initCache()
	defer DB.Db.CloseDB()
	defer cache.Ca.CloseCache()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}
	grpcServer := grpc.NewServer([]grpc.ServerOption{})
	pb.RegisterAuthenticationServer(grpcServer, login.NewServer())
	grpcServer.Serve(lis)
}
