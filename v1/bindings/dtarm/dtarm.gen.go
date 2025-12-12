// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dtarm

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

// IDTADistributorDistributorRequest is an auto generated low-level Go binding around an user-defined struct.
type IDTADistributorDistributorRequest struct {
	Shares          *big.Int
	Amount          *big.Int
	FundTokenId     [32]byte
	FundAdminAddr   common.Address
	DistributorAddr common.Address
	CreatedAt       *big.Int
	RequestType     uint8
	Status          uint8
}

// IDTAMessageDTAPayment is an auto generated low-level Go binding around an user-defined struct.
type IDTAMessageDTAPayment struct {
	OffChainPaymentCurrency uint8
	PaymentTokenSourceAddr  common.Address
	PaymentTokenDestAddr    common.Address
}

// IDTAMessageDtaRequestStatusMessage is an auto generated low-level Go binding around an user-defined struct.
type IDTAMessageDtaRequestStatusMessage struct {
	RequestId [32]byte
	Status    uint8
	Err       []byte
}

// IFundTokenRegistryFundTokenData is an auto generated low-level Go binding around an user-defined struct.
type IFundTokenRegistryFundTokenData struct {
	FundTokenAddr                 common.Address
	NavFeedDecimals               uint8
	PurchaseTokenRoundingDecimals uint8
	PurchaseTokenDecimals         uint8
	FundRoundingDecimals          uint8
	FundTokenDecimals             uint8
	RequestsPerDay                uint8
	NavAddr                       common.Address
	TokenChainSelector            uint64
	DtaRequestSettlementAddr      common.Address
	TimezoneOffsetSecs            *big.Int
	NavTTL                        *big.Int
	PaymentInfo                   IDTAMessageDTAPayment
}

// IOpenFundDistributorRegistryDistributorData is an auto generated low-level Go binding around an user-defined struct.
type IOpenFundDistributorRegistryDistributorData struct {
	DistributorWalletAddr     common.Address
	IsWalletOwnershipVerified bool
}

// DtarmMetaData contains all meta data concerning the Dtarm contract.
var DtarmMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"ccipRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"allowDistributorForToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDistributorRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDistributorRequest\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dtaMessage\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestStatusMessage\",\"components\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"directCompleteDistributorRequest\",\"inputs\":[{\"name\":\"dtaMessage\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DtaRequestStatusMessage\",\"components\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableFundToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disallowDistributorForToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableFundToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceAllowDistributorForToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCCIPGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributor\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIOpenFundDistributorRegistry.DistributorData\",\"components\":[{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isWalletOwnershipVerified\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributorRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDTADistributor.DistributorRequest\",\"components\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"requestType\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.DistributorRequestType\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributors\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"distributorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFundAdmins\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fundAdminAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFundToken\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFundTokenRegistry.FundTokenData\",\"components\":[{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"navFeedDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"purchaseTokenRoundingDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"purchaseTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fundRoundingDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fundTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"requestsPerDay\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"navAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dtaRequestSettlementAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timezoneOffsetSecs\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"navTTL\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFundTokens\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLocalChainSelector\",\"inputs\":[],\"outputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenRequestsByDate\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"date\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDistributorAllowedForToken\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isTokenEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFundAdminRegistered\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processDistributorRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverFunds\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerDistributor\",\"inputs\":[{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerFundAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerFundToken\",\"inputs\":[{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenData\",\"type\":\"tuple\",\"internalType\":\"structIFundTokenRegistry.FundTokenData\",\"components\":[{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"navFeedDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"purchaseTokenRoundingDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"purchaseTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fundRoundingDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fundTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"requestsPerDay\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"navAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dtaRequestSettlementAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timezoneOffsetSecs\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"navTTL\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"paymentInfo\",\"type\":\"tuple\",\"internalType\":\"structIDTAMessage.DTAPayment\",\"components\":[{\"name\":\"offChainPaymentCurrency\",\"type\":\"uint8\",\"internalType\":\"enumCurrency\"},{\"name\":\"paymentTokenSourceAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentTokenDestAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestRedemption\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestSubscription\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyDistributorWallet\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTokens\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DistributorRegistered\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributorRequestCanceled\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributorRequestProcessed\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIDTAMessage.RequestStatus\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributorRequestProcessing\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundAdminRegistered\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundTokenAllowlistUpdated\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundTokenRegistered\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fundTokenAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"navAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvalidDTARequestSettlement\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"actualChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"actualDTARequestSettlementAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageFailed\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeFundsRecovered\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedemptionRequested\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubscriptionRequested\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenWithdrawn\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressDoesNotExist\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressExists\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CancelNotAllowed\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestStatus\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"}]},{\"type\":\"error\",\"name\":\"DistributorNotAllowed\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DistributorWalletOwnershipNotVerified\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DoesNotExist\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ErrEscrowPaymentToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Exists\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FailedToRecoverFunds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDistributorWallet\",\"inputs\":[{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributorWalletAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidFeeManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFundTokenData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidLocalChainSelector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNAV\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nav\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"type\":\"error\",\"name\":\"InvalidPaymentInfo\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenDecimals\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWithdrawInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RateLimitExceeded\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"distributorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestType\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.DistributorRequestType\"}]},{\"type\":\"error\",\"name\":\"RequestNotInStatus\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedStatus\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"},{\"name\":\"actualStatus\",\"type\":\"uint8\",\"internalType\":\"enumIDTAMessage.RequestStatus\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StaleNAV\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"requestCreatedAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"navUpdatedAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"fundIssuerTimeZoneOffsetSecs\",\"type\":\"int24\",\"internalType\":\"int24\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedDistributorNotAllowedForToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedDistributorNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedNotDistributor\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedNotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedStatus\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedStatus\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"actualStatus\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"UnexpectedTokenStatus\",\"inputs\":[{\"name\":\"fundAdminAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fundTokenId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actualStatus\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]",
}

// DtarmABI is the input ABI used to generate the binding from.
// Deprecated: Use DtarmMetaData.ABI instead.
var DtarmABI = DtarmMetaData.ABI

// Dtarm is an auto generated Go binding around an Ethereum contract.
type Dtarm struct {
	DtarmCaller     // Read-only binding to the contract
	DtarmTransactor // Write-only binding to the contract
	DtarmFilterer   // Log filterer for contract events
}

