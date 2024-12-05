package service

import (
	"bridge/pkg/bridge"
	pb "bridge/proto"
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/google/uuid"
)

type BridgeService struct {
	pb.UnimplementedBridgeServiceServer
	bridge *bridge.Bridge

	mu            sync.RWMutex
	bridgeUpdates map[string]chan *pb.BridgeStatusUpdate
}

func NewBridgeService(b *bridge.Bridge) *BridgeService {
	return &BridgeService{
		bridge:        b,
		bridgeUpdates: make(map[string]chan *pb.BridgeStatusUpdate),
	}
}

func (s *BridgeService) SendBridge(
	ctx context.Context,
	req *pb.SendBridgeRequest,
) (*pb.SendBridgeResponse, error) {
	amount, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid amount format")
	}

	bridgeID := uuid.New().String()

	updates, err := s.bridge.Send(
		req.FromChainId,
		req.ToChainId,
		bridge.TokenName(req.TokenName),
		amount,
	)
	if err != nil {
		return nil, err
	}

	s.mu.Lock()
	s.bridgeUpdates[bridgeID] = updates
	s.mu.Unlock()

	return &pb.SendBridgeResponse{BridgeId: bridgeID}, nil
}

func (s *BridgeService) GetBridgeStatus(
	req *pb.GetBridgeStatusRequest,
	stream pb.BridgeService_GetBridgeStatusServer,
) error {
	s.mu.RLock()
	updates, ok := s.bridgeUpdates[req.BridgeId]
	s.mu.RUnlock()

	if !ok {
		return fmt.Errorf("bridge ID not found")
	}

	for update := range updates {
		if err := stream.Send(update); err != nil {
			return err
		}

		if update.Status == pb.BridgeStatusUpdate_RECEIVED ||
			update.Status == pb.BridgeStatusUpdate_FAILED {
			s.mu.Lock()
			delete(s.bridgeUpdates, req.BridgeId)
			s.mu.Unlock()
			break
		}
	}

	return nil
}
