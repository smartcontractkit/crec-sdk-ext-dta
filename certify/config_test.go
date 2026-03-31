package certify

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConfigValidateSupportsVersionedDTAServices(t *testing.T) {
	tests := []struct {
		name        string
		service     string
		wantService string
		wantErr     string
	}{
		{
			name:        "v1 canonical service",
			service:     dtaServiceV1,
			wantService: dtaServiceV1,
		},
		{
			name:        "v2 mixed case normalizes",
			service:     " DTA.V2 ",
			wantService: dtaServiceV2,
		},
		{
			name:    "unknown service rejected",
			service: "dta.v3",
			wantErr: "DTA_SERVICE must be one of dta.v1, dta.v2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := validCertificationConfig()
			cfg.DTAService = tt.service

			err := cfg.Validate()
			if tt.wantErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.wantErr)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.wantService, cfg.DTAService)
		})
	}
}

func TestConfigServiceHelpersNormalizeCase(t *testing.T) {
	cfg := &Config{DTAService: " DTA.V1 "}
	require.True(t, cfg.UsesDTAServiceV1())
	require.False(t, cfg.UsesDTAServiceV2())

	cfg.DTAService = "dta.v2"
	require.False(t, cfg.UsesDTAServiceV1())
	require.True(t, cfg.UsesDTAServiceV2())
}

func validCertificationConfig() *Config {
	return &Config{
		Run:                         true,
		CREOrgID:                    "test-org",
		CREAPIKey:                   "test-api-key",
		CourierURL:                  defaultCourierURL,
		ChainSelector:               "16015286601757825753",
		AccountSignerPrivateKey:     "deadbeef",
		DTAService:                  defaultDTAService,
		DTARequestManagementAddress: "0x1111111111111111111111111111111111111111",
		DTARequestSettlementAddress: "0x2222222222222222222222222222222222222222",
		WalletDeployTimeout:         time.Minute,
		WatcherActiveTimeout:        time.Minute,
		WatcherPropagationWait:      0,
		OperationConfirmTimeout:     time.Minute,
		EventPollInterval:           time.Second,
		CleanupTimeout:              time.Minute,
		OperationDeadlineLead:       time.Minute,
	}
}
