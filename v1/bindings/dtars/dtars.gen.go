// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dtars

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

// ClientAny2EVMMessage is an auto generated low-level Go binding around an user-defined struct.
type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

// ClientEVMTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

// IDTAMessageDTAPayment is an auto generated low-level Go binding around an user-defined struct.
type IDTAMessageDTAPayment struct {
	OffChainPaymentCurrency uint8
	PaymentTokenSourceAddr  common.Address
	PaymentTokenDestAddr    common.Address
}

// IDTAMessageDtaRequestMessage is an auto generated low-level Go binding around an user-defined struct.
type IDTAMessageDtaRequestMessage struct {
	FundTokenId           [32]byte
	RequestId             [32]byte
	Shares                *big.Int
	Amount                *big.Int
	PaymentInfo           IDTAMessageDTAPayment
	FundAdminAddr         common.Address
	DistributorWalletAddr common.Address
	DistributorAddr       common.Address
}

// IDTARequestSettlementDTAData is an auto generated low-level Go binding around an user-defined struct.
type IDTARequestSettlementDTAData struct {
	DtaAddr          common.Address
	DtaChainSelector uint64
	FundAdminAddr    common.Address
	FundTokenId      [32]byte
	FundTokenAddr    common.Address
	MintType         uint8
	BurnType         uint8
	Enabled          bool
}

// DtarsMetaData contains all meta data concerning the Dtars contract.
var DtarsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"ccipRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"_executeSettlement\",\"inputs\":[{\"name\":\"dtaRequest\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestMessage\",\"components\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dtaData\",\"type\":\"tuple\",\"internalType\":\"structIDTARequestSettlement.DTAData\",\"components\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mintType\",\"type\":\"uint8\",\"internalType\":\"enumIDTARequestSettlement.TokenMintType\"},{\"name\":\"burnType\",\"type\":\"uint8\",\"internalType\":\"enumIDTARequestSettlement.TokenBurnType\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"settlementType\",\"type\":\"uint8\",\"internalType\":\"enumDTARequestSettlementU.SettlementType\"},{\"name\":\"paymentTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mintType\",\"type\":\"uint8\",\"internalType\":\"enumIDTARequestSettlement.TokenMintType\"},{\"name\":\"burnType\",\"type\":\"uint8\",\"internalType\":\"enumIDTARequestSettlement.TokenBurnType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipHandleDTAMessage\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dtaMessage\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestMessage\",\"components\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"ccipDestTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeRequestProcessing\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"directHandleDTAMessage\",\"inputs\":[{\"name\":\"dtaMessage\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestMessage\",\"components\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disallowDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedDTAs\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"dtaKeys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCIPGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDTAData\",\"inputs\":[{\"name\":\"dtaKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"dtaData\",\"type\":\"tuple\",\"internalType\":\"structIDTARequestSettlement.DTAData\",\"components\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mintType\",\"type\":\"uint8\",\"internalType\":\"enumIDTARequestSettlement.TokenMintType\"},{\"name\":\"burnType\",\"type\":\"uint8\",\"internalType\":\"enumIDTARequestSettlement.TokenBurnType\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAllowedDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverFunds\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTokens\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageRecvFailed\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DTAAdded\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DTARemoved\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DTASettlementClosed\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requestType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumIDTAMessage.DistributorRequestType\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"err\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DTASettlementOpened\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requestType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumIDTAMessage.DistributorRequestType\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumCurrency\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmptyRequestType\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvalidSubscriptionCrossChainPayment\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"ccipDestTokenAmountsLength\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ccipPaymentTokenAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeFundsRecovered\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SettlementFailed\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymentTokenAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"errData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenWithdrawn\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnauthorizedSenderDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"reqType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIDTAMessage.DistributorRequestType\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DoesNotExist\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"Exists\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FailedToRecoverFunds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDTAKey\",\"inputs\":[{\"name\":\"dtaKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPaymentInfo\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWithdrawInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LowBalanceForCCIPSend\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"currentBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ccipFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LowBalanceForCrossChainRedemption\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedNotAllowedDTA\",\"inputs\":[{\"name\":\"dtaAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dtaChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// DtarsABI is the input ABI used to generate the binding from.
// Deprecated: Use DtarsMetaData.ABI instead.
var DtarsABI = DtarsMetaData.ABI

// Dtars is an auto generated Go binding around an Ethereum contract.
type Dtars struct {
	DtarsCaller     // Read-only binding to the contract
	DtarsTransactor // Write-only binding to the contract
	DtarsFilterer   // Log filterer for contract events
}

// DtarsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DtarsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtarsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DtarsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtarsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DtarsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtarsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DtarsSession struct {
	Contract     *Dtars            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DtarsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DtarsCallerSession struct {
	Contract *DtarsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DtarsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DtarsTransactorSession struct {
	Contract     *DtarsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DtarsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DtarsRaw struct {
	Contract *Dtars // Generic contract binding to access the raw methods on
}

// DtarsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DtarsCallerRaw struct {
	Contract *DtarsCaller // Generic read-only contract binding to access the raw methods on
}

// DtarsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DtarsTransactorRaw struct {
	Contract *DtarsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDtars creates a new instance of Dtars, bound to a specific deployed contract.
func NewDtars(address common.Address, backend bind.ContractBackend) (*Dtars, error) {
	contract, err := bindDtars(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dtars{DtarsCaller: DtarsCaller{contract: contract}, DtarsTransactor: DtarsTransactor{contract: contract}, DtarsFilterer: DtarsFilterer{contract: contract}}, nil
}

// NewDtarsCaller creates a new read-only instance of Dtars, bound to a specific deployed contract.
func NewDtarsCaller(address common.Address, caller bind.ContractCaller) (*DtarsCaller, error) {
	contract, err := bindDtars(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DtarsCaller{contract: contract}, nil
}

// NewDtarsTransactor creates a new write-only instance of Dtars, bound to a specific deployed contract.
func NewDtarsTransactor(address common.Address, transactor bind.ContractTransactor) (*DtarsTransactor, error) {
	contract, err := bindDtars(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DtarsTransactor{contract: contract}, nil
}

// NewDtarsFilterer creates a new log filterer instance of Dtars, bound to a specific deployed contract.
func NewDtarsFilterer(address common.Address, filterer bind.ContractFilterer) (*DtarsFilterer, error) {
	contract, err := bindDtars(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DtarsFilterer{contract: contract}, nil
}

// bindDtars binds a generic wrapper to an already deployed contract.
func bindDtars(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DtarsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dtars *DtarsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dtars.Contract.DtarsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dtars *DtarsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtars.Contract.DtarsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dtars *DtarsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dtars.Contract.DtarsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dtars *DtarsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dtars.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dtars *DtarsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtars.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dtars *DtarsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dtars.Contract.contract.Transact(opts, method, params...)
}

// GetAllowedDTAs is a free data retrieval call binding the contract method 0xe4688a46.
//
// Solidity: function getAllowedDTAs(uint256 offset) view returns(bytes32[] dtaKeys)
func (_Dtars *DtarsCaller) GetAllowedDTAs(opts *bind.CallOpts, offset *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Dtars.contract.Call(opts, &out, "getAllowedDTAs", offset)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllowedDTAs is a free data retrieval call binding the contract method 0xe4688a46.
//
// Solidity: function getAllowedDTAs(uint256 offset) view returns(bytes32[] dtaKeys)
func (_Dtars *DtarsSession) GetAllowedDTAs(offset *big.Int) ([][32]byte, error) {
	return _Dtars.Contract.GetAllowedDTAs(&_Dtars.CallOpts, offset)
}

// GetAllowedDTAs is a free data retrieval call binding the contract method 0xe4688a46.
//
// Solidity: function getAllowedDTAs(uint256 offset) view returns(bytes32[] dtaKeys)
func (_Dtars *DtarsCallerSession) GetAllowedDTAs(offset *big.Int) ([][32]byte, error) {
	return _Dtars.Contract.GetAllowedDTAs(&_Dtars.CallOpts, offset)
}

// GetCCIPGasLimit is a free data retrieval call binding the contract method 0xd83ce949.
//
// Solidity: function getCCIPGasLimit() view returns(uint256)
func (_Dtars *DtarsCaller) GetCCIPGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dtars.contract.Call(opts, &out, "getCCIPGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCCIPGasLimit is a free data retrieval call binding the contract method 0xd83ce949.
//
// Solidity: function getCCIPGasLimit() view returns(uint256)
func (_Dtars *DtarsSession) GetCCIPGasLimit() (*big.Int, error) {
	return _Dtars.Contract.GetCCIPGasLimit(&_Dtars.CallOpts)
}

// GetCCIPGasLimit is a free data retrieval call binding the contract method 0xd83ce949.
//
// Solidity: function getCCIPGasLimit() view returns(uint256)
func (_Dtars *DtarsCallerSession) GetCCIPGasLimit() (*big.Int, error) {
	return _Dtars.Contract.GetCCIPGasLimit(&_Dtars.CallOpts)
}

// GetDTAData is a free data retrieval call binding the contract method 0x633e2eb4.
//
// Solidity: function getDTAData(bytes32 dtaKey) view returns((address,uint64,address,bytes32,address,uint8,uint8,bool) dtaData)
func (_Dtars *DtarsCaller) GetDTAData(opts *bind.CallOpts, dtaKey [32]byte) (IDTARequestSettlementDTAData, error) {
	var out []interface{}
	err := _Dtars.contract.Call(opts, &out, "getDTAData", dtaKey)

	if err != nil {
		return *new(IDTARequestSettlementDTAData), err
	}

	out0 := *abi.ConvertType(out[0], new(IDTARequestSettlementDTAData)).(*IDTARequestSettlementDTAData)

	return out0, err

}

// GetDTAData is a free data retrieval call binding the contract method 0x633e2eb4.
//
// Solidity: function getDTAData(bytes32 dtaKey) view returns((address,uint64,address,bytes32,address,uint8,uint8,bool) dtaData)
func (_Dtars *DtarsSession) GetDTAData(dtaKey [32]byte) (IDTARequestSettlementDTAData, error) {
	return _Dtars.Contract.GetDTAData(&_Dtars.CallOpts, dtaKey)
}

// GetDTAData is a free data retrieval call binding the contract method 0x633e2eb4.
//
// Solidity: function getDTAData(bytes32 dtaKey) view returns((address,uint64,address,bytes32,address,uint8,uint8,bool) dtaData)
func (_Dtars *DtarsCallerSession) GetDTAData(dtaKey [32]byte) (IDTARequestSettlementDTAData, error) {
	return _Dtars.Contract.GetDTAData(&_Dtars.CallOpts, dtaKey)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtars *DtarsCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dtars.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtars *DtarsSession) GetRouter() (common.Address, error) {
	return _Dtars.Contract.GetRouter(&_Dtars.CallOpts)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtars *DtarsCallerSession) GetRouter() (common.Address, error) {
	return _Dtars.Contract.GetRouter(&_Dtars.CallOpts)
}

// IsAllowedDTA is a free data retrieval call binding the contract method 0xa2de36fc.
//
// Solidity: function isAllowedDTA(address dtaAddr, uint64 dtaChainSelector, address fundAdminAddr, bytes32 fundTokenId) view returns(bool)
func (_Dtars *DtarsCaller) IsAllowedDTA(opts *bind.CallOpts, dtaAddr common.Address, dtaChainSelector uint64, fundAdminAddr common.Address, fundTokenId [32]byte) (bool, error) {
	var out []interface{}
	err := _Dtars.contract.Call(opts, &out, "isAllowedDTA", dtaAddr, dtaChainSelector, fundAdminAddr, fundTokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedDTA is a free data retrieval call binding the contract method 0xa2de36fc.
//
// Solidity: function isAllowedDTA(address dtaAddr, uint64 dtaChainSelector, address fundAdminAddr, bytes32 fundTokenId) view returns(bool)
func (_Dtars *DtarsSession) IsAllowedDTA(dtaAddr common.Address, dtaChainSelector uint64, fundAdminAddr common.Address, fundTokenId [32]byte) (bool, error) {
	return _Dtars.Contract.IsAllowedDTA(&_Dtars.CallOpts, dtaAddr, dtaChainSelector, fundAdminAddr, fundTokenId)
}

// IsAllowedDTA is a free data retrieval call binding the contract method 0xa2de36fc.
//
// Solidity: function isAllowedDTA(address dtaAddr, uint64 dtaChainSelector, address fundAdminAddr, bytes32 fundTokenId) view returns(bool)
func (_Dtars *DtarsCallerSession) IsAllowedDTA(dtaAddr common.Address, dtaChainSelector uint64, fundAdminAddr common.Address, fundTokenId [32]byte) (bool, error) {
	return _Dtars.Contract.IsAllowedDTA(&_Dtars.CallOpts, dtaAddr, dtaChainSelector, fundAdminAddr, fundTokenId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtars *DtarsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dtars.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtars *DtarsSession) Owner() (common.Address, error) {
	return _Dtars.Contract.Owner(&_Dtars.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtars *DtarsCallerSession) Owner() (common.Address, error) {
	return _Dtars.Contract.Owner(&_Dtars.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtars *DtarsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Dtars.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtars *DtarsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dtars.Contract.SupportsInterface(&_Dtars.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtars *DtarsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dtars.Contract.SupportsInterface(&_Dtars.CallOpts, interfaceId)
}

// ExecuteSettlement is a paid mutator transaction binding the contract method 0xb40cb60e.
//
// Solidity: function _executeSettlement((bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaRequest, (address,uint64,address,bytes32,address,uint8,uint8,bool) dtaData, uint8 settlementType, address paymentTokenAddr) returns()
func (_Dtars *DtarsTransactor) ExecuteSettlement(opts *bind.TransactOpts, dtaRequest IDTAMessageDtaRequestMessage, dtaData IDTARequestSettlementDTAData, settlementType uint8, paymentTokenAddr common.Address) (*types.Transaction, error) {
	return _Dtars.contract.Transact(opts, "_executeSettlement", dtaRequest, dtaData, settlementType, paymentTokenAddr)
}

// ExecuteSettlement is a paid mutator transaction binding the contract method 0xb40cb60e.
//
// Solidity: function _executeSettlement((bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaRequest, (address,uint64,address,bytes32,address,uint8,uint8,bool) dtaData, uint8 settlementType, address paymentTokenAddr) returns()
func (_Dtars *DtarsSession) ExecuteSettlement(dtaRequest IDTAMessageDtaRequestMessage, dtaData IDTARequestSettlementDTAData, settlementType uint8, paymentTokenAddr common.Address) (*types.Transaction, error) {
	return _Dtars.Contract.ExecuteSettlement(&_Dtars.TransactOpts, dtaRequest, dtaData, settlementType, paymentTokenAddr)
}

// ExecuteSettlement is a paid mutator transaction binding the contract method 0xb40cb60e.
//
// Solidity: function _executeSettlement((bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaRequest, (address,uint64,address,bytes32,address,uint8,uint8,bool) dtaData, uint8 settlementType, address paymentTokenAddr) returns()
func (_Dtars *DtarsTransactorSession) ExecuteSettlement(dtaRequest IDTAMessageDtaRequestMessage, dtaData IDTARequestSettlementDTAData, settlementType uint8, paymentTokenAddr common.Address) (*types.Transaction, error) {
	return _Dtars.Contract.ExecuteSettlement(&_Dtars.TransactOpts, dtaRequest, dtaData, settlementType, paymentTokenAddr)
}

// AllowDTA is a paid mutator transaction binding the contract method 0x9a33f7e7.
//
// Solidity: function allowDTA(address dtaAddr, uint64 dtaChainSelector, address fundAdminAddr, bytes32 fundTokenId, address fundTokenAddr, uint8 mintType, uint8 burnType) returns()
func (_Dtars *DtarsTransactor) AllowDTA(opts *bind.TransactOpts, dtaAddr common.Address, dtaChainSelector uint64, fundAdminAddr common.Address, fundTokenId [32]byte, fundTokenAddr common.Address, mintType uint8, burnType uint8) (*types.Transaction, error) {
	return _Dtars.contract.Transact(opts, "allowDTA", dtaAddr, dtaChainSelector, fundAdminAddr, fundTokenId, fundTokenAddr, mintType, burnType)
}

// AllowDTA is a paid mutator transaction binding the contract method 0x9a33f7e7.
//
// Solidity: function allowDTA(address dtaAddr, uint64 dtaChainSelector, address fundAdminAddr, bytes32 fundTokenId, address fundTokenAddr, uint8 mintType, uint8 burnType) returns()
func (_Dtars *DtarsSession) AllowDTA(dtaAddr common.Address, dtaChainSelector uint64, fundAdminAddr common.Address, fundTokenId [32]byte, fundTokenAddr common.Address, mintType uint8, burnType uint8) (*types.Transaction, error) {
	return _Dtars.Contract.AllowDTA(&_Dtars.TransactOpts, dtaAddr, dtaChainSelector, fundAdminAddr, fundTokenId, fundTokenAddr, mintType, burnType)
}

// AllowDTA is a paid mutator transaction binding the contract method 0x9a33f7e7.
//
// Solidity: function allowDTA(address dtaAddr, uint64 dtaChainSelector, address fundAdminAddr, bytes32 fundTokenId, address fundTokenAddr, uint8 mintType, uint8 burnType) returns()
func (_Dtars *DtarsTransactorSession) AllowDTA(dtaAddr common.Address, dtaChainSelector uint64, fundAdminAddr common.Address, fundTokenId [32]byte, fundTokenAddr common.Address, mintType uint8, burnType uint8) (*types.Transaction, error) {
	return _Dtars.Contract.AllowDTA(&_Dtars.TransactOpts, dtaAddr, dtaChainSelector, fundAdminAddr, fundTokenId, fundTokenAddr, mintType, burnType)
}

// CcipHandleDTAMessage is a paid mutator transaction binding the contract method 0x0e7f0a40.
//
// Solidity: function ccipHandleDTAMessage(address dtaAddr, uint64 dtaChainSelector, (bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaMessage, (address,uint256)[] ccipDestTokenAmounts) returns()
func (_Dtars *DtarsTransactor) CcipHandleDTAMessage(opts *bind.TransactOpts, dtaAddr common.Address, dtaChainSelector uint64, dtaMessage IDTAMessageDtaRequestMessage, ccipDestTokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _Dtars.contract.Transact(opts, "ccipHandleDTAMessage", dtaAddr, dtaChainSelector, dtaMessage, ccipDestTokenAmounts)
}

// CcipHandleDTAMessage is a paid mutator transaction binding the contract method 0x0e7f0a40.
//
// Solidity: function ccipHandleDTAMessage(address dtaAddr, uint64 dtaChainSelector, (bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaMessage, (address,uint256)[] ccipDestTokenAmounts) returns()
func (_Dtars *DtarsSession) CcipHandleDTAMessage(dtaAddr common.Address, dtaChainSelector uint64, dtaMessage IDTAMessageDtaRequestMessage, ccipDestTokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _Dtars.Contract.CcipHandleDTAMessage(&_Dtars.TransactOpts, dtaAddr, dtaChainSelector, dtaMessage, ccipDestTokenAmounts)
}

// CcipHandleDTAMessage is a paid mutator transaction binding the contract method 0x0e7f0a40.
//
// Solidity: function ccipHandleDTAMessage(address dtaAddr, uint64 dtaChainSelector, (bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaMessage, (address,uint256)[] ccipDestTokenAmounts) returns()
func (_Dtars *DtarsTransactorSession) CcipHandleDTAMessage(dtaAddr common.Address, dtaChainSelector uint64, dtaMessage IDTAMessageDtaRequestMessage, ccipDestTokenAmounts []ClientEVMTokenAmount) (*types.Transaction, error) {
	return _Dtars.Contract.CcipHandleDTAMessage(&_Dtars.TransactOpts, dtaAddr, dtaChainSelector, dtaMessage, ccipDestTokenAmounts)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtars *DtarsTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtars.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtars *DtarsSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtars.Contract.CcipReceive(&_Dtars.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtars *DtarsTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtars.Contract.CcipReceive(&_Dtars.TransactOpts, message)
}

// CompleteRequestProcessing is a paid mutator transaction binding the contract method 0x17d7e379.
//
// Solidity: function completeRequestProcessing(bytes32 requestId, bool success, bytes err) returns()
func (_Dtars *DtarsTransactor) CompleteRequestProcessing(opts *bind.TransactOpts, requestId [32]byte, success bool, err []byte) (*types.Transaction, error) {
	return _Dtars.contract.Transact(opts, "completeRequestProcessing", requestId, success, err)
}

// CompleteRequestProcessing is a paid mutator transaction binding the contract method 0x17d7e379.
//
// Solidity: function completeRequestProcessing(bytes32 requestId, bool success, bytes err) returns()
func (_Dtars *DtarsSession) CompleteRequestProcessing(requestId [32]byte, success bool, err []byte) (*types.Transaction, error) {
	return _Dtars.Contract.CompleteRequestProcessing(&_Dtars.TransactOpts, requestId, success, err)
}

// CompleteRequestProcessing is a paid mutator transaction binding the contract method 0x17d7e379.
//
// Solidity: function completeRequestProcessing(bytes32 requestId, bool success, bytes err) returns()
func (_Dtars *DtarsTransactorSession) CompleteRequestProcessing(requestId [32]byte, success bool, err []byte) (*types.Transaction, error) {
	return _Dtars.Contract.CompleteRequestProcessing(&_Dtars.TransactOpts, requestId, success, err)
}

// DirectHandleDTAMessage is a paid mutator transaction binding the contract method 0x44310975.
//
// Solidity: function directHandleDTAMessage((bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaMessage) returns()
func (_Dtars *DtarsTransactor) DirectHandleDTAMessage(opts *bind.TransactOpts, dtaMessage IDTAMessageDtaRequestMessage) (*types.Transaction, error) {
	return _Dtars.contract.Transact(opts, "directHandleDTAMessage", dtaMessage)
}

// DirectHandleDTAMessage is a paid mutator transaction binding the contract method 0x44310975.
//
// Solidity: function directHandleDTAMessage((bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaMessage) returns()
func (_Dtars *DtarsSession) DirectHandleDTAMessage(dtaMessage IDTAMessageDtaRequestMessage) (*types.Transaction, error) {
	return _Dtars.Contract.DirectHandleDTAMessage(&_Dtars.TransactOpts, dtaMessage)
}

// DirectHandleDTAMessage is a paid mutator transaction binding the contract method 0x44310975.
//
// Solidity: function directHandleDTAMessage((bytes32,bytes32,uint256,uint256,(uint8,address,address),address,address,address) dtaMessage) returns()
func (_Dtars *DtarsTransactorSession) DirectHandleDTAMessage(dtaMessage IDTAMessageDtaRequestMessage) (*types.Transaction, error) {
	return _Dtars.Contract.DirectHandleDTAMessage(&_Dtars.TransactOpts, dtaMessage)
}

// DisallowDTA is a paid mutator transaction binding the contract method 0x7b067bd6.
//
// Solidity: function disallowDTA(address dtaAddr, uint64 dtaChainSelector, address fundAdminAddr, bytes32 fundTokenId) returns()
func (_Dtars *DtarsTransactor) DisallowDTA(opts *bind.TransactOpts, dtaAddr common.Address, dtaChainSelector uint64, fundAdminAddr common.Address, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtars.contract.Transact(opts, "disallowDTA", dtaAddr, dtaChainSelector, fundAdminAddr, fundTokenId)
}

// DisallowDTA is a paid mutator transaction binding the contract method 0x7b067bd6.
//
// Solidity: function disallowDTA(address dtaAddr, uint64 dtaChainSelector, address fundAdminAddr, bytes32 fundTokenId) returns()
func (_Dtars *DtarsSession) DisallowDTA(dtaAddr common.Address, dtaChainSelector uint64, fundAdminAddr common.Address, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtars.Contract.DisallowDTA(&_Dtars.TransactOpts, dtaAddr, dtaChainSelector, fundAdminAddr, fundTokenId)
}

// DisallowDTA is a paid mutator transaction binding the contract method 0x7b067bd6.
//
// Solidity: function disallowDTA(address dtaAddr, uint64 dtaChainSelector, address fundAdminAddr, bytes32 fundTokenId) returns()
func (_Dtars *DtarsTransactorSession) DisallowDTA(dtaAddr common.Address, dtaChainSelector uint64, fundAdminAddr common.Address, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtars.Contract.DisallowDTA(&_Dtars.TransactOpts, dtaAddr, dtaChainSelector, fundAdminAddr, fundTokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb92db27.
//
// Solidity: function initialize(uint64 localChainSelector) returns()
func (_Dtars *DtarsTransactor) Initialize(opts *bind.TransactOpts, localChainSelector uint64) (*types.Transaction, error) {
	return _Dtars.contract.Transact(opts, "initialize", localChainSelector)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb92db27.
//
// Solidity: function initialize(uint64 localChainSelector) returns()
func (_Dtars *DtarsSession) Initialize(localChainSelector uint64) (*types.Transaction, error) {
	return _Dtars.Contract.Initialize(&_Dtars.TransactOpts, localChainSelector)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb92db27.
//
// Solidity: function initialize(uint64 localChainSelector) returns()
func (_Dtars *DtarsTransactorSession) Initialize(localChainSelector uint64) (*types.Transaction, error) {
	return _Dtars.Contract.Initialize(&_Dtars.TransactOpts, localChainSelector)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address recipient) returns()
func (_Dtars *DtarsTransactor) RecoverFunds(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Dtars.contract.Transact(opts, "recoverFunds", recipient)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address recipient) returns()
func (_Dtars *DtarsSession) RecoverFunds(recipient common.Address) (*types.Transaction, error) {
	return _Dtars.Contract.RecoverFunds(&_Dtars.TransactOpts, recipient)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address recipient) returns()
func (_Dtars *DtarsTransactorSession) RecoverFunds(recipient common.Address) (*types.Transaction, error) {
	return _Dtars.Contract.RecoverFunds(&_Dtars.TransactOpts, recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtars *DtarsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtars.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtars *DtarsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dtars.Contract.RenounceOwnership(&_Dtars.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtars *DtarsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dtars.Contract.RenounceOwnership(&_Dtars.TransactOpts)
}

// SetCCIPGasLimit is a paid mutator transaction binding the contract method 0xca0880cf.
//
// Solidity: function setCCIPGasLimit(uint256 gasLimit) returns()
func (_Dtars *DtarsTransactor) SetCCIPGasLimit(opts *bind.TransactOpts, gasLimit *big.Int) (*types.Transaction, error) {
	return _Dtars.contract.Transact(opts, "setCCIPGasLimit", gasLimit)
}

// SetCCIPGasLimit is a paid mutator transaction binding the contract method 0xca0880cf.
//
// Solidity: function setCCIPGasLimit(uint256 gasLimit) returns()
func (_Dtars *DtarsSession) SetCCIPGasLimit(gasLimit *big.Int) (*types.Transaction, error) {
	return _Dtars.Contract.SetCCIPGasLimit(&_Dtars.TransactOpts, gasLimit)
}

// SetCCIPGasLimit is a paid mutator transaction binding the contract method 0xca0880cf.
//
// Solidity: function setCCIPGasLimit(uint256 gasLimit) returns()
func (_Dtars *DtarsTransactorSession) SetCCIPGasLimit(gasLimit *big.Int) (*types.Transaction, error) {
	return _Dtars.Contract.SetCCIPGasLimit(&_Dtars.TransactOpts, gasLimit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtars *DtarsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dtars.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtars *DtarsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dtars.Contract.TransferOwnership(&_Dtars.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtars *DtarsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dtars.Contract.TransferOwnership(&_Dtars.TransactOpts, newOwner)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address token, address recipient, uint256 amount) returns()
func (_Dtars *DtarsTransactor) WithdrawTokens(opts *bind.TransactOpts, token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dtars.contract.Transact(opts, "withdrawTokens", token, recipient, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address token, address recipient, uint256 amount) returns()
func (_Dtars *DtarsSession) WithdrawTokens(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dtars.Contract.WithdrawTokens(&_Dtars.TransactOpts, token, recipient, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address token, address recipient, uint256 amount) returns()
func (_Dtars *DtarsTransactorSession) WithdrawTokens(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dtars.Contract.WithdrawTokens(&_Dtars.TransactOpts, token, recipient, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtars *DtarsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtars.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtars *DtarsSession) Receive() (*types.Transaction, error) {
	return _Dtars.Contract.Receive(&_Dtars.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtars *DtarsTransactorSession) Receive() (*types.Transaction, error) {
	return _Dtars.Contract.Receive(&_Dtars.TransactOpts)
}

// DtarsCCIPMessageRecvFailedIterator is returned from FilterCCIPMessageRecvFailed and is used to iterate over the raw logs and unpacked data for CCIPMessageRecvFailed events raised by the Dtars contract.
type DtarsCCIPMessageRecvFailedIterator struct {
	Event *DtarsCCIPMessageRecvFailed // Event containing the contract specifics and raw log

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
func (it *DtarsCCIPMessageRecvFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarsCCIPMessageRecvFailed)
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
		it.Event = new(DtarsCCIPMessageRecvFailed)
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
func (it *DtarsCCIPMessageRecvFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarsCCIPMessageRecvFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarsCCIPMessageRecvFailed represents a CCIPMessageRecvFailed event raised by the Dtars contract.
type DtarsCCIPMessageRecvFailed struct {
	MessageId [32]byte
	Reason    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCCIPMessageRecvFailed is a free log retrieval operation binding the contract event 0x55f7fbddf1abc1bd0e3d869bb2ddf2b3351f7bdd1cba843b1634102ed60afdcb.
//
// Solidity: event CCIPMessageRecvFailed(bytes32 indexed messageId, bytes reason)
func (_Dtars *DtarsFilterer) FilterCCIPMessageRecvFailed(opts *bind.FilterOpts, messageId [][32]byte) (*DtarsCCIPMessageRecvFailedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Dtars.contract.FilterLogs(opts, "CCIPMessageRecvFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarsCCIPMessageRecvFailedIterator{contract: _Dtars.contract, event: "CCIPMessageRecvFailed", logs: logs, sub: sub}, nil
}

// WatchCCIPMessageRecvFailed is a free log subscription operation binding the contract event 0x55f7fbddf1abc1bd0e3d869bb2ddf2b3351f7bdd1cba843b1634102ed60afdcb.
//
// Solidity: event CCIPMessageRecvFailed(bytes32 indexed messageId, bytes reason)
func (_Dtars *DtarsFilterer) WatchCCIPMessageRecvFailed(opts *bind.WatchOpts, sink chan<- *DtarsCCIPMessageRecvFailed, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Dtars.contract.WatchLogs(opts, "CCIPMessageRecvFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarsCCIPMessageRecvFailed)
				if err := _Dtars.contract.UnpackLog(event, "CCIPMessageRecvFailed", log); err != nil {
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

// ParseCCIPMessageRecvFailed is a log parse operation binding the contract event 0x55f7fbddf1abc1bd0e3d869bb2ddf2b3351f7bdd1cba843b1634102ed60afdcb.
//
// Solidity: event CCIPMessageRecvFailed(bytes32 indexed messageId, bytes reason)
func (_Dtars *DtarsFilterer) ParseCCIPMessageRecvFailed(log types.Log) (*DtarsCCIPMessageRecvFailed, error) {
	event := new(DtarsCCIPMessageRecvFailed)
	if err := _Dtars.contract.UnpackLog(event, "CCIPMessageRecvFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarsDTAAddedIterator is returned from FilterDTAAdded and is used to iterate over the raw logs and unpacked data for DTAAdded events raised by the Dtars contract.
type DtarsDTAAddedIterator struct {
	Event *DtarsDTAAdded // Event containing the contract specifics and raw log

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
func (it *DtarsDTAAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarsDTAAdded)
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
		it.Event = new(DtarsDTAAdded)
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
func (it *DtarsDTAAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarsDTAAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarsDTAAdded represents a DTAAdded event raised by the Dtars contract.
type DtarsDTAAdded struct {
	DtaAddr          common.Address
	DtaChainSelector uint64
	FundAdminAddr    common.Address
	FundTokenId      [32]byte
	FundTokenAddr    common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDTAAdded is a free log retrieval operation binding the contract event 0x0d18ffbbabe35515bd4698694df785c6e4dce11673b50cd4bbd046d8a55be435.
//
// Solidity: event DTAAdded(address indexed dtaAddr, uint64 indexed dtaChainSelector, address fundAdminAddr, bytes32 indexed fundTokenId, address fundTokenAddr)
func (_Dtars *DtarsFilterer) FilterDTAAdded(opts *bind.FilterOpts, dtaAddr []common.Address, dtaChainSelector []uint64, fundTokenId [][32]byte) (*DtarsDTAAddedIterator, error) {

	var dtaAddrRule []interface{}
	for _, dtaAddrItem := range dtaAddr {
		dtaAddrRule = append(dtaAddrRule, dtaAddrItem)
	}
	var dtaChainSelectorRule []interface{}
	for _, dtaChainSelectorItem := range dtaChainSelector {
		dtaChainSelectorRule = append(dtaChainSelectorRule, dtaChainSelectorItem)
	}

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtars.contract.FilterLogs(opts, "DTAAdded", dtaAddrRule, dtaChainSelectorRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarsDTAAddedIterator{contract: _Dtars.contract, event: "DTAAdded", logs: logs, sub: sub}, nil
}

// WatchDTAAdded is a free log subscription operation binding the contract event 0x0d18ffbbabe35515bd4698694df785c6e4dce11673b50cd4bbd046d8a55be435.
//
// Solidity: event DTAAdded(address indexed dtaAddr, uint64 indexed dtaChainSelector, address fundAdminAddr, bytes32 indexed fundTokenId, address fundTokenAddr)
func (_Dtars *DtarsFilterer) WatchDTAAdded(opts *bind.WatchOpts, sink chan<- *DtarsDTAAdded, dtaAddr []common.Address, dtaChainSelector []uint64, fundTokenId [][32]byte) (event.Subscription, error) {

	var dtaAddrRule []interface{}
	for _, dtaAddrItem := range dtaAddr {
		dtaAddrRule = append(dtaAddrRule, dtaAddrItem)
	}
	var dtaChainSelectorRule []interface{}
	for _, dtaChainSelectorItem := range dtaChainSelector {
		dtaChainSelectorRule = append(dtaChainSelectorRule, dtaChainSelectorItem)
	}

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtars.contract.WatchLogs(opts, "DTAAdded", dtaAddrRule, dtaChainSelectorRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarsDTAAdded)
				if err := _Dtars.contract.UnpackLog(event, "DTAAdded", log); err != nil {
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

// ParseDTAAdded is a log parse operation binding the contract event 0x0d18ffbbabe35515bd4698694df785c6e4dce11673b50cd4bbd046d8a55be435.
//
// Solidity: event DTAAdded(address indexed dtaAddr, uint64 indexed dtaChainSelector, address fundAdminAddr, bytes32 indexed fundTokenId, address fundTokenAddr)
func (_Dtars *DtarsFilterer) ParseDTAAdded(log types.Log) (*DtarsDTAAdded, error) {
	event := new(DtarsDTAAdded)
	if err := _Dtars.contract.UnpackLog(event, "DTAAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarsDTARemovedIterator is returned from FilterDTARemoved and is used to iterate over the raw logs and unpacked data for DTARemoved events raised by the Dtars contract.
type DtarsDTARemovedIterator struct {
	Event *DtarsDTARemoved // Event containing the contract specifics and raw log

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
func (it *DtarsDTARemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarsDTARemoved)
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
		it.Event = new(DtarsDTARemoved)
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
func (it *DtarsDTARemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarsDTARemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarsDTARemoved represents a DTARemoved event raised by the Dtars contract.
type DtarsDTARemoved struct {
	DtaAddr          common.Address
	DtaChainSelector uint64
	FundAdminAddr    common.Address
	FundTokenId      [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDTARemoved is a free log retrieval operation binding the contract event 0x10ff155cc426567ffe468a7ea2929661be58b9f5c959b1ff6bc70fd68a0d71a6.
//
// Solidity: event DTARemoved(address indexed dtaAddr, uint64 indexed dtaChainSelector, address fundAdminAddr, bytes32 indexed fundTokenId)
func (_Dtars *DtarsFilterer) FilterDTARemoved(opts *bind.FilterOpts, dtaAddr []common.Address, dtaChainSelector []uint64, fundTokenId [][32]byte) (*DtarsDTARemovedIterator, error) {

	var dtaAddrRule []interface{}
	for _, dtaAddrItem := range dtaAddr {
		dtaAddrRule = append(dtaAddrRule, dtaAddrItem)
	}
	var dtaChainSelectorRule []interface{}
	for _, dtaChainSelectorItem := range dtaChainSelector {
		dtaChainSelectorRule = append(dtaChainSelectorRule, dtaChainSelectorItem)
	}

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtars.contract.FilterLogs(opts, "DTARemoved", dtaAddrRule, dtaChainSelectorRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarsDTARemovedIterator{contract: _Dtars.contract, event: "DTARemoved", logs: logs, sub: sub}, nil
}

// WatchDTARemoved is a free log subscription operation binding the contract event 0x10ff155cc426567ffe468a7ea2929661be58b9f5c959b1ff6bc70fd68a0d71a6.
//
// Solidity: event DTARemoved(address indexed dtaAddr, uint64 indexed dtaChainSelector, address fundAdminAddr, bytes32 indexed fundTokenId)
func (_Dtars *DtarsFilterer) WatchDTARemoved(opts *bind.WatchOpts, sink chan<- *DtarsDTARemoved, dtaAddr []common.Address, dtaChainSelector []uint64, fundTokenId [][32]byte) (event.Subscription, error) {

	var dtaAddrRule []interface{}
	for _, dtaAddrItem := range dtaAddr {
		dtaAddrRule = append(dtaAddrRule, dtaAddrItem)
	}
	var dtaChainSelectorRule []interface{}
	for _, dtaChainSelectorItem := range dtaChainSelector {
		dtaChainSelectorRule = append(dtaChainSelectorRule, dtaChainSelectorItem)
	}

	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtars.contract.WatchLogs(opts, "DTARemoved", dtaAddrRule, dtaChainSelectorRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarsDTARemoved)
				if err := _Dtars.contract.UnpackLog(event, "DTARemoved", log); err != nil {
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

// ParseDTARemoved is a log parse operation binding the contract event 0x10ff155cc426567ffe468a7ea2929661be58b9f5c959b1ff6bc70fd68a0d71a6.
//
// Solidity: event DTARemoved(address indexed dtaAddr, uint64 indexed dtaChainSelector, address fundAdminAddr, bytes32 indexed fundTokenId)
func (_Dtars *DtarsFilterer) ParseDTARemoved(log types.Log) (*DtarsDTARemoved, error) {
	event := new(DtarsDTARemoved)
	if err := _Dtars.contract.UnpackLog(event, "DTARemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarsDTASettlementClosedIterator is returned from FilterDTASettlementClosed and is used to iterate over the raw logs and unpacked data for DTASettlementClosed events raised by the Dtars contract.
type DtarsDTASettlementClosedIterator struct {
	Event *DtarsDTASettlementClosed // Event containing the contract specifics and raw log

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
func (it *DtarsDTASettlementClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarsDTASettlementClosed)
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
		it.Event = new(DtarsDTASettlementClosed)
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
func (it *DtarsDTASettlementClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarsDTASettlementClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarsDTASettlementClosed represents a DTASettlementClosed event raised by the Dtars contract.
type DtarsDTASettlementClosed struct {
	FundAdminAddr    common.Address
	FundTokenId      [32]byte
	RequestType      uint8
	DistributorAddr  common.Address
	DtaChainSelector uint64
	DtaAddr          common.Address
	RequestId        [32]byte
	Success          bool
	Err              []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDTASettlementClosed is a free log retrieval operation binding the contract event 0x5c71a76872d1e0dc218857abf9f3e2db734cb991ee4a70393f9c99be273e6c2a.
//
// Solidity: event DTASettlementClosed(address indexed fundAdminAddr, bytes32 indexed fundTokenId, uint8 indexed requestType, address distributorAddr, uint64 dtaChainSelector, address dtaAddr, bytes32 requestId, bool success, bytes err)
func (_Dtars *DtarsFilterer) FilterDTASettlementClosed(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte, requestType []uint8) (*DtarsDTASettlementClosedIterator, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var requestTypeRule []interface{}
	for _, requestTypeItem := range requestType {
		requestTypeRule = append(requestTypeRule, requestTypeItem)
	}

	logs, sub, err := _Dtars.contract.FilterLogs(opts, "DTASettlementClosed", fundAdminAddrRule, fundTokenIdRule, requestTypeRule)
	if err != nil {
		return nil, err
	}
	return &DtarsDTASettlementClosedIterator{contract: _Dtars.contract, event: "DTASettlementClosed", logs: logs, sub: sub}, nil
}

// WatchDTASettlementClosed is a free log subscription operation binding the contract event 0x5c71a76872d1e0dc218857abf9f3e2db734cb991ee4a70393f9c99be273e6c2a.
//
// Solidity: event DTASettlementClosed(address indexed fundAdminAddr, bytes32 indexed fundTokenId, uint8 indexed requestType, address distributorAddr, uint64 dtaChainSelector, address dtaAddr, bytes32 requestId, bool success, bytes err)
func (_Dtars *DtarsFilterer) WatchDTASettlementClosed(opts *bind.WatchOpts, sink chan<- *DtarsDTASettlementClosed, fundAdminAddr []common.Address, fundTokenId [][32]byte, requestType []uint8) (event.Subscription, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var requestTypeRule []interface{}
	for _, requestTypeItem := range requestType {
		requestTypeRule = append(requestTypeRule, requestTypeItem)
	}

	logs, sub, err := _Dtars.contract.WatchLogs(opts, "DTASettlementClosed", fundAdminAddrRule, fundTokenIdRule, requestTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarsDTASettlementClosed)
				if err := _Dtars.contract.UnpackLog(event, "DTASettlementClosed", log); err != nil {
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

// ParseDTASettlementClosed is a log parse operation binding the contract event 0x5c71a76872d1e0dc218857abf9f3e2db734cb991ee4a70393f9c99be273e6c2a.
//
// Solidity: event DTASettlementClosed(address indexed fundAdminAddr, bytes32 indexed fundTokenId, uint8 indexed requestType, address distributorAddr, uint64 dtaChainSelector, address dtaAddr, bytes32 requestId, bool success, bytes err)
func (_Dtars *DtarsFilterer) ParseDTASettlementClosed(log types.Log) (*DtarsDTASettlementClosed, error) {
	event := new(DtarsDTASettlementClosed)
	if err := _Dtars.contract.UnpackLog(event, "DTASettlementClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarsDTASettlementOpenedIterator is returned from FilterDTASettlementOpened and is used to iterate over the raw logs and unpacked data for DTASettlementOpened events raised by the Dtars contract.
type DtarsDTASettlementOpenedIterator struct {
	Event *DtarsDTASettlementOpened // Event containing the contract specifics and raw log

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
func (it *DtarsDTASettlementOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarsDTASettlementOpened)
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
		it.Event = new(DtarsDTASettlementOpened)
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
func (it *DtarsDTASettlementOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarsDTASettlementOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarsDTASettlementOpened represents a DTASettlementOpened event raised by the Dtars contract.
type DtarsDTASettlementOpened struct {
	FundAdminAddr         common.Address
	FundTokenId           [32]byte
	RequestType           uint8
	DistributorAddr       common.Address
	DtaChainSelector      uint64
	DtaAddr               common.Address
	RequestId             [32]byte
	DistributorWalletAddr common.Address
	Shares                *big.Int
	Amount                *big.Int
	Currency              uint8
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDTASettlementOpened is a free log retrieval operation binding the contract event 0xb3c9f6a7de5ff1522c00816154846198e5a1a4c2919fedbfa5ef96b07195c8d2.
//
// Solidity: event DTASettlementOpened(address indexed fundAdminAddr, bytes32 indexed fundTokenId, uint8 indexed requestType, address distributorAddr, uint64 dtaChainSelector, address dtaAddr, bytes32 requestId, address distributorWalletAddr, uint256 shares, uint256 amount, uint8 currency)
func (_Dtars *DtarsFilterer) FilterDTASettlementOpened(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte, requestType []uint8) (*DtarsDTASettlementOpenedIterator, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var requestTypeRule []interface{}
	for _, requestTypeItem := range requestType {
		requestTypeRule = append(requestTypeRule, requestTypeItem)
	}

	logs, sub, err := _Dtars.contract.FilterLogs(opts, "DTASettlementOpened", fundAdminAddrRule, fundTokenIdRule, requestTypeRule)
	if err != nil {
		return nil, err
	}
	return &DtarsDTASettlementOpenedIterator{contract: _Dtars.contract, event: "DTASettlementOpened", logs: logs, sub: sub}, nil
}

// WatchDTASettlementOpened is a free log subscription operation binding the contract event 0xb3c9f6a7de5ff1522c00816154846198e5a1a4c2919fedbfa5ef96b07195c8d2.
//
// Solidity: event DTASettlementOpened(address indexed fundAdminAddr, bytes32 indexed fundTokenId, uint8 indexed requestType, address distributorAddr, uint64 dtaChainSelector, address dtaAddr, bytes32 requestId, address distributorWalletAddr, uint256 shares, uint256 amount, uint8 currency)
func (_Dtars *DtarsFilterer) WatchDTASettlementOpened(opts *bind.WatchOpts, sink chan<- *DtarsDTASettlementOpened, fundAdminAddr []common.Address, fundTokenId [][32]byte, requestType []uint8) (event.Subscription, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var requestTypeRule []interface{}
	for _, requestTypeItem := range requestType {
		requestTypeRule = append(requestTypeRule, requestTypeItem)
	}

	logs, sub, err := _Dtars.contract.WatchLogs(opts, "DTASettlementOpened", fundAdminAddrRule, fundTokenIdRule, requestTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarsDTASettlementOpened)
				if err := _Dtars.contract.UnpackLog(event, "DTASettlementOpened", log); err != nil {
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

// ParseDTASettlementOpened is a log parse operation binding the contract event 0xb3c9f6a7de5ff1522c00816154846198e5a1a4c2919fedbfa5ef96b07195c8d2.
//
// Solidity: event DTASettlementOpened(address indexed fundAdminAddr, bytes32 indexed fundTokenId, uint8 indexed requestType, address distributorAddr, uint64 dtaChainSelector, address dtaAddr, bytes32 requestId, address distributorWalletAddr, uint256 shares, uint256 amount, uint8 currency)
func (_Dtars *DtarsFilterer) ParseDTASettlementOpened(log types.Log) (*DtarsDTASettlementOpened, error) {
	event := new(DtarsDTASettlementOpened)
	if err := _Dtars.contract.UnpackLog(event, "DTASettlementOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarsEmptyRequestTypeIterator is returned from FilterEmptyRequestType and is used to iterate over the raw logs and unpacked data for EmptyRequestType events raised by the Dtars contract.
type DtarsEmptyRequestTypeIterator struct {
	Event *DtarsEmptyRequestType // Event containing the contract specifics and raw log

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
func (it *DtarsEmptyRequestTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarsEmptyRequestType)
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
		it.Event = new(DtarsEmptyRequestType)
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
func (it *DtarsEmptyRequestTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarsEmptyRequestTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarsEmptyRequestType represents a EmptyRequestType event raised by the Dtars contract.
type DtarsEmptyRequestType struct {
	MessageId [32]byte
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEmptyRequestType is a free log retrieval operation binding the contract event 0xcd5c141a2c935a10029d428cf166c562516f4d59acc952185c1768152f112608.
//
// Solidity: event EmptyRequestType(bytes32 indexed messageId, bytes32 indexed requestId)
func (_Dtars *DtarsFilterer) FilterEmptyRequestType(opts *bind.FilterOpts, messageId [][32]byte, requestId [][32]byte) (*DtarsEmptyRequestTypeIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Dtars.contract.FilterLogs(opts, "EmptyRequestType", messageIdRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarsEmptyRequestTypeIterator{contract: _Dtars.contract, event: "EmptyRequestType", logs: logs, sub: sub}, nil
}

// WatchEmptyRequestType is a free log subscription operation binding the contract event 0xcd5c141a2c935a10029d428cf166c562516f4d59acc952185c1768152f112608.
//
// Solidity: event EmptyRequestType(bytes32 indexed messageId, bytes32 indexed requestId)
func (_Dtars *DtarsFilterer) WatchEmptyRequestType(opts *bind.WatchOpts, sink chan<- *DtarsEmptyRequestType, messageId [][32]byte, requestId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Dtars.contract.WatchLogs(opts, "EmptyRequestType", messageIdRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarsEmptyRequestType)
				if err := _Dtars.contract.UnpackLog(event, "EmptyRequestType", log); err != nil {
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

// ParseEmptyRequestType is a log parse operation binding the contract event 0xcd5c141a2c935a10029d428cf166c562516f4d59acc952185c1768152f112608.
//
// Solidity: event EmptyRequestType(bytes32 indexed messageId, bytes32 indexed requestId)
func (_Dtars *DtarsFilterer) ParseEmptyRequestType(log types.Log) (*DtarsEmptyRequestType, error) {
	event := new(DtarsEmptyRequestType)
	if err := _Dtars.contract.UnpackLog(event, "EmptyRequestType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dtars contract.
type DtarsInitializedIterator struct {
	Event *DtarsInitialized // Event containing the contract specifics and raw log

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
func (it *DtarsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarsInitialized)
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
		it.Event = new(DtarsInitialized)
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
func (it *DtarsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarsInitialized represents a Initialized event raised by the Dtars contract.
type DtarsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Dtars *DtarsFilterer) FilterInitialized(opts *bind.FilterOpts) (*DtarsInitializedIterator, error) {

	logs, sub, err := _Dtars.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DtarsInitializedIterator{contract: _Dtars.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Dtars *DtarsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DtarsInitialized) (event.Subscription, error) {

	logs, sub, err := _Dtars.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarsInitialized)
				if err := _Dtars.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Dtars *DtarsFilterer) ParseInitialized(log types.Log) (*DtarsInitialized, error) {
	event := new(DtarsInitialized)
	if err := _Dtars.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarsInvalidSubscriptionCrossChainPaymentIterator is returned from FilterInvalidSubscriptionCrossChainPayment and is used to iterate over the raw logs and unpacked data for InvalidSubscriptionCrossChainPayment events raised by the Dtars contract.
type DtarsInvalidSubscriptionCrossChainPaymentIterator struct {
	Event *DtarsInvalidSubscriptionCrossChainPayment // Event containing the contract specifics and raw log

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
func (it *DtarsInvalidSubscriptionCrossChainPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarsInvalidSubscriptionCrossChainPayment)
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
		it.Event = new(DtarsInvalidSubscriptionCrossChainPayment)
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
func (it *DtarsInvalidSubscriptionCrossChainPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarsInvalidSubscriptionCrossChainPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarsInvalidSubscriptionCrossChainPayment represents a InvalidSubscriptionCrossChainPayment event raised by the Dtars contract.
type DtarsInvalidSubscriptionCrossChainPayment struct {
	FundAdminAddr              common.Address
	FundTokenId                [32]byte
	RequestId                  [32]byte
	PaymentTokenDestAddr       common.Address
	CcipDestTokenAmountsLength *big.Int
	CcipPaymentTokenAddr       common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterInvalidSubscriptionCrossChainPayment is a free log retrieval operation binding the contract event 0x22c6bd14d208b4294c75d93e7eb52a4396891c57b9ea18e4da15444c5acb8904.
//
// Solidity: event InvalidSubscriptionCrossChainPayment(address indexed fundAdminAddr, bytes32 indexed fundTokenId, bytes32 requestId, address paymentTokenDestAddr, uint256 ccipDestTokenAmountsLength, address ccipPaymentTokenAddr)
func (_Dtars *DtarsFilterer) FilterInvalidSubscriptionCrossChainPayment(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte) (*DtarsInvalidSubscriptionCrossChainPaymentIterator, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtars.contract.FilterLogs(opts, "InvalidSubscriptionCrossChainPayment", fundAdminAddrRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarsInvalidSubscriptionCrossChainPaymentIterator{contract: _Dtars.contract, event: "InvalidSubscriptionCrossChainPayment", logs: logs, sub: sub}, nil
}

// WatchInvalidSubscriptionCrossChainPayment is a free log subscription operation binding the contract event 0x22c6bd14d208b4294c75d93e7eb52a4396891c57b9ea18e4da15444c5acb8904.
//
// Solidity: event InvalidSubscriptionCrossChainPayment(address indexed fundAdminAddr, bytes32 indexed fundTokenId, bytes32 requestId, address paymentTokenDestAddr, uint256 ccipDestTokenAmountsLength, address ccipPaymentTokenAddr)
func (_Dtars *DtarsFilterer) WatchInvalidSubscriptionCrossChainPayment(opts *bind.WatchOpts, sink chan<- *DtarsInvalidSubscriptionCrossChainPayment, fundAdminAddr []common.Address, fundTokenId [][32]byte) (event.Subscription, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtars.contract.WatchLogs(opts, "InvalidSubscriptionCrossChainPayment", fundAdminAddrRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarsInvalidSubscriptionCrossChainPayment)
				if err := _Dtars.contract.UnpackLog(event, "InvalidSubscriptionCrossChainPayment", log); err != nil {
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

// ParseInvalidSubscriptionCrossChainPayment is a log parse operation binding the contract event 0x22c6bd14d208b4294c75d93e7eb52a4396891c57b9ea18e4da15444c5acb8904.
//
// Solidity: event InvalidSubscriptionCrossChainPayment(address indexed fundAdminAddr, bytes32 indexed fundTokenId, bytes32 requestId, address paymentTokenDestAddr, uint256 ccipDestTokenAmountsLength, address ccipPaymentTokenAddr)
func (_Dtars *DtarsFilterer) ParseInvalidSubscriptionCrossChainPayment(log types.Log) (*DtarsInvalidSubscriptionCrossChainPayment, error) {
	event := new(DtarsInvalidSubscriptionCrossChainPayment)
	if err := _Dtars.contract.UnpackLog(event, "InvalidSubscriptionCrossChainPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarsNativeFundsRecoveredIterator is returned from FilterNativeFundsRecovered and is used to iterate over the raw logs and unpacked data for NativeFundsRecovered events raised by the Dtars contract.
type DtarsNativeFundsRecoveredIterator struct {
	Event *DtarsNativeFundsRecovered // Event containing the contract specifics and raw log

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
func (it *DtarsNativeFundsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarsNativeFundsRecovered)
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
		it.Event = new(DtarsNativeFundsRecovered)
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
func (it *DtarsNativeFundsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarsNativeFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarsNativeFundsRecovered represents a NativeFundsRecovered event raised by the Dtars contract.
type DtarsNativeFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeFundsRecovered is a free log retrieval operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_Dtars *DtarsFilterer) FilterNativeFundsRecovered(opts *bind.FilterOpts) (*DtarsNativeFundsRecoveredIterator, error) {

	logs, sub, err := _Dtars.contract.FilterLogs(opts, "NativeFundsRecovered")
	if err != nil {
		return nil, err
	}
	return &DtarsNativeFundsRecoveredIterator{contract: _Dtars.contract, event: "NativeFundsRecovered", logs: logs, sub: sub}, nil
}

// WatchNativeFundsRecovered is a free log subscription operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_Dtars *DtarsFilterer) WatchNativeFundsRecovered(opts *bind.WatchOpts, sink chan<- *DtarsNativeFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _Dtars.contract.WatchLogs(opts, "NativeFundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarsNativeFundsRecovered)
				if err := _Dtars.contract.UnpackLog(event, "NativeFundsRecovered", log); err != nil {
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

// ParseNativeFundsRecovered is a log parse operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_Dtars *DtarsFilterer) ParseNativeFundsRecovered(log types.Log) (*DtarsNativeFundsRecovered, error) {
	event := new(DtarsNativeFundsRecovered)
	if err := _Dtars.contract.UnpackLog(event, "NativeFundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dtars contract.
type DtarsOwnershipTransferredIterator struct {
	Event *DtarsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DtarsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarsOwnershipTransferred)
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
		it.Event = new(DtarsOwnershipTransferred)
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
func (it *DtarsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarsOwnershipTransferred represents a OwnershipTransferred event raised by the Dtars contract.
type DtarsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtars *DtarsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DtarsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dtars.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DtarsOwnershipTransferredIterator{contract: _Dtars.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtars *DtarsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DtarsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dtars.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarsOwnershipTransferred)
				if err := _Dtars.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dtars *DtarsFilterer) ParseOwnershipTransferred(log types.Log) (*DtarsOwnershipTransferred, error) {
	event := new(DtarsOwnershipTransferred)
	if err := _Dtars.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarsSettlementFailedIterator is returned from FilterSettlementFailed and is used to iterate over the raw logs and unpacked data for SettlementFailed events raised by the Dtars contract.
type DtarsSettlementFailedIterator struct {
	Event *DtarsSettlementFailed // Event containing the contract specifics and raw log

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
func (it *DtarsSettlementFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarsSettlementFailed)
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
		it.Event = new(DtarsSettlementFailed)
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
func (it *DtarsSettlementFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarsSettlementFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarsSettlementFailed represents a SettlementFailed event raised by the Dtars contract.
type DtarsSettlementFailed struct {
	FundAdminAddr         common.Address
	FundTokenId           [32]byte
	DistributorAddr       common.Address
	PaymentTokenAddr      common.Address
	DistributorWalletAddr common.Address
	RequestId             [32]byte
	Shares                *big.Int
	Amount                *big.Int
	ErrData               []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSettlementFailed is a free log retrieval operation binding the contract event 0x8192ce999af5379d3a955c27840222f7cd179a6b2636283d6563c6e8a69a75cd.
//
// Solidity: event SettlementFailed(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, address paymentTokenAddr, address distributorWalletAddr, bytes32 requestId, uint256 shares, uint256 amount, bytes errData)
func (_Dtars *DtarsFilterer) FilterSettlementFailed(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtarsSettlementFailedIterator, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtars.contract.FilterLogs(opts, "SettlementFailed", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtarsSettlementFailedIterator{contract: _Dtars.contract, event: "SettlementFailed", logs: logs, sub: sub}, nil
}

// WatchSettlementFailed is a free log subscription operation binding the contract event 0x8192ce999af5379d3a955c27840222f7cd179a6b2636283d6563c6e8a69a75cd.
//
// Solidity: event SettlementFailed(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, address paymentTokenAddr, address distributorWalletAddr, bytes32 requestId, uint256 shares, uint256 amount, bytes errData)
func (_Dtars *DtarsFilterer) WatchSettlementFailed(opts *bind.WatchOpts, sink chan<- *DtarsSettlementFailed, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var distributorAddrRule []interface{}
	for _, distributorAddrItem := range distributorAddr {
		distributorAddrRule = append(distributorAddrRule, distributorAddrItem)
	}

	logs, sub, err := _Dtars.contract.WatchLogs(opts, "SettlementFailed", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarsSettlementFailed)
				if err := _Dtars.contract.UnpackLog(event, "SettlementFailed", log); err != nil {
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

// ParseSettlementFailed is a log parse operation binding the contract event 0x8192ce999af5379d3a955c27840222f7cd179a6b2636283d6563c6e8a69a75cd.
//
// Solidity: event SettlementFailed(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, address paymentTokenAddr, address distributorWalletAddr, bytes32 requestId, uint256 shares, uint256 amount, bytes errData)
func (_Dtars *DtarsFilterer) ParseSettlementFailed(log types.Log) (*DtarsSettlementFailed, error) {
	event := new(DtarsSettlementFailed)
	if err := _Dtars.contract.UnpackLog(event, "SettlementFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarsTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the Dtars contract.
type DtarsTokenWithdrawnIterator struct {
	Event *DtarsTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *DtarsTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarsTokenWithdrawn)
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
		it.Event = new(DtarsTokenWithdrawn)
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
func (it *DtarsTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarsTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarsTokenWithdrawn represents a TokenWithdrawn event raised by the Dtars contract.
type DtarsTokenWithdrawn struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed recipient, uint256 amount)
func (_Dtars *DtarsFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*DtarsTokenWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Dtars.contract.FilterLogs(opts, "TokenWithdrawn", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &DtarsTokenWithdrawnIterator{contract: _Dtars.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed recipient, uint256 amount)
func (_Dtars *DtarsFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *DtarsTokenWithdrawn, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Dtars.contract.WatchLogs(opts, "TokenWithdrawn", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarsTokenWithdrawn)
				if err := _Dtars.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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

// ParseTokenWithdrawn is a log parse operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed recipient, uint256 amount)
func (_Dtars *DtarsFilterer) ParseTokenWithdrawn(log types.Log) (*DtarsTokenWithdrawn, error) {
	event := new(DtarsTokenWithdrawn)
	if err := _Dtars.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarsUnauthorizedSenderDTAIterator is returned from FilterUnauthorizedSenderDTA and is used to iterate over the raw logs and unpacked data for UnauthorizedSenderDTA events raised by the Dtars contract.
type DtarsUnauthorizedSenderDTAIterator struct {
	Event *DtarsUnauthorizedSenderDTA // Event containing the contract specifics and raw log

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
func (it *DtarsUnauthorizedSenderDTAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarsUnauthorizedSenderDTA)
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
		it.Event = new(DtarsUnauthorizedSenderDTA)
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
func (it *DtarsUnauthorizedSenderDTAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarsUnauthorizedSenderDTAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarsUnauthorizedSenderDTA represents a UnauthorizedSenderDTA event raised by the Dtars contract.
type DtarsUnauthorizedSenderDTA struct {
	DtaAddr          common.Address
	DtaChainSelector uint64
	FundAdminAddr    common.Address
	FundTokenId      [32]byte
	DistributorAddr  common.Address
	RequestId        [32]byte
	ReqType          uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUnauthorizedSenderDTA is a free log retrieval operation binding the contract event 0x84c39baab915017fa80ad6633f3bad811f77bd8cdf4214e37c93fd980806e04b.
//
// Solidity: event UnauthorizedSenderDTA(address dtaAddr, uint64 dtaChainSelector, address fundAdminAddr, bytes32 fundTokenId, address distributorAddr, bytes32 requestId, uint8 reqType)
func (_Dtars *DtarsFilterer) FilterUnauthorizedSenderDTA(opts *bind.FilterOpts) (*DtarsUnauthorizedSenderDTAIterator, error) {

	logs, sub, err := _Dtars.contract.FilterLogs(opts, "UnauthorizedSenderDTA")
	if err != nil {
		return nil, err
	}
	return &DtarsUnauthorizedSenderDTAIterator{contract: _Dtars.contract, event: "UnauthorizedSenderDTA", logs: logs, sub: sub}, nil
}

// WatchUnauthorizedSenderDTA is a free log subscription operation binding the contract event 0x84c39baab915017fa80ad6633f3bad811f77bd8cdf4214e37c93fd980806e04b.
//
// Solidity: event UnauthorizedSenderDTA(address dtaAddr, uint64 dtaChainSelector, address fundAdminAddr, bytes32 fundTokenId, address distributorAddr, bytes32 requestId, uint8 reqType)
func (_Dtars *DtarsFilterer) WatchUnauthorizedSenderDTA(opts *bind.WatchOpts, sink chan<- *DtarsUnauthorizedSenderDTA) (event.Subscription, error) {

	logs, sub, err := _Dtars.contract.WatchLogs(opts, "UnauthorizedSenderDTA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarsUnauthorizedSenderDTA)
				if err := _Dtars.contract.UnpackLog(event, "UnauthorizedSenderDTA", log); err != nil {
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

// ParseUnauthorizedSenderDTA is a log parse operation binding the contract event 0x84c39baab915017fa80ad6633f3bad811f77bd8cdf4214e37c93fd980806e04b.
//
// Solidity: event UnauthorizedSenderDTA(address dtaAddr, uint64 dtaChainSelector, address fundAdminAddr, bytes32 fundTokenId, address distributorAddr, bytes32 requestId, uint8 reqType)
func (_Dtars *DtarsFilterer) ParseUnauthorizedSenderDTA(log types.Log) (*DtarsUnauthorizedSenderDTA, error) {
	event := new(DtarsUnauthorizedSenderDTA)
	if err := _Dtars.contract.UnpackLog(event, "UnauthorizedSenderDTA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
