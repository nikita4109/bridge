package bridge

import (
	"bridge/proto"
	"context"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestGetTxFromLogs(t *testing.T) {
	client, err := ethclient.Dial("https://rpc.mantle.xyz")
	assert.True(t, err == nil)

	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(72636823),
		ToBlock:   big.NewInt(72636823),
	})

	guid := common.HexToHash("0xdc917291ca8fe940e76977e228ad74eb164854e6ca6bfbc587f8a43a62eb818d")

	txHash, found := getDestinationTxFromLogs(logs, guid)
	log.Println(txHash.Hex())
	assert.True(t, found)
}

func TestMonitor(t *testing.T) {
	client, err := ethclient.Dial("https://rpc.mantle.xyz")
	assert.True(t, err == nil)

	err = waitForDestinationTx(context.Background(), 72637780, common.HexToHash("0x94F8174D1CFB963E19DBAC911A4EBCC54694BEB96DC40E93990584DEFC40544E"), client, make(chan<- *proto.BridgeStatusUpdate, 3))
	assert.True(t, err == nil)
}
