package main

import (
	"bridge/pkg/bridge"
	"bridge/pkg/config"
	"bridge/pkg/service"
	pb "bridge/proto"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	privateKey := os.Getenv("PKEY")
	if privateKey == "" {
		log.Fatal("PKEY is empty")
	}

	lzBridge, err := bridge.NewBridge(config.ChainRPCs, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	defer lzBridge.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	bridgeService := service.NewBridgeService(lzBridge)
	pb.RegisterBridgeServiceServer(grpcServer, bridgeService)

	log.Printf("Starting gRPC server on :%v", config.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
