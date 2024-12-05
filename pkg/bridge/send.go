package bridge

import (
	erc20 "bridge/pkg/erc"
	"bridge/pkg/oftwrapper"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	pb "bridge/proto"
)

func (b *Bridge) Send(
	fromChainId uint64,
	toChainId uint64,
	tokenName TokenName,
	amount *big.Int,
) (chan *pb.BridgeStatusUpdate, error) {
	if err := b.validateChains(fromChainId, toChainId); err != nil {
		return nil, err
	}

	token, ok := b.tokens[tokenName]
	if !ok {
		return nil, fmt.Errorf("there is no token: %v", tokenName)
	}

	if err := b.validateToken(token, fromChainId, toChainId); err != nil {
		return nil, err
	}

	if err := b.checkBalance(fromChainId, token, amount); err != nil {
		return nil, err
	}

	return b.send(fromChainId, toChainId, token, amount), nil
}

func approveInfToken(
	client *ethclient.Client,
	auth *bind.TransactOpts,
	tokenAddress common.Address,
	spender common.Address,
) (*types.Transaction, error) {
	token, err := erc20.NewErc20(tokenAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate token contract: %w", err)
	}

	maxUint256 := new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))
	tx, err := token.Approve(auth, spender, maxUint256)
	if err != nil {
		return nil, fmt.Errorf("failed to approve tokens: %w", err)
	}

	return tx, nil
}

func (b *Bridge) send(
	fromChainId uint64,
	toChainId uint64,
	token Token,
	amount *big.Int,
) chan *pb.BridgeStatusUpdate {
	updates := make(chan *pb.BridgeStatusUpdate, 10)
	go b.processSend(
		fromChainId,
		toChainId,
		token,
		amount,
		updates,
	)

	return updates
}

func (b *Bridge) processSend(
	fromChainId uint64,
	toChainId uint64,
	token Token,
	amount *big.Int,
	updates chan *pb.BridgeStatusUpdate,
) {
	defer close(updates)

	updates <- &pb.BridgeStatusUpdate{Status: pb.BridgeStatusUpdate_PENDING}

	srcClient := b.chainClients[fromChainId]
	dstClient := b.chainClients[toChainId]

	auth, err := bind.NewKeyedTransactorWithChainID(b.privateKey, big.NewInt(int64(fromChainId)))
	if err != nil {
		updates <- &pb.BridgeStatusUpdate{
			Status: pb.BridgeStatusUpdate_FAILED,
			Error:  fmt.Sprintf("failed to create auth: %s", err),
		}
		return
	}

	currentBlockNumber, err := dstClient.BlockNumber(context.Background())
	if err != nil {
		updates <- &pb.BridgeStatusUpdate{
			Status: pb.BridgeStatusUpdate_FAILED,
			Error:  fmt.Sprintf("failed to get block number: %s", err),
		}
		return
	}

	var tx *types.Transaction
	tx, err = b.sendBridgeTx(
		auth,
		fromChainId,
		toChainId,
		token,
		srcClient,
		amount,
	)
	if err != nil {
		updates <- &pb.BridgeStatusUpdate{
			Status: pb.BridgeStatusUpdate_FAILED,
			Error:  fmt.Sprintf("failed to send transaction: %s", err),
		}
		return
	}

	receipt, err := bind.WaitMined(context.Background(), srcClient, tx)
	if err != nil {
		updates <- &pb.BridgeStatusUpdate{
			Status: pb.BridgeStatusUpdate_FAILED,
			Error:  fmt.Sprintf("failed to get receipt: %s", err),
		}
		return
	}

	if receipt.Status == types.ReceiptStatusSuccessful {
		guid, err := extractGuidFromReceipt(receipt)
		if err != nil {
			updates <- &pb.BridgeStatusUpdate{
				Status:          pb.BridgeStatusUpdate_FAILED,
				Error:           fmt.Sprintf("failed to extract GUID: %s", err),
				TransactionHash: tx.Hash().Hex(),
			}
			return
		}

		updates <- &pb.BridgeStatusUpdate{
			Status:          pb.BridgeStatusUpdate_SENT,
			TransactionHash: tx.Hash().Hex(),
		}

		err = waitForDestinationTx(
			context.Background(),
			currentBlockNumber,
			guid,
			dstClient,
			updates,
		)
		if err != nil {
			updates <- &pb.BridgeStatusUpdate{
				Status:          pb.BridgeStatusUpdate_FAILED,
				Error:           fmt.Sprintf("failed to confirm destination transaction: %s", err),
				TransactionHash: tx.Hash().Hex(),
			}
			return
		}
	} else {
		updates <- &pb.BridgeStatusUpdate{
			Status:          pb.BridgeStatusUpdate_FAILED,
			Error:           fmt.Sprintf("transaction failed on source chain"),
			TransactionHash: tx.Hash().Hex(),
		}
		return
	}
}

func (b *Bridge) sendBridgeTx(
	auth *bind.TransactOpts,
	fromChainId uint64,
	toChainId uint64,
	token Token,
	srcClient *ethclient.Client,
	amount *big.Int,
) (*types.Transaction, error) {
	minAmount := big.NewInt(0)

	wrapper, err := oftwrapper.NewOftwrapper(token.OFTAddress[fromChainId], srcClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create wrapper: %s", err)
	}

	recipient := auth.From.Bytes()
	var recipientBytes32 [32]byte
	copy(recipientBytes32[:], recipient)

	if token.OFTType == OFTMNT {
		sendParam := oftwrapper.SendParam{
			DstEid:       uint32(b.lzChainIDs[toChainId]),
			To:           recipientBytes32,
			AmountLD:     amount,
			MinAmountLD:  minAmount,
			ExtraOptions: []byte{},
			ComposeMsg:   []byte{},
			OftCmd:       []byte{},
		}

		fee, err := wrapper.QuoteSend(
			&bind.CallOpts{},
			sendParam,
			false,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to quote send fees: %s", err)
		}

		auth.Value = fee.NativeFee

		tx, err := wrapper.Send(
			auth,
			sendParam,
			fee,
			auth.From,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to send: %s", err)
		}

		return tx, nil
	}

	return nil, fmt.Errorf("this type OFT does not supported")
}