// DtarmCaller is an auto generated read-only Go binding around an Ethereum contract.
type DtarmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtarmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DtarmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtarmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DtarmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtarmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DtarmSession struct {
	Contract     *Dtarm            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DtarmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DtarmCallerSession struct {
	Contract *DtarmCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DtarmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DtarmTransactorSession struct {
	Contract     *DtarmTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DtarmRaw is an auto generated low-level Go binding around an Ethereum contract.
type DtarmRaw struct {
	Contract *Dtarm // Generic contract binding to access the raw methods on
}

// DtarmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DtarmCallerRaw struct {
	Contract *DtarmCaller // Generic read-only contract binding to access the raw methods on
}

// DtarmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DtarmTransactorRaw struct {
	Contract *DtarmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDtarm creates a new instance of Dtarm, bound to a specific deployed contract.
func NewDtarm(address common.Address, backend bind.ContractBackend) (*Dtarm, error) {
	contract, err := bindDtarm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dtarm{DtarmCaller: DtarmCaller{contract: contract}, DtarmTransactor: DtarmTransactor{contract: contract}, DtarmFilterer: DtarmFilterer{contract: contract}}, nil
}

// NewDtarmCaller creates a new read-only instance of Dtarm, bound to a specific deployed contract.
func NewDtarmCaller(address common.Address, caller bind.ContractCaller) (*DtarmCaller, error) {
	contract, err := bindDtarm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DtarmCaller{contract: contract}, nil
}

// NewDtarmTransactor creates a new write-only instance of Dtarm, bound to a specific deployed contract.
func NewDtarmTransactor(address common.Address, transactor bind.ContractTransactor) (*DtarmTransactor, error) {
	contract, err := bindDtarm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DtarmTransactor{contract: contract}, nil
}

// NewDtarmFilterer creates a new log filterer instance of Dtarm, bound to a specific deployed contract.
func NewDtarmFilterer(address common.Address, filterer bind.ContractFilterer) (*DtarmFilterer, error) {
	contract, err := bindDtarm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DtarmFilterer{contract: contract}, nil
}

// bindDtarm binds a generic wrapper to an already deployed contract.
func bindDtarm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DtarmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dtarm *DtarmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dtarm.Contract.DtarmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dtarm *DtarmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarm.Contract.DtarmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dtarm *DtarmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dtarm.Contract.DtarmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dtarm *DtarmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dtarm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dtarm *DtarmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dtarm *DtarmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dtarm.Contract.contract.Transact(opts, method, params...)
}

// GetCCIPGasLimit is a free data retrieval call binding the contract method 0xd83ce949.
//
// Solidity: function getCCIPGasLimit() view returns(uint256)
func (_Dtarm *DtarmCaller) GetCCIPGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dtarm.contract.Call(opts, &out, "getCCIPGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCCIPGasLimit is a free data retrieval call binding the contract method 0xd83ce949.
//
// Solidity: function getCCIPGasLimit() view returns(uint256)
func (_Dtarm *DtarmSession) GetCCIPGasLimit() (*big.Int, error) {
	return _Dtarm.Contract.GetCCIPGasLimit(&_Dtarm.CallOpts)
}

// GetCCIPGasLimit is a free data retrieval call binding the contract method 0xd83ce949.
//
// Solidity: function getCCIPGasLimit() view returns(uint256)
func (_Dtarm *DtarmCallerSession) GetCCIPGasLimit() (*big.Int, error) {
	return _Dtarm.Contract.GetCCIPGasLimit(&_Dtarm.CallOpts)
}

// GetDistributor is a free data retrieval call binding the contract method 0x993df8c6.
//
// Solidity: function getDistributor(address distributorAddr) view returns((address,bool))
func (_Dtarm *DtarmCaller) GetDistributor(opts *bind.CallOpts, distributorAddr common.Address) (IOpenFundDistributorRegistryDistributorData, error) {
	var out []interface{}
	err := _Dtarm.contract.Call(opts, &out, "getDistributor", distributorAddr)

	if err != nil {
		return *new(IOpenFundDistributorRegistryDistributorData), err
	}

	out0 := *abi.ConvertType(out[0], new(IOpenFundDistributorRegistryDistributorData)).(*IOpenFundDistributorRegistryDistributorData)

	return out0, err

}

// GetDistributor is a free data retrieval call binding the contract method 0x993df8c6.
//
// Solidity: function getDistributor(address distributorAddr) view returns((address,bool))
func (_Dtarm *DtarmSession) GetDistributor(distributorAddr common.Address) (IOpenFundDistributorRegistryDistributorData, error) {
	return _Dtarm.Contract.GetDistributor(&_Dtarm.CallOpts, distributorAddr)
}

// GetDistributor is a free data retrieval call binding the contract method 0x993df8c6.
//
// Solidity: function getDistributor(address distributorAddr) view returns((address,bool))
func (_Dtarm *DtarmCallerSession) GetDistributor(distributorAddr common.Address) (IOpenFundDistributorRegistryDistributorData, error) {
	return _Dtarm.Contract.GetDistributor(&_Dtarm.CallOpts, distributorAddr)
}

// GetDistributorRequest is a free data retrieval call binding the contract method 0x3cb723c5.
//
// Solidity: function getDistributorRequest(bytes32 requestId) view returns((uint256,uint256,bytes32,address,address,uint40,uint8,uint8))
func (_Dtarm *DtarmCaller) GetDistributorRequest(opts *bind.CallOpts, requestId [32]byte) (IDTADistributorDistributorRequest, error) {
	var out []interface{}
	err := _Dtarm.contract.Call(opts, &out, "getDistributorRequest", requestId)

	if err != nil {
		return *new(IDTADistributorDistributorRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IDTADistributorDistributorRequest)).(*IDTADistributorDistributorRequest)

	return out0, err

}

// GetDistributorRequest is a free data retrieval call binding the contract method 0x3cb723c5.
//
// Solidity: function getDistributorRequest(bytes32 requestId) view returns((uint256,uint256,bytes32,address,address,uint40,uint8,uint8))
func (_Dtarm *DtarmSession) GetDistributorRequest(requestId [32]byte) (IDTADistributorDistributorRequest, error) {
	return _Dtarm.Contract.GetDistributorRequest(&_Dtarm.CallOpts, requestId)
}

// GetDistributorRequest is a free data retrieval call binding the contract method 0x3cb723c5.
//
// Solidity: function getDistributorRequest(bytes32 requestId) view returns((uint256,uint256,bytes32,address,address,uint40,uint8,uint8))
func (_Dtarm *DtarmCallerSession) GetDistributorRequest(requestId [32]byte) (IDTADistributorDistributorRequest, error) {
	return _Dtarm.Contract.GetDistributorRequest(&_Dtarm.CallOpts, requestId)
}

// GetDistributors is a free data retrieval call binding the contract method 0xdafe1ef7.
//
// Solidity: function getDistributors(uint256 offset) view returns(address[] distributorAddrs)
func (_Dtarm *DtarmCaller) GetDistributors(opts *bind.CallOpts, offset *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Dtarm.contract.Call(opts, &out, "getDistributors", offset)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDistributors is a free data retrieval call binding the contract method 0xdafe1ef7.
//
// Solidity: function getDistributors(uint256 offset) view returns(address[] distributorAddrs)
func (_Dtarm *DtarmSession) GetDistributors(offset *big.Int) ([]common.Address, error) {
	return _Dtarm.Contract.GetDistributors(&_Dtarm.CallOpts, offset)
}

// GetDistributors is a free data retrieval call binding the contract method 0xdafe1ef7.
//
// Solidity: function getDistributors(uint256 offset) view returns(address[] distributorAddrs)
func (_Dtarm *DtarmCallerSession) GetDistributors(offset *big.Int) ([]common.Address, error) {
	return _Dtarm.Contract.GetDistributors(&_Dtarm.CallOpts, offset)
}

// GetFundAdmins is a free data retrieval call binding the contract method 0x26dd451e.
//
// Solidity: function getFundAdmins(uint256 offset) view returns(address[] fundAdminAddrs)
func (_Dtarm *DtarmCaller) GetFundAdmins(opts *bind.CallOpts, offset *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Dtarm.contract.Call(opts, &out, "getFundAdmins", offset)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFundAdmins is a free data retrieval call binding the contract method 0x26dd451e.
//
// Solidity: function getFundAdmins(uint256 offset) view returns(address[] fundAdminAddrs)
func (_Dtarm *DtarmSession) GetFundAdmins(offset *big.Int) ([]common.Address, error) {
	return _Dtarm.Contract.GetFundAdmins(&_Dtarm.CallOpts, offset)
}

// GetFundAdmins is a free data retrieval call binding the contract method 0x26dd451e.
//
// Solidity: function getFundAdmins(uint256 offset) view returns(address[] fundAdminAddrs)
func (_Dtarm *DtarmCallerSession) GetFundAdmins(offset *big.Int) ([]common.Address, error) {
	return _Dtarm.Contract.GetFundAdmins(&_Dtarm.CallOpts, offset)
}

// GetFundToken is a free data retrieval call binding the contract method 0x5f364f07.
//
// Solidity: function getFundToken(address fundAdminAddr, bytes32 fundTokenId) view returns(bool enabled, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)))
func (_Dtarm *DtarmCaller) GetFundToken(opts *bind.CallOpts, fundAdminAddr common.Address, fundTokenId [32]byte) (bool, IFundTokenRegistryFundTokenData, error) {
	var out []interface{}
	err := _Dtarm.contract.Call(opts, &out, "getFundToken", fundAdminAddr, fundTokenId)

	if err != nil {
		return *new(bool), *new(IFundTokenRegistryFundTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(IFundTokenRegistryFundTokenData)).(*IFundTokenRegistryFundTokenData)

	return out0, out1, err

}

// GetFundToken is a free data retrieval call binding the contract method 0x5f364f07.
//
// Solidity: function getFundToken(address fundAdminAddr, bytes32 fundTokenId) view returns(bool enabled, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)))
func (_Dtarm *DtarmSession) GetFundToken(fundAdminAddr common.Address, fundTokenId [32]byte) (bool, IFundTokenRegistryFundTokenData, error) {
	return _Dtarm.Contract.GetFundToken(&_Dtarm.CallOpts, fundAdminAddr, fundTokenId)
}

// GetFundToken is a free data retrieval call binding the contract method 0x5f364f07.
//
// Solidity: function getFundToken(address fundAdminAddr, bytes32 fundTokenId) view returns(bool enabled, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)))
func (_Dtarm *DtarmCallerSession) GetFundToken(fundAdminAddr common.Address, fundTokenId [32]byte) (bool, IFundTokenRegistryFundTokenData, error) {
	return _Dtarm.Contract.GetFundToken(&_Dtarm.CallOpts, fundAdminAddr, fundTokenId)
}

// GetFundTokens is a free data retrieval call binding the contract method 0xe81e32ff.
//
// Solidity: function getFundTokens(address fundAdminAddr, uint256 offset) view returns(bytes32[])
func (_Dtarm *DtarmCaller) GetFundTokens(opts *bind.CallOpts, fundAdminAddr common.Address, offset *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Dtarm.contract.Call(opts, &out, "getFundTokens", fundAdminAddr, offset)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetFundTokens is a free data retrieval call binding the contract method 0xe81e32ff.
//
// Solidity: function getFundTokens(address fundAdminAddr, uint256 offset) view returns(bytes32[])
func (_Dtarm *DtarmSession) GetFundTokens(fundAdminAddr common.Address, offset *big.Int) ([][32]byte, error) {
	return _Dtarm.Contract.GetFundTokens(&_Dtarm.CallOpts, fundAdminAddr, offset)
}

// GetFundTokens is a free data retrieval call binding the contract method 0xe81e32ff.
//
// Solidity: function getFundTokens(address fundAdminAddr, uint256 offset) view returns(bytes32[])
func (_Dtarm *DtarmCallerSession) GetFundTokens(fundAdminAddr common.Address, offset *big.Int) ([][32]byte, error) {
	return _Dtarm.Contract.GetFundTokens(&_Dtarm.CallOpts, fundAdminAddr, offset)
}

// GetLocalChainSelector is a free data retrieval call binding the contract method 0xeaa83ddd.
//
// Solidity: function getLocalChainSelector() view returns(uint64 localChainSelector)
func (_Dtarm *DtarmCaller) GetLocalChainSelector(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Dtarm.contract.Call(opts, &out, "getLocalChainSelector")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLocalChainSelector is a free data retrieval call binding the contract method 0xeaa83ddd.
//
// Solidity: function getLocalChainSelector() view returns(uint64 localChainSelector)
func (_Dtarm *DtarmSession) GetLocalChainSelector() (uint64, error) {
	return _Dtarm.Contract.GetLocalChainSelector(&_Dtarm.CallOpts)
}

// GetLocalChainSelector is a free data retrieval call binding the contract method 0xeaa83ddd.
//
// Solidity: function getLocalChainSelector() view returns(uint64 localChainSelector)
func (_Dtarm *DtarmCallerSession) GetLocalChainSelector() (uint64, error) {
	return _Dtarm.Contract.GetLocalChainSelector(&_Dtarm.CallOpts)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtarm *DtarmCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dtarm.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtarm *DtarmSession) GetRouter() (common.Address, error) {
	return _Dtarm.Contract.GetRouter(&_Dtarm.CallOpts)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_Dtarm *DtarmCallerSession) GetRouter() (common.Address, error) {
	return _Dtarm.Contract.GetRouter(&_Dtarm.CallOpts)
}

// GetTokenRequestsByDate is a free data retrieval call binding the contract method 0xc39c2fec.
//
// Solidity: function getTokenRequestsByDate(address fundAdminAddr, bytes32 fundTokenId, uint40 date, uint256 offset) view returns(bytes32[])
func (_Dtarm *DtarmCaller) GetTokenRequestsByDate(opts *bind.CallOpts, fundAdminAddr common.Address, fundTokenId [32]byte, date *big.Int, offset *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Dtarm.contract.Call(opts, &out, "getTokenRequestsByDate", fundAdminAddr, fundTokenId, date, offset)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetTokenRequestsByDate is a free data retrieval call binding the contract method 0xc39c2fec.
//
// Solidity: function getTokenRequestsByDate(address fundAdminAddr, bytes32 fundTokenId, uint40 date, uint256 offset) view returns(bytes32[])
func (_Dtarm *DtarmSession) GetTokenRequestsByDate(fundAdminAddr common.Address, fundTokenId [32]byte, date *big.Int, offset *big.Int) ([][32]byte, error) {
	return _Dtarm.Contract.GetTokenRequestsByDate(&_Dtarm.CallOpts, fundAdminAddr, fundTokenId, date, offset)
}

// GetTokenRequestsByDate is a free data retrieval call binding the contract method 0xc39c2fec.
//
// Solidity: function getTokenRequestsByDate(address fundAdminAddr, bytes32 fundTokenId, uint40 date, uint256 offset) view returns(bytes32[])
func (_Dtarm *DtarmCallerSession) GetTokenRequestsByDate(fundAdminAddr common.Address, fundTokenId [32]byte, date *big.Int, offset *big.Int) ([][32]byte, error) {
	return _Dtarm.Contract.GetTokenRequestsByDate(&_Dtarm.CallOpts, fundAdminAddr, fundTokenId, date, offset)
}

// IsDistributorAllowedForToken is a free data retrieval call binding the contract method 0xb4253193.
//
// Solidity: function isDistributorAllowedForToken(address fundAdminAddr, bytes32 fundTokenId, address distributorAddr) view returns(bool isAllowed, bool isTokenEnabled)
func (_Dtarm *DtarmCaller) IsDistributorAllowedForToken(opts *bind.CallOpts, fundAdminAddr common.Address, fundTokenId [32]byte, distributorAddr common.Address) (struct {
	IsAllowed      bool
	IsTokenEnabled bool
}, error) {
	var out []interface{}
	err := _Dtarm.contract.Call(opts, &out, "isDistributorAllowedForToken", fundAdminAddr, fundTokenId, distributorAddr)

	outstruct := new(struct {
		IsAllowed      bool
		IsTokenEnabled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsAllowed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsTokenEnabled = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// IsDistributorAllowedForToken is a free data retrieval call binding the contract method 0xb4253193.
//
// Solidity: function isDistributorAllowedForToken(address fundAdminAddr, bytes32 fundTokenId, address distributorAddr) view returns(bool isAllowed, bool isTokenEnabled)
func (_Dtarm *DtarmSession) IsDistributorAllowedForToken(fundAdminAddr common.Address, fundTokenId [32]byte, distributorAddr common.Address) (struct {
	IsAllowed      bool
	IsTokenEnabled bool
}, error) {
	return _Dtarm.Contract.IsDistributorAllowedForToken(&_Dtarm.CallOpts, fundAdminAddr, fundTokenId, distributorAddr)
}

// IsDistributorAllowedForToken is a free data retrieval call binding the contract method 0xb4253193.
//
// Solidity: function isDistributorAllowedForToken(address fundAdminAddr, bytes32 fundTokenId, address distributorAddr) view returns(bool isAllowed, bool isTokenEnabled)
func (_Dtarm *DtarmCallerSession) IsDistributorAllowedForToken(fundAdminAddr common.Address, fundTokenId [32]byte, distributorAddr common.Address) (struct {
	IsAllowed      bool
	IsTokenEnabled bool
}, error) {
	return _Dtarm.Contract.IsDistributorAllowedForToken(&_Dtarm.CallOpts, fundAdminAddr, fundTokenId, distributorAddr)
}

// IsFundAdminRegistered is a free data retrieval call binding the contract method 0xf3e28ee4.
//
// Solidity: function isFundAdminRegistered(address fundAdminAddr) view returns(bool)
func (_Dtarm *DtarmCaller) IsFundAdminRegistered(opts *bind.CallOpts, fundAdminAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Dtarm.contract.Call(opts, &out, "isFundAdminRegistered", fundAdminAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFundAdminRegistered is a free data retrieval call binding the contract method 0xf3e28ee4.
//
// Solidity: function isFundAdminRegistered(address fundAdminAddr) view returns(bool)
func (_Dtarm *DtarmSession) IsFundAdminRegistered(fundAdminAddr common.Address) (bool, error) {
	return _Dtarm.Contract.IsFundAdminRegistered(&_Dtarm.CallOpts, fundAdminAddr)
}

// IsFundAdminRegistered is a free data retrieval call binding the contract method 0xf3e28ee4.
//
// Solidity: function isFundAdminRegistered(address fundAdminAddr) view returns(bool)
func (_Dtarm *DtarmCallerSession) IsFundAdminRegistered(fundAdminAddr common.Address) (bool, error) {
	return _Dtarm.Contract.IsFundAdminRegistered(&_Dtarm.CallOpts, fundAdminAddr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtarm *DtarmCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dtarm.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtarm *DtarmSession) Owner() (common.Address, error) {
	return _Dtarm.Contract.Owner(&_Dtarm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dtarm *DtarmCallerSession) Owner() (common.Address, error) {
	return _Dtarm.Contract.Owner(&_Dtarm.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtarm *DtarmCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Dtarm.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtarm *DtarmSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dtarm.Contract.SupportsInterface(&_Dtarm.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dtarm *DtarmCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dtarm.Contract.SupportsInterface(&_Dtarm.CallOpts, interfaceId)
}

// AllowDistributorForToken is a paid mutator transaction binding the contract method 0xb1fbb355.
//
// Solidity: function allowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarm *DtarmTransactor) AllowDistributorForToken(opts *bind.TransactOpts, fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "allowDistributorForToken", fundTokenId, distributorAddr)
}

// AllowDistributorForToken is a paid mutator transaction binding the contract method 0xb1fbb355.
//
// Solidity: function allowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarm *DtarmSession) AllowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.AllowDistributorForToken(&_Dtarm.TransactOpts, fundTokenId, distributorAddr)
}

// AllowDistributorForToken is a paid mutator transaction binding the contract method 0xb1fbb355.
//
// Solidity: function allowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarm *DtarmTransactorSession) AllowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.AllowDistributorForToken(&_Dtarm.TransactOpts, fundTokenId, distributorAddr)
}

// CancelDistributorRequest is a paid mutator transaction binding the contract method 0xac17bfd8.
//
// Solidity: function cancelDistributorRequest(bytes32 requestId) returns()
func (_Dtarm *DtarmTransactor) CancelDistributorRequest(opts *bind.TransactOpts, requestId [32]byte) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "cancelDistributorRequest", requestId)
}

// CancelDistributorRequest is a paid mutator transaction binding the contract method 0xac17bfd8.
//
// Solidity: function cancelDistributorRequest(bytes32 requestId) returns()
func (_Dtarm *DtarmSession) CancelDistributorRequest(requestId [32]byte) (*types.Transaction, error) {
	return _Dtarm.Contract.CancelDistributorRequest(&_Dtarm.TransactOpts, requestId)
}

// CancelDistributorRequest is a paid mutator transaction binding the contract method 0xac17bfd8.
//
// Solidity: function cancelDistributorRequest(bytes32 requestId) returns()
func (_Dtarm *DtarmTransactorSession) CancelDistributorRequest(requestId [32]byte) (*types.Transaction, error) {
	return _Dtarm.Contract.CancelDistributorRequest(&_Dtarm.TransactOpts, requestId)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtarm *DtarmTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtarm *DtarmSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtarm.Contract.CcipReceive(&_Dtarm.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_Dtarm *DtarmTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _Dtarm.Contract.CcipReceive(&_Dtarm.TransactOpts, message)
}

// CompleteDistributorRequest is a paid mutator transaction binding the contract method 0xe9246fc2.
//
// Solidity: function completeDistributorRequest(address sender, uint64 sourceChainSelector, (bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtarm *DtarmTransactor) CompleteDistributorRequest(opts *bind.TransactOpts, sender common.Address, sourceChainSelector uint64, dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "completeDistributorRequest", sender, sourceChainSelector, dtaMessage)
}

// CompleteDistributorRequest is a paid mutator transaction binding the contract method 0xe9246fc2.
//
// Solidity: function completeDistributorRequest(address sender, uint64 sourceChainSelector, (bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtarm *DtarmSession) CompleteDistributorRequest(sender common.Address, sourceChainSelector uint64, dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtarm.Contract.CompleteDistributorRequest(&_Dtarm.TransactOpts, sender, sourceChainSelector, dtaMessage)
}

// CompleteDistributorRequest is a paid mutator transaction binding the contract method 0xe9246fc2.
//
// Solidity: function completeDistributorRequest(address sender, uint64 sourceChainSelector, (bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtarm *DtarmTransactorSession) CompleteDistributorRequest(sender common.Address, sourceChainSelector uint64, dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtarm.Contract.CompleteDistributorRequest(&_Dtarm.TransactOpts, sender, sourceChainSelector, dtaMessage)
}

// DirectCompleteDistributorRequest is a paid mutator transaction binding the contract method 0x4286d9db.
//
// Solidity: function directCompleteDistributorRequest((bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtarm *DtarmTransactor) DirectCompleteDistributorRequest(opts *bind.TransactOpts, dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "directCompleteDistributorRequest", dtaMessage)
}

// DirectCompleteDistributorRequest is a paid mutator transaction binding the contract method 0x4286d9db.
//
// Solidity: function directCompleteDistributorRequest((bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtarm *DtarmSession) DirectCompleteDistributorRequest(dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtarm.Contract.DirectCompleteDistributorRequest(&_Dtarm.TransactOpts, dtaMessage)
}

// DirectCompleteDistributorRequest is a paid mutator transaction binding the contract method 0x4286d9db.
//
// Solidity: function directCompleteDistributorRequest((bytes32,uint8,bytes) dtaMessage) returns()
func (_Dtarm *DtarmTransactorSession) DirectCompleteDistributorRequest(dtaMessage IDTAMessageDtaRequestStatusMessage) (*types.Transaction, error) {
	return _Dtarm.Contract.DirectCompleteDistributorRequest(&_Dtarm.TransactOpts, dtaMessage)
}

// DisableFundToken is a paid mutator transaction binding the contract method 0x8198b420.
//
// Solidity: function disableFundToken(bytes32 fundTokenId) returns()
func (_Dtarm *DtarmTransactor) DisableFundToken(opts *bind.TransactOpts, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "disableFundToken", fundTokenId)
}

// DisableFundToken is a paid mutator transaction binding the contract method 0x8198b420.
//
// Solidity: function disableFundToken(bytes32 fundTokenId) returns()
func (_Dtarm *DtarmSession) DisableFundToken(fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarm.Contract.DisableFundToken(&_Dtarm.TransactOpts, fundTokenId)
}

// DisableFundToken is a paid mutator transaction binding the contract method 0x8198b420.
//
// Solidity: function disableFundToken(bytes32 fundTokenId) returns()
func (_Dtarm *DtarmTransactorSession) DisableFundToken(fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarm.Contract.DisableFundToken(&_Dtarm.TransactOpts, fundTokenId)
}

// DisallowDistributorForToken is a paid mutator transaction binding the contract method 0xf9848667.
//
// Solidity: function disallowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarm *DtarmTransactor) DisallowDistributorForToken(opts *bind.TransactOpts, fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "disallowDistributorForToken", fundTokenId, distributorAddr)
}

// DisallowDistributorForToken is a paid mutator transaction binding the contract method 0xf9848667.
//
// Solidity: function disallowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarm *DtarmSession) DisallowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.DisallowDistributorForToken(&_Dtarm.TransactOpts, fundTokenId, distributorAddr)
}

// DisallowDistributorForToken is a paid mutator transaction binding the contract method 0xf9848667.
//
// Solidity: function disallowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarm *DtarmTransactorSession) DisallowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.DisallowDistributorForToken(&_Dtarm.TransactOpts, fundTokenId, distributorAddr)
}

// EnableFundToken is a paid mutator transaction binding the contract method 0x31bb8270.
//
// Solidity: function enableFundToken(bytes32 fundTokenId) returns()
func (_Dtarm *DtarmTransactor) EnableFundToken(opts *bind.TransactOpts, fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "enableFundToken", fundTokenId)
}

// EnableFundToken is a paid mutator transaction binding the contract method 0x31bb8270.
//
// Solidity: function enableFundToken(bytes32 fundTokenId) returns()
func (_Dtarm *DtarmSession) EnableFundToken(fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarm.Contract.EnableFundToken(&_Dtarm.TransactOpts, fundTokenId)
}

// EnableFundToken is a paid mutator transaction binding the contract method 0x31bb8270.
//
// Solidity: function enableFundToken(bytes32 fundTokenId) returns()
func (_Dtarm *DtarmTransactorSession) EnableFundToken(fundTokenId [32]byte) (*types.Transaction, error) {
	return _Dtarm.Contract.EnableFundToken(&_Dtarm.TransactOpts, fundTokenId)
}

// ForceAllowDistributorForToken is a paid mutator transaction binding the contract method 0xd612b214.
//
// Solidity: function forceAllowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarm *DtarmTransactor) ForceAllowDistributorForToken(opts *bind.TransactOpts, fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "forceAllowDistributorForToken", fundTokenId, distributorAddr)
}

// ForceAllowDistributorForToken is a paid mutator transaction binding the contract method 0xd612b214.
//
// Solidity: function forceAllowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarm *DtarmSession) ForceAllowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.ForceAllowDistributorForToken(&_Dtarm.TransactOpts, fundTokenId, distributorAddr)
}

// ForceAllowDistributorForToken is a paid mutator transaction binding the contract method 0xd612b214.
//
// Solidity: function forceAllowDistributorForToken(bytes32 fundTokenId, address distributorAddr) returns()
func (_Dtarm *DtarmTransactorSession) ForceAllowDistributorForToken(fundTokenId [32]byte, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.ForceAllowDistributorForToken(&_Dtarm.TransactOpts, fundTokenId, distributorAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 localChainSelector, address feeManager) returns()
func (_Dtarm *DtarmTransactor) Initialize(opts *bind.TransactOpts, localChainSelector uint64, feeManager common.Address) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "initialize", localChainSelector, feeManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 localChainSelector, address feeManager) returns()
func (_Dtarm *DtarmSession) Initialize(localChainSelector uint64, feeManager common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.Initialize(&_Dtarm.TransactOpts, localChainSelector, feeManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 localChainSelector, address feeManager) returns()
func (_Dtarm *DtarmTransactorSession) Initialize(localChainSelector uint64, feeManager common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.Initialize(&_Dtarm.TransactOpts, localChainSelector, feeManager)
}

// ProcessDistributorRequest is a paid mutator transaction binding the contract method 0x53a103bc.
//
// Solidity: function processDistributorRequest(bytes32 requestId) returns()
func (_Dtarm *DtarmTransactor) ProcessDistributorRequest(opts *bind.TransactOpts, requestId [32]byte) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "processDistributorRequest", requestId)
}

