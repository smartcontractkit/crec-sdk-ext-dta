package v1

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	workflows "github.com/smartcontractkit/cre-workflow-utils"
	apiClient "github.com/smartcontractkit/crec-api-go/client"
)

// ConcreteEvent represents any decoded concrete event payload.
type ConcreteEvent interface{}

// DecodedEvent wraps WatcherEventPayload with a decoded ConcreteEvent.
type DecodedEvent struct {
	apiClient.WatcherEventPayload
	ConcreteEvent   ConcreteEvent
	FundTokenData   FundTokenData
	PaymentRequests []workflows.PaymentRequest
}

// EventName returns the parsed event name from the payload.
func (e DecodedEvent) EventName() EventName {
	name, ok := parseEventName(e.Name)
	if !ok {
		return EventUnknown
	}
	return name
}

// DecodeFromEvent extracts the WatcherEventPayload from an apiClient.Event and decodes
// the ConcreteEvent based on the event type.
func DecodeFromEvent(_ context.Context, event apiClient.Event) (DecodedEvent, error) {
	payload, err := event.Payload.AsWatcherEventPayload()
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("extract watcher payload: %w", err)
	}

	name, ok := parseEventName(payload.Name)
	if !ok {
		return DecodedEvent{}, fmt.Errorf("unsupported event name: %s", payload.Name)
	}
	decoder, ok := eventDecoders[name]
	if !ok {
		return DecodedEvent{}, fmt.Errorf("unsupported event decoder for event name: %s", name)
	}

	verifiableEventBytes, err := base64.StdEncoding.DecodeString(payload.VerifiableEvent)
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("failed to decode verifiable event: %w", err)
	}

	var verifiableEvent workflows.VerifiableEvent
	err = json.Unmarshal(verifiableEventBytes, &verifiableEvent)
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("failed to unmarshal verifiable event: %w", err)
	}

	params := make(map[string]string, len(verifiableEvent.Event.Args))
	for k, v := range verifiableEvent.Event.Args {
		params[k] = fmt.Sprintf("%v", v)
	}

	concrete, err := decoder(params, verifiableEvent.Trigger.TxHash)
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("decode %s: %w", name, err)
	}

	referenceData, err := decodeReferenceData(payload.VerifiableEvent)
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("decode reference data: %w", err)
	}

	fundTokenData, err := decodeFundTokenData(referenceData)
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("decode fundTokenData: %w", err)
	}

	paymentRequests, err := decodePaymentRequests(referenceData)
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("decode paymentRequest: %w", err)
	}

	return DecodedEvent{
		WatcherEventPayload: payload,
		ConcreteEvent:       concrete,
		FundTokenData:       fundTokenData,
		PaymentRequests:     paymentRequests,
	}, nil
}

func decodeReferenceData(verifiableEventB64 string) (workflows.ReferenceData, error) {
	var referenceData workflows.ReferenceData

	verifiableEventBytes, err := base64.StdEncoding.DecodeString(verifiableEventB64)
	if err != nil {
		return referenceData, fmt.Errorf("failed to base64 decode verifiable event: %w", err)
	}

	var verifiableEvent workflows.VerifiableEvent
	err = json.Unmarshal(verifiableEventBytes, &verifiableEvent)
	if err != nil {
		return referenceData, fmt.Errorf("failed to unmarshal verifiable event: %w", err)
	}

	if verifiableEvent.ReferenceData != nil && verifiableEvent.ReferenceData.Type == workflows.RawMessageTypeReferenceData {
		err = json.Unmarshal(verifiableEvent.ReferenceData.Value, &referenceData)
		if err != nil {
			return referenceData, fmt.Errorf("failed to unmarshal reference data: %w", err)
		}
	}

	return referenceData, nil
}

func decodeFundTokenData(referenceData workflows.ReferenceData) (FundTokenData, error) {
	for _, onChainReferenceData := range referenceData.OnChain {
		if onChainReferenceData.Source.ContractFunctionSignature == DTARequestManagementABI().Methods["getFundToken"].Sig {
			fundTokenDataRaw, ok := onChainReferenceData.Data["fund_token_data"]
			if !ok {
				return FundTokenData{}, fmt.Errorf("fund_token_data key not found in reference data")
			}

			fundTokenDataBytes, err := json.Marshal(fundTokenDataRaw)
			if err != nil {
				return FundTokenData{}, fmt.Errorf("failed to marshal fund_token_data: %w", err)
			}

			var fundTokenData FundTokenData
			err = json.Unmarshal(fundTokenDataBytes, &fundTokenData)
			if err != nil {
				return FundTokenData{}, fmt.Errorf("failed to unmarshal fund_token_data: %w", err)
			}

			return fundTokenData, nil
		}
	}

	return FundTokenData{}, fmt.Errorf("fundTokenData not found")
}

func decodePaymentRequests(referenceData workflows.ReferenceData) ([]workflows.PaymentRequest, error) {
	paymentRequests := make([]workflows.PaymentRequest, 0)
	for _, request := range referenceData.Requests {
		if request.Type == workflows.RawMessageTypePaymentRequest {
			var paymentRequest workflows.PaymentRequest
			err := json.Unmarshal(request.Value, &paymentRequest)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal paymentRequest: %w", err)
			}
			paymentRequests = append(paymentRequests, paymentRequest)
		}
	}

	return paymentRequests, nil
}
