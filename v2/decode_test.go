package v2_test

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	workflows "github.com/smartcontractkit/crec-workflow-utils"
	apiClient "github.com/smartcontractkit/crec-api-go/client"
	apiModels "github.com/smartcontractkit/crec-api-go/models"
	v2 "github.com/smartcontractkit/crec-sdk-ext-dta/v2"
	"github.com/smartcontractkit/crec-sdk-ext-dta/v2/events"
	"github.com/stretchr/testify/require"
)

type referenceDataOpt func(*workflows.ReferenceData)

func withFundTokenData(ftd events.FundTokenData) referenceDataOpt {
	return func(rd *workflows.ReferenceData) {
		rd.OnChain = append(rd.OnChain, workflows.OnChainReferenceData{
			Source: workflows.OnChainReferenceDataSource{
				ContractAddress:           "0x1234567890123456789012345678901234567890",
				ContractFunctionSignature: "getFundToken(address,bytes32)",
				CallData:                  "0x1234",
				Block:                     "latest",
			},
			Data: map[string]any{
				"enabled":         true,
				"fund_token_data": ftd,
			},
		})
	}
}

func withDistributorRequest(dr events.DistributorRequest) referenceDataOpt {
	return func(rd *workflows.ReferenceData) {
		rd.OnChain = append(rd.OnChain, workflows.OnChainReferenceData{
			Source: workflows.OnChainReferenceDataSource{
				ContractAddress:           "0x1234567890123456789012345678901234567890",
				ContractFunctionSignature: "getDistributorRequest(bytes32)",
				CallData:                  "0x5678",
				Block:                     "latest",
			},
			Data: map[string]any{
				"distributor_request": dr,
			},
		})
	}
}

func withMatchingSignatureButMissingKey(sig string) referenceDataOpt {
	return func(rd *workflows.ReferenceData) {
		rd.OnChain = append(rd.OnChain, workflows.OnChainReferenceData{
			Source: workflows.OnChainReferenceDataSource{
				ContractAddress:           "0x1234567890123456789012345678901234567890",
				ContractFunctionSignature: sig,
				CallData:                  "0x0000",
				Block:                     "latest",
			},
			Data: map[string]any{
				"some_other_field": "irrelevant",
			},
		})
	}
}

func testFundTokenData() events.FundTokenData {
	return events.FundTokenData{
		FundTokenAddr:                 common.HexToAddress("0x1234567890123456789012345678901234567890"),
		NavFeedDecimals:               18,
		PurchaseTokenRoundingDecimals: 18,
		PurchaseTokenDecimals:         18,
		FundRoundingDecimals:          18,
		FundTokenDecimals:             18,
		RequestsPerDay:                10,
		NavAddr:                       common.HexToAddress("0x1234567890123456789012345678901234567890"),
		TokenChainSelector:            1,
		DtaRequestSettlementAddr:      common.HexToAddress("0x1234567890123456789012345678901234567890"),
		TimezoneOffsetSecs:            big.NewInt(0),
		NavTTL:                        big.NewInt(0),
		PaymentInfo: events.DTAPayment{
			OffChainPaymentCurrency: 147,
			PaymentTokenSourceAddr:  common.HexToAddress("0x0"),
			PaymentTokenDestAddr:    common.HexToAddress("0x0"),
		},
	}
}