// ProcessDistributorRequest is a paid mutator transaction binding the contract method 0x53a103bc.
//
// Solidity: function processDistributorRequest(bytes32 requestId) returns()
func (_Dtarm *DtarmSession) ProcessDistributorRequest(requestId [32]byte) (*types.Transaction, error) {
	return _Dtarm.Contract.ProcessDistributorRequest(&_Dtarm.TransactOpts, requestId)
}

// ProcessDistributorRequest is a paid mutator transaction binding the contract method 0x53a103bc.
//
// Solidity: function processDistributorRequest(bytes32 requestId) returns()
func (_Dtarm *DtarmTransactorSession) ProcessDistributorRequest(requestId [32]byte) (*types.Transaction, error) {
	return _Dtarm.Contract.ProcessDistributorRequest(&_Dtarm.TransactOpts, requestId)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address recipient) returns()
func (_Dtarm *DtarmTransactor) RecoverFunds(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "recoverFunds", recipient)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address recipient) returns()
func (_Dtarm *DtarmSession) RecoverFunds(recipient common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.RecoverFunds(&_Dtarm.TransactOpts, recipient)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address recipient) returns()
func (_Dtarm *DtarmTransactorSession) RecoverFunds(recipient common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.RecoverFunds(&_Dtarm.TransactOpts, recipient)
}

