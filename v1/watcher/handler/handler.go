package handler

import (
	"fmt"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/cre"

	wfcommon "github.com/smartcontractkit/crec-courier-service/workflows/common"
)

// OnLog processes EVM log events from DTA contracts.
// It decodes event parameters, composes workflow metadata, and posts signed events.
func OnLog(cfg *wfcommon.Config, rt cre.Runtime, payload *evm.Log) (string, error) {
	rt.Logger().Info("OnLog", "payload", fmt.Sprintf("%+v", payload))
	ts := wfcommon.GetBlockTimestamp(rt, wfcommon.EnsureChainSelector(cfg, "3379446385462418246"), payload.BlockNumber)

	abiJSON, _ := wfcommon.GetContractABI(cfg, cfg.DetectEventTriggerConfig.ContractName)
	params, _ := wfcommon.DecodeEventParams(abiJSON, cfg.DetectEventTriggerConfig.ContractEventName, payload)

	metadata := wfcommon.ComposeWorkflowEventMetadata(
		"watcher-dta-v1",
		cfg.ChainID,
		cfg.DetectEventTriggerConfig.ContractEventName,
		params,
	)
	rt.Logger().Info("ComposeWorkflowEventMetadata", "metadata", fmt.Sprintf("%+v", metadata))

	pre, err := wfcommon.BuildAndSignEventEnvelope(
		cfg.Service,
		cfg.DetectEventTriggerConfig.ContractEventName,
		cfg.DetectEventTriggerConfig.ContractAddress,
		abiJSON,
		cfg.ChainID,
		wfcommon.PBToUint64(payload.BlockNumber),
		uint64(payload.Index),
		wfcommon.TxHashFromLog(payload),
		ts,
		params,
		metadata,
	)
	if err != nil {
		return "", err
	}

	return wfcommon.PostSignedEvent(
		cfg,
		rt,
		cfg.DetectEventTriggerConfig.ContractEventName,
		cfg.DetectEventTriggerConfig.ContractAddress,
		pre,
	)
}

