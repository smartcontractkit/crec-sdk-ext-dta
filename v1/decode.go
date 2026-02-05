package v1

import (
	"context"
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
	ConcreteEvent      ConcreteEvent
	FundTokenData      *FundTokenData
	DistributorRequest *DistributorRequest
	PaymentRequests    []workflows.PaymentRequest
}

// EventName returns the parsed event name from the payload.
func (e DecodedEvent) EventName() EventName {
	verifiableEvent, err := workflows.DecodeVerifiableEvent(e.WatcherEventPayload.VerifiableEvent)
	if err != nil || verifiableEvent == nil {
		return EventUnknown
	}
	name, ok := parseEventName(verifiableEvent.Name)
	if !ok {
		return EventUnknown
	}
	return name
}

// EventDecoders returns the map of event decoders for external use.
func EventDecoders() map[EventName]eventDecoder {
	return eventDecoders
}

// DecodeFromEvent extracts the WatcherEventPayload from an apiClient.Event and decodes
// the ConcreteEvent based on the event type.
func DecodeFromEvent(ctx context.Context, event apiClient.Event) (DecodedEvent, error) {
	payload, err := event.Payload.AsWatcherEventPayload()
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("extract watcher payload: %w", err)
	}

	verifiableEvent, err := workflows.DecodeVerifiableEvent(payload.VerifiableEvent)
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("failed to decode verifiable event: %w", err)
	}
	if verifiableEvent == nil {
		return DecodedEvent{}, fmt.Errorf("verifiable event is nil")
	}

	name, ok := parseEventName(verifiableEvent.Name)
	if !ok {
		return DecodedEvent{}, fmt.Errorf("unsupported event name: %s", verifiableEvent.Name)
	}
	decoder, ok := eventDecoders[name]
	if !ok {
		return DecodedEvent{}, fmt.Errorf("unsupported event decoder for event name: %s", name)
	}

	evmEvent, err := verifiableEvent.ChainEvent.AsEVMEvent()
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("failed to convert chain event to evm event: %w", err)
	}

	params := make(map[string]string, len(*evmEvent.Params))
	for k, v := range *evmEvent.Params {
		params[k] = fmt.Sprintf("%v", v)
	}

	concrete, err := decoder(params, evmEvent.TxHash)
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("decode %s: %w", name, err)
	}

	referenceData, err := workflows.GetReferenceDataFromVerifiableEvent(*verifiableEvent)
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("decode reference data: %w", err)
	}

	var fundTokenData *FundTokenData
	var distributorRequest *DistributorRequest
	var paymentRequests []workflows.PaymentRequest
	if referenceData != nil {
		fundTokenData, err = decodeFundTokenData(*referenceData)
		if err != nil {
			return DecodedEvent{}, fmt.Errorf("decode fundTokenData: %w", err)
		}
		distributorRequest, err = decodeDistributorRequest(*referenceData)
		if err != nil {
			return DecodedEvent{}, fmt.Errorf("decode distributorRequest: %w", err)
		}
		paymentRequests, err = decodePaymentRequests(*referenceData)
		if err != nil {
			return DecodedEvent{}, fmt.Errorf("decode paymentRequest: %w", err)
		}
	}

	return DecodedEvent{
		WatcherEventPayload: payload,
		ConcreteEvent:       concrete,
		FundTokenData:       fundTokenData,
		DistributorRequest:  distributorRequest,
		PaymentRequests:     paymentRequests,
	}, nil
}

func decodeFundTokenData(referenceData workflows.ReferenceData) (*FundTokenData, error) {
	for _, onChainReferenceData := range referenceData.OnChain {
		if onChainReferenceData.Source.ContractFunctionSignature == DTARequestManagementABI().Methods["getFundToken"].Sig {
			fundTokenDataRaw, ok := onChainReferenceData.Data["fund_token_data"]
			if !ok {
				return nil, fmt.Errorf("fund_token_data key not found in reference data")
			}

			fundTokenDataBytes, err := json.Marshal(fundTokenDataRaw)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal fund_token_data: %w", err)
			}

			var fundTokenData FundTokenData
			err = json.Unmarshal(fundTokenDataBytes, &fundTokenData)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal fund_token_data: %w", err)
			}

			return &fundTokenData, nil
		}
	}

	return nil, fmt.Errorf("fundTokenData not found")
}

func decodeDistributorRequest(referenceData workflows.ReferenceData) (*DistributorRequest, error) {
	for _, onChainReferenceData := range referenceData.OnChain {
		if onChainReferenceData.Source.ContractFunctionSignature == DTARequestManagementABI().Methods["getDistributorRequest"].Sig {
			distributorRequestRaw, ok := onChainReferenceData.Data["distributor_request"]
			if !ok {
				return nil, fmt.Errorf("distributor_request key not found in reference data")
			}

			distributorRequestBytes, err := json.Marshal(distributorRequestRaw)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal distributor_request: %w", err)
			}

			var distributorRequest DistributorRequest
			err = json.Unmarshal(distributorRequestBytes, &distributorRequest)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal distributor_request: %w", err)
			}

			return &distributorRequest, nil
		}
	}
	return nil, nil // distributor request is optional
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
