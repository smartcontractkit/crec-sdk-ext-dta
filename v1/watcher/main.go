//go:build wasip1

package main

import (
	"log/slog"

	"github.com/smartcontractkit/cre-sdk-go/cre"
	"github.com/smartcontractkit/cre-sdk-go/cre/wasm"

	wfcommon "github.com/smartcontractkit/crec-courier-service/workflows/common"
	wf "github.com/smartcontractkit/crec-sdk-ext-dta/v1/watcher/handler"
)

func main() {
	r := wasm.NewRunner(wfcommon.ParseWorkflowConfig)
	r.Run(func(cfg *wfcommon.Config, _ *slog.Logger, _ cre.SecretsProvider) (cre.Workflow[*wfcommon.Config], error) {
		return wfcommon.InitEventListenerWorkflow(cfg, wf.OnLog)
	})
}

