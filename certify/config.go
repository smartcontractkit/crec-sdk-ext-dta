package certify

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	defaultCourierURL              = "http://localhost:8088"
	defaultDTAService              = dtaServiceV2
	defaultWalletDeployTimeout     = 10 * time.Minute
	defaultWatcherActiveTimeout    = 10 * time.Minute
	defaultWatcherPropagationWait  = 2 * time.Minute
	defaultOperationConfirmTimeout = 5 * time.Minute
	defaultEventPollInterval       = 5 * time.Second
	defaultCleanupTimeout          = 2 * time.Minute
	defaultOperationDeadlineLead   = 30 * time.Minute
)

// Config is the runtime configuration for the deployed-CREC certification test.
type Config struct {
	Run bool

	CREOrgID  string
	CREAPIKey string
	CourierURL string

	ChainSelector           string
	AccountSignerPrivateKey string

	CRECValidSigners        []string
	MinRequiredSignatures   int

	DTAService                  string
	DTARequestManagementAddress string
	DTARequestSettlementAddress string

	WalletDeployTimeout     time.Duration
	WatcherActiveTimeout    time.Duration
	WatcherPropagationWait  time.Duration
	OperationConfirmTimeout time.Duration
	EventPollInterval       time.Duration
	CleanupTimeout          time.Duration
	OperationDeadlineLead   time.Duration
}

