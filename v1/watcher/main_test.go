package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"strconv"
	"testing"

	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/crec-api-go/models"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-protos/cre/go/values/pb"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	evmmock "github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/mock"
	httpcap "github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	httpmock "github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http/mock"
	"github.com/smartcontractkit/cre-sdk-go/cre/testutils"

	workflows "github.com/smartcontractkit/crec-workflow-utils"

	wf "github.com/smartcontractkit/crec-sdk-ext-dta/v1/watcher/handler"
)

func TestWatcherV1_DTA_SimpleFlow_Post(t *testing.T) {
	rt := testutils.NewRuntime(t, testutils.Secrets{})
	service := "dta.v1"

	cfg := &workflows.Config{
		Network:       "evm",
		ChainSelector: "3379446385462418246",
		Service:       &service,
		CourierURL:    "http://example.com",
		DetectEventTriggerConfig: workflows.DetectEventTriggerConfig{
			ContractName:       "TransparentUpgradeableProxy",
			ContractAddress:    "0x84eA74d481Ee0A5332c457a4d796187F6Ba67fEB",
			ContractEventNames: []string{"FundAdminRegistered"},
			ContractReaderConfig: workflows.ContractReaderConfig{
				Contracts: map[string]workflows.ContractDef{
					"TransparentUpgradeableProxy": {
						ContractABI: `[{"type":"event","name":"FundAdminRegistered","inputs":[{"name":"fundAdminAddr","type":"address","indexed":false,"internalType":"address"}],"anonymous":false},{"type":"function","name":"getFundToken","inputs":[{"name":"fundAdminAddr","type":"address","internalType":"address"},{"name":"fundTokenId","type":"bytes32","internalType":"bytes32"}],"outputs":[{"name":"enabled","type":"bool","internalType":"bool"},{"name":"","type":"tuple","internalType":"struct IFundTokenRegistry.FundTokenData","components":[{"name":"fundTokenAddr","type":"address","internalType":"address"},{"name":"navFeedDecimals","type":"uint8","internalType":"uint8"},{"name":"purchaseTokenRoundingDecimals","type":"uint8","internalType":"uint8"},{"name":"purchaseTokenDecimals","type":"uint8","internalType":"uint8"},{"name":"fundRoundingDecimals","type":"uint8","internalType":"uint8"},{"name":"fundTokenDecimals","type":"uint8","internalType":"uint8"},{"name":"requestsPerDay","type":"uint8","internalType":"uint8"},{"name":"navAddr","type":"address","internalType":"address"},{"name":"tokenChainSelector","type":"uint64","internalType":"uint64"},{"name":"dtaRequestSettlementAddr","type":"address","internalType":"address"},{"name":"timezoneOffsetSecs","type":"int24","internalType":"int24"},{"name":"navTTL","type":"uint24","internalType":"uint24"},{"name":"paymentInfo","type":"tuple","internalType":"struct IDTAMessage.DTAPayment","components":[{"name":"offChainPaymentCurrency","type":"uint8","internalType":"enum Currency"},{"name":"paymentTokenSourceAddr","type":"address","internalType":"address"},{"name":"paymentTokenDestAddr","type":"address","internalType":"address"}]}]}],"stateMutability":"view"}]`,
					},
				},
			},
		},
	}

	chainSelector, err := strconv.ParseUint(cfg.ChainSelector, 10, 64)
	require.NoError(t, err)
	evmCap, err := evmmock.NewClientCapability(chainSelector, t)
	require.NoError(t, err)
	evmCap.HeaderByNumber = func(_ context.Context, _ *evm.HeaderByNumberRequest) (*evm.HeaderByNumberReply, error) {
		return &evm.HeaderByNumberReply{Header: &evm.Header{Timestamp: 42}}, nil
	}

	httpCap, err := httpmock.NewClientCapability(t)
	require.NoError(t, err)
	var httpCapResponse []byte
	httpCap.SendRequest = func(_ context.Context, req *httpcap.Request) (*httpcap.Response, error) {
		httpCapResponse = req.Body
		return &httpcap.Response{StatusCode: 200}, nil
	}

	eventSig := crypto.Keccak256Hash([]byte("FundAdminRegistered(address)")).Bytes()
	addr := gethCommon.HexToAddress("0x0000000000000000000000000000000000000001")
	data := make([]byte, 32)
	copy(data[12:], addr.Bytes())

	log := &evm.Log{
		Address:     gethCommon.HexToAddress(cfg.DetectEventTriggerConfig.ContractAddress).Bytes(),
		Topics:      [][]byte{eventSig},
		Data:        data,
		BlockNumber: &pb.BigInt{AbsVal: big.NewInt(100).Bytes(), Sign: 1},
		Index:       0,
	}

	_, err = wf.OnLog(cfg, rt, log, models.Latest)
	require.NoError(t, err)

	var body map[string]any
	err = json.Unmarshal(httpCapResponse, &body)
	require.NoError(t, err)

	verStr, ok := body["verifiable_event"].(string)
	require.True(t, ok)
	verBytes, err := base64.StdEncoding.DecodeString(verStr)
	require.NoError(t, err)
	var embedded map[string]any
	err = json.Unmarshal(verBytes, &embedded)
	require.NoError(t, err)
	require.Equal(t, service, embedded["service"])

	ev, ok := embedded["chain_event"].(map[string]any)
	require.True(t, ok)
	require.Equal(t, "FundAdminRegistered(address)", ev["event_signature"])

	params, ok := ev["params"].(map[string]any)
	require.True(t, ok)
	require.Equal(t, "0x0000000000000000000000000000000000000001", params["fund_admin_addr"])
}
