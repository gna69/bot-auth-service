package main

import (
	"context"
	"fmt"
	"net"

	"gitgub.com/gna69/bot-auth-service/internal/adapter/pg"
	"gitgub.com/gna69/bot-auth-service/internal/adapter/services"
	"gitgub.com/gna69/bot-auth-service/internal/driver/grpc_server"
	"gitgub.com/gna69/bot-auth-service/proto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/auth.proto

func main() {
	ctx := context.Background()

	pgConn, err := pg.NewPgClient(ctx)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Str("pg", "success connection").Send()

	userService := services.NewUserService(pgConn)
	groupService := services.NewGroupService(pgConn)
	log.Debug().Str("services", "success creating")

	grpcSrv := grpc_server.NewGrpcServer(userService, groupService)

	srv := grpc.NewServer()
	proto.RegisterAuthServiceServer(srv, grpcSrv)
	log.Debug().Str("server", "success registration")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", cfg.GrpcPort))
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	log.Info().Msgf("start grpc server on %s port", cfg.GrpcPort)
	err = srv.Serve(lis)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
}
