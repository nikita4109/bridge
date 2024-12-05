package bridge

import (
	pb "bridge/proto"
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func extractGuidFromReceipt(receipt *types.Receipt) (common.Hash, error) {
	oftSentSig := crypto.Keccak256Hash([]byte("OFTSent(bytes32,uint32,address,uint256,uint256)"))

	for _, log := range receipt.Logs {
		if len(log.Topics) >= 3 && log.Topics[0] == oftSentSig {
			return log.Topics[1], nil
		}
	}
	return common.Hash{}, fmt.Errorf("OFTSent event not found in receipt")
}

func waitForDestinationTx(
	ctx context.Context,
	currentBlockNumber uint64,
	guid common.Hash,
	dstClient *ethclient.Client,
	updates chan<- *pb.BridgeStatusUpdate,
) error {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for destination transaction")
		case <-ticker.C:
			latestBlock, err := dstClient.BlockNumber(ctx)
			if err != nil {
				log.Print(err)
				continue
			}

			latestBlock = min(latestBlock, currentBlockNumber+10)

			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(currentBlockNumber)),
				ToBlock:   big.NewInt(int64(latestBlock)),
			}

			logs, err := dstClient.FilterLogs(ctx, query)
			if err != nil {
				log.Print(err)
				continue
			}

			txHash, found := getDestinationTxFromLogs(logs, guid)
			if found {
				updates <- &pb.BridgeStatusUpdate{
					Status:          pb.BridgeStatusUpdate_RECEIVED,
					TransactionHash: txHash.Hex(),
				}

				return nil
			}

			currentBlockNumber = latestBlock + 1
		}
	}
}

func getDestinationTxFromLogs(
	logs []types.Log,
	guid common.Hash,
) (common.Hash, bool) {
	oftReceivedSig := crypto.Keccak256Hash([]byte("OFTReceived(bytes32,uint32,address,uint256)"))

	for _, log := range logs {
		if len(log.Topics) >= 2 && log.Topics[0] == oftReceivedSig && log.Topics[1] == guid {
			return log.TxHash, true
		}
	}

	return common.Hash{}, false
}