func buildWatcherEventPayload(eventName string, data map[string]any, opts ...referenceDataOpt) apiClient.WatcherEventPayload {
	var refData *workflows.ReferenceData
	if len(opts) > 0 {
		rd := &workflows.ReferenceData{}
		for _, opt := range opts {
			opt(rd)
		}
		refData = rd
	}

	var eventData map[string]any
	if refData != nil {
		refDataBytes, err := json.Marshal(refData)
		if err != nil {
			panic(err)
		}
		tv := workflows.TypeAndValue{
			Type:  workflows.RawMessageTypeReferenceData,
			Value: json.RawMessage(refDataBytes),
		}
		tvBytes, err := json.Marshal(tv)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(tvBytes, &eventData)
		if err != nil {
			panic(err)
		}
	}

	contractAddress := "0x1234567890123456789012345678901234567890"
	txHash := common.HexToHash("0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890").Hex()
	evmEvent := apiModels.EVMEvent{
		Address:        contractAddress,
		BlockNumber:    12345,
		BlockTimestamp: uint64(time.Now().Unix()),
		ChainId:        "1",
		EventSignature: eventName + "()",
		LogIndex:       0,
		Params:         &data,
		TopicHash:      common.HexToHash("0x1234567890123456789012345678901234567890123456789012345678901234").Hex(),
		TxHash:         txHash,
	}

	chainEvent := &apiModels.VerifiableEvent_ChainEvent{}
	if err := chainEvent.FromEVMEvent(evmEvent); err != nil {
		panic(err)
	}

	chainFamily := "evm"
	chainSelector := "1"
	service := "test-service"
	timestamp := time.Now()
	event := apiModels.VerifiableEvent{
		ChainEvent:    chainEvent,
		ChainFamily:   &chainFamily,
		ChainSelector: &chainSelector,
		Data:          &eventData,
		Name:          eventName,
		Service:       &service,
		Timestamp:     timestamp,
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}

	eventHash := common.BytesToHash(eventBytes[:32]).Hex()

	return apiClient.WatcherEventPayload{
		EventHash:       eventHash,
		VerifiableEvent: base64.StdEncoding.EncodeToString(eventBytes),
		WatcherId:       "test-watcher",
	}
}

func buildEvent(t *testing.T, payload apiClient.WatcherEventPayload) apiClient.Event {
	t.Helper()
	var event apiClient.Event
	err := event.Payload.FromWatcherEventPayload(payload)
	require.NoError(t, err)
	return event
}

func TestDecodeFromEvent_DecodeEvents(t *testing.T) {
	tests := []struct {
		name           string
		eventName      string
		data           map[string]any
		refDataOpts    []referenceDataOpt
		assertConcrete func(t *testing.T, concrete events.ConcreteEvent)
	}{
		{
			name:      "DistributorRegistered",
			eventName: events.EventDistributorRegistered.String(),
			data: map[string]any{
				"distributor_addr": "0x00000000000000000000000000000000000000aa",
			},
			refDataOpts: []referenceDataOpt{withFundTokenData(testFundTokenData())},
			assertConcrete: func(t *testing.T, concrete events.ConcreteEvent) {
				c, ok := concrete.(*events.DistributorRegistered)
				require.True(t, ok, "expected *DistributorRegistered, got %T", concrete)
				require.Equal(t, common.HexToAddress("0xaa"), c.DistributorAddr)
			},
		},
		{
			name:      "DistributorRequestProcessed",
			eventName: events.EventDistributorRequestProcessed.String(),
			data: map[string]any{
				"request_id": common.HexToHash("0x01").Hex(),
				"shares":     "12345678901234567890",
				"status":     "7",
				"error":      "some-bytes",
			},
			refDataOpts: []referenceDataOpt{withFundTokenData(testFundTokenData())},
			assertConcrete: func(t *testing.T, concrete events.ConcreteEvent) {
				c, ok := concrete.(*events.DistributorRequestProcessed)
				require.True(t, ok, "expected *DistributorRequestProcessed, got %T", concrete)
				require.Equal(t, common.HexToHash("0x01"), c.RequestId)
				expectedShares, _ := new(big.Int).SetString("12345678901234567890", 10)
				require.Equal(t, expectedShares, c.Shares)
				require.Equal(t, events.RequestStatus(7), c.Status)
				require.Equal(t, []byte("some-bytes"), c.Error)
			},
		},
		{
			name:      "SubscriptionRequested with v2 ReferenceID",
			eventName: events.EventSubscriptionRequested.String(),
			data: map[string]any{
				"fund_admin_addr":  common.HexToAddress("0xcc").Hex(),
				"fund_token_id":    common.HexToHash("0x02").Hex(),
				"distributor_addr": common.HexToAddress("0xaa").Hex(),
				"reference_id":     common.HexToHash("0xff").Hex(),
				"request_id":       common.HexToHash("0x03").Hex(),
				"amount":           "42",
				"created_at":       "12345",
			},
			refDataOpts: []referenceDataOpt{withFundTokenData(testFundTokenData())},
			assertConcrete: func(t *testing.T, concrete events.ConcreteEvent) {
				c, ok := concrete.(*events.SubscriptionRequested)
				require.True(t, ok, "expected *SubscriptionRequested, got %T", concrete)
				require.Equal(t, common.HexToAddress("0xcc"), c.FundAdminAddr)
				require.Equal(t, common.HexToHash("0x02"), c.FundTokenId)
				require.Equal(t, common.HexToHash("0xff"), c.ReferenceID)
				require.Equal(t, common.HexToHash("0x03"), c.RequestId)
				require.Equal(t, big.NewInt(42), c.Amount)
				require.Equal(t, uint64(12345), c.CreatedAt)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := buildWatcherEventPayload(tt.eventName, tt.data, tt.refDataOpts...)
			event := buildEvent(t, payload)

			result, err := v2.DecodeFromEvent(context.Background(), event)
			require.NoError(t, err)

			name, ok := events.ParseEventName(tt.eventName)
			require.True(t, ok)
			require.Equal(t, name, result.EventName())

			tt.assertConcrete(t, result.ConcreteEvent)
		})
	}
}

