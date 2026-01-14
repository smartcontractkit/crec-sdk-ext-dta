package v1_test

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	workflows "github.com/smartcontractkit/cre-workflow-utils"
	apiClient "github.com/smartcontractkit/crec-api-go/client"
	v1 "github.com/smartcontractkit/crec-sdk-ext-dta/v1"
	"github.com/stretchr/testify/require"
)

// buildWatcherEventPayload creates a WatcherEventPayload with the given data for testing
func buildWatcherEventPayload(eventName string, data map[string]any) apiClient.WatcherEventPayload {
	referenceData := workflows.ReferenceData{
		OnChain: []workflows.OnChainReferenceData{
			{
				Source: workflows.OnChainReferenceDataSource{
					ContractAddress:           "0x1234567890123456789012345678901234567890",
					ContractFunctionSignature: v1.DTARequestManagementABI().Methods["getFundToken"].Sig,
					CallData:                  "0x1234567890123456789012345678901234567890",
					Block:                     "latest",
				},
				Data: map[string]any{
					"enabled": true,
					"fund_token_data": v1.FundTokenData{
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
						PaymentInfo: v1.DTAPayment{
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
	event := workflows.VerifiableEvent{
		Event: workflows.Event{
			EventName: eventName,
			Args:      data,
		},
		Trigger: workflows.Trigger{
			TxHash: "0xdeadbeef",
		},
		ReferenceData: &refDataTypeAndValue,
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}
	return apiClient.WatcherEventPayload{
		Address:         "0x1234567890123456789012345678901234567890",
		ChainSelector:   "1",
		Name:            eventName,
		VerifiableEvent: base64.StdEncoding.EncodeToString(eventBytes),
		Type:            apiClient.WatcherEvent,
		WatcherId:       "test-watcher",
	}
}

// buildEvent creates an apiClient.Event from a WatcherEventPayload
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

	payload := buildWatcherEventPayload(v1.EventDistributorRegistered.String(), data)
	event := buildEvent(t, payload)

	result, err := v1.DecodeFromEvent(context.Background(), event)
	require.NoError(t, err)

	require.Equal(t, v1.EventDistributorRegistered, result.EventName())

	concrete, ok := result.ConcreteEvent.(*v1.DistributorRegistered)
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

	payload := buildWatcherEventPayload(v1.EventDistributorRequestProcessed.String(), data)
	event := buildEvent(t, payload)

	result, err := v1.DecodeFromEvent(context.Background(), event)
	require.NoError(t, err)

	require.Equal(t, v1.EventDistributorRequestProcessed, result.EventName())

	concrete, ok := result.ConcreteEvent.(*v1.DistributorRequestProcessed)
	require.True(t, ok, "expected *DistributorRequestProcessed, got %T", result.ConcreteEvent)
	require.Equal(t, common.HexToHash("0x01"), concrete.RequestId)

	expectedShares, _ := new(big.Int).SetString("12345678901234567890", 10)
	require.Equal(t, expectedShares, concrete.Shares)
	require.Equal(t, v1.RequestStatus(7), concrete.Status)
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

	payload := buildWatcherEventPayload(v1.EventSubscriptionRequested.String(), data)
	event := buildEvent(t, payload)

	result, err := v1.DecodeFromEvent(context.Background(), event)
	require.NoError(t, err)

	require.Equal(t, v1.EventSubscriptionRequested, result.EventName())

	concrete, ok := result.ConcreteEvent.(*v1.SubscriptionRequested)
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
			eventType:      v1.EventSubscriptionRequested.String(),
			amountValue:    "1.2e+21",
			expectedAmount: "1200000000000000000000",
			expectError:    false,
		},
		{
			name:           "Large scientific notation",
			eventType:      v1.EventSubscriptionRequested.String(),
			amountValue:    "5e18",
			expectedAmount: "5000000000000000000",
			expectError:    false,
		},
		{
			name:           "Decimal amount with zeros",
			eventType:      v1.EventSubscriptionRequested.String(),
			amountValue:    "600000000000000000000.000000",
			expectedAmount: "600000000000000000000",
			expectError:    false,
		},
		{
			name:           "Simple decimal (non-zero fractional) should fail",
			eventType:      v1.EventSubscriptionRequested.String(),
			amountValue:    "123.456",
			expectedAmount: "",
			expectError:    true,
		},
		{
			name:           "Invalid amount should fail",
			eventType:      v1.EventSubscriptionRequested.String(),
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

			concrete, ok := result.ConcreteEvent.(*v1.SubscriptionRequested)
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
	// Create an event with an invalid payload type
	var event apiClient.Event
	// Set an invalid/empty payload
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

	payload := buildWatcherEventPayload(v1.EventRedemptionRequested.String(), data)
	event := buildEvent(t, payload)

	result, err := v1.DecodeFromEvent(context.Background(), event)
	require.NoError(t, err)

	require.Equal(t, v1.EventRedemptionRequested, result.EventName())

	concrete, ok := result.ConcreteEvent.(*v1.RedemptionRequested)
	require.True(t, ok, "expected *RedemptionRequested, got %T", result.ConcreteEvent)

	expectedShares := new(big.Int)
	expectedShares.SetString("314159000000000000000", 10)
	require.Equal(t, expectedShares, concrete.Shares)
}

func TestDecodeFromEvent_DecodedEventFields(t *testing.T) {
	data := map[string]interface{}{
		"distributor_addr": "0x00000000000000000000000000000000000000cc",
	}

	payload := buildWatcherEventPayload(v1.EventDistributorRegistered.String(), data)
	event := buildEvent(t, payload)

	result, err := v1.DecodeFromEvent(context.Background(), event)
	require.NoError(t, err)

	// Verify WatcherEventPayload fields are accessible directly
	require.Equal(t, "1", result.ChainSelector)
	require.Equal(t, payload.Address, result.Address)
	require.Equal(t, payload.Name, result.Name)
	require.Equal(t, payload.Domain, result.Domain)
	require.Equal(t, payload.WatcherId, result.WatcherId)
}

// TestEventPayloadRoundTrip verifies that event payloads can be marshalled and unmarshalled
func TestEventPayloadRoundTrip(t *testing.T) {
	data := map[string]interface{}{
		"distributor_addr": "0x00000000000000000000000000000000000000aa",
	}

	originalPayload := buildWatcherEventPayload(v1.EventDistributorRegistered.String(), data)

	// Marshal to JSON
	jsonBytes, err := json.Marshal(originalPayload)
	require.NoError(t, err)

	// Unmarshal back
	var decodedPayload apiClient.WatcherEventPayload
	err = json.Unmarshal(jsonBytes, &decodedPayload)
	require.NoError(t, err)

	require.Equal(t, originalPayload.Name, decodedPayload.Name)
	require.Equal(t, originalPayload.Address, decodedPayload.Address)
}
