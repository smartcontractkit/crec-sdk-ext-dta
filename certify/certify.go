package certify

import (
	"context"
	crand "crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	crecapi "github.com/smartcontractkit/crec-api-go/client"
	crec "github.com/smartcontractkit/crec-sdk"
	"github.com/smartcontractkit/crec-sdk/channels"
	sdkevents "github.com/smartcontractkit/crec-sdk/events"
	"github.com/smartcontractkit/crec-sdk/transact"
	"github.com/smartcontractkit/crec-sdk/transact/signer/local"
	transactTypes "github.com/smartcontractkit/crec-sdk/transact/types"
	"github.com/smartcontractkit/crec-sdk/wallets"
	"github.com/smartcontractkit/crec-sdk/watchers"
)

const (
	apiCallTimeout = time.Minute

	mockManagementABIJSON = `[
		{
			"type": "function",
			"name": "emitDistributorRequestProcessing",
			"inputs": [],
			"outputs": [
				{"name": "requestId", "type": "bytes32"}
			],
			"stateMutability": "nonpayable"
		}
	]`

	mockSettlementABIJSON = `[
		{
			"type": "function",
			"name": "emitDTASettlementOpened",
			"inputs": [],
			"outputs": [
				{"name": "requestId", "type": "bytes32"}
			],
			"stateMutability": "nonpayable"
		}
	]`
)