func TestDecodeFromEvent_ReferenceDataLenient(t *testing.T) {
	baseData := map[string]any{
		"distributor_addr": "0x00000000000000000000000000000000000000aa",
	}

	tests := []struct {
		name              string
		refDataOpts       []referenceDataOpt
		expectFundToken   bool
		expectDistributor bool
	}{
		{
			name:              "no reference data at all",
			refDataOpts:       nil,
			expectFundToken:   false,
			expectDistributor: false,
		},
		{
			name:              "empty reference data",
			refDataOpts:       []referenceDataOpt{},
			expectFundToken:   false,
			expectDistributor: false,
		},
		{
			name: "matching getFundToken signature but missing fund_token_data key",
			refDataOpts: []referenceDataOpt{
				withMatchingSignatureButMissingKey("getFundToken(address,bytes32)"),
			},
			expectFundToken:   false,
			expectDistributor: false,
		},
		{
			name: "matching getDistributorRequest signature but missing distributor_request key",
			refDataOpts: []referenceDataOpt{
				withFundTokenData(testFundTokenData()),
				withMatchingSignatureButMissingKey("getDistributorRequest(bytes32)"),
			},
			expectFundToken:   true,
			expectDistributor: false,
		},
		{
			name: "full reference data present",
			refDataOpts: []referenceDataOpt{
				withFundTokenData(testFundTokenData()),
				withDistributorRequest(events.DistributorRequest{
					Shares:          big.NewInt(100),
					Amount:          big.NewInt(200),
					FundTokenId:     [32]byte{1},
					FundAdminAddr:   common.HexToAddress("0xaa"),
					DistributorAddr: common.HexToAddress("0xbb"),
					CreatedAt:       big.NewInt(12345),
					RequestType:     1,
					Status:          2,
				}),
			},
			expectFundToken:   true,
			expectDistributor: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := buildWatcherEventPayload(
				events.EventDistributorRegistered.String(),
				baseData,
				tt.refDataOpts...,
			)
			event := buildEvent(t, payload)

			result, err := v2.DecodeFromEvent(context.Background(), event)
			require.NoError(t, err)
			require.Equal(t, events.EventDistributorRegistered, result.EventName())
			require.NotNil(t, result.ConcreteEvent)

			if tt.expectFundToken {
				require.NotNil(t, result.FundTokenData)
			} else {
				require.Nil(t, result.FundTokenData)
			}

			if tt.expectDistributor {
				require.NotNil(t, result.DistributorRequest)
			} else {
				require.Nil(t, result.DistributorRequest)
			}
		})
	}
}

