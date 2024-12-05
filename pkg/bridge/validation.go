package bridge

import (
	erc20 "bridge/pkg/erc"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (b *Bridge) validateChains(fromChainId, toChainId uint64) error {
	if _, ok := b.chainClients[fromChainId]; !ok {
		return fmt.Errorf("no client configured for source chain %d", fromChainId)
	}
	if _, ok := b.chainClients[toChainId]; !ok {
		return fmt.Errorf("no client configured for destination chain %d", toChainId)
	}
	if _, ok := b.lzChainIDs[fromChainId]; !ok {
		return fmt.Errorf("no lz chains configured for source chain %d", toChainId)
	}
	if _, ok := b.lzChainIDs[toChainId]; !ok {
		return fmt.Errorf("no lz chains configured for destination chain %d", toChainId)
	}
	return nil
}

func (b *Bridge) validateToken(token Token, fromChainId, toChainId uint64) error {
	if _, ok := token.OFTAddress[fromChainId]; !ok {
		return fmt.Errorf("token not supported on source chain %d", fromChainId)
	}
	if _, ok := token.OFTAddress[toChainId]; !ok {
		return fmt.Errorf("token not supported on destination chain %d", toChainId)
	}
	return nil
}

func (b *Bridge) checkBalance(
	fromChainId uint64,
	token Token,
	amount *big.Int,
) error {
	erc20, err := erc20.NewErc20(token.Address[fromChainId], b.chainClients[fromChainId])
	if err != nil {
		return fmt.Errorf("failed to create wrapper: %w", err)
	}

	balance, err := erc20.BalanceOf(&bind.CallOpts{}, b.ownerAddress)
	if err != nil {
		return fmt.Errorf("failed to get balance: %w", err)
	}

	if balance.Cmp(amount) < 0 {
		return fmt.Errorf("insufficient balance: have %v, need %v", balance, amount)
	}

	return nil
}

func (b *Bridge) checkAllowance(
	fromChainId uint64,
	token Token,
	amount *big.Int,
	spender common.Address,
) error {
	erc20, err := erc20.NewErc20(token.Address[fromChainId], b.chainClients[fromChainId])
	if err != nil {
		return fmt.Errorf("failed to create wrapper: %w", err)
	}

	allowance, err := erc20.Allowance(&bind.CallOpts{}, b.ownerAddress, spender)
	if err != nil {
		return fmt.Errorf("failed to get allowance: %w", err)
	}

	if allowance.Cmp(amount) < 0 {
		return fmt.Errorf("insufficient allowance: have %v, need %v", allowance, amount)
	}

	return nil
}
