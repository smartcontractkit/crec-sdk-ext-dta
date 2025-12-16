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
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-protos/cre/go/values/pb"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	evmmock "github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/mock"
	httpcap "github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	httpmock "github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http/mock"
	workflows "github.com/smartcontractkit/cre-workflow-utils"
	wf "github.com/smartcontractkit/crec-sdk-ext-dta/v1/watcher/handler"
)

func TestWatcherV1_DTA_SimpleFlow_Post(t *testing.T) {
	rt := workflows.PrepareTestingRuntime(t)

	cfg := &workflows.Config{
		Network:       "evm",
		ChainID:       "31337",
		ChainSelector: "3379446385462418246",
		Service:       "dta",
		CourierURL:    "http://example.com",
		ApiKeySecret:  "courier",
		DetectEventTriggerConfig: workflows.DetectEventTriggerConfig{
			ContractName:      "TransparentUpgradeableProxy",
			ContractAddress:   "0x84eA74d481Ee0A5332c457a4d796187F6Ba67fEB",
			ContractEventName: "FundAdminRegistered",
			ContractReaderConfig: workflows.ContractReaderConfig{
				Contracts: map[string]workflows.ContractDef{
					"TransparentUpgradeableProxy": {
						ContractABI: `[{"type":"event","name":"FundAdminRegistered","inputs":[{"name":"fundAdminAddr","type":"address","indexed":false,"internalType":"address"}],"anonymous":false}]`,
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
	httpCap.SendRequest = func(_ context.Context, req *httpcap.Request) (*httpcap.Response, error) {
		var body map[string]any
		require.NoError(t, json.Unmarshal(req.Body, &body))
		verStr := body["verifiable_event"].(string)
		verBytes, _ := base64.StdEncoding.DecodeString(verStr)
		var embedded map[string]any
		require.NoError(t, json.Unmarshal(verBytes, &embedded))

		ev := embedded["event"].(map[string]any)
		require.Equal(t, "dta", ev["service"])
		require.Equal(t, "FundAdminRegistered", ev["name"])

		params := embedded["parameters"].(map[string]any)
		require.Equal(t, "0x0000000000000000000000000000000000000001", params["fund_admin_addr"])

		meta := embedded["metadata"].(map[string]any)
		wfe := meta["workflowEvent"].(map[string]any)
		attrs := wfe["attributes"].(map[string]any)
		fa := attrs["fund_admin_addr"].(map[string]any)
		require.Equal(t, "0x0000000000000000000000000000000000000001", fa["value"])

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

	_, err = wf.OnLog(cfg, rt, log)
	require.NoError(t, err)
}

