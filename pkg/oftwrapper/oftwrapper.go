// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oftwrapper

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ICommonOFTLzCallParams is an auto generated low-level Go binding around an user-defined struct.
type ICommonOFTLzCallParams struct {
	RefundAddress     common.Address
	ZroPaymentAddress common.Address
	AdapterParams     []byte
}

// IOFTWrapperFeeObj is an auto generated low-level Go binding around an user-defined struct.
type IOFTWrapperFeeObj struct {
	CallerBps *big.Int
	Caller    common.Address
	PartnerId [2]byte
}

// MessagingFee is an auto generated low-level Go binding around an user-defined struct.
type MessagingFee struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}

// MessagingReceipt is an auto generated low-level Go binding around an user-defined struct.
type MessagingReceipt struct {
	Guid  [32]byte
	Nonce uint64
	Fee   MessagingFee
}

// OFTReceipt is an auto generated low-level Go binding around an user-defined struct.
type OFTReceipt struct {
	AmountSentLD     *big.Int
	AmountReceivedLD *big.Int
}

// SendParam is an auto generated low-level Go binding around an user-defined struct.
type SendParam struct {
	DstEid       uint32
	To           [32]byte
	AmountLD     *big.Int
	MinAmountLD  *big.Int
	ExtraOptions []byte
	ComposeMsg   []byte
	OftCmd       []byte
}

// OftwrapperMetaData contains all meta data concerning the Oftwrapper contract.
var OftwrapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_defaultBps\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WrapperFeeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wrapperFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"callerFee\",\"type\":\"uint256\"}],\"name\":\"WrapperFees\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BPS_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_UINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_useZro\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"estimateSendFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_useZro\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"estimateSendFeeV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_callerBps\",\"type\":\"uint256\"}],\"name\":\"getAmountAndFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wrapperFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callerFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oftBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nativeOft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"sendNativeOFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nativeOft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"}],\"internalType\":\"structICommonOFT.LzCallParams\",\"name\":\"_callParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"sendNativeOFTFeeV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"sendOFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"}],\"internalType\":\"structICommonOFT.LzCallParams\",\"name\":\"_callParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"sendOFTFeeV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"}],\"internalType\":\"structICommonOFT.LzCallParams\",\"name\":\"_callParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"sendOFTV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxyOft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"sendProxyOFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxyOft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"}],\"internalType\":\"structICommonOFT.LzCallParams\",\"name\":\"_callParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"sendProxyOFTFeeV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxyOft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"}],\"internalType\":\"structICommonOFT.LzCallParams\",\"name\":\"_callParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"sendProxyOFTV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_defaultBps\",\"type\":\"uint256\"}],\"name\":\"setDefaultBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bps\",\"type\":\"uint256\"}],\"name\":\"setOFTBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_payInLzToken\",\"type\":\"bool\"}],\"name\":\"quoteSend\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"msgFee\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"_fee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"fee\",\"type\":\"tuple\"}],\"internalType\":\"structMessagingReceipt\",\"name\":\"msgReceipt\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTReceipt\",\"name\":\"oftReceipt\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// OftwrapperABI is the input ABI used to generate the binding from.
// Deprecated: Use OftwrapperMetaData.ABI instead.
var OftwrapperABI = OftwrapperMetaData.ABI

// Oftwrapper is an auto generated Go binding around an Ethereum contract.
type Oftwrapper struct {
	OftwrapperCaller     // Read-only binding to the contract
	OftwrapperTransactor // Write-only binding to the contract
	OftwrapperFilterer   // Log filterer for contract events
}

// OftwrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type OftwrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OftwrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OftwrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OftwrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OftwrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OftwrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OftwrapperSession struct {
	Contract     *Oftwrapper       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OftwrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OftwrapperCallerSession struct {
	Contract *OftwrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OftwrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OftwrapperTransactorSession struct {
	Contract     *OftwrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OftwrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type OftwrapperRaw struct {
	Contract *Oftwrapper // Generic contract binding to access the raw methods on
}

// OftwrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OftwrapperCallerRaw struct {
	Contract *OftwrapperCaller // Generic read-only contract binding to access the raw methods on
}

// OftwrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OftwrapperTransactorRaw struct {
	Contract *OftwrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOftwrapper creates a new instance of Oftwrapper, bound to a specific deployed contract.
func NewOftwrapper(address common.Address, backend bind.ContractBackend) (*Oftwrapper, error) {
	contract, err := bindOftwrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oftwrapper{OftwrapperCaller: OftwrapperCaller{contract: contract}, OftwrapperTransactor: OftwrapperTransactor{contract: contract}, OftwrapperFilterer: OftwrapperFilterer{contract: contract}}, nil
}

// NewOftwrapperCaller creates a new read-only instance of Oftwrapper, bound to a specific deployed contract.
func NewOftwrapperCaller(address common.Address, caller bind.ContractCaller) (*OftwrapperCaller, error) {
	contract, err := bindOftwrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OftwrapperCaller{contract: contract}, nil
}

// NewOftwrapperTransactor creates a new write-only instance of Oftwrapper, bound to a specific deployed contract.
func NewOftwrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*OftwrapperTransactor, error) {
	contract, err := bindOftwrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OftwrapperTransactor{contract: contract}, nil
}

// NewOftwrapperFilterer creates a new log filterer instance of Oftwrapper, bound to a specific deployed contract.
func NewOftwrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*OftwrapperFilterer, error) {
	contract, err := bindOftwrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OftwrapperFilterer{contract: contract}, nil
}

