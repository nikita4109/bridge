package bridge

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Bridge struct {
	chainClients map[uint64]*ethclient.Client
	privateKey   *ecdsa.PrivateKey

	ownerAddress common.Address
	tokens       map[TokenName]Token
	lzChainIDs   map[uint64]uint64
}

func NewBridge(chainRPCs map[uint64]string, privateKeyHex string) (*Bridge, error) {
	if len(privateKeyHex) > 2 && privateKeyHex[:2] == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	chainClients := make(map[uint64]*ethclient.Client)
	for chainID, rpc := range chainRPCs {
		client, err := ethclient.Dial(rpc)
		if err != nil {
			for _, c := range chainClients {
				c.Close()
			}
			return nil, fmt.Errorf("failed to connect to chain %d: %w", chainID, err)
		}
		chainClients[chainID] = client
	}

	tokens := map[TokenName]Token{
		PUFF: {
			Name: PUFF,
			Address: map[uint64]common.Address{
				EthChainID:    common.HexToAddress("0x31b6100F5F4466e6dAeb1edb2f2cE6e548cF8938"),
				MantleChainId: common.HexToAddress("0x26a6b0dcdCfb981362aFA56D581e4A7dBA3Be140"),
			},
			OFTAddress: map[uint64]common.Address{
				EthChainID:    common.HexToAddress("0x31b6100f5f4466e6daeb1edb2f2ce6e548cf8938"),
				MantleChainId: common.HexToAddress("0xe6fa1Be9dAa660C99268F3F47fc193ba066A3AAc"),
			},
			OFTType: OFTMNT,
		},
		CMETH: {
			Name: CMETH,
			Address: map[uint64]common.Address{
				EthChainID:    common.HexToAddress("0xe6829d9a7ee3040e1276fa75293bde931859e8fa"),
				MantleChainId: common.HexToAddress("0xE6829d9a7eE3040e1276Fa75293Bde931859e8fA"),
			},
			OFTAddress: map[uint64]common.Address{
				EthChainID:    common.HexToAddress("0x4aFA9620D0B79137383A7A9AB3477837d475e948"),
				MantleChainId: common.HexToAddress("0xE6829d9a7eE3040e1276Fa75293Bde931859e8fA"),
			},
			OFTType: OFTMNT,
		},
	}

	return &Bridge{
		chainClients: chainClients,
		privateKey:   privateKey,
		ownerAddress: crypto.PubkeyToAddress(*publicKeyECDSA),
		tokens:       tokens,
		lzChainIDs: map[uint64]uint64{
			1:    30101,
			5000: 30181,
		},
	}, nil
}

func (b *Bridge) Close() {
	for _, client := range b.chainClients {
		client.Close()
	}
}
