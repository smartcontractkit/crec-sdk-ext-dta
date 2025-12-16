package handler

import (
	"fmt"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/cre"
	workflows "github.com/smartcontractkit/cre-workflow-utils"
)

// OnLog processes EVM log events from DTA contracts.
// It decodes event parameters, composes workflow metadata, and posts signed events.
func OnLog(cfg *workflows.Config, rt cre.Runtime, payload *evm.Log) (string, error) {
	rt.Logger().Info("OnLog", "payload", fmt.Sprintf("%+v", payload))
	ts := workflows.GetBlockTimestamp(rt, workflows.EnsureChainSelector(cfg, cfg.ChainSelector), payload.BlockNumber)

	abiJSON, _ := workflows.GetContractABI(cfg, cfg.DetectEventTriggerConfig.ContractName)
	params, _ := workflows.DecodeEventParams(abiJSON, cfg.DetectEventTriggerConfig.ContractEventName, payload)

	metadata := workflows.ComposeWorkflowEventMetadata(
		cfg.WorkflowName,
		cfg.ChainID,
		cfg.DetectEventTriggerConfig.ContractEventName,
		params,
	)
	rt.Logger().Info("ComposeWorkflowEventMetadata", "metadata", fmt.Sprintf("%+v", metadata))

	pre, err := workflows.BuildAndHashEventEnvelope(
		cfg.Service,
		cfg.DetectEventTriggerConfig.ContractEventName,
		cfg.DetectEventTriggerConfig.ContractAddress,
		abiJSON,
		cfg.ChainID,
		workflows.PBToUint64(payload.BlockNumber),
		uint64(payload.Index),
		workflows.TxHashFromLog(payload),
		ts,
		params,
		metadata,
	)
	if err != nil {
		return "", err
	}

	return workflows.PostSignedEvent(
		cfg,
		rt,
		cfg.DetectEventTriggerConfig.ContractEventName,
		cfg.DetectEventTriggerConfig.ContractAddress,
		pre,
	)
}