func TestDecodeFromEvent_DistributorRequestReferenceIDRoundTrip(t *testing.T) {
	referenceID := [32]byte{9}
	data := map[string]any{
		"distributor_addr": "0x00000000000000000000000000000000000000aa",
	}

	payload := buildWatcherEventPayload(
		events.EventDistributorRegistered.String(),
		data,
		withFundTokenData(testFundTokenData()),
		withDistributorRequest(events.DistributorRequest{
			Shares:          big.NewInt(100),
			Amount:          big.NewInt(200),
			FundTokenId:     [32]byte{1},
			ReferenceID:     referenceID,
			FundAdminAddr:   common.HexToAddress("0xaa"),
			DistributorAddr: common.HexToAddress("0xbb"),
			CreatedAt:       big.NewInt(12345),
			RequestType:     1,
			Status:          2,
		}),
	)
	event := buildEvent(t, payload)

	result, err := v2.DecodeFromEvent(context.Background(), event)
	require.NoError(t, err)
	require.NotNil(t, result.DistributorRequest)
	require.Equal(t, referenceID, result.DistributorRequest.ReferenceID)
}

func TestDecodeFromEvent_UnsupportedEvent(t *testing.T) {
	data := map[string]any{"some_field": "some_value"}
	payload := buildWatcherEventPayload("SomeUnknownEvent", data)
	event := buildEvent(t, payload)

	_, err := v2.DecodeFromEvent(context.Background(), event)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported event")
}

func TestDecodeFromEvent_InvalidPayload(t *testing.T) {
	var event apiClient.Event
	event.Payload = apiClient.Event_Payload{}

	_, err := v2.DecodeFromEvent(context.Background(), event)
	require.Error(t, err)
}

func TestDecodeFromEvent_DecodedEventFields(t *testing.T) {
	data := map[string]any{
		"distributor_addr": "0x00000000000000000000000000000000000000cc",
	}
	payload := buildWatcherEventPayload(
		events.EventDistributorRegistered.String(),
		data,
		withFundTokenData(testFundTokenData()),
	)
	event := buildEvent(t, payload)

	result, err := v2.DecodeFromEvent(context.Background(), event)
	require.NoError(t, err)

	require.Equal(t, payload.WatcherId, result.WatcherId)
	require.Equal(t, payload.VerifiableEvent, result.VerifiableEvent)
	require.Equal(t, payload.EventHash, result.EventHash)
}

func TestDecodeFromEvent_ScientificNotation(t *testing.T) {
	tests := []struct {
		name           string
		amountValue    string
		expectedAmount string
		expectError    bool
	}{
		{
			name:           "scientific notation amount",
			amountValue:    "1.2e+21",
			expectedAmount: "1200000000000000000000",
		},
		{
			name:           "large scientific notation",
			amountValue:    "5e18",
			expectedAmount: "5000000000000000000",
		},
		{
			name:           "decimal with trailing zeros",
			amountValue:    "600000000000000000000.000000",
			expectedAmount: "600000000000000000000",
		},
		{
			name:        "non-zero fractional fails",
			amountValue: "123.456",
			expectError: true,
		},
		{
			name:        "invalid amount fails",
			amountValue: "invalid-amount",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := map[string]any{
				"fund_admin_addr":  common.HexToAddress("0xcc").Hex(),
				"fund_token_id":    common.HexToHash("0x02").Hex(),
				"distributor_addr": common.HexToAddress("0xaa").Hex(),
				"reference_id":     common.HexToHash("0x00").Hex(),
				"request_id":       common.HexToHash("0x03").Hex(),
				"amount":           tt.amountValue,
				"created_at":       "12345",
			}
			payload := buildWatcherEventPayload(
				events.EventSubscriptionRequested.String(),
				data,
				withFundTokenData(testFundTokenData()),
			)
			event := buildEvent(t, payload)

			result, err := v2.DecodeFromEvent(context.Background(), event)
			if tt.expectError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			concrete, ok := result.ConcreteEvent.(*events.SubscriptionRequested)
			require.True(t, ok)

			expected := new(big.Int)
			expected.SetString(tt.expectedAmount, 10)
			require.Equal(t, expected, concrete.Amount)
		})
	}
}
