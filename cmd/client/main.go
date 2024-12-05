package main

import (
	"bridge/pkg/bridge"
	"bridge/pkg/config"
	pb "bridge/proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", config.GRPCPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBridgeServiceClient(conn)

	resp, err := client.SendBridge(context.Background(), &pb.SendBridgeRequest{
		FromChainId: bridge.EthChainID,
		ToChainId:   bridge.MantleChainId,
		TokenName:   string(bridge.PUFF),
		Amount:      "1845415744840951057",
	})
	if err != nil {
		log.Fatalf("failed to send bridge: %v", err)
	}

	stream, err := client.GetBridgeStatus(context.Background(),
		&pb.GetBridgeStatusRequest{BridgeId: resp.BridgeId})
	if err != nil {
		log.Fatalf("failed to get status: %v", err)
	}

	for {
		update, err := stream.Recv()
		if err != nil {
			log.Printf("stream ended: %v", err)
			break
		}

		log.Printf("Status: %v, Hash: %s, Error: %s",
			update.Status, update.TransactionHash, update.Error)
	}
}