// LoadConfig loads certification configuration from environment variables.
// If RUN_CREC_CERTIFICATION is false, only that flag is parsed and the returned
// config contains defaults so the package safely skips under normal `go test ./...`.
func LoadConfig() (*Config, error) {
	run, err := envBool("RUN_CREC_CERTIFICATION", false)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Run:                     run,
		CourierURL:              defaultCourierURL,
		DTAService:              defaultDTAService,
		WalletDeployTimeout:     defaultWalletDeployTimeout,
		WatcherActiveTimeout:    defaultWatcherActiveTimeout,
		WatcherPropagationWait:  defaultWatcherPropagationWait,
		OperationConfirmTimeout: defaultOperationConfirmTimeout,
		EventPollInterval:       defaultEventPollInterval,
		CleanupTimeout:          defaultCleanupTimeout,
		OperationDeadlineLead:   defaultOperationDeadlineLead,
	}

	if !cfg.Run {
		return cfg, nil
	}

	cfg.CREOrgID = strings.TrimSpace(os.Getenv("CRE_ORG_ID"))
	cfg.CREAPIKey = strings.TrimSpace(os.Getenv("CRE_API_KEY"))
	cfg.CourierURL = envString("COURIER_URL", defaultCourierURL)

	cfg.ChainSelector = strings.TrimSpace(os.Getenv("CHAIN_SELECTOR"))
	cfg.AccountSignerPrivateKey = strings.TrimSpace(os.Getenv("ACCOUNT_SIGNER_PRIVATE_KEY"))

	cfg.CRECValidSigners = envCSV("CREC_VALID_SIGNERS")

	cfg.MinRequiredSignatures, err = envInt("MIN_REQUIRED_SIGNATURES", 2)
	if err != nil {
		return nil, err
	}

	cfg.DTAService = normalizeDTAService(envString("DTA_SERVICE", defaultDTAService))
	cfg.DTARequestManagementAddress = strings.TrimSpace(os.Getenv("DTA_REQUEST_MANAGEMENT_ADDRESS"))
	cfg.DTARequestSettlementAddress = strings.TrimSpace(os.Getenv("DTA_REQUEST_SETTLEMENT_ADDRESS"))

	cfg.WalletDeployTimeout, err = envDuration("WALLET_DEPLOY_TIMEOUT", defaultWalletDeployTimeout)
	if err != nil {
		return nil, err
	}
	cfg.WatcherActiveTimeout, err = envDuration("WATCHER_ACTIVE_TIMEOUT", defaultWatcherActiveTimeout)
	if err != nil {
		return nil, err
	}
	cfg.WatcherPropagationWait, err = envDuration("WATCHER_PROPAGATION_WAIT", defaultWatcherPropagationWait)
	if err != nil {
		return nil, err
	}
	cfg.OperationConfirmTimeout, err = envDuration("OPERATION_CONFIRM_TIMEOUT", defaultOperationConfirmTimeout)
	if err != nil {
		return nil, err
	}
	cfg.EventPollInterval, err = envDuration("EVENT_POLL_INTERVAL", defaultEventPollInterval)
	if err != nil {
		return nil, err
	}
	cfg.CleanupTimeout, err = envDuration("CLEANUP_TIMEOUT", defaultCleanupTimeout)
	if err != nil {
		return nil, err
	}
	cfg.OperationDeadlineLead, err = envDuration("OPERATION_DEADLINE_LEAD", defaultOperationDeadlineLead)
	if err != nil {
		return nil, err
	}

	if len(cfg.CRECValidSigners) == 0 {
		cfg.MinRequiredSignatures = 0
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// Validate validates certification configuration.
func (c *Config) Validate() error {
	var missing []string

	if c.CREOrgID == "" {
		missing = append(missing, "CRE_ORG_ID")
	}
	if c.CREAPIKey == "" {
		missing = append(missing, "CRE_API_KEY")
	}
	if c.ChainSelector == "" {
		missing = append(missing, "CHAIN_SELECTOR")
	}
	if c.AccountSignerPrivateKey == "" {
		missing = append(missing, "ACCOUNT_SIGNER_PRIVATE_KEY")
	}
	if c.DTARequestManagementAddress == "" {
		missing = append(missing, "DTA_REQUEST_MANAGEMENT_ADDRESS")
	}
	if c.DTARequestSettlementAddress == "" {
		missing = append(missing, "DTA_REQUEST_SETTLEMENT_ADDRESS")
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required certification env vars: %s", strings.Join(missing, ", "))
	}

	if _, ok := new(big.Int).SetString(c.ChainSelector, 10); !ok {
		return fmt.Errorf("invalid CHAIN_SELECTOR: %q", c.ChainSelector)
	}

	if !common.IsHexAddress(c.DTARequestManagementAddress) {
		return fmt.Errorf("invalid DTA_REQUEST_MANAGEMENT_ADDRESS: %q", c.DTARequestManagementAddress)
	}
	if !common.IsHexAddress(c.DTARequestSettlementAddress) {
		return fmt.Errorf("invalid DTA_REQUEST_SETTLEMENT_ADDRESS: %q", c.DTARequestSettlementAddress)
	}

	for _, signer := range c.CRECValidSigners {
		if !common.IsHexAddress(signer) {
			return fmt.Errorf("invalid CREC_VALID_SIGNERS entry: %q", signer)
		}
	}

	if c.MinRequiredSignatures < 0 {
		return fmt.Errorf("MIN_REQUIRED_SIGNATURES must be >= 0")
	}
	if len(c.CRECValidSigners) > 0 && c.MinRequiredSignatures > len(c.CRECValidSigners) {
		return fmt.Errorf(
			"MIN_REQUIRED_SIGNATURES (%d) exceeds number of CREC_VALID_SIGNERS (%d)",
			c.MinRequiredSignatures,
			len(c.CRECValidSigners),
		)
	}

	if strings.TrimSpace(c.CourierURL) == "" {
		return fmt.Errorf("COURIER_URL must not be empty")
	}
	c.DTAService = normalizeDTAService(c.DTAService)
	switch c.DTAService {
	case dtaServiceV1, dtaServiceV2:
	default:
		return fmt.Errorf("DTA_SERVICE must be one of %s, got %q", strings.Join(supportedDTAServices(), ", "), c.DTAService)
	}

	if c.WalletDeployTimeout <= 0 {
		return fmt.Errorf("WALLET_DEPLOY_TIMEOUT must be > 0")
	}
	if c.WatcherActiveTimeout <= 0 {
		return fmt.Errorf("WATCHER_ACTIVE_TIMEOUT must be > 0")
	}
	if c.OperationConfirmTimeout <= 0 {
		return fmt.Errorf("OPERATION_CONFIRM_TIMEOUT must be > 0")
	}
	if c.EventPollInterval <= 0 {
		return fmt.Errorf("EVENT_POLL_INTERVAL must be > 0")
	}
	if c.CleanupTimeout <= 0 {
		return fmt.Errorf("CLEANUP_TIMEOUT must be > 0")
	}
	if c.WatcherPropagationWait < 0 {
		return fmt.Errorf("WATCHER_PROPAGATION_WAIT must be >= 0")
	}
	if c.OperationDeadlineLead < 0 {
		return fmt.Errorf("OPERATION_DEADLINE_LEAD must be >= 0")
	}

	return nil
}

func (c *Config) UsesDTAServiceV1() bool {
	return c != nil && normalizeDTAService(c.DTAService) == dtaServiceV1
}

func (c *Config) UsesDTAServiceV2() bool {
	return c != nil && normalizeDTAService(c.DTAService) == dtaServiceV2
}

func supportedDTAServices() []string {
	return []string{dtaServiceV1, dtaServiceV2}
}

func normalizeDTAService(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func envString(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func envCSV(key string) []string {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return nil
	}

	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}

func envBool(key string, fallback bool) (bool, error) {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return fallback, nil
	}

	value, err := strconv.ParseBool(raw)
	if err != nil {
		return false, fmt.Errorf("parse %s as bool: %w", key, err)
	}
	return value, nil
}

func envInt(key string, fallback int) (int, error) {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return fallback, nil
	}

	value, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("parse %s as int: %w", key, err)
	}
	return value, nil
}

func envDuration(key string, fallback time.Duration) (time.Duration, error) {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return fallback, nil
	}

	value, err := time.ParseDuration(raw)
	if err != nil {
		return 0, fmt.Errorf("parse %s as duration: %w", key, err)
	}
	return value, nil
}
