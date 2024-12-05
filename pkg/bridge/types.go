package bridge

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	EthChainID    = 1
	MantleChainId = 5000
)

type TokenName string

const (
	PUFF  TokenName = "Puff"
	METH  TokenName = "mETH"
	CMETH TokenName = "cmETH"
)

type OFTType string

const (
	OFTMNT     OFTType = "mnt"
	OFTV1      OFTType = "v1"
	ProxyOFTV1 OFTType = "proxy_v1"
	OFTV2      OFTType = "v2"
	ProxyOFTV2 OFTType = "proxy_v2"
)

type Token struct {
	Name       TokenName
	Address    map[uint64]common.Address
	OFTAddress map[uint64]common.Address
	OFTType    OFTType
}