// RegisterDistributor is a paid mutator transaction binding the contract method 0x31ab0518.
//
// Solidity: function registerDistributor(address distributorWalletAddr) returns()
func (_Dtarm *DtarmTransactor) RegisterDistributor(opts *bind.TransactOpts, distributorWalletAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "registerDistributor", distributorWalletAddr)
}

// RegisterDistributor is a paid mutator transaction binding the contract method 0x31ab0518.
//
// Solidity: function registerDistributor(address distributorWalletAddr) returns()
func (_Dtarm *DtarmSession) RegisterDistributor(distributorWalletAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.RegisterDistributor(&_Dtarm.TransactOpts, distributorWalletAddr)
}

// RegisterDistributor is a paid mutator transaction binding the contract method 0x31ab0518.
//
// Solidity: function registerDistributor(address distributorWalletAddr) returns()
func (_Dtarm *DtarmTransactorSession) RegisterDistributor(distributorWalletAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.RegisterDistributor(&_Dtarm.TransactOpts, distributorWalletAddr)
}

// RegisterFundAdmin is a paid mutator transaction binding the contract method 0xfcd961f9.
//
// Solidity: function registerFundAdmin() returns()
func (_Dtarm *DtarmTransactor) RegisterFundAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "registerFundAdmin")
}

// RegisterFundAdmin is a paid mutator transaction binding the contract method 0xfcd961f9.
//
// Solidity: function registerFundAdmin() returns()
func (_Dtarm *DtarmSession) RegisterFundAdmin() (*types.Transaction, error) {
	return _Dtarm.Contract.RegisterFundAdmin(&_Dtarm.TransactOpts)
}

// RegisterFundAdmin is a paid mutator transaction binding the contract method 0xfcd961f9.
//
// Solidity: function registerFundAdmin() returns()
func (_Dtarm *DtarmTransactorSession) RegisterFundAdmin() (*types.Transaction, error) {
	return _Dtarm.Contract.RegisterFundAdmin(&_Dtarm.TransactOpts)
}

// RegisterFundToken is a paid mutator transaction binding the contract method 0x309f21e9.
//
// Solidity: function registerFundToken(bytes32 fundTokenId, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)) tokenData) returns()
func (_Dtarm *DtarmTransactor) RegisterFundToken(opts *bind.TransactOpts, fundTokenId [32]byte, tokenData IFundTokenRegistryFundTokenData) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "registerFundToken", fundTokenId, tokenData)
}

// RegisterFundToken is a paid mutator transaction binding the contract method 0x309f21e9.
//
// Solidity: function registerFundToken(bytes32 fundTokenId, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)) tokenData) returns()
func (_Dtarm *DtarmSession) RegisterFundToken(fundTokenId [32]byte, tokenData IFundTokenRegistryFundTokenData) (*types.Transaction, error) {
	return _Dtarm.Contract.RegisterFundToken(&_Dtarm.TransactOpts, fundTokenId, tokenData)
}

// RegisterFundToken is a paid mutator transaction binding the contract method 0x309f21e9.
//
// Solidity: function registerFundToken(bytes32 fundTokenId, (address,uint8,uint8,uint8,uint8,uint8,uint8,address,uint64,address,int24,uint24,(uint8,address,address)) tokenData) returns()
func (_Dtarm *DtarmTransactorSession) RegisterFundToken(fundTokenId [32]byte, tokenData IFundTokenRegistryFundTokenData) (*types.Transaction, error) {
	return _Dtarm.Contract.RegisterFundToken(&_Dtarm.TransactOpts, fundTokenId, tokenData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtarm *DtarmTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtarm *DtarmSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dtarm.Contract.RenounceOwnership(&_Dtarm.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtarm *DtarmTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dtarm.Contract.RenounceOwnership(&_Dtarm.TransactOpts)
}

// RequestRedemption is a paid mutator transaction binding the contract method 0x56226014.
//
// Solidity: function requestRedemption(address fundAdminAddr, bytes32 fundTokenId, uint256 shares) returns(bytes32 requestId)
func (_Dtarm *DtarmTransactor) RequestRedemption(opts *bind.TransactOpts, fundAdminAddr common.Address, fundTokenId [32]byte, shares *big.Int) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "requestRedemption", fundAdminAddr, fundTokenId, shares)
}