// RunCertificationScenario runs a deployed-CREC certification flow for the configured DTA extension service.
// Use version-matched management mocks:
//   - dta.v1 -> test/MockDTARequestManagement.sol
//   - dta.v2 -> test/MockDTARequestManagementV2.sol
// The settlement mock for both services is test/MockDTARequestSettlement.sol,
// and it should be wired to the selected management mock address.
func RunCertificationScenario(t *testing.T, cfg *Config) {
	t.Helper()

	t.Logf(
		"running DTA certification against %s (service=%s, chain_selector=%s, org=%s)",
		cfg.CourierURL,
		cfg.DTAService,
		cfg.ChainSelector,
		cfg.CREOrgID,
	)

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(cfg.AccountSignerPrivateKey, "0x"))
	require.NoError(t, err, "failed to parse ACCOUNT_SIGNER_PRIVATE_KEY")

	operationSigner := local.NewSigner(privateKey)
	signerAddress := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	apiClient, err := crec.NewAPIClient(cfg.CourierURL, cfg.CREAPIKey)
	require.NoError(t, err, "failed to create CREC API client")

	channelEventsClient, err := sdkevents.NewClient(&sdkevents.Options{
		CRECClient:            apiClient,
		MinRequiredSignatures: cfg.MinRequiredSignatures,
		ValidSigners:          cfg.CRECValidSigners,
	})
	require.NoError(t, err, "failed to create events client")

	transactClient, err := transact.NewClient(&transact.Options{
		CRECClient: apiClient,
	})
	require.NoError(t, err, "failed to create transact client")

	channelsClient, err := channels.NewClient(&channels.Options{
		APIClient: apiClient,
	})
	require.NoError(t, err, "failed to create channels client")

	watchersClient, err := watchers.NewClient(&watchers.Options{
		APIClient: apiClient,
	})
	require.NoError(t, err, "failed to create watchers client")

	walletsClient, err := wallets.NewClient(&wallets.Options{
		APIClient: apiClient,
	})
	require.NoError(t, err, "failed to create wallets client")

	channelID := createChannel(t, channelsClient)
	t.Cleanup(func() {
		archiveChannel(t, cfg, channelsClient, channelID)
	})

	walletID, walletAddress := createWallet(t, cfg, walletsClient, signerAddress)
	t.Cleanup(func() {
		archiveWallet(t, cfg, walletsClient, walletID)
	})

	managementWatcherID := createWatcher(
		t,
		cfg,
		watchersClient,
		channelID,
		cfg.DTARequestManagementAddress,
		"DistributorRequestProcessing",
	)
	t.Cleanup(func() {
		archiveWatcher(t, cfg, watchersClient, channelID, managementWatcherID)
	})

	settlementWatcherID := createWatcher(
		t,
		cfg,
		watchersClient,
		channelID,
		cfg.DTARequestSettlementAddress,
		"DTASettlementOpened",
	)
	t.Cleanup(func() {
		archiveWatcher(t, cfg, watchersClient, channelID, settlementWatcherID)
	})

	// Local SDK deadline probe: this makes the certification explicitly check that the
	// shipped DTA operations SDK emits and clones deadlines correctly.
	sdkDeadline := big.NewInt(time.Now().Add(15 * time.Minute).Unix())
	operationsExt, err := newRegisterFundAdminOperationBuilder(cfg, walletAddress, sdkDeadline)
	require.NoError(t, err, "failed to construct DTA operations extension")

	deadlineProbeOp1, err := operationsExt.PrepareRegisterFundAdminOperation()
	require.NoError(t, err, "failed to build DTA deadline probe operation")
	require.NotNil(t, deadlineProbeOp1.Deadline, "prepared DTA operation should always include deadline")
	require.Equal(t, sdkDeadline.String(), deadlineProbeOp1.Deadline.String(), "configured deadline must be propagated")

	deadlineProbeOp1.Deadline.SetInt64(0)

	deadlineProbeOp2, err := operationsExt.PrepareRegisterFundAdminOperation()
	require.NoError(t, err, "failed to build second DTA deadline probe operation")
	require.NotNil(t, deadlineProbeOp2.Deadline, "prepared DTA operation should always include deadline")
	require.Equal(t, sdkDeadline.String(), deadlineProbeOp2.Deadline.String(), "prepared operations must clone the default deadline")

	if cfg.WatcherPropagationWait > 0 {
		t.Logf("waiting %s for watcher propagation and polling initialization", cfg.WatcherPropagationWait)
		time.Sleep(cfg.WatcherPropagationWait)
	}

	var offset int64

	managementCalldata, err := packNoArgMethod(mockManagementABIJSON, "emitDistributorRequestProcessing")
	require.NoError(t, err, "failed to pack management mock trigger calldata")

	managementOp := buildTriggerOperation(
		walletAddress,
		cfg.DTARequestManagementAddress,
		managementCalldata,
		cfg.OperationDeadlineLead,
	)
	require.NotNil(t, managementOp.Deadline, "trigger operation should always include a deadline")
	if cfg.OperationDeadlineLead > 0 {
		require.True(t, managementOp.Deadline.Sign() > 0, "non-zero deadline lead should produce non-zero deadline")
	}

	managementSendCtx, cancel := context.WithTimeout(context.Background(), apiCallTimeout)
	managementHash, managementSignature, err := transactClient.SignOperation(
		managementSendCtx,
		managementOp,
		operationSigner,
		cfg.ChainSelector,
	)
	cancel()
	require.NoError(t, err, "failed to sign management trigger operation")
	t.Logf("signed management trigger operation hash: %s", managementHash.Hex())

	managementSendCtx, cancel = context.WithTimeout(context.Background(), apiCallTimeout)
	sentManagementOperation, err := transactClient.SendSignedOperation(
		managementSendCtx,
		channelID,
		managementOp,
		managementSignature,
		cfg.ChainSelector,
	)
	cancel()
	require.NoError(t, err, "failed to send management trigger operation")

	managementWaitCtx, cancel := context.WithTimeout(context.Background(), cfg.OperationConfirmTimeout)
	waitForOperationConfirmed(
		managementWaitCtx,
		t,
		transactClient,
		channelID,
		sentManagementOperation.OperationId,
		cfg.EventPollInterval,
	)
	cancel()

	managementEventCtx, cancel := context.WithTimeout(context.Background(), cfg.OperationConfirmTimeout)
	managementWatcherEvent := waitForWatcherEvent(
		managementEventCtx,
		t,
		channelEventsClient,
		channelID,
		managementWatcherID,
		&offset,
		cfg.EventPollInterval,
	)
	cancel()

	verifyWatcherEvent(t, cfg, channelEventsClient, managementWatcherEvent)

	decodedManagement, err := decodeManagementWatcherEvent(context.Background(), cfg, *managementWatcherEvent)
	require.NoError(t, err, "failed to decode DistributorRequestProcessing watcher event")
	require.Equal(t, managementEventDistributorRequestProcessing, decodedManagement.EventName)
	require.True(t, decodedManagement.HasFundTokenData, "DistributorRequestProcessing should include fund token enrichment")
	require.True(t, decodedManagement.HasDistributorRequest, "DistributorRequestProcessing should include distributor request enrichment")
	require.Equal(t, 0, decodedManagement.PaymentRequestCount, "management watcher path should not emit payment requests")
	require.NotNil(t, decodedManagement.Shares, "DistributorRequestProcessing should include shares")
	require.NotNil(t, decodedManagement.Amount, "DistributorRequestProcessing should include amount")
	require.True(t, decodedManagement.Shares.Sign() > 0, "mock shares should be > 0")
	require.True(t, decodedManagement.Amount.Sign() > 0, "mock amount should be > 0")
	require.True(
		t,
		strings.EqualFold(decodedManagement.FundTokenSettlementAddr.Hex(), cfg.DTARequestSettlementAddress),
		"fund token enrichment should point to configured settlement mock",
	)
	t.Logf("management watcher event decoded successfully for request %s", decodedManagement.RequestID.Hex())

	settlementCalldata, err := packNoArgMethod(mockSettlementABIJSON, "emitDTASettlementOpened")
	require.NoError(t, err, "failed to pack settlement mock trigger calldata")

	settlementOp := buildTriggerOperation(
		walletAddress,
		cfg.DTARequestSettlementAddress,
		settlementCalldata,
		cfg.OperationDeadlineLead,
	)
	require.NotNil(t, settlementOp.Deadline, "trigger operation should always include a deadline")
	if cfg.OperationDeadlineLead > 0 {
		require.True(t, settlementOp.Deadline.Sign() > 0, "non-zero deadline lead should produce non-zero deadline")
	}

	settlementSendCtx, cancel := context.WithTimeout(context.Background(), apiCallTimeout)
	settlementHash, settlementSignature, err := transactClient.SignOperation(
		settlementSendCtx,
		settlementOp,
		operationSigner,
		cfg.ChainSelector,
	)
	cancel()
	require.NoError(t, err, "failed to sign settlement trigger operation")
	t.Logf("signed settlement trigger operation hash: %s", settlementHash.Hex())

	settlementSendCtx, cancel = context.WithTimeout(context.Background(), apiCallTimeout)
	sentSettlementOperation, err := transactClient.SendSignedOperation(
		settlementSendCtx,
		channelID,
		settlementOp,
		settlementSignature,
		cfg.ChainSelector,
	)
	cancel()
	require.NoError(t, err, "failed to send settlement trigger operation")

	settlementWaitCtx, cancel := context.WithTimeout(context.Background(), cfg.OperationConfirmTimeout)
	waitForOperationConfirmed(
		settlementWaitCtx,
		t,
		transactClient,
		channelID,
		sentSettlementOperation.OperationId,
		cfg.EventPollInterval,
	)
	cancel()

	settlementEventCtx, cancel := context.WithTimeout(context.Background(), cfg.OperationConfirmTimeout)
	settlementWatcherEvent := waitForWatcherEvent(
		settlementEventCtx,
		t,
		channelEventsClient,
		channelID,
		settlementWatcherID,
		&offset,
		cfg.EventPollInterval,
	)
	cancel()

	verifyWatcherEvent(t, cfg, channelEventsClient, settlementWatcherEvent)

	decodedSettlement, err := decodeSettlementWatcherEvent(context.Background(), cfg, *settlementWatcherEvent)
	require.NoError(t, err, "failed to decode DTASettlementOpened watcher event")
	require.Equal(t, settlementEventDTASettlementOpened, decodedSettlement.EventName)
	require.True(t, decodedSettlement.HasFundTokenData, "DTASettlementOpened should include fund token enrichment")
	require.True(t, decodedSettlement.HasDistributorRequest, "DTASettlementOpened should include distributor request enrichment")
	require.Equal(t, 1, decodedSettlement.PaymentRequestCount, "DTASettlementOpened should emit exactly one payment request")
	require.NotNil(t, decodedSettlement.Shares, "DTASettlementOpened should include shares")
	require.NotNil(t, decodedSettlement.Amount, "DTASettlementOpened should include amount")
	require.True(t, decodedSettlement.Shares.Sign() > 0, "mock shares should be > 0")
	require.True(t, decodedSettlement.Amount.Sign() > 0, "mock amount should be > 0")
	require.True(
		t,
		strings.EqualFold(decodedSettlement.FundTokenSettlementAddr.Hex(), cfg.DTARequestSettlementAddress),
		"fund token enrichment should point to configured settlement mock",
	)
	t.Logf("settlement watcher event decoded successfully for request %s", decodedSettlement.RequestID.Hex())
}

