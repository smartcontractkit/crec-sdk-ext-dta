package certify

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	crecapi "github.com/smartcontractkit/crec-api-go/client"
	dtav1 "github.com/smartcontractkit/crec-sdk-ext-dta/v1"
	dtav1events "github.com/smartcontractkit/crec-sdk-ext-dta/v1/events"
	dtav1ops "github.com/smartcontractkit/crec-sdk-ext-dta/v1/operations"
	dtav2 "github.com/smartcontractkit/crec-sdk-ext-dta/v2"
	dtav2events "github.com/smartcontractkit/crec-sdk-ext-dta/v2/events"
	dtav2ops "github.com/smartcontractkit/crec-sdk-ext-dta/v2/operations"
	transactTypes "github.com/smartcontractkit/crec-sdk/transact/types"
)

const (
	dtaServiceV1 = "dta.v1"
	dtaServiceV2 = "dta.v2"

	managementEventDistributorRequestProcessing = "DistributorRequestProcessing"
	settlementEventDTASettlementOpened          = "DTASettlementOpened"
)

type registerFundAdminOperationBuilder interface {
	PrepareRegisterFundAdminOperation() (*transactTypes.Operation, error)
}

type decodedWatcherEvent struct {
	EventName string

	RequestID common.Hash
	Shares    *big.Int
	Amount    *big.Int

	FundTokenSettlementAddr common.Address
	HasFundTokenData        bool
	HasDistributorRequest   bool
	PaymentRequestCount     int
}

func newRegisterFundAdminOperationBuilder(
	cfg *Config,
	walletAddress string,
	deadline *big.Int,
) (registerFundAdminOperationBuilder, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is required")
	}

	switch {
	case cfg.UsesDTAServiceV1():
		return dtav1ops.New(&dtav1ops.Options{
			AccountAddress:              walletAddress,
			Deadline:                    deadline,
			DTARequestManagementAddress: cfg.DTARequestManagementAddress,
			DTARequestSettlementAddress: cfg.DTARequestSettlementAddress,
		})
	case cfg.UsesDTAServiceV2():
		return dtav2ops.New(&dtav2ops.Options{
			AccountAddress:              walletAddress,
			Deadline:                    deadline,
			DTARequestManagementAddress: cfg.DTARequestManagementAddress,
			DTARequestSettlementAddress: cfg.DTARequestSettlementAddress,
		})
	default:
		return nil, fmt.Errorf("unsupported DTA_SERVICE: %q", cfg.DTAService)
	}
}

func decodeManagementWatcherEvent(
	ctx context.Context,
	cfg *Config,
	event crecapi.Event,
) (decodedWatcherEvent, error) {
	if cfg == nil {
		return decodedWatcherEvent{}, fmt.Errorf("config is required")
	}

	switch {
	case cfg.UsesDTAServiceV1():
		decoded, err := dtav1.DecodeFromEvent(ctx, event)
		if err != nil {
			return decodedWatcherEvent{}, fmt.Errorf("decode dta.v1 management watcher event: %w", err)
		}

		concrete, ok := decoded.ConcreteEvent.(*dtav1events.DistributorRequestProcessing)
		if !ok {
			return decodedWatcherEvent{}, fmt.Errorf(
				"expected *v1events.DistributorRequestProcessing, got %T",
				decoded.ConcreteEvent,
			)
		}

		return decodedWatcherEvent{
			EventName:               decoded.EventName().String(),
			RequestID:               concrete.RequestId,
			Shares:                  concrete.Shares,
			Amount:                  concrete.Amount,
			FundTokenSettlementAddr: v1SettlementAddr(decoded.FundTokenData),
			HasFundTokenData:        decoded.FundTokenData != nil,
			HasDistributorRequest:   decoded.DistributorRequest != nil,
			PaymentRequestCount:     len(decoded.PaymentRequests),
		}, nil

	case cfg.UsesDTAServiceV2():
		decoded, err := dtav2.DecodeFromEvent(ctx, event)
		if err != nil {
			return decodedWatcherEvent{}, fmt.Errorf("decode dta.v2 management watcher event: %w", err)
		}

		concrete, ok := decoded.ConcreteEvent.(*dtav2events.DistributorRequestProcessing)
		if !ok {
			return decodedWatcherEvent{}, fmt.Errorf(
				"expected *v2events.DistributorRequestProcessing, got %T",
				decoded.ConcreteEvent,
			)
		}

		return decodedWatcherEvent{
			EventName:               decoded.EventName().String(),
			RequestID:               concrete.RequestId,
			Shares:                  concrete.Shares,
			Amount:                  concrete.Amount,
			FundTokenSettlementAddr: v2SettlementAddr(decoded.FundTokenData),
			HasFundTokenData:        decoded.FundTokenData != nil,
			HasDistributorRequest:   decoded.DistributorRequest != nil,
			PaymentRequestCount:     len(decoded.PaymentRequests),
		}, nil

	default:
		return decodedWatcherEvent{}, fmt.Errorf("unsupported DTA_SERVICE: %q", cfg.DTAService)
	}
}