// RequestRedemption is a paid mutator transaction binding the contract method 0x56226014.
//
// Solidity: function requestRedemption(address fundAdminAddr, bytes32 fundTokenId, uint256 shares) returns(bytes32 requestId)
func (_Dtarm *DtarmSession) RequestRedemption(fundAdminAddr common.Address, fundTokenId [32]byte, shares *big.Int) (*types.Transaction, error) {
	return _Dtarm.Contract.RequestRedemption(&_Dtarm.TransactOpts, fundAdminAddr, fundTokenId, shares)
}

// RequestRedemption is a paid mutator transaction binding the contract method 0x56226014.
//
// Solidity: function requestRedemption(address fundAdminAddr, bytes32 fundTokenId, uint256 shares) returns(bytes32 requestId)
func (_Dtarm *DtarmTransactorSession) RequestRedemption(fundAdminAddr common.Address, fundTokenId [32]byte, shares *big.Int) (*types.Transaction, error) {
	return _Dtarm.Contract.RequestRedemption(&_Dtarm.TransactOpts, fundAdminAddr, fundTokenId, shares)
}

// RequestSubscription is a paid mutator transaction binding the contract method 0x6aad53da.
//
// Solidity: function requestSubscription(address fundAdminAddr, bytes32 fundTokenId, uint256 amount) returns(bytes32 requestId)
func (_Dtarm *DtarmTransactor) RequestSubscription(opts *bind.TransactOpts, fundAdminAddr common.Address, fundTokenId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "requestSubscription", fundAdminAddr, fundTokenId, amount)
}

// RequestSubscription is a paid mutator transaction binding the contract method 0x6aad53da.
//
// Solidity: function requestSubscription(address fundAdminAddr, bytes32 fundTokenId, uint256 amount) returns(bytes32 requestId)
func (_Dtarm *DtarmSession) RequestSubscription(fundAdminAddr common.Address, fundTokenId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Dtarm.Contract.RequestSubscription(&_Dtarm.TransactOpts, fundAdminAddr, fundTokenId, amount)
}

// RequestSubscription is a paid mutator transaction binding the contract method 0x6aad53da.
//
// Solidity: function requestSubscription(address fundAdminAddr, bytes32 fundTokenId, uint256 amount) returns(bytes32 requestId)
func (_Dtarm *DtarmTransactorSession) RequestSubscription(fundAdminAddr common.Address, fundTokenId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Dtarm.Contract.RequestSubscription(&_Dtarm.TransactOpts, fundAdminAddr, fundTokenId, amount)
}

// SetCCIPGasLimit is a paid mutator transaction binding the contract method 0xca0880cf.
//
// Solidity: function setCCIPGasLimit(uint256 gasLimit) returns()
func (_Dtarm *DtarmTransactor) SetCCIPGasLimit(opts *bind.TransactOpts, gasLimit *big.Int) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "setCCIPGasLimit", gasLimit)
}

// SetCCIPGasLimit is a paid mutator transaction binding the contract method 0xca0880cf.
//
// Solidity: function setCCIPGasLimit(uint256 gasLimit) returns()
func (_Dtarm *DtarmSession) SetCCIPGasLimit(gasLimit *big.Int) (*types.Transaction, error) {
	return _Dtarm.Contract.SetCCIPGasLimit(&_Dtarm.TransactOpts, gasLimit)
}

