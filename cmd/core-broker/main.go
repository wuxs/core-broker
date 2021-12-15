/*
 * Filename: /Users/sc/liuzhen/core-broker/cmd/core-broker/main.go
 * Path: /Users/sc/liuzhen/core-broker/cmd/core-broker
 * Created Date: Thursday, December 9th 2021, 2:31:00 pm
 * Author: sc
 *
 * Copyright (c) 2021 Your Company
 */

package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/tkeel-io/core-broker/pkg/server"
	"github.com/tkeel-io/core-broker/pkg/service"
	"github.com/tkeel-io/kit/app"
	"github.com/tkeel-io/kit/log"
	"github.com/tkeel-io/kit/transport"

	// User import.

	Topic_v1 "github.com/tkeel-io/core-broker/api/topic/v1"
	Entity_v1 "github.com/tkeel-io/core-broker/api/ws/v1"

	openapi "github.com/tkeel-io/core-broker/api/openapi/v1"
)

var (
	// Name app.
	Name string
	// HTTPAddr string.
	HTTPAddr string
	// GRPCAddr string.
	GRPCAddr string
)

func init() {
	flag.StringVar(&Name, "name", "core-broker", "app name.")
	flag.StringVar(&HTTPAddr, "http_addr", ":31234", "http listen address.")
	flag.StringVar(&GRPCAddr, "grpc_addr", ":31233", "grpc listen address.")
}

func main() {
	flag.Parse()

	httpSrv := server.NewHTTPServer(HTTPAddr)
	grpcSrv := server.NewGRPCServer(GRPCAddr)
	serverList := []transport.Server{httpSrv, grpcSrv}

	app := app.New(Name,
		&log.Conf{
			App:   Name,
			Level: "debug",
			Dev:   true,
		},
		serverList...,
	)

	{ // User service
		OpenapiSrv := service.NewOpenapiService()
		openapi.RegisterOpenapiHTTPServer(httpSrv.Container, OpenapiSrv)
		openapi.RegisterOpenapiServer(grpcSrv.GetServe(), OpenapiSrv)

		EntitySrv := service.NewEntityService()
		go EntitySrv.Run()
		Entity_v1.RegisterEntityHTTPServer(httpSrv.Container, EntitySrv)

		TopicSrv := service.NewTopicService()
		Topic_v1.RegisterTopicHTTPServer(httpSrv.Container, TopicSrv)
		Topic_v1.RegisterTopicServer(grpcSrv.GetServe(), TopicSrv)
	}

	if err := app.Run(context.TODO()); err != nil {
		panic(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
	<-stop

	if err := app.Stop(context.TODO()); err != nil {
		panic(err)
	}
}
