package v1_test

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
	v1 "github.com/smartcontractkit/crec-sdk-ext-dta/v1"
	"github.com/smartcontractkit/crec-sdk-ext-dta/v1/events"
	"github.com/stretchr/testify/require"
)

// buildWatcherEventPayload creates a WatcherEventPayload with the given data for testing
func buildWatcherEventPayload(eventName string, data map[string]any) apiClient.WatcherEventPayload {
	referenceData := workflows.ReferenceData{
		OnChain: []workflows.OnChainReferenceData{
			{
				Source: workflows.OnChainReferenceDataSource{
					ContractAddress:           "0x1234567890123456789012345678901234567890",
					ContractFunctionSignature: "getFundToken(address,bytes32)",
					CallData:                  "0x1234567890123456789012345678901234567890",
					Block:                     "latest",
				},
				Data: map[string]any{
					"enabled": true,
					"fund_token_data": events.FundTokenData{
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
							PaymentTokenSourceAddr:  common.HexToAddress("0x0000000000000000000000000000000000000000"),
							PaymentTokenDestAddr:    common.HexToAddress("0x0000000000000000000000000000000000000000"),
						},
					},
				},
			},
		},
	}
	refDataBytes, err := json.Marshal(referenceData)
	if err != nil {
		panic(err)
	}
	refDataTypeAndValue := workflows.TypeAndValue{
		Type:  workflows.RawMessageTypeReferenceData,
		Value: json.RawMessage(refDataBytes),
	}
	typeAndValueBytes, err := json.Marshal(refDataTypeAndValue)
	if err != nil {
		panic(err)
	}
	var eventData map[string]any
	err = json.Unmarshal(typeAndValueBytes, &eventData)
	if err != nil {
		panic(err)
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

func TestDecodeFromEvent_DistributorRegistered(t *testing.T) {
	data := map[string]interface{}{
		"distributor_addr": "0x00000000000000000000000000000000000000aa",
	}

	payload := buildWatcherEventPayload(events.EventDistributorRegistered.String(), data)
	event := buildEvent(t, payload)

	result, err := v1.DecodeFromEvent(context.Background(), event)
	require.NoError(t, err)

	require.Equal(t, events.EventDistributorRegistered, result.EventName())

	concrete, ok := result.ConcreteEvent.(*events.DistributorRegistered)
	require.True(t, ok, "expected *DistributorRegistered, got %T", result.ConcreteEvent)
	require.Equal(t, common.HexToAddress("0x00000000000000000000000000000000000000aa"), concrete.DistributorAddr)
}

func TestDecodeFromEvent_DistributorRequestProcessed(t *testing.T) {
	data := map[string]interface{}{
		"request_id": common.HexToHash("0x01").Hex(),
		"shares":     "12345678901234567890",
		"status":     "7",
		"error":      "some-bytes",
	}

	payload := buildWatcherEventPayload(events.EventDistributorRequestProcessed.String(), data)
	event := buildEvent(t, payload)

	result, err := v1.DecodeFromEvent(context.Background(), event)
	require.NoError(t, err)

	require.Equal(t, events.EventDistributorRequestProcessed, result.EventName())

	concrete, ok := result.ConcreteEvent.(*events.DistributorRequestProcessed)
	require.True(t, ok, "expected *DistributorRequestProcessed, got %T", result.ConcreteEvent)
	require.Equal(t, common.HexToHash("0x01"), concrete.RequestId)

	expectedShares, _ := new(big.Int).SetString("12345678901234567890", 10)
	require.Equal(t, expectedShares, concrete.Shares)
	require.Equal(t, events.RequestStatus(7), concrete.Status)
	require.Equal(t, []byte("some-bytes"), concrete.Error)
}

func TestDecodeFromEvent_SubscriptionRequested(t *testing.T) {
	data := map[string]interface{}{
		"fund_token_id":    common.HexToHash("0x02").Hex(),
		"distributor_addr": common.HexToAddress("0x00000000000000000000000000000000000000aa").Hex(),
		"request_id":       common.HexToHash("0x03").Hex(),
		"amount":           "42",
		"created_at":       "12345",
	}

	payload := buildWatcherEventPayload(events.EventSubscriptionRequested.String(), data)
	event := buildEvent(t, payload)

	result, err := v1.DecodeFromEvent(context.Background(), event)
	require.NoError(t, err)

	require.Equal(t, events.EventSubscriptionRequested, result.EventName())

	concrete, ok := result.ConcreteEvent.(*events.SubscriptionRequested)
	require.True(t, ok, "expected *SubscriptionRequested, got %T", result.ConcreteEvent)
	require.Equal(t, common.HexToHash("0x02"), concrete.FundTokenId)
	require.Equal(t, common.HexToAddress("0x00000000000000000000000000000000000000aa"), concrete.DistributorAddr)
	require.Equal(t, common.HexToHash("0x03"), concrete.RequestId)
	require.Equal(t, big.NewInt(42), concrete.Amount)
	require.Equal(t, uint64(12345), concrete.CreatedAt)
}

func TestDecodeFromEvent_ScientificNotation(t *testing.T) {
	tests := []struct {
		name           string
		eventType      string
		amountValue    string
		expectedAmount string
		expectError    bool
	}{
		{
			name:           "Scientific notation amount",
			eventType:      events.EventSubscriptionRequested.String(),
			amountValue:    "1.2e+21",
			expectedAmount: "1200000000000000000000",
			expectError:    false,
		},
		{
			name:           "Large scientific notation",
			eventType:      events.EventSubscriptionRequested.String(),
			amountValue:    "5e18",
			expectedAmount: "5000000000000000000",
			expectError:    false,
		},
		{
			name:           "Decimal amount with zeros",
			eventType:      events.EventSubscriptionRequested.String(),
			amountValue:    "600000000000000000000.000000",
			expectedAmount: "600000000000000000000",
			expectError:    false,
		},
		{
			name:           "Simple decimal (non-zero fractional) should fail",
			eventType:      events.EventSubscriptionRequested.String(),
			amountValue:    "123.456",
			expectedAmount: "",
			expectError:    true,
		},
		{
			name:           "Invalid amount should fail",
			eventType:      events.EventSubscriptionRequested.String(),
			amountValue:    "invalid-amount",
			expectedAmount: "",
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := map[string]interface{}{
				"fund_token_id":    common.HexToHash("0x02").Hex(),
				"distributor_addr": common.HexToAddress("0xaa").Hex(),
				"request_id":       common.HexToHash("0x03").Hex(),
				"amount":           tt.amountValue,
				"created_at":       "12345",
			}

			payload := buildWatcherEventPayload(tt.eventType, data)
			event := buildEvent(t, payload)

			result, err := v1.DecodeFromEvent(context.Background(), event)

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

func TestDecodeFromEvent_UnsupportedEvent(t *testing.T) {
	data := map[string]interface{}{
		"some_field": "some_value",
	}

	payload := buildWatcherEventPayload("SomeUnknownEvent", data)
	event := buildEvent(t, payload)

	_, err := v1.DecodeFromEvent(context.Background(), event)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported event")
}

func TestDecodeFromEvent_InvalidPayload(t *testing.T) {
	var event apiClient.Event
	event.Payload = apiClient.Event_Payload{}

	_, err := v1.DecodeFromEvent(context.Background(), event)
	require.Error(t, err)
}

func TestDecodeFromEvent_RedemptionRequested(t *testing.T) {
	data := map[string]interface{}{
		"fund_token_id":    common.HexToHash("0x02").Hex(),
		"distributor_addr": common.HexToAddress("0x00000000000000000000000000000000000000bb").Hex(),
		"request_id":       common.HexToHash("0x04").Hex(),
		"shares":           "3.14159e+20",
		"created_at":       "54321",
	}

	payload := buildWatcherEventPayload(events.EventRedemptionRequested.String(), data)
	event := buildEvent(t, payload)

	result, err := v1.DecodeFromEvent(context.Background(), event)
	require.NoError(t, err)

	require.Equal(t, events.EventRedemptionRequested, result.EventName())

	concrete, ok := result.ConcreteEvent.(*events.RedemptionRequested)
	require.True(t, ok, "expected *RedemptionRequested, got %T", result.ConcreteEvent)

	expectedShares := new(big.Int)
	expectedShares.SetString("314159000000000000000", 10)
	require.Equal(t, expectedShares, concrete.Shares)
}

func TestDecodeFromEvent_DecodedEventFields(t *testing.T) {
	data := map[string]interface{}{
		"distributor_addr": "0x00000000000000000000000000000000000000cc",
	}

	payload := buildWatcherEventPayload(events.EventDistributorRegistered.String(), data)
	event := buildEvent(t, payload)

	result, err := v1.DecodeFromEvent(context.Background(), event)
	require.NoError(t, err)

	require.Equal(t, payload.WatcherId, result.WatcherId)
	require.Equal(t, payload.VerifiableEvent, result.VerifiableEvent)
	require.Equal(t, payload.EventHash, result.EventHash)
}

func TestDecodeFromEvent_NoReferenceData(t *testing.T) {
	data := map[string]interface{}{
		"distributor_addr": "0x00000000000000000000000000000000000000aa",
	}

	payload := buildWatcherEventPayloadWithoutReferenceData(events.EventDistributorRegistered.String(), data)
	event := buildEvent(t, payload)

	result, err := v1.DecodeFromEvent(context.Background(), event)
	require.NoError(t, err)

	require.Equal(t, events.EventDistributorRegistered, result.EventName())
	require.NotNil(t, result.ConcreteEvent)
	require.Nil(t, result.FundTokenData)
	require.Nil(t, result.DistributorRequest)
	require.Empty(t, result.PaymentRequests)
}

func TestDecodeFromEvent_MissingReferenceDataKeys(t *testing.T) {
	data := map[string]interface{}{
		"distributor_addr": "0x00000000000000000000000000000000000000aa",
	}

	referenceData := workflows.ReferenceData{
		OnChain: []workflows.OnChainReferenceData{
			{
				Source: workflows.OnChainReferenceDataSource{
					ContractAddress:           "0x1234567890123456789012345678901234567890",
					ContractFunctionSignature: "getFundToken(address,bytes32)",
					CallData:                  "0x1234",
					Block:                     "latest",
				},
				Data: map[string]any{
					"some_other_key": "irrelevant",
				},
			},
			{
				Source: workflows.OnChainReferenceDataSource{
					ContractAddress:           "0x1234567890123456789012345678901234567890",
					ContractFunctionSignature: "getDistributorRequest(bytes32)",
					CallData:                  "0x5678",
					Block:                     "latest",
				},
				Data: map[string]any{
					"wrong_key": "irrelevant",
				},
			},
		},
	}

	payload := buildWatcherEventPayloadWithReferenceData(events.EventDistributorRegistered.String(), data, referenceData)
	event := buildEvent(t, payload)

	result, err := v1.DecodeFromEvent(context.Background(), event)
	require.NoError(t, err)

	require.Equal(t, events.EventDistributorRegistered, result.EventName())
	require.NotNil(t, result.ConcreteEvent)
	require.Nil(t, result.FundTokenData)
	require.Nil(t, result.DistributorRequest)
}

func buildWatcherEventPayloadWithoutReferenceData(eventName string, data map[string]any) apiClient.WatcherEventPayload {
	return buildWatcherEventPayloadWithData(eventName, data, nil)
}

func buildWatcherEventPayloadWithReferenceData(eventName string, data map[string]any, rd workflows.ReferenceData) apiClient.WatcherEventPayload {
	return buildWatcherEventPayloadWithData(eventName, data, &rd)
}

func buildWatcherEventPayloadWithData(eventName string, data map[string]any, refData *workflows.ReferenceData) apiClient.WatcherEventPayload {
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

func TestEventPayloadRoundTrip(t *testing.T) {
	data := map[string]interface{}{
		"distributor_addr": "0x00000000000000000000000000000000000000aa",
	}

	originalPayload := buildWatcherEventPayload(events.EventDistributorRegistered.String(), data)

	jsonBytes, err := json.Marshal(originalPayload)
	require.NoError(t, err)

	var decodedPayload apiClient.WatcherEventPayload
	err = json.Unmarshal(jsonBytes, &decodedPayload)
	require.NoError(t, err)

	require.Equal(t, originalPayload.WatcherId, decodedPayload.WatcherId)
	require.Equal(t, originalPayload.VerifiableEvent, decodedPayload.VerifiableEvent)
	require.Equal(t, originalPayload.EventHash, decodedPayload.EventHash)
}