func createChannel(t *testing.T, channelsClient *channels.Client) uuid.UUID {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), apiCallTimeout)
	defer cancel()

	channel, err := channelsClient.Create(ctx, channels.CreateInput{
		Name: uniqueName("dta-certify-channel"),
	})
	require.NoError(t, err, "failed to create certification channel")

	t.Logf("created channel %s", channel.ChannelId.String())

	return channel.ChannelId
}

func createWallet(
	t *testing.T,
	cfg *Config,
	walletsClient *wallets.Client,
	signerAddress string,
) (uuid.UUID, string) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), apiCallTimeout)
	defer cancel()

	allowedSigners := []string{signerAddress}

	wallet, err := walletsClient.Create(ctx, wallets.CreateInput{
		Name:                uniqueName("dta-certify-wallet"),
		ChainSelector:       cfg.ChainSelector,
		WalletOwnerAddress:  signerAddress,
		WalletType:          crecapi.Ecdsa,
		AllowedEcdsaSigners: &allowedSigners,
	})
	require.NoError(t, err, "failed to create certification wallet")

	t.Logf("created wallet %s, waiting for deployment", wallet.WalletId.String())

	waitCtx, waitCancel := context.WithTimeout(context.Background(), cfg.WalletDeployTimeout)
	defer waitCancel()

	walletAddress := waitForWalletStatus(
		waitCtx,
		t,
		walletsClient,
		wallet.WalletId,
		"deployed",
		cfg.EventPollInterval,
	)
	require.NotEmpty(t, walletAddress, "wallet should have address after deployment")

	t.Logf("wallet %s deployed at %s", wallet.WalletId.String(), walletAddress)

	return wallet.WalletId, walletAddress
}

