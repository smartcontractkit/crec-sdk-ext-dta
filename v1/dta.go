// Package v1 provides CREC SDK extension for DTA (Digital Token Asset) v1.0 operations.
//
// This version works with DTARequestManagement and DTARequestSettlement contracts.
//
// # Usage
//
//	ext, err := v1.New(&v1.Options{
//		DTARequestManagementAddress: "0x...",
//		DTARequestSettlementAddress: "0x...",
//		AccountAddress:              "0x...",
//	})
package v1

import (
	"fmt"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
)

// Options defines the configuration for creating a new CREC DTA v1 extension.
type Options struct {
	// Logger is an optional logger instance. If nil, a default nop logger is used.
	Logger *slog.Logger

	// DTARequestManagementAddress is the address of the DTARequestManagement contract.
	DTARequestManagementAddress string

	// DTARequestSettlementAddress is the address of the DTARequestSettlement contract.
	DTARequestSettlementAddress string

	// AccountAddress is the address of the account performing the DTA operations.
	AccountAddress string
}

// validate checks that all required options are valid.
func (o *Options) validate() error {
	if !common.IsHexAddress(o.DTARequestManagementAddress) {
		return fmt.Errorf("invalid DTARequestManagementAddress: %q", o.DTARequestManagementAddress)
	}
	if !common.IsHexAddress(o.DTARequestSettlementAddress) {
		return fmt.Errorf("invalid DTARequestSettlementAddress: %q", o.DTARequestSettlementAddress)
	}
	if !common.IsHexAddress(o.AccountAddress) {
		return fmt.Errorf("invalid AccountAddress: %q", o.AccountAddress)
	}
	return nil
}

// Extension provides methods for preparing DTA v1 operations.
type Extension struct {
	logger                      *slog.Logger
	dtaRequestManagementAddress common.Address
	dtaRequestSettlementAddress common.Address
	accountAddress              common.Address
}

// New creates a new CREC DTA v1 extension with the provided options.
// Returns a pointer to the Extension and an error if any issues occur during initialization.
func New(opts *Options) (*Extension, error) {
	if opts == nil {
		return nil, fmt.Errorf("options is required")
	}

	if err := opts.validate(); err != nil {
		return nil, fmt.Errorf("invalid options: %w", err)
	}

	logger := opts.Logger
	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("creating CREC DTA v1 extension")

	return &Extension{
		logger:                      logger,
		dtaRequestManagementAddress: common.HexToAddress(opts.DTARequestManagementAddress),
		dtaRequestSettlementAddress: common.HexToAddress(opts.DTARequestSettlementAddress),
		accountAddress:              common.HexToAddress(opts.AccountAddress),
	}, nil
}