// bindOftwrapper binds a generic wrapper to an already deployed contract.
func bindOftwrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OftwrapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oftwrapper *OftwrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oftwrapper.Contract.OftwrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oftwrapper *OftwrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oftwrapper.Contract.OftwrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oftwrapper *OftwrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oftwrapper.Contract.OftwrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oftwrapper *OftwrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oftwrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oftwrapper *OftwrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oftwrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oftwrapper *OftwrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oftwrapper.Contract.contract.Transact(opts, method, params...)
}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_Oftwrapper *OftwrapperCaller) BPSDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oftwrapper.contract.Call(opts, &out, "BPS_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_Oftwrapper *OftwrapperSession) BPSDENOMINATOR() (*big.Int, error) {
	return _Oftwrapper.Contract.BPSDENOMINATOR(&_Oftwrapper.CallOpts)
}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_Oftwrapper *OftwrapperCallerSession) BPSDENOMINATOR() (*big.Int, error) {
	return _Oftwrapper.Contract.BPSDENOMINATOR(&_Oftwrapper.CallOpts)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Oftwrapper *OftwrapperCaller) MAXUINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oftwrapper.contract.Call(opts, &out, "MAX_UINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Oftwrapper *OftwrapperSession) MAXUINT() (*big.Int, error) {
	return _Oftwrapper.Contract.MAXUINT(&_Oftwrapper.CallOpts)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Oftwrapper *OftwrapperCallerSession) MAXUINT() (*big.Int, error) {
	return _Oftwrapper.Contract.MAXUINT(&_Oftwrapper.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Oftwrapper *OftwrapperCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oftwrapper.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Oftwrapper *OftwrapperSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Oftwrapper.Contract.BalanceOf(&_Oftwrapper.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Oftwrapper *OftwrapperCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Oftwrapper.Contract.BalanceOf(&_Oftwrapper.CallOpts, account)
}

// DefaultBps is a free data retrieval call binding the contract method 0xd1b308dc.
//
// Solidity: function defaultBps() view returns(uint256)
func (_Oftwrapper *OftwrapperCaller) DefaultBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oftwrapper.contract.Call(opts, &out, "defaultBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultBps is a free data retrieval call binding the contract method 0xd1b308dc.
//
// Solidity: function defaultBps() view returns(uint256)
func (_Oftwrapper *OftwrapperSession) DefaultBps() (*big.Int, error) {
	return _Oftwrapper.Contract.DefaultBps(&_Oftwrapper.CallOpts)
}

// DefaultBps is a free data retrieval call binding the contract method 0xd1b308dc.
//
// Solidity: function defaultBps() view returns(uint256)
func (_Oftwrapper *OftwrapperCallerSession) DefaultBps() (*big.Int, error) {
	return _Oftwrapper.Contract.DefaultBps(&_Oftwrapper.CallOpts)
}

// EstimateSendFee is a free data retrieval call binding the contract method 0xe1bafc80.
//
// Solidity: function estimateSendFee(address _oft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams, (uint256,address,bytes2) _feeObj) view returns(uint256 nativeFee, uint256 zroFee)
func (_Oftwrapper *OftwrapperCaller) EstimateSendFee(opts *bind.CallOpts, _oft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _useZro bool, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Oftwrapper.contract.Call(opts, &out, "estimateSendFee", _oft, _dstChainId, _toAddress, _amount, _useZro, _adapterParams, _feeObj)

	outstruct := new(struct {
		NativeFee *big.Int
		ZroFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ZroFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateSendFee is a free data retrieval call binding the contract method 0xe1bafc80.
//
// Solidity: function estimateSendFee(address _oft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams, (uint256,address,bytes2) _feeObj) view returns(uint256 nativeFee, uint256 zroFee)
func (_Oftwrapper *OftwrapperSession) EstimateSendFee(_oft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _useZro bool, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Oftwrapper.Contract.EstimateSendFee(&_Oftwrapper.CallOpts, _oft, _dstChainId, _toAddress, _amount, _useZro, _adapterParams, _feeObj)
}

// EstimateSendFee is a free data retrieval call binding the contract method 0xe1bafc80.
//
// Solidity: function estimateSendFee(address _oft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams, (uint256,address,bytes2) _feeObj) view returns(uint256 nativeFee, uint256 zroFee)
func (_Oftwrapper *OftwrapperCallerSession) EstimateSendFee(_oft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _useZro bool, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Oftwrapper.Contract.EstimateSendFee(&_Oftwrapper.CallOpts, _oft, _dstChainId, _toAddress, _amount, _useZro, _adapterParams, _feeObj)
}

// EstimateSendFeeV2 is a free data retrieval call binding the contract method 0x8d8c915c.
//
// Solidity: function estimateSendFeeV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams, (uint256,address,bytes2) _feeObj) view returns(uint256 nativeFee, uint256 zroFee)
func (_Oftwrapper *OftwrapperCaller) EstimateSendFeeV2(opts *bind.CallOpts, _oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _useZro bool, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Oftwrapper.contract.Call(opts, &out, "estimateSendFeeV2", _oft, _dstChainId, _toAddress, _amount, _useZro, _adapterParams, _feeObj)

	outstruct := new(struct {
		NativeFee *big.Int
		ZroFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ZroFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateSendFeeV2 is a free data retrieval call binding the contract method 0x8d8c915c.
//
// Solidity: function estimateSendFeeV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams, (uint256,address,bytes2) _feeObj) view returns(uint256 nativeFee, uint256 zroFee)
func (_Oftwrapper *OftwrapperSession) EstimateSendFeeV2(_oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _useZro bool, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Oftwrapper.Contract.EstimateSendFeeV2(&_Oftwrapper.CallOpts, _oft, _dstChainId, _toAddress, _amount, _useZro, _adapterParams, _feeObj)
}

// EstimateSendFeeV2 is a free data retrieval call binding the contract method 0x8d8c915c.
//
// Solidity: function estimateSendFeeV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams, (uint256,address,bytes2) _feeObj) view returns(uint256 nativeFee, uint256 zroFee)
func (_Oftwrapper *OftwrapperCallerSession) EstimateSendFeeV2(_oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _useZro bool, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Oftwrapper.Contract.EstimateSendFeeV2(&_Oftwrapper.CallOpts, _oft, _dstChainId, _toAddress, _amount, _useZro, _adapterParams, _feeObj)
}

// GetAmountAndFees is a free data retrieval call binding the contract method 0x17696f64.
//
// Solidity: function getAmountAndFees(address _token, uint256 _amount, uint256 _callerBps) view returns(uint256 amount, uint256 wrapperFee, uint256 callerFee)
func (_Oftwrapper *OftwrapperCaller) GetAmountAndFees(opts *bind.CallOpts, _token common.Address, _amount *big.Int, _callerBps *big.Int) (struct {
	Amount     *big.Int
	WrapperFee *big.Int
	CallerFee  *big.Int
}, error) {
	var out []interface{}
	err := _Oftwrapper.contract.Call(opts, &out, "getAmountAndFees", _token, _amount, _callerBps)

	outstruct := new(struct {
		Amount     *big.Int
		WrapperFee *big.Int
		CallerFee  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WrapperFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CallerFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAmountAndFees is a free data retrieval call binding the contract method 0x17696f64.
//
// Solidity: function getAmountAndFees(address _token, uint256 _amount, uint256 _callerBps) view returns(uint256 amount, uint256 wrapperFee, uint256 callerFee)
func (_Oftwrapper *OftwrapperSession) GetAmountAndFees(_token common.Address, _amount *big.Int, _callerBps *big.Int) (struct {
	Amount     *big.Int
	WrapperFee *big.Int
	CallerFee  *big.Int
}, error) {
	return _Oftwrapper.Contract.GetAmountAndFees(&_Oftwrapper.CallOpts, _token, _amount, _callerBps)
}

// GetAmountAndFees is a free data retrieval call binding the contract method 0x17696f64.
//
// Solidity: function getAmountAndFees(address _token, uint256 _amount, uint256 _callerBps) view returns(uint256 amount, uint256 wrapperFee, uint256 callerFee)
func (_Oftwrapper *OftwrapperCallerSession) GetAmountAndFees(_token common.Address, _amount *big.Int, _callerBps *big.Int) (struct {
	Amount     *big.Int
	WrapperFee *big.Int
	CallerFee  *big.Int
}, error) {
	return _Oftwrapper.Contract.GetAmountAndFees(&_Oftwrapper.CallOpts, _token, _amount, _callerBps)
}

// OftBps is a free data retrieval call binding the contract method 0x7a751182.
//
// Solidity: function oftBps(address ) view returns(uint256)
func (_Oftwrapper *OftwrapperCaller) OftBps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oftwrapper.contract.Call(opts, &out, "oftBps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OftBps is a free data retrieval call binding the contract method 0x7a751182.
//
// Solidity: function oftBps(address ) view returns(uint256)
func (_Oftwrapper *OftwrapperSession) OftBps(arg0 common.Address) (*big.Int, error) {
	return _Oftwrapper.Contract.OftBps(&_Oftwrapper.CallOpts, arg0)
}

// OftBps is a free data retrieval call binding the contract method 0x7a751182.
//
// Solidity: function oftBps(address ) view returns(uint256)
func (_Oftwrapper *OftwrapperCallerSession) OftBps(arg0 common.Address) (*big.Int, error) {
	return _Oftwrapper.Contract.OftBps(&_Oftwrapper.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oftwrapper *OftwrapperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oftwrapper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oftwrapper *OftwrapperSession) Owner() (common.Address, error) {
	return _Oftwrapper.Contract.Owner(&_Oftwrapper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oftwrapper *OftwrapperCallerSession) Owner() (common.Address, error) {
	return _Oftwrapper.Contract.Owner(&_Oftwrapper.CallOpts)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Oftwrapper *OftwrapperCaller) QuoteSend(opts *bind.CallOpts, _sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	var out []interface{}
	err := _Oftwrapper.contract.Call(opts, &out, "quoteSend", _sendParam, _payInLzToken)

	if err != nil {
		return *new(MessagingFee), err
	}

	out0 := *abi.ConvertType(out[0], new(MessagingFee)).(*MessagingFee)

	return out0, err

}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Oftwrapper *OftwrapperSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _Oftwrapper.Contract.QuoteSend(&_Oftwrapper.CallOpts, _sendParam, _payInLzToken)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Oftwrapper *OftwrapperCallerSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _Oftwrapper.Contract.QuoteSend(&_Oftwrapper.CallOpts, _sendParam, _payInLzToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oftwrapper *OftwrapperTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oftwrapper.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oftwrapper *OftwrapperSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oftwrapper.Contract.RenounceOwnership(&_Oftwrapper.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oftwrapper *OftwrapperTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oftwrapper.Contract.RenounceOwnership(&_Oftwrapper.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Oftwrapper *OftwrapperTransactor) Send(opts *bind.TransactOpts, _sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oftwrapper.contract.Transact(opts, "send", _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Oftwrapper *OftwrapperSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oftwrapper.Contract.Send(&_Oftwrapper.TransactOpts, _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Oftwrapper *OftwrapperTransactorSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Oftwrapper.Contract.Send(&_Oftwrapper.TransactOpts, _sendParam, _fee, _refundAddress)
}

// SendNativeOFT is a paid mutator transaction binding the contract method 0xfdb80c81.
//
// Solidity: function sendNativeOFT(address _nativeOft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactor) SendNativeOFT(opts *bind.TransactOpts, _nativeOft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.contract.Transact(opts, "sendNativeOFT", _nativeOft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendNativeOFT is a paid mutator transaction binding the contract method 0xfdb80c81.
//
// Solidity: function sendNativeOFT(address _nativeOft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperSession) SendNativeOFT(_nativeOft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendNativeOFT(&_Oftwrapper.TransactOpts, _nativeOft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendNativeOFT is a paid mutator transaction binding the contract method 0xfdb80c81.
//
// Solidity: function sendNativeOFT(address _nativeOft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactorSession) SendNativeOFT(_nativeOft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendNativeOFT(&_Oftwrapper.TransactOpts, _nativeOft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendNativeOFTFeeV2 is a paid mutator transaction binding the contract method 0xca3e534c.
//
// Solidity: function sendNativeOFTFeeV2(address _nativeOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactor) SendNativeOFTFeeV2(opts *bind.TransactOpts, _nativeOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.contract.Transact(opts, "sendNativeOFTFeeV2", _nativeOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendNativeOFTFeeV2 is a paid mutator transaction binding the contract method 0xca3e534c.
//
// Solidity: function sendNativeOFTFeeV2(address _nativeOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperSession) SendNativeOFTFeeV2(_nativeOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendNativeOFTFeeV2(&_Oftwrapper.TransactOpts, _nativeOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendNativeOFTFeeV2 is a paid mutator transaction binding the contract method 0xca3e534c.
//
// Solidity: function sendNativeOFTFeeV2(address _nativeOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactorSession) SendNativeOFTFeeV2(_nativeOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendNativeOFTFeeV2(&_Oftwrapper.TransactOpts, _nativeOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendOFT is a paid mutator transaction binding the contract method 0x498eff64.
//
// Solidity: function sendOFT(address _oft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactor) SendOFT(opts *bind.TransactOpts, _oft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.contract.Transact(opts, "sendOFT", _oft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendOFT is a paid mutator transaction binding the contract method 0x498eff64.
//
// Solidity: function sendOFT(address _oft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperSession) SendOFT(_oft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendOFT(&_Oftwrapper.TransactOpts, _oft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendOFT is a paid mutator transaction binding the contract method 0x498eff64.
//
// Solidity: function sendOFT(address _oft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactorSession) SendOFT(_oft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendOFT(&_Oftwrapper.TransactOpts, _oft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendOFTFeeV2 is a paid mutator transaction binding the contract method 0x85154849.
//
// Solidity: function sendOFTFeeV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactor) SendOFTFeeV2(opts *bind.TransactOpts, _oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.contract.Transact(opts, "sendOFTFeeV2", _oft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendOFTFeeV2 is a paid mutator transaction binding the contract method 0x85154849.
//
// Solidity: function sendOFTFeeV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperSession) SendOFTFeeV2(_oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendOFTFeeV2(&_Oftwrapper.TransactOpts, _oft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendOFTFeeV2 is a paid mutator transaction binding the contract method 0x85154849.
//
// Solidity: function sendOFTFeeV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactorSession) SendOFTFeeV2(_oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendOFTFeeV2(&_Oftwrapper.TransactOpts, _oft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendOFTV2 is a paid mutator transaction binding the contract method 0xa8198c00.
//
// Solidity: function sendOFTV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactor) SendOFTV2(opts *bind.TransactOpts, _oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.contract.Transact(opts, "sendOFTV2", _oft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendOFTV2 is a paid mutator transaction binding the contract method 0xa8198c00.
//
// Solidity: function sendOFTV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperSession) SendOFTV2(_oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendOFTV2(&_Oftwrapper.TransactOpts, _oft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendOFTV2 is a paid mutator transaction binding the contract method 0xa8198c00.
//
// Solidity: function sendOFTV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactorSession) SendOFTV2(_oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendOFTV2(&_Oftwrapper.TransactOpts, _oft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendProxyOFT is a paid mutator transaction binding the contract method 0xc3c8032a.
//
// Solidity: function sendProxyOFT(address _proxyOft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactor) SendProxyOFT(opts *bind.TransactOpts, _proxyOft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.contract.Transact(opts, "sendProxyOFT", _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendProxyOFT is a paid mutator transaction binding the contract method 0xc3c8032a.
//
// Solidity: function sendProxyOFT(address _proxyOft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperSession) SendProxyOFT(_proxyOft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendProxyOFT(&_Oftwrapper.TransactOpts, _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendProxyOFT is a paid mutator transaction binding the contract method 0xc3c8032a.
//
// Solidity: function sendProxyOFT(address _proxyOft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactorSession) SendProxyOFT(_proxyOft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendProxyOFT(&_Oftwrapper.TransactOpts, _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendProxyOFTFeeV2 is a paid mutator transaction binding the contract method 0x8bcb586c.
//
// Solidity: function sendProxyOFTFeeV2(address _proxyOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactor) SendProxyOFTFeeV2(opts *bind.TransactOpts, _proxyOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.contract.Transact(opts, "sendProxyOFTFeeV2", _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendProxyOFTFeeV2 is a paid mutator transaction binding the contract method 0x8bcb586c.
//
// Solidity: function sendProxyOFTFeeV2(address _proxyOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperSession) SendProxyOFTFeeV2(_proxyOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendProxyOFTFeeV2(&_Oftwrapper.TransactOpts, _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendProxyOFTFeeV2 is a paid mutator transaction binding the contract method 0x8bcb586c.
//
// Solidity: function sendProxyOFTFeeV2(address _proxyOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactorSession) SendProxyOFTFeeV2(_proxyOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendProxyOFTFeeV2(&_Oftwrapper.TransactOpts, _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendProxyOFTV2 is a paid mutator transaction binding the contract method 0xdda16a10.
//
// Solidity: function sendProxyOFTV2(address _proxyOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactor) SendProxyOFTV2(opts *bind.TransactOpts, _proxyOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.contract.Transact(opts, "sendProxyOFTV2", _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendProxyOFTV2 is a paid mutator transaction binding the contract method 0xdda16a10.
//
// Solidity: function sendProxyOFTV2(address _proxyOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperSession) SendProxyOFTV2(_proxyOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendProxyOFTV2(&_Oftwrapper.TransactOpts, _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendProxyOFTV2 is a paid mutator transaction binding the contract method 0xdda16a10.
//
// Solidity: function sendProxyOFTV2(address _proxyOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Oftwrapper *OftwrapperTransactorSession) SendProxyOFTV2(_proxyOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SendProxyOFTV2(&_Oftwrapper.TransactOpts, _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SetDefaultBps is a paid mutator transaction binding the contract method 0xa46d74bc.
//
// Solidity: function setDefaultBps(uint256 _defaultBps) returns()
func (_Oftwrapper *OftwrapperTransactor) SetDefaultBps(opts *bind.TransactOpts, _defaultBps *big.Int) (*types.Transaction, error) {
	return _Oftwrapper.contract.Transact(opts, "setDefaultBps", _defaultBps)
}

// SetDefaultBps is a paid mutator transaction binding the contract method 0xa46d74bc.
//
// Solidity: function setDefaultBps(uint256 _defaultBps) returns()
func (_Oftwrapper *OftwrapperSession) SetDefaultBps(_defaultBps *big.Int) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SetDefaultBps(&_Oftwrapper.TransactOpts, _defaultBps)
}

// SetDefaultBps is a paid mutator transaction binding the contract method 0xa46d74bc.
//
// Solidity: function setDefaultBps(uint256 _defaultBps) returns()
func (_Oftwrapper *OftwrapperTransactorSession) SetDefaultBps(_defaultBps *big.Int) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SetDefaultBps(&_Oftwrapper.TransactOpts, _defaultBps)
}

// SetOFTBps is a paid mutator transaction binding the contract method 0x0c3d2756.
//
// Solidity: function setOFTBps(address _token, uint256 _bps) returns()
func (_Oftwrapper *OftwrapperTransactor) SetOFTBps(opts *bind.TransactOpts, _token common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _Oftwrapper.contract.Transact(opts, "setOFTBps", _token, _bps)
}

// SetOFTBps is a paid mutator transaction binding the contract method 0x0c3d2756.
//
// Solidity: function setOFTBps(address _token, uint256 _bps) returns()
func (_Oftwrapper *OftwrapperSession) SetOFTBps(_token common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SetOFTBps(&_Oftwrapper.TransactOpts, _token, _bps)
}

// SetOFTBps is a paid mutator transaction binding the contract method 0x0c3d2756.
//
// Solidity: function setOFTBps(address _token, uint256 _bps) returns()
func (_Oftwrapper *OftwrapperTransactorSession) SetOFTBps(_token common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _Oftwrapper.Contract.SetOFTBps(&_Oftwrapper.TransactOpts, _token, _bps)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oftwrapper *OftwrapperTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Oftwrapper.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oftwrapper *OftwrapperSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oftwrapper.Contract.TransferOwnership(&_Oftwrapper.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oftwrapper *OftwrapperTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oftwrapper.Contract.TransferOwnership(&_Oftwrapper.TransactOpts, newOwner)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xe55dc4e6.
//
// Solidity: function withdrawFees(address _oft, address _to, uint256 _amount) returns()
func (_Oftwrapper *OftwrapperTransactor) WithdrawFees(opts *bind.TransactOpts, _oft common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Oftwrapper.contract.Transact(opts, "withdrawFees", _oft, _to, _amount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xe55dc4e6.
//
// Solidity: function withdrawFees(address _oft, address _to, uint256 _amount) returns()
func (_Oftwrapper *OftwrapperSession) WithdrawFees(_oft common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Oftwrapper.Contract.WithdrawFees(&_Oftwrapper.TransactOpts, _oft, _to, _amount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xe55dc4e6.
//
// Solidity: function withdrawFees(address _oft, address _to, uint256 _amount) returns()
func (_Oftwrapper *OftwrapperTransactorSession) WithdrawFees(_oft common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Oftwrapper.Contract.WithdrawFees(&_Oftwrapper.TransactOpts, _oft, _to, _amount)
}

// OftwrapperOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oftwrapper contract.
type OftwrapperOwnershipTransferredIterator struct {
	Event *OftwrapperOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OftwrapperOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftwrapperOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OftwrapperOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OftwrapperOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftwrapperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftwrapperOwnershipTransferred represents a OwnershipTransferred event raised by the Oftwrapper contract.
type OftwrapperOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oftwrapper *OftwrapperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OftwrapperOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oftwrapper.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OftwrapperOwnershipTransferredIterator{contract: _Oftwrapper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oftwrapper *OftwrapperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OftwrapperOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oftwrapper.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftwrapperOwnershipTransferred)
				if err := _Oftwrapper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oftwrapper *OftwrapperFilterer) ParseOwnershipTransferred(log types.Log) (*OftwrapperOwnershipTransferred, error) {
	event := new(OftwrapperOwnershipTransferred)
	if err := _Oftwrapper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftwrapperWrapperFeeWithdrawnIterator is returned from FilterWrapperFeeWithdrawn and is used to iterate over the raw logs and unpacked data for WrapperFeeWithdrawn events raised by the Oftwrapper contract.
type OftwrapperWrapperFeeWithdrawnIterator struct {
	Event *OftwrapperWrapperFeeWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OftwrapperWrapperFeeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftwrapperWrapperFeeWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OftwrapperWrapperFeeWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OftwrapperWrapperFeeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftwrapperWrapperFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftwrapperWrapperFeeWithdrawn represents a WrapperFeeWithdrawn event raised by the Oftwrapper contract.
type OftwrapperWrapperFeeWithdrawn struct {
	Oft    common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWrapperFeeWithdrawn is a free log retrieval operation binding the contract event 0xf6514f9f283faac4cf3f3a6a702c116227ad0f2c727fb336e4c10b418bc6d613.
//
// Solidity: event WrapperFeeWithdrawn(address indexed oft, address to, uint256 amount)
func (_Oftwrapper *OftwrapperFilterer) FilterWrapperFeeWithdrawn(opts *bind.FilterOpts, oft []common.Address) (*OftwrapperWrapperFeeWithdrawnIterator, error) {

	var oftRule []interface{}
	for _, oftItem := range oft {
		oftRule = append(oftRule, oftItem)
	}

	logs, sub, err := _Oftwrapper.contract.FilterLogs(opts, "WrapperFeeWithdrawn", oftRule)
	if err != nil {
		return nil, err
	}
	return &OftwrapperWrapperFeeWithdrawnIterator{contract: _Oftwrapper.contract, event: "WrapperFeeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchWrapperFeeWithdrawn is a free log subscription operation binding the contract event 0xf6514f9f283faac4cf3f3a6a702c116227ad0f2c727fb336e4c10b418bc6d613.
//
// Solidity: event WrapperFeeWithdrawn(address indexed oft, address to, uint256 amount)
func (_Oftwrapper *OftwrapperFilterer) WatchWrapperFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *OftwrapperWrapperFeeWithdrawn, oft []common.Address) (event.Subscription, error) {

	var oftRule []interface{}
	for _, oftItem := range oft {
		oftRule = append(oftRule, oftItem)
	}

	logs, sub, err := _Oftwrapper.contract.WatchLogs(opts, "WrapperFeeWithdrawn", oftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftwrapperWrapperFeeWithdrawn)
				if err := _Oftwrapper.contract.UnpackLog(event, "WrapperFeeWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWrapperFeeWithdrawn is a log parse operation binding the contract event 0xf6514f9f283faac4cf3f3a6a702c116227ad0f2c727fb336e4c10b418bc6d613.
//
// Solidity: event WrapperFeeWithdrawn(address indexed oft, address to, uint256 amount)
func (_Oftwrapper *OftwrapperFilterer) ParseWrapperFeeWithdrawn(log types.Log) (*OftwrapperWrapperFeeWithdrawn, error) {
	event := new(OftwrapperWrapperFeeWithdrawn)
	if err := _Oftwrapper.contract.UnpackLog(event, "WrapperFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OftwrapperWrapperFeesIterator is returned from FilterWrapperFees and is used to iterate over the raw logs and unpacked data for WrapperFees events raised by the Oftwrapper contract.
type OftwrapperWrapperFeesIterator struct {
	Event *OftwrapperWrapperFees // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OftwrapperWrapperFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OftwrapperWrapperFees)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OftwrapperWrapperFees)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OftwrapperWrapperFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OftwrapperWrapperFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OftwrapperWrapperFees represents a WrapperFees event raised by the Oftwrapper contract.
type OftwrapperWrapperFees struct {
	PartnerId  [2]byte
	Token      common.Address
	WrapperFee *big.Int
	CallerFee  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWrapperFees is a free log retrieval operation binding the contract event 0x97bcdc1dd7ab82ef93280983f23d391afea463d0333fddd1a4617693b9ccfeea.
//
// Solidity: event WrapperFees(bytes2 indexed partnerId, address token, uint256 wrapperFee, uint256 callerFee)
func (_Oftwrapper *OftwrapperFilterer) FilterWrapperFees(opts *bind.FilterOpts, partnerId [][2]byte) (*OftwrapperWrapperFeesIterator, error) {

	var partnerIdRule []interface{}
	for _, partnerIdItem := range partnerId {
		partnerIdRule = append(partnerIdRule, partnerIdItem)
	}

	logs, sub, err := _Oftwrapper.contract.FilterLogs(opts, "WrapperFees", partnerIdRule)
	if err != nil {
		return nil, err
	}
	return &OftwrapperWrapperFeesIterator{contract: _Oftwrapper.contract, event: "WrapperFees", logs: logs, sub: sub}, nil
}

// WatchWrapperFees is a free log subscription operation binding the contract event 0x97bcdc1dd7ab82ef93280983f23d391afea463d0333fddd1a4617693b9ccfeea.
//
// Solidity: event WrapperFees(bytes2 indexed partnerId, address token, uint256 wrapperFee, uint256 callerFee)
func (_Oftwrapper *OftwrapperFilterer) WatchWrapperFees(opts *bind.WatchOpts, sink chan<- *OftwrapperWrapperFees, partnerId [][2]byte) (event.Subscription, error) {

	var partnerIdRule []interface{}
	for _, partnerIdItem := range partnerId {
		partnerIdRule = append(partnerIdRule, partnerIdItem)
	}

	logs, sub, err := _Oftwrapper.contract.WatchLogs(opts, "WrapperFees", partnerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OftwrapperWrapperFees)
				if err := _Oftwrapper.contract.UnpackLog(event, "WrapperFees", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWrapperFees is a log parse operation binding the contract event 0x97bcdc1dd7ab82ef93280983f23d391afea463d0333fddd1a4617693b9ccfeea.
//
// Solidity: event WrapperFees(bytes2 indexed partnerId, address token, uint256 wrapperFee, uint256 callerFee)
func (_Oftwrapper *OftwrapperFilterer) ParseWrapperFees(log types.Log) (*OftwrapperWrapperFees, error) {
	event := new(OftwrapperWrapperFees)
	if err := _Oftwrapper.contract.UnpackLog(event, "WrapperFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