func decodeSettlementWatcherEvent(
	ctx context.Context,
	cfg *Config,
	event crecapi.Event,
) (decodedWatcherEvent, error) {
	if cfg == nil {
		return decodedWatcherEvent{}, fmt.Errorf("config is required")
	}

	switch {
	case cfg.UsesDTAServiceV1():
		decoded, err := dtav1.DecodeFromEvent(ctx, event)
		if err != nil {
			return decodedWatcherEvent{}, fmt.Errorf("decode dta.v1 settlement watcher event: %w", err)
		}

		concrete, ok := decoded.ConcreteEvent.(*dtav1events.DTASettlementOpened)
		if !ok {
			return decodedWatcherEvent{}, fmt.Errorf(
				"expected *v1events.DTASettlementOpened, got %T",
				decoded.ConcreteEvent,
			)
		}

		return decodedWatcherEvent{
			EventName:               decoded.EventName().String(),
			RequestID:               concrete.RequestId,
			Shares:                  concrete.Shares,
			Amount:                  concrete.Amount,
			FundTokenSettlementAddr: v1SettlementAddr(decoded.FundTokenData),
			HasFundTokenData:        decoded.FundTokenData != nil,
			HasDistributorRequest:   decoded.DistributorRequest != nil,
			PaymentRequestCount:     len(decoded.PaymentRequests),
		}, nil

	case cfg.UsesDTAServiceV2():
		decoded, err := dtav2.DecodeFromEvent(ctx, event)
		if err != nil {
			return decodedWatcherEvent{}, fmt.Errorf("decode dta.v2 settlement watcher event: %w", err)
		}

		concrete, ok := decoded.ConcreteEvent.(*dtav2events.DTASettlementOpened)
		if !ok {
			return decodedWatcherEvent{}, fmt.Errorf(
				"expected *v2events.DTASettlementOpened, got %T",
				decoded.ConcreteEvent,
			)
		}

		return decodedWatcherEvent{
			EventName:               decoded.EventName().String(),
			RequestID:               concrete.RequestId,
			Shares:                  concrete.Shares,
			Amount:                  concrete.Amount,
			FundTokenSettlementAddr: v2SettlementAddr(decoded.FundTokenData),
			HasFundTokenData:        decoded.FundTokenData != nil,
			HasDistributorRequest:   decoded.DistributorRequest != nil,
			PaymentRequestCount:     len(decoded.PaymentRequests),
		}, nil

	default:
		return decodedWatcherEvent{}, fmt.Errorf("unsupported DTA_SERVICE: %q", cfg.DTAService)
	}
}

func v1SettlementAddr(data *dtav1events.FundTokenData) common.Address {
	if data == nil {
		return common.Address{}
	}
	return data.DtaRequestSettlementAddr
}

func v2SettlementAddr(data *dtav2events.FundTokenData) common.Address {
	if data == nil {
		return common.Address{}
	}
	return data.DtaRequestSettlementAddr
}