// SetCCIPGasLimit is a paid mutator transaction binding the contract method 0xca0880cf.
//
// Solidity: function setCCIPGasLimit(uint256 gasLimit) returns()
func (_Dtarm *DtarmTransactorSession) SetCCIPGasLimit(gasLimit *big.Int) (*types.Transaction, error) {
	return _Dtarm.Contract.SetCCIPGasLimit(&_Dtarm.TransactOpts, gasLimit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtarm *DtarmTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtarm *DtarmSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.TransferOwnership(&_Dtarm.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtarm *DtarmTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.TransferOwnership(&_Dtarm.TransactOpts, newOwner)
}

// VerifyDistributorWallet is a paid mutator transaction binding the contract method 0x1024cd9c.
//
// Solidity: function verifyDistributorWallet(address distributorAddr) returns()
func (_Dtarm *DtarmTransactor) VerifyDistributorWallet(opts *bind.TransactOpts, distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "verifyDistributorWallet", distributorAddr)
}

// VerifyDistributorWallet is a paid mutator transaction binding the contract method 0x1024cd9c.
//
// Solidity: function verifyDistributorWallet(address distributorAddr) returns()
func (_Dtarm *DtarmSession) VerifyDistributorWallet(distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.VerifyDistributorWallet(&_Dtarm.TransactOpts, distributorAddr)
}

// VerifyDistributorWallet is a paid mutator transaction binding the contract method 0x1024cd9c.
//
// Solidity: function verifyDistributorWallet(address distributorAddr) returns()
func (_Dtarm *DtarmTransactorSession) VerifyDistributorWallet(distributorAddr common.Address) (*types.Transaction, error) {
	return _Dtarm.Contract.VerifyDistributorWallet(&_Dtarm.TransactOpts, distributorAddr)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address token, address recipient, uint256 amount) returns()
func (_Dtarm *DtarmTransactor) WithdrawTokens(opts *bind.TransactOpts, token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dtarm.contract.Transact(opts, "withdrawTokens", token, recipient, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address token, address recipient, uint256 amount) returns()
func (_Dtarm *DtarmSession) WithdrawTokens(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dtarm.Contract.WithdrawTokens(&_Dtarm.TransactOpts, token, recipient, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address token, address recipient, uint256 amount) returns()
func (_Dtarm *DtarmTransactorSession) WithdrawTokens(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dtarm.Contract.WithdrawTokens(&_Dtarm.TransactOpts, token, recipient, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtarm *DtarmTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtarm.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtarm *DtarmSession) Receive() (*types.Transaction, error) {
	return _Dtarm.Contract.Receive(&_Dtarm.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dtarm *DtarmTransactorSession) Receive() (*types.Transaction, error) {
	return _Dtarm.Contract.Receive(&_Dtarm.TransactOpts)
}

// DtarmDistributorRegisteredIterator is returned from FilterDistributorRegistered and is used to iterate over the raw logs and unpacked data for DistributorRegistered events raised by the Dtarm contract.
type DtarmDistributorRegisteredIterator struct {
	Event *DtarmDistributorRegistered // Event containing the contract specifics and raw log

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
func (it *DtarmDistributorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmDistributorRegistered)
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
		it.Event = new(DtarmDistributorRegistered)
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
func (it *DtarmDistributorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmDistributorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmDistributorRegistered represents a DistributorRegistered event raised by the Dtarm contract.
type DtarmDistributorRegistered struct {
	DistributorAddr common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributorRegistered is a free log retrieval operation binding the contract event 0x3b2c2671a4556fca8f080b2c4469c9041a5019e2d24ddc37cfa48c95f542ac55.
//
// Solidity: event DistributorRegistered(address distributorAddr)
func (_Dtarm *DtarmFilterer) FilterDistributorRegistered(opts *bind.FilterOpts) (*DtarmDistributorRegisteredIterator, error) {

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "DistributorRegistered")
	if err != nil {
		return nil, err
	}
	return &DtarmDistributorRegisteredIterator{contract: _Dtarm.contract, event: "DistributorRegistered", logs: logs, sub: sub}, nil
}

// WatchDistributorRegistered is a free log subscription operation binding the contract event 0x3b2c2671a4556fca8f080b2c4469c9041a5019e2d24ddc37cfa48c95f542ac55.
//
// Solidity: event DistributorRegistered(address distributorAddr)
func (_Dtarm *DtarmFilterer) WatchDistributorRegistered(opts *bind.WatchOpts, sink chan<- *DtarmDistributorRegistered) (event.Subscription, error) {

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "DistributorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmDistributorRegistered)
				if err := _Dtarm.contract.UnpackLog(event, "DistributorRegistered", log); err != nil {
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

// ParseDistributorRegistered is a log parse operation binding the contract event 0x3b2c2671a4556fca8f080b2c4469c9041a5019e2d24ddc37cfa48c95f542ac55.
//
// Solidity: event DistributorRegistered(address distributorAddr)
func (_Dtarm *DtarmFilterer) ParseDistributorRegistered(log types.Log) (*DtarmDistributorRegistered, error) {
	event := new(DtarmDistributorRegistered)
	if err := _Dtarm.contract.UnpackLog(event, "DistributorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarmDistributorRequestCanceledIterator is returned from FilterDistributorRequestCanceled and is used to iterate over the raw logs and unpacked data for DistributorRequestCanceled events raised by the Dtarm contract.
type DtarmDistributorRequestCanceledIterator struct {
	Event *DtarmDistributorRequestCanceled // Event containing the contract specifics and raw log

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
func (it *DtarmDistributorRequestCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmDistributorRequestCanceled)
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
		it.Event = new(DtarmDistributorRequestCanceled)
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
func (it *DtarmDistributorRequestCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmDistributorRequestCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmDistributorRequestCanceled represents a DistributorRequestCanceled event raised by the Dtarm contract.
type DtarmDistributorRequestCanceled struct {
	FundAdminAddr   common.Address
	FundTokenId     [32]byte
	DistributorAddr common.Address
	RequestId       [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributorRequestCanceled is a free log retrieval operation binding the contract event 0xa7a5bac33f20d98f2f5d14a6c646812e2adc85c780523174e9634e49f89920cb.
//
// Solidity: event DistributorRequestCanceled(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId)
func (_Dtarm *DtarmFilterer) FilterDistributorRequestCanceled(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtarmDistributorRequestCanceledIterator, error) {

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

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "DistributorRequestCanceled", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtarmDistributorRequestCanceledIterator{contract: _Dtarm.contract, event: "DistributorRequestCanceled", logs: logs, sub: sub}, nil
}

// WatchDistributorRequestCanceled is a free log subscription operation binding the contract event 0xa7a5bac33f20d98f2f5d14a6c646812e2adc85c780523174e9634e49f89920cb.
//
// Solidity: event DistributorRequestCanceled(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId)
func (_Dtarm *DtarmFilterer) WatchDistributorRequestCanceled(opts *bind.WatchOpts, sink chan<- *DtarmDistributorRequestCanceled, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "DistributorRequestCanceled", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmDistributorRequestCanceled)
				if err := _Dtarm.contract.UnpackLog(event, "DistributorRequestCanceled", log); err != nil {
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

// ParseDistributorRequestCanceled is a log parse operation binding the contract event 0xa7a5bac33f20d98f2f5d14a6c646812e2adc85c780523174e9634e49f89920cb.
//
// Solidity: event DistributorRequestCanceled(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId)
func (_Dtarm *DtarmFilterer) ParseDistributorRequestCanceled(log types.Log) (*DtarmDistributorRequestCanceled, error) {
	event := new(DtarmDistributorRequestCanceled)
	if err := _Dtarm.contract.UnpackLog(event, "DistributorRequestCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarmDistributorRequestProcessedIterator is returned from FilterDistributorRequestProcessed and is used to iterate over the raw logs and unpacked data for DistributorRequestProcessed events raised by the Dtarm contract.
type DtarmDistributorRequestProcessedIterator struct {
	Event *DtarmDistributorRequestProcessed // Event containing the contract specifics and raw log

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
func (it *DtarmDistributorRequestProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmDistributorRequestProcessed)
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
		it.Event = new(DtarmDistributorRequestProcessed)
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
func (it *DtarmDistributorRequestProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmDistributorRequestProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmDistributorRequestProcessed represents a DistributorRequestProcessed event raised by the Dtarm contract.
type DtarmDistributorRequestProcessed struct {
	RequestId [32]byte
	Shares    *big.Int
	Status    uint8
	Error     []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributorRequestProcessed is a free log retrieval operation binding the contract event 0x196b304e6ee346f728835d1a56f547b4b2b6673b303827a1963fde2992257272.
//
// Solidity: event DistributorRequestProcessed(bytes32 requestId, uint256 shares, uint8 status, bytes error)
func (_Dtarm *DtarmFilterer) FilterDistributorRequestProcessed(opts *bind.FilterOpts) (*DtarmDistributorRequestProcessedIterator, error) {

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "DistributorRequestProcessed")
	if err != nil {
		return nil, err
	}
	return &DtarmDistributorRequestProcessedIterator{contract: _Dtarm.contract, event: "DistributorRequestProcessed", logs: logs, sub: sub}, nil
}

// WatchDistributorRequestProcessed is a free log subscription operation binding the contract event 0x196b304e6ee346f728835d1a56f547b4b2b6673b303827a1963fde2992257272.
//
// Solidity: event DistributorRequestProcessed(bytes32 requestId, uint256 shares, uint8 status, bytes error)
func (_Dtarm *DtarmFilterer) WatchDistributorRequestProcessed(opts *bind.WatchOpts, sink chan<- *DtarmDistributorRequestProcessed) (event.Subscription, error) {

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "DistributorRequestProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmDistributorRequestProcessed)
				if err := _Dtarm.contract.UnpackLog(event, "DistributorRequestProcessed", log); err != nil {
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

// ParseDistributorRequestProcessed is a log parse operation binding the contract event 0x196b304e6ee346f728835d1a56f547b4b2b6673b303827a1963fde2992257272.
//
// Solidity: event DistributorRequestProcessed(bytes32 requestId, uint256 shares, uint8 status, bytes error)
func (_Dtarm *DtarmFilterer) ParseDistributorRequestProcessed(log types.Log) (*DtarmDistributorRequestProcessed, error) {
	event := new(DtarmDistributorRequestProcessed)
	if err := _Dtarm.contract.UnpackLog(event, "DistributorRequestProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarmDistributorRequestProcessingIterator is returned from FilterDistributorRequestProcessing and is used to iterate over the raw logs and unpacked data for DistributorRequestProcessing events raised by the Dtarm contract.
type DtarmDistributorRequestProcessingIterator struct {
	Event *DtarmDistributorRequestProcessing // Event containing the contract specifics and raw log

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
func (it *DtarmDistributorRequestProcessingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmDistributorRequestProcessing)
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
		it.Event = new(DtarmDistributorRequestProcessing)
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
func (it *DtarmDistributorRequestProcessingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmDistributorRequestProcessingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmDistributorRequestProcessing represents a DistributorRequestProcessing event raised by the Dtarm contract.
type DtarmDistributorRequestProcessing struct {
	FundAdminAddr   common.Address
	FundTokenId     [32]byte
	DistributorAddr common.Address
	RequestId       [32]byte
	Shares          *big.Int
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributorRequestProcessing is a free log retrieval operation binding the contract event 0x0fe79093a6d3f5159806fb45e2f50562fffbec87d3356fda646d3a466cc58783.
//
// Solidity: event DistributorRequestProcessing(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId, uint256 shares, uint256 amount)
func (_Dtarm *DtarmFilterer) FilterDistributorRequestProcessing(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtarmDistributorRequestProcessingIterator, error) {

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

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "DistributorRequestProcessing", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtarmDistributorRequestProcessingIterator{contract: _Dtarm.contract, event: "DistributorRequestProcessing", logs: logs, sub: sub}, nil
}

// WatchDistributorRequestProcessing is a free log subscription operation binding the contract event 0x0fe79093a6d3f5159806fb45e2f50562fffbec87d3356fda646d3a466cc58783.
//
// Solidity: event DistributorRequestProcessing(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId, uint256 shares, uint256 amount)
func (_Dtarm *DtarmFilterer) WatchDistributorRequestProcessing(opts *bind.WatchOpts, sink chan<- *DtarmDistributorRequestProcessing, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "DistributorRequestProcessing", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmDistributorRequestProcessing)
				if err := _Dtarm.contract.UnpackLog(event, "DistributorRequestProcessing", log); err != nil {
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

// ParseDistributorRequestProcessing is a log parse operation binding the contract event 0x0fe79093a6d3f5159806fb45e2f50562fffbec87d3356fda646d3a466cc58783.
//
// Solidity: event DistributorRequestProcessing(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId, uint256 shares, uint256 amount)
func (_Dtarm *DtarmFilterer) ParseDistributorRequestProcessing(log types.Log) (*DtarmDistributorRequestProcessing, error) {
	event := new(DtarmDistributorRequestProcessing)
	if err := _Dtarm.contract.UnpackLog(event, "DistributorRequestProcessing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarmFundAdminRegisteredIterator is returned from FilterFundAdminRegistered and is used to iterate over the raw logs and unpacked data for FundAdminRegistered events raised by the Dtarm contract.
type DtarmFundAdminRegisteredIterator struct {
	Event *DtarmFundAdminRegistered // Event containing the contract specifics and raw log

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
func (it *DtarmFundAdminRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmFundAdminRegistered)
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
		it.Event = new(DtarmFundAdminRegistered)
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
func (it *DtarmFundAdminRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmFundAdminRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmFundAdminRegistered represents a FundAdminRegistered event raised by the Dtarm contract.
type DtarmFundAdminRegistered struct {
	FundAdminAddr common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFundAdminRegistered is a free log retrieval operation binding the contract event 0xa5989e8c3c95d94721e03ce0e0d7f247ffe2272bd032bd6d68c91aff5b68b8c1.
//
// Solidity: event FundAdminRegistered(address fundAdminAddr)
func (_Dtarm *DtarmFilterer) FilterFundAdminRegistered(opts *bind.FilterOpts) (*DtarmFundAdminRegisteredIterator, error) {

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "FundAdminRegistered")
	if err != nil {
		return nil, err
	}
	return &DtarmFundAdminRegisteredIterator{contract: _Dtarm.contract, event: "FundAdminRegistered", logs: logs, sub: sub}, nil
}

// WatchFundAdminRegistered is a free log subscription operation binding the contract event 0xa5989e8c3c95d94721e03ce0e0d7f247ffe2272bd032bd6d68c91aff5b68b8c1.
//
// Solidity: event FundAdminRegistered(address fundAdminAddr)
func (_Dtarm *DtarmFilterer) WatchFundAdminRegistered(opts *bind.WatchOpts, sink chan<- *DtarmFundAdminRegistered) (event.Subscription, error) {

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "FundAdminRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmFundAdminRegistered)
				if err := _Dtarm.contract.UnpackLog(event, "FundAdminRegistered", log); err != nil {
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

// ParseFundAdminRegistered is a log parse operation binding the contract event 0xa5989e8c3c95d94721e03ce0e0d7f247ffe2272bd032bd6d68c91aff5b68b8c1.
//
// Solidity: event FundAdminRegistered(address fundAdminAddr)
func (_Dtarm *DtarmFilterer) ParseFundAdminRegistered(log types.Log) (*DtarmFundAdminRegistered, error) {
	event := new(DtarmFundAdminRegistered)
	if err := _Dtarm.contract.UnpackLog(event, "FundAdminRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarmFundTokenAllowlistUpdatedIterator is returned from FilterFundTokenAllowlistUpdated and is used to iterate over the raw logs and unpacked data for FundTokenAllowlistUpdated events raised by the Dtarm contract.
type DtarmFundTokenAllowlistUpdatedIterator struct {
	Event *DtarmFundTokenAllowlistUpdated // Event containing the contract specifics and raw log

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
func (it *DtarmFundTokenAllowlistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmFundTokenAllowlistUpdated)
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
		it.Event = new(DtarmFundTokenAllowlistUpdated)
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
func (it *DtarmFundTokenAllowlistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmFundTokenAllowlistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmFundTokenAllowlistUpdated represents a FundTokenAllowlistUpdated event raised by the Dtarm contract.
type DtarmFundTokenAllowlistUpdated struct {
	FundAdminAddr   common.Address
	FundTokenId     [32]byte
	DistributorAddr common.Address
	Allowed         bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFundTokenAllowlistUpdated is a free log retrieval operation binding the contract event 0x13dc013c86d37e8239ff3de6c79a07b274a5ca98a3d2500738b3f841b39fd286.
//
// Solidity: event FundTokenAllowlistUpdated(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bool allowed)
func (_Dtarm *DtarmFilterer) FilterFundTokenAllowlistUpdated(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtarmFundTokenAllowlistUpdatedIterator, error) {

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

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "FundTokenAllowlistUpdated", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtarmFundTokenAllowlistUpdatedIterator{contract: _Dtarm.contract, event: "FundTokenAllowlistUpdated", logs: logs, sub: sub}, nil
}

// WatchFundTokenAllowlistUpdated is a free log subscription operation binding the contract event 0x13dc013c86d37e8239ff3de6c79a07b274a5ca98a3d2500738b3f841b39fd286.
//
// Solidity: event FundTokenAllowlistUpdated(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bool allowed)
func (_Dtarm *DtarmFilterer) WatchFundTokenAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *DtarmFundTokenAllowlistUpdated, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "FundTokenAllowlistUpdated", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmFundTokenAllowlistUpdated)
				if err := _Dtarm.contract.UnpackLog(event, "FundTokenAllowlistUpdated", log); err != nil {
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

// ParseFundTokenAllowlistUpdated is a log parse operation binding the contract event 0x13dc013c86d37e8239ff3de6c79a07b274a5ca98a3d2500738b3f841b39fd286.
//
// Solidity: event FundTokenAllowlistUpdated(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bool allowed)
func (_Dtarm *DtarmFilterer) ParseFundTokenAllowlistUpdated(log types.Log) (*DtarmFundTokenAllowlistUpdated, error) {
	event := new(DtarmFundTokenAllowlistUpdated)
	if err := _Dtarm.contract.UnpackLog(event, "FundTokenAllowlistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarmFundTokenRegisteredIterator is returned from FilterFundTokenRegistered and is used to iterate over the raw logs and unpacked data for FundTokenRegistered events raised by the Dtarm contract.
type DtarmFundTokenRegisteredIterator struct {
	Event *DtarmFundTokenRegistered // Event containing the contract specifics and raw log

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
func (it *DtarmFundTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmFundTokenRegistered)
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
		it.Event = new(DtarmFundTokenRegistered)
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
func (it *DtarmFundTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmFundTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmFundTokenRegistered represents a FundTokenRegistered event raised by the Dtarm contract.
type DtarmFundTokenRegistered struct {
	FundAdminAddr      common.Address
	FundTokenId        [32]byte
	FundTokenAddr      common.Address
	NavAddr            common.Address
	TokenChainSelector uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFundTokenRegistered is a free log retrieval operation binding the contract event 0x3fdb34b0cf6792d3099047d32c8b39f6970cbe15a8deafebafc91ed6d124a58c.
//
// Solidity: event FundTokenRegistered(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed fundTokenAddr, address navAddr, uint64 tokenChainSelector)
func (_Dtarm *DtarmFilterer) FilterFundTokenRegistered(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte, fundTokenAddr []common.Address) (*DtarmFundTokenRegisteredIterator, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var fundTokenAddrRule []interface{}
	for _, fundTokenAddrItem := range fundTokenAddr {
		fundTokenAddrRule = append(fundTokenAddrRule, fundTokenAddrItem)
	}

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "FundTokenRegistered", fundAdminAddrRule, fundTokenIdRule, fundTokenAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtarmFundTokenRegisteredIterator{contract: _Dtarm.contract, event: "FundTokenRegistered", logs: logs, sub: sub}, nil
}

// WatchFundTokenRegistered is a free log subscription operation binding the contract event 0x3fdb34b0cf6792d3099047d32c8b39f6970cbe15a8deafebafc91ed6d124a58c.
//
// Solidity: event FundTokenRegistered(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed fundTokenAddr, address navAddr, uint64 tokenChainSelector)
func (_Dtarm *DtarmFilterer) WatchFundTokenRegistered(opts *bind.WatchOpts, sink chan<- *DtarmFundTokenRegistered, fundAdminAddr []common.Address, fundTokenId [][32]byte, fundTokenAddr []common.Address) (event.Subscription, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}
	var fundTokenAddrRule []interface{}
	for _, fundTokenAddrItem := range fundTokenAddr {
		fundTokenAddrRule = append(fundTokenAddrRule, fundTokenAddrItem)
	}

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "FundTokenRegistered", fundAdminAddrRule, fundTokenIdRule, fundTokenAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmFundTokenRegistered)
				if err := _Dtarm.contract.UnpackLog(event, "FundTokenRegistered", log); err != nil {
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

// ParseFundTokenRegistered is a log parse operation binding the contract event 0x3fdb34b0cf6792d3099047d32c8b39f6970cbe15a8deafebafc91ed6d124a58c.
//
// Solidity: event FundTokenRegistered(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed fundTokenAddr, address navAddr, uint64 tokenChainSelector)
func (_Dtarm *DtarmFilterer) ParseFundTokenRegistered(log types.Log) (*DtarmFundTokenRegistered, error) {
	event := new(DtarmFundTokenRegistered)
	if err := _Dtarm.contract.UnpackLog(event, "FundTokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarmInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dtarm contract.
type DtarmInitializedIterator struct {
	Event *DtarmInitialized // Event containing the contract specifics and raw log

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
func (it *DtarmInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmInitialized)
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
		it.Event = new(DtarmInitialized)
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
func (it *DtarmInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmInitialized represents a Initialized event raised by the Dtarm contract.
type DtarmInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Dtarm *DtarmFilterer) FilterInitialized(opts *bind.FilterOpts) (*DtarmInitializedIterator, error) {

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DtarmInitializedIterator{contract: _Dtarm.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Dtarm *DtarmFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DtarmInitialized) (event.Subscription, error) {

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmInitialized)
				if err := _Dtarm.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Dtarm *DtarmFilterer) ParseInitialized(log types.Log) (*DtarmInitialized, error) {
	event := new(DtarmInitialized)
	if err := _Dtarm.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarmInvalidDTARequestSettlementIterator is returned from FilterInvalidDTARequestSettlement and is used to iterate over the raw logs and unpacked data for InvalidDTARequestSettlement events raised by the Dtarm contract.
type DtarmInvalidDTARequestSettlementIterator struct {
	Event *DtarmInvalidDTARequestSettlement // Event containing the contract specifics and raw log

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
func (it *DtarmInvalidDTARequestSettlementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmInvalidDTARequestSettlement)
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
		it.Event = new(DtarmInvalidDTARequestSettlement)
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
func (it *DtarmInvalidDTARequestSettlementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmInvalidDTARequestSettlementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmInvalidDTARequestSettlement represents a InvalidDTARequestSettlement event raised by the Dtarm contract.
type DtarmInvalidDTARequestSettlement struct {
	FundAdminAddr                  common.Address
	FundTokenId                    [32]byte
	RequestId                      [32]byte
	ActualChainSelector            uint64
	ActualDTARequestSettlementAddr common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterInvalidDTARequestSettlement is a free log retrieval operation binding the contract event 0x50e987e07231ee41c1815c2568b18c138fa98a4431c64f09baea9125702d9762.
//
// Solidity: event InvalidDTARequestSettlement(address indexed fundAdminAddr, bytes32 indexed fundTokenId, bytes32 requestId, uint64 actualChainSelector, address actualDTARequestSettlementAddr)
func (_Dtarm *DtarmFilterer) FilterInvalidDTARequestSettlement(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte) (*DtarmInvalidDTARequestSettlementIterator, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "InvalidDTARequestSettlement", fundAdminAddrRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarmInvalidDTARequestSettlementIterator{contract: _Dtarm.contract, event: "InvalidDTARequestSettlement", logs: logs, sub: sub}, nil
}

// WatchInvalidDTARequestSettlement is a free log subscription operation binding the contract event 0x50e987e07231ee41c1815c2568b18c138fa98a4431c64f09baea9125702d9762.
//
// Solidity: event InvalidDTARequestSettlement(address indexed fundAdminAddr, bytes32 indexed fundTokenId, bytes32 requestId, uint64 actualChainSelector, address actualDTARequestSettlementAddr)
func (_Dtarm *DtarmFilterer) WatchInvalidDTARequestSettlement(opts *bind.WatchOpts, sink chan<- *DtarmInvalidDTARequestSettlement, fundAdminAddr []common.Address, fundTokenId [][32]byte) (event.Subscription, error) {

	var fundAdminAddrRule []interface{}
	for _, fundAdminAddrItem := range fundAdminAddr {
		fundAdminAddrRule = append(fundAdminAddrRule, fundAdminAddrItem)
	}
	var fundTokenIdRule []interface{}
	for _, fundTokenIdItem := range fundTokenId {
		fundTokenIdRule = append(fundTokenIdRule, fundTokenIdItem)
	}

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "InvalidDTARequestSettlement", fundAdminAddrRule, fundTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmInvalidDTARequestSettlement)
				if err := _Dtarm.contract.UnpackLog(event, "InvalidDTARequestSettlement", log); err != nil {
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

// ParseInvalidDTARequestSettlement is a log parse operation binding the contract event 0x50e987e07231ee41c1815c2568b18c138fa98a4431c64f09baea9125702d9762.
//
// Solidity: event InvalidDTARequestSettlement(address indexed fundAdminAddr, bytes32 indexed fundTokenId, bytes32 requestId, uint64 actualChainSelector, address actualDTARequestSettlementAddr)
func (_Dtarm *DtarmFilterer) ParseInvalidDTARequestSettlement(log types.Log) (*DtarmInvalidDTARequestSettlement, error) {
	event := new(DtarmInvalidDTARequestSettlement)
	if err := _Dtarm.contract.UnpackLog(event, "InvalidDTARequestSettlement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarmMessageFailedIterator is returned from FilterMessageFailed and is used to iterate over the raw logs and unpacked data for MessageFailed events raised by the Dtarm contract.
type DtarmMessageFailedIterator struct {
	Event *DtarmMessageFailed // Event containing the contract specifics and raw log

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
func (it *DtarmMessageFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmMessageFailed)
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
		it.Event = new(DtarmMessageFailed)
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
func (it *DtarmMessageFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmMessageFailed represents a MessageFailed event raised by the Dtarm contract.
type DtarmMessageFailed struct {
	MessageId [32]byte
	Reason    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageFailed is a free log retrieval operation binding the contract event 0x55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f.
//
// Solidity: event MessageFailed(bytes32 indexed messageId, bytes reason)
func (_Dtarm *DtarmFilterer) FilterMessageFailed(opts *bind.FilterOpts, messageId [][32]byte) (*DtarmMessageFailedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "MessageFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &DtarmMessageFailedIterator{contract: _Dtarm.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

// WatchMessageFailed is a free log subscription operation binding the contract event 0x55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f.
//
// Solidity: event MessageFailed(bytes32 indexed messageId, bytes reason)
func (_Dtarm *DtarmFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *DtarmMessageFailed, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "MessageFailed", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmMessageFailed)
				if err := _Dtarm.contract.UnpackLog(event, "MessageFailed", log); err != nil {
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

// ParseMessageFailed is a log parse operation binding the contract event 0x55bc02a9ef6f146737edeeb425738006f67f077e7138de3bf84a15bde1a5b56f.
//
// Solidity: event MessageFailed(bytes32 indexed messageId, bytes reason)
func (_Dtarm *DtarmFilterer) ParseMessageFailed(log types.Log) (*DtarmMessageFailed, error) {
	event := new(DtarmMessageFailed)
	if err := _Dtarm.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarmNativeFundsRecoveredIterator is returned from FilterNativeFundsRecovered and is used to iterate over the raw logs and unpacked data for NativeFundsRecovered events raised by the Dtarm contract.
type DtarmNativeFundsRecoveredIterator struct {
	Event *DtarmNativeFundsRecovered // Event containing the contract specifics and raw log

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
func (it *DtarmNativeFundsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmNativeFundsRecovered)
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
		it.Event = new(DtarmNativeFundsRecovered)
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
func (it *DtarmNativeFundsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmNativeFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmNativeFundsRecovered represents a NativeFundsRecovered event raised by the Dtarm contract.
type DtarmNativeFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeFundsRecovered is a free log retrieval operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_Dtarm *DtarmFilterer) FilterNativeFundsRecovered(opts *bind.FilterOpts) (*DtarmNativeFundsRecoveredIterator, error) {

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "NativeFundsRecovered")
	if err != nil {
		return nil, err
	}
	return &DtarmNativeFundsRecoveredIterator{contract: _Dtarm.contract, event: "NativeFundsRecovered", logs: logs, sub: sub}, nil
}

// WatchNativeFundsRecovered is a free log subscription operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_Dtarm *DtarmFilterer) WatchNativeFundsRecovered(opts *bind.WatchOpts, sink chan<- *DtarmNativeFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "NativeFundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmNativeFundsRecovered)
				if err := _Dtarm.contract.UnpackLog(event, "NativeFundsRecovered", log); err != nil {
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
func (_Dtarm *DtarmFilterer) ParseNativeFundsRecovered(log types.Log) (*DtarmNativeFundsRecovered, error) {
	event := new(DtarmNativeFundsRecovered)
	if err := _Dtarm.contract.UnpackLog(event, "NativeFundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarmOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dtarm contract.
type DtarmOwnershipTransferredIterator struct {
	Event *DtarmOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DtarmOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmOwnershipTransferred)
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
		it.Event = new(DtarmOwnershipTransferred)
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
func (it *DtarmOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmOwnershipTransferred represents a OwnershipTransferred event raised by the Dtarm contract.
type DtarmOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtarm *DtarmFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DtarmOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DtarmOwnershipTransferredIterator{contract: _Dtarm.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtarm *DtarmFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DtarmOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmOwnershipTransferred)
				if err := _Dtarm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dtarm *DtarmFilterer) ParseOwnershipTransferred(log types.Log) (*DtarmOwnershipTransferred, error) {
	event := new(DtarmOwnershipTransferred)
	if err := _Dtarm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarmRedemptionRequestedIterator is returned from FilterRedemptionRequested and is used to iterate over the raw logs and unpacked data for RedemptionRequested events raised by the Dtarm contract.
type DtarmRedemptionRequestedIterator struct {
	Event *DtarmRedemptionRequested // Event containing the contract specifics and raw log

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
func (it *DtarmRedemptionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmRedemptionRequested)
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
		it.Event = new(DtarmRedemptionRequested)
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
func (it *DtarmRedemptionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmRedemptionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmRedemptionRequested represents a RedemptionRequested event raised by the Dtarm contract.
type DtarmRedemptionRequested struct {
	FundAdminAddr   common.Address
	FundTokenId     [32]byte
	DistributorAddr common.Address
	RequestId       [32]byte
	Shares          *big.Int
	CreatedAt       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRedemptionRequested is a free log retrieval operation binding the contract event 0x064109c697574218a390e788e6cd1a98d603bfbf5fad2c37cd90c70a1c150d31.
//
// Solidity: event RedemptionRequested(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId, uint256 shares, uint40 createdAt)
func (_Dtarm *DtarmFilterer) FilterRedemptionRequested(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtarmRedemptionRequestedIterator, error) {

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

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "RedemptionRequested", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtarmRedemptionRequestedIterator{contract: _Dtarm.contract, event: "RedemptionRequested", logs: logs, sub: sub}, nil
}

// WatchRedemptionRequested is a free log subscription operation binding the contract event 0x064109c697574218a390e788e6cd1a98d603bfbf5fad2c37cd90c70a1c150d31.
//
// Solidity: event RedemptionRequested(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId, uint256 shares, uint40 createdAt)
func (_Dtarm *DtarmFilterer) WatchRedemptionRequested(opts *bind.WatchOpts, sink chan<- *DtarmRedemptionRequested, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "RedemptionRequested", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmRedemptionRequested)
				if err := _Dtarm.contract.UnpackLog(event, "RedemptionRequested", log); err != nil {
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

// ParseRedemptionRequested is a log parse operation binding the contract event 0x064109c697574218a390e788e6cd1a98d603bfbf5fad2c37cd90c70a1c150d31.
//
// Solidity: event RedemptionRequested(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId, uint256 shares, uint40 createdAt)
func (_Dtarm *DtarmFilterer) ParseRedemptionRequested(log types.Log) (*DtarmRedemptionRequested, error) {
	event := new(DtarmRedemptionRequested)
	if err := _Dtarm.contract.UnpackLog(event, "RedemptionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarmSubscriptionRequestedIterator is returned from FilterSubscriptionRequested and is used to iterate over the raw logs and unpacked data for SubscriptionRequested events raised by the Dtarm contract.
type DtarmSubscriptionRequestedIterator struct {
	Event *DtarmSubscriptionRequested // Event containing the contract specifics and raw log

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
func (it *DtarmSubscriptionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmSubscriptionRequested)
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
		it.Event = new(DtarmSubscriptionRequested)
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
func (it *DtarmSubscriptionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmSubscriptionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmSubscriptionRequested represents a SubscriptionRequested event raised by the Dtarm contract.
type DtarmSubscriptionRequested struct {
	FundAdminAddr   common.Address
	FundTokenId     [32]byte
	DistributorAddr common.Address
	RequestId       [32]byte
	Amount          *big.Int
	CreatedAt       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionRequested is a free log retrieval operation binding the contract event 0x38c91dbc981d3a6ddeb0aaf314c6aa1ffee67d81c7fc7b15bfa1a5a65fd4fff9.
//
// Solidity: event SubscriptionRequested(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId, uint256 amount, uint40 createdAt)
func (_Dtarm *DtarmFilterer) FilterSubscriptionRequested(opts *bind.FilterOpts, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (*DtarmSubscriptionRequestedIterator, error) {

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

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "SubscriptionRequested", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DtarmSubscriptionRequestedIterator{contract: _Dtarm.contract, event: "SubscriptionRequested", logs: logs, sub: sub}, nil
}

// WatchSubscriptionRequested is a free log subscription operation binding the contract event 0x38c91dbc981d3a6ddeb0aaf314c6aa1ffee67d81c7fc7b15bfa1a5a65fd4fff9.
//
// Solidity: event SubscriptionRequested(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId, uint256 amount, uint40 createdAt)
func (_Dtarm *DtarmFilterer) WatchSubscriptionRequested(opts *bind.WatchOpts, sink chan<- *DtarmSubscriptionRequested, fundAdminAddr []common.Address, fundTokenId [][32]byte, distributorAddr []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "SubscriptionRequested", fundAdminAddrRule, fundTokenIdRule, distributorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmSubscriptionRequested)
				if err := _Dtarm.contract.UnpackLog(event, "SubscriptionRequested", log); err != nil {
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

// ParseSubscriptionRequested is a log parse operation binding the contract event 0x38c91dbc981d3a6ddeb0aaf314c6aa1ffee67d81c7fc7b15bfa1a5a65fd4fff9.
//
// Solidity: event SubscriptionRequested(address indexed fundAdminAddr, bytes32 indexed fundTokenId, address indexed distributorAddr, bytes32 requestId, uint256 amount, uint40 createdAt)
func (_Dtarm *DtarmFilterer) ParseSubscriptionRequested(log types.Log) (*DtarmSubscriptionRequested, error) {
	event := new(DtarmSubscriptionRequested)
	if err := _Dtarm.contract.UnpackLog(event, "SubscriptionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DtarmTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the Dtarm contract.
type DtarmTokenWithdrawnIterator struct {
	Event *DtarmTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *DtarmTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtarmTokenWithdrawn)
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
		it.Event = new(DtarmTokenWithdrawn)
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
func (it *DtarmTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtarmTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtarmTokenWithdrawn represents a TokenWithdrawn event raised by the Dtarm contract.
type DtarmTokenWithdrawn struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed recipient, uint256 amount)
func (_Dtarm *DtarmFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*DtarmTokenWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Dtarm.contract.FilterLogs(opts, "TokenWithdrawn", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &DtarmTokenWithdrawnIterator{contract: _Dtarm.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed recipient, uint256 amount)
func (_Dtarm *DtarmFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *DtarmTokenWithdrawn, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Dtarm.contract.WatchLogs(opts, "TokenWithdrawn", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtarmTokenWithdrawn)
				if err := _Dtarm.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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
func (_Dtarm *DtarmFilterer) ParseTokenWithdrawn(log types.Log) (*DtarmTokenWithdrawn, error) {
	event := new(DtarmTokenWithdrawn)
	if err := _Dtarm.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
