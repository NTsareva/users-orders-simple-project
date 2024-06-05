package main

import (
	"context"
	"fmt"
	"net"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"user-service/ent"
	"user-service/proto/user-service/proto"
)

type Config struct {
	Database struct {
		User     string
		Password string
		Host     string
		Port     int
		DBName   string
	}
	Server struct {
		Port int
	}
}

type server struct {
	proto.UnimplementedUserServiceServer
	client *ent.Client
	logger *zap.Logger
}

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	var config Config
	if _, err := toml.DecodeFile("../configs/config.toml", &config); err != nil {
		logger.Fatal("error loading config", zap.Error(err))
	}

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName)
	client, err := ent.Open("postgres", dataSourceName)
	if err != nil {
		logger.Fatal("error connecting to the database", zap.Error(err))
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatal("error creating schema", zap.Error(err))
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Database.Port))
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}

	serv := grpc.NewServer()
	proto.RegisterUserServiceServer(serv, &server{client: client, logger: logger})
	logger.Info("user server started", zap.Int("port", config.Server.Port))
	if err := serv.Serve(listener); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
}