package v1

import (
	"context"
	"encoding/json"
	"fmt"

	apiClient "github.com/smartcontractkit/crec-api-go/client"
	"github.com/smartcontractkit/crec-sdk-ext-dta/v1/events"
	workflows "github.com/smartcontractkit/crec-workflow-utils"
)

// Solidity method signatures used to match on-chain reference data.
// These are deterministic from the ABI and kept as constants to avoid
// importing the heavy go-ethereum/accounts/abi package.
const (
	getFundTokenSig          = "getFundToken(address,bytes32)"
	getDistributorRequestSig = "getDistributorRequest(bytes32)"
)

// DecodedEvent wraps WatcherEventPayload with a decoded ConcreteEvent and enrichment data.
type DecodedEvent struct {
	apiClient.WatcherEventPayload
	// ConcreteEvent is the decoded event struct matching the blockchain event type.
	ConcreteEvent events.ConcreteEvent
	// FundTokenData holds fund token configuration from on-chain reference data, when present.
	FundTokenData *events.FundTokenData
	// DistributorRequest holds the distributor request from on-chain reference data, when present.
	DistributorRequest *events.DistributorRequest
	// PaymentRequests holds payment requests from the verifiable event reference data.
	PaymentRequests []workflows.PaymentRequest
}

// EventName returns the parsed event name from the payload.
func (e DecodedEvent) EventName() events.EventName {
	verifiableEvent, err := workflows.DecodeVerifiableEvent(e.VerifiableEvent)
	if err != nil || verifiableEvent == nil {
		return events.EventUnknown
	}
	name, ok := events.ParseEventName(verifiableEvent.Name)
	if !ok {
		return events.EventUnknown
	}
	return name
}

// DecodeFromEvent extracts the WatcherEventPayload from an apiClient.Event and decodes
// the ConcreteEvent and enrichment data (FundTokenData, DistributorRequest, PaymentRequests) based on the event type.
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

	name, ok := events.ParseEventName(verifiableEvent.Name)
	if !ok {
		return DecodedEvent{}, fmt.Errorf("unsupported event name: %s", verifiableEvent.Name)
	}
	decoders := events.EventDecoders()
	decoder, ok := decoders[name]
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

	var fundTokenData *events.FundTokenData
	var distributorRequest *events.DistributorRequest
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

func decodeFundTokenData(referenceData workflows.ReferenceData) (*events.FundTokenData, error) {
	for _, onChainReferenceData := range referenceData.OnChain {
		if onChainReferenceData.Source.ContractFunctionSignature == getFundTokenSig {
			fundTokenDataRaw, ok := onChainReferenceData.Data["fund_token_data"]
			if !ok {
				continue
			}

			fundTokenDataBytes, err := json.Marshal(fundTokenDataRaw)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal fund_token_data: %w", err)
			}

			var fundTokenData events.FundTokenData
			err = json.Unmarshal(fundTokenDataBytes, &fundTokenData)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal fund_token_data: %w", err)
			}

			return &fundTokenData, nil
		}
	}

	return nil, nil
}

func decodeDistributorRequest(referenceData workflows.ReferenceData) (*events.DistributorRequest, error) {
	for _, onChainReferenceData := range referenceData.OnChain {
		if onChainReferenceData.Source.ContractFunctionSignature == getDistributorRequestSig {
			distributorRequestRaw, ok := onChainReferenceData.Data["distributor_request"]
			if !ok {
				continue
			}

			distributorRequestBytes, err := json.Marshal(distributorRequestRaw)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal distributor_request: %w", err)
			}

			var distributorRequest events.DistributorRequest
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