func createWatcher(
	t *testing.T,
	cfg *Config,
	watchersClient *watchers.Client,
	channelID uuid.UUID,
	contractAddress string,
	eventName string,
) uuid.UUID {
	t.Helper()

	name := uniqueName("dta-certify-watcher-" + strings.ToLower(eventName))

	ctx, cancel := context.WithTimeout(context.Background(), apiCallTimeout)
	defer cancel()

	watcher, err := watchersClient.CreateWithService(
		ctx,
		channelID,
		watchers.CreateWithServiceInput{
			Name:          &name,
			ChainSelector: cfg.ChainSelector,
			Address:       contractAddress,
			Service:       cfg.DTAService,
			Events:        []string{eventName},
		},
	)
	require.NoError(t, err, "failed to create watcher for %s", eventName)

	t.Logf("created watcher %s for %s on %s", watcher.WatcherId.String(), eventName, contractAddress)

	waitCtx, waitCancel := context.WithTimeout(context.Background(), cfg.WatcherActiveTimeout)
	defer waitCancel()

	waitForWatcherStatus(
		waitCtx,
		t,
		watchersClient,
		channelID,
		watcher.WatcherId,
		"active",
		cfg.EventPollInterval,
	)

	t.Logf("watcher %s for %s is active", watcher.WatcherId.String(), eventName)

	return watcher.WatcherId
}

func waitForWalletStatus(
	ctx context.Context,
	t *testing.T,
	walletsClient *wallets.Client,
	walletID uuid.UUID,
	expectedStatus string,
	pollInterval time.Duration,
) string {
	t.Helper()

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	expected := strings.ToLower(expectedStatus)
	lastStatus := "unknown"

	for {
		wallet, err := walletsClient.Get(ctx, walletID)
		if err == nil {
			lastStatus = statusString(wallet.Status)

			if lastStatus == expected {
				return wallet.Address
			}

			if lastStatus == "failed" {
				require.FailNowf(
					t,
					"wallet deployment failed",
					"wallet %s entered failed status before reaching %s",
					walletID.String(),
					expectedStatus,
				)
				return ""
			}
		} else {
			t.Logf("wallet %s get error while waiting for %s: %v", walletID.String(), expectedStatus, err)
		}

		select {
		case <-ctx.Done():
			require.FailNowf(
				t,
				"wallet deployment timeout",
				"wallet %s did not reach %s before timeout; last status=%s: %v",
				walletID.String(),
				expectedStatus,
				lastStatus,
				ctx.Err(),
			)
			return ""
		case <-ticker.C:
		}
	}
}

func waitForWatcherStatus(
	ctx context.Context,
	t *testing.T,
	watchersClient *watchers.Client,
	channelID uuid.UUID,
	watcherID uuid.UUID,
	expectedStatus string,
	pollInterval time.Duration,
) {
	t.Helper()

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	expected := strings.ToLower(expectedStatus)
	lastStatus := "unknown"

	for {
		watcher, err := watchersClient.Get(ctx, channelID, watcherID)
		if err == nil {
			lastStatus = statusString(watcher.Status)

			if lastStatus == expected {
				return
			}

			if lastStatus == "failed" || lastStatus == "error" {
				require.FailNowf(
					t,
					"watcher activation failed",
					"watcher %s entered terminal status %s before reaching %s",
					watcherID.String(),
					lastStatus,
					expectedStatus,
				)
				return
			}
		} else {
			t.Logf("watcher %s get error while waiting for %s: %v", watcherID.String(), expectedStatus, err)
		}

		select {
		case <-ctx.Done():
			require.FailNowf(
				t,
				"watcher activation timeout",
				"watcher %s did not reach %s before timeout; last status=%s: %v",
				watcherID.String(),
				expectedStatus,
				lastStatus,
				ctx.Err(),
			)
			return
		case <-ticker.C:
		}
	}
}

func waitForOperationConfirmed(
	ctx context.Context,
	t *testing.T,
	transactClient *transact.Client,
	channelID uuid.UUID,
	operationID uuid.UUID,
	pollInterval time.Duration,
) {
	t.Helper()

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	lastStatus := "unknown"

	for {
		operation, err := transactClient.GetOperation(ctx, channelID, operationID)
		if err == nil {
			lastStatus = statusString(operation.Status)

			if lastStatus == "confirmed" {
				return
			}

			if lastStatus == "failed" {
				require.FailNowf(
					t,
					"operation failed",
					"operation %s entered failed status",
					operationID.String(),
				)
				return
			}
		} else {
			t.Logf("operation %s get error while waiting for confirmation: %v", operationID.String(), err)
		}

		select {
		case <-ctx.Done():
			require.FailNowf(
				t,
				"operation confirmation timeout",
				"operation %s did not reach confirmed before timeout; last status=%s: %v",
				operationID.String(),
				lastStatus,
				ctx.Err(),
			)
			return
		case <-ticker.C:
		}
	}
}

func waitForWatcherEvent(
	ctx context.Context,
	t *testing.T,
	eventsClient *sdkevents.Client,
	channelID uuid.UUID,
	watcherID uuid.UUID,
	offset *int64,
	pollInterval time.Duration,
) *crecapi.Event {
	t.Helper()

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		limit := 100
		params := crecapi.GetChannelsChannelIdEventsParams{
			Offset: offset,
			Limit:  &limit,
		}

		polledEvents, _, err := eventsClient.Poll(ctx, channelID, &params)
		if err == nil {
			for i := range polledEvents {
				current := polledEvents[i]
				*offset = *offset + 1

				if statusString(current.Headers.Type) != "watcher.event" {
					continue
				}

				payload, err := current.Payload.AsWatcherEventPayload()
				if err != nil {
					t.Logf("failed to parse watcher payload at offset %d: %v", *offset-1, err)
					continue
				}

				if payload.WatcherId == watcherID.String() {
					return &current
				}
			}
		} else {
			t.Logf("poll channel events failed while waiting for watcher %s: %v", watcherID.String(), err)
		}

		select {
		case <-ctx.Done():
			require.FailNowf(
				t,
				"watcher event timeout",
				"did not receive watcher.event for watcher %s before timeout: %v",
				watcherID.String(),
				ctx.Err(),
			)
			return nil
		case <-ticker.C:
		}
	}
}

func verifyWatcherEvent(
	t *testing.T,
	cfg *Config,
	eventsClient *sdkevents.Client,
	watcherEvent *crecapi.Event,
) {
	t.Helper()

	if watcherEvent == nil {
		require.FailNow(t, "watcher event is nil")
		return
	}

	if cfg.MinRequiredSignatures == 0 || len(cfg.CRECValidSigners) == 0 {
		t.Log("CREC_VALID_SIGNERS not configured; skipping watcher.event signature verification")
		return
	}

	valid, err := eventsClient.Verify(watcherEvent)
	require.NoError(t, err, "watcher.event signature verification returned error")
	require.True(t, valid, "watcher.event signature verification failed")
}

func archiveWatcher(
	t *testing.T,
	cfg *Config,
	watchersClient *watchers.Client,
	channelID uuid.UUID,
	watcherID uuid.UUID,
) {
	t.Helper()

	if watcherID == uuid.Nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CleanupTimeout)
	defer cancel()

	_, err := watchersClient.Archive(ctx, channelID, watcherID)
	if err != nil && !isNotFoundError(err) {
		t.Logf("watcher %s archive request failed: %v", watcherID.String(), err)
		return
	}

	waitCtx, waitCancel := context.WithTimeout(context.Background(), cfg.CleanupTimeout)
	defer waitCancel()

	ticker := time.NewTicker(cfg.EventPollInterval)
	defer ticker.Stop()

	for {
		watcher, err := watchersClient.Get(waitCtx, channelID, watcherID)
		if err != nil {
			if isNotFoundError(err) {
				return
			}
			t.Logf("watcher %s get during cleanup failed: %v", watcherID.String(), err)
		} else if statusString(watcher.Status) == "archived" {
			return
		}

		select {
		case <-waitCtx.Done():
			t.Logf("watcher %s did not reach archived before cleanup timeout", watcherID.String())
			return
		case <-ticker.C:
		}
	}
}

func archiveWallet(
	t *testing.T,
	cfg *Config,
	walletsClient *wallets.Client,
	walletID uuid.UUID,
) {
	t.Helper()

	if walletID == uuid.Nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CleanupTimeout)
	defer cancel()

	if err := walletsClient.Archive(ctx, walletID); err != nil && !isNotFoundError(err) {
		t.Logf("wallet %s archive failed: %v", walletID.String(), err)
	}
}

func archiveChannel(
	t *testing.T,
	cfg *Config,
	channelsClient *channels.Client,
	channelID uuid.UUID,
) {
	t.Helper()

	if channelID == uuid.Nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CleanupTimeout)
	defer cancel()

	if _, err := channelsClient.Archive(ctx, channelID); err != nil && !isNotFoundError(err) {
		t.Logf("channel %s archive failed: %v", channelID.String(), err)
	}
}

func packNoArgMethod(jsonABI string, method string) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(jsonABI))
	if err != nil {
		return nil, fmt.Errorf("parse ABI for %s: %w", method, err)
	}

	calldata, err := parsedABI.Pack(method)
	if err != nil {
		return nil, fmt.Errorf("pack %s: %w", method, err)
	}

	return calldata, nil
}

func buildTriggerOperation(
	walletAddress string,
	targetAddress string,
	calldata []byte,
	deadlineLead time.Duration,
) *transactTypes.Operation {
	deadline := big.NewInt(0)
	if deadlineLead > 0 {
		deadline = big.NewInt(time.Now().Add(deadlineLead).Unix())
	}

	return &transactTypes.Operation{
		ID:       newRandomOperationID(),
		Deadline: deadline,
		Account:  common.HexToAddress(walletAddress),
		Transactions: []transactTypes.Transaction{
			{
				To:    common.HexToAddress(targetAddress),
				Value: big.NewInt(0),
				Data:  calldata,
			},
		},
	}
}

func newRandomOperationID() *big.Int {
	b := make([]byte, 16)
	if _, err := crand.Read(b); err != nil {
		return big.NewInt(time.Now().UnixNano() + 1)
	}

	id := new(big.Int).SetBytes(b)
	if id.Sign() == 0 {
		return big.NewInt(time.Now().UnixNano() + 1)
	}

	return id
}

func uniqueName(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
}

func statusString(v any) string {
	return strings.ToLower(strings.TrimSpace(fmt.Sprint(v)))
}

func isNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	lower := strings.ToLower(err.Error())
	return strings.Contains(lower, "not found") || strings.Contains(lower, "404")
}
