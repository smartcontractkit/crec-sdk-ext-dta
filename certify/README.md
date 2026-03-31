# DTA certification runbook

This folder contains the deployed-CREC certification harness for the DTA watcher services:

- `dta.v1`
- `dta.v2`

The certification entrypoint is the Go test in this folder:

```bash
go test ./certify -run TestCertify -count=1 -v
```

There are also Taskfile shortcuts:

```bash
task certify:v1
task certify:v2
```

---

## What this certification actually tests

At a high level, the `certify` package proves that a **deployed CRE watcher service** for DTA can:

1. create a CRE channel,
2. create and deploy a CRE wallet,
3. create management and settlement watchers against configured contract addresses,
4. receive watcher events from the selected DTA service (`dta.v1` or `dta.v2`),
5. optionally verify watcher-event signatures,
6. decode the watcher payload with the matching versioned Go SDK,
7. verify that expected on-chain enrichment is present, and
8. verify that settlement events produce a payment request.

### Important scope limitation

This harness is **not** a generic “point at any live DTA deployment and run certification” test.

It sends trigger operations to contracts exposing these methods:

- `emitDistributorRequestProcessing()`
- `emitDTASettlementOpened()`

Those methods exist in the **mock contracts in this repo**, not in a normal live DTA deployment.

That means the current certification flow is designed around these local mock contracts:

- `test/MockDTARequestManagement.sol` for `dta.v1`
- `test/MockDTARequestManagementV2.sol` for `dta.v2`
- `test/MockDTARequestSettlement.sol` for both

So, for the current code as written, the minimal on-chain setup is the mock pair, not the full production DTA stack.

---

## What the test does step by step

When `TestCertify` runs, it does the following:

1. loads env from `certify/.env` and then `.env`,
2. builds CRE clients for:
   - API,
   - channels,
   - wallets,
   - watchers,
   - transact,
   - events,
3. creates a temporary channel,
4. creates a temporary CRE wallet using `ACCOUNT_SIGNER_PRIVATE_KEY` as the allowed ECDSA signer,
5. creates **two watchers** against the configured contracts:
   - management watcher on `DTA_REQUEST_MANAGEMENT_ADDRESS` for event `DistributorRequestProcessing`,
   - settlement watcher on `DTA_REQUEST_SETTLEMENT_ADDRESS` for event `DTASettlementOpened`,
6. waits for the wallet to be deployed and both watchers to become active,
7. performs a **local SDK deadline probe** by building `PrepareRegisterFundAdminOperation()` twice and proving deadlines are cloned correctly,
8. signs and sends an operation to call `emitDistributorRequestProcessing()` on the management mock,
9. waits for operation confirmation,
10. polls the channel until the matching `watcher.event` arrives,
11. optionally verifies the watcher-event signatures,
12. decodes the event using the selected versioned SDK,
13. asserts:
    - event name is correct,
    - fund token enrichment exists,
    - distributor request enrichment exists,
    - shares and amount are present and `> 0`,
    - payment request count is `0` for the management event,
14. signs and sends an operation to call `emitDTASettlementOpened()` on the settlement mock,
15. waits for confirmation and watcher event,
16. decodes again,
17. asserts:
    - event name is correct,
    - enrichment exists,
    - payment request count is `1` for the settlement event,
    - settlement address inside enrichment matches the configured settlement address,
18. archives watchers, wallet, and channel.

---

## Before you run this

You need **all** of the following:

1. a CRE/CREC environment you can talk to,
2. a valid CRE API key,
3. the DTA watcher service already deployed in CRE under either:
   - `dta.v1`, or
   - `dta.v2`,
4. deployed mock contracts on the target chain,
5. the correct chain selector for that chain,
6. a funded private key able to sign and send operations to the mock contracts.

### The certification does **not** deploy the watcher service for you

The certification creates watchers that reference a service ID. It does **not** publish the workflow bundle itself.

If the service is not already deployed, deploy it first from this repo:

```bash
task watcher:release VERSION=v1
task watcher:release VERSION=v2
```

Use the appropriate `CRE_TARGET`, API credentials, and watcher values for your environment.

---

## Environment configuration

The test loads env in this order:

1. existing shell environment,
2. `certify/.env`,
3. `.env`

Already-exported env vars win; the loader does not overwrite them.

`RUN_CREC_CERTIFICATION` defaults to `false`, so plain `go test ./...` does not accidentally hit a live CRE environment.

---

## Required certification env vars

The main runtime settings are defined in `certify/.env`.

### Core settings

| Variable | Required | Meaning |
|---|---:|---|
| `RUN_CREC_CERTIFICATION` | yes | Must be `true` to run instead of skip |
| `CRE_ORG_ID` | yes | Required by config validation |
| `CRE_API_KEY` | yes | CRE/CREC API key |
| `COURIER_URL` | yes | CRE/CREC endpoint; default is `http://localhost:8088` |
| `CHAIN_SELECTOR` | yes | **CCIP chain selector**, not EVM chain ID |
| `ACCOUNT_SIGNER_PRIVATE_KEY` | yes | Private key used to sign operations for the temporary CRE wallet |
| `DTA_SERVICE` | yes | Must be `dta.v1` or `dta.v2` |
| `DTA_REQUEST_MANAGEMENT_ADDRESS` | yes | Address of the version-matched management mock |
| `DTA_REQUEST_SETTLEMENT_ADDRESS` | yes | Address of the shared settlement mock |

### Optional signature verification settings

| Variable | Required | Meaning |
|---|---:|---|
| `CREC_VALID_SIGNERS` | no | Comma-separated watcher signer addresses |
| `MIN_REQUIRED_SIGNATURES` | no | Minimum signatures required for verification |

If `CREC_VALID_SIGNERS` is empty, signature verification is skipped.

### Timeout / polling settings

| Variable | Default |
|---|---:|
| `WALLET_DEPLOY_TIMEOUT` | `10m` |
| `WATCHER_ACTIVE_TIMEOUT` | `10m` |
| `WATCHER_PROPAGATION_WAIT` | `2m` |
| `OPERATION_CONFIRM_TIMEOUT` | `5m` |
| `EVENT_POLL_INTERVAL` | `5s` |
| `CLEANUP_TIMEOUT` | `2m` |
| `OPERATION_DEADLINE_LEAD` | `30m` |

---

## `CHAIN_SELECTOR` is not chain ID

This is the most common configuration mistake.

For certification, `CHAIN_SELECTOR` must match the selector used by CRE/transact/watchers for the target chain.  
It is **not** the EVM chain ID like `11155111`.

### Common selectors used in this repo

These values are visible in `helper/HelperConfig.s.sol` from the DTA contracts tree you provided:

| Network | EVM chain ID | Chain selector |
|---|---:|---:|
| Ethereum Sepolia | `11155111` | `16015286601757825753` |
| Arbitrum Sepolia | `421614` | `3478487238524512106` |
| Local Geth | `1337` | `3379446385462418246` |
| Local Anvil | `31337` | `31337` |

For `certify`, only the selector matters.

---

## Example env blocks for different chains

### Ethereum Sepolia

```bash
RUN_CREC_CERTIFICATION=true
CRE_ORG_ID=<your-org-id>
CRE_API_KEY=<your-api-key>
COURIER_URL=<your-cre-url>

CHAIN_SELECTOR=16015286601757825753
ACCOUNT_SIGNER_PRIVATE_KEY=<funded-private-key>

DTA_SERVICE=dta.v1
DTA_REQUEST_MANAGEMENT_ADDRESS=<v1-management-mock-address>
DTA_REQUEST_SETTLEMENT_ADDRESS=<settlement-mock-address>
```

### Arbitrum Sepolia

```bash
RUN_CREC_CERTIFICATION=true
CRE_ORG_ID=<your-org-id>
CRE_API_KEY=<your-api-key>
COURIER_URL=<your-cre-url>

CHAIN_SELECTOR=3478487238524512106
ACCOUNT_SIGNER_PRIVATE_KEY=<funded-private-key>

DTA_SERVICE=dta.v2
DTA_REQUEST_MANAGEMENT_ADDRESS=<v2-management-mock-address>
DTA_REQUEST_SETTLEMENT_ADDRESS=<settlement-mock-address>
```

### Local Geth

```bash
RUN_CREC_CERTIFICATION=true
CRE_ORG_ID=local-org
CRE_API_KEY=local-api-key
COURIER_URL=http://localhost:8088

CHAIN_SELECTOR=3379446385462418246
ACCOUNT_SIGNER_PRIVATE_KEY=<local-funded-key>

DTA_SERVICE=dta.v1
DTA_REQUEST_MANAGEMENT_ADDRESS=<mock-address>
DTA_REQUEST_SETTLEMENT_ADDRESS=<mock-address>
```

### Local Anvil

```bash
RUN_CREC_CERTIFICATION=true
CRE_ORG_ID=local-org
CRE_API_KEY=local-api-key
COURIER_URL=http://localhost:8088

CHAIN_SELECTOR=31337
ACCOUNT_SIGNER_PRIVATE_KEY=<anvil-funded-key>

DTA_SERVICE=dta.v2
DTA_REQUEST_MANAGEMENT_ADDRESS=<mock-address>
DTA_REQUEST_SETTLEMENT_ADDRESS=<mock-address>
```

Use Anvil only if your CRE/CREC environment is actually wired to it.

---

## What contracts are required to run `certify`

### Minimal contract set required by the current certification code

| Purpose | Required for `dta.v1` | Required for `dta.v2` |
|---|---|---|
| Management trigger + enrichment mock | `test/MockDTARequestManagement.sol` | `test/MockDTARequestManagementV2.sol` |
| Settlement trigger mock | `test/MockDTARequestSettlement.sol` | `test/MockDTARequestSettlement.sol` |

### Why only these contracts?

Because the certification code explicitly calls:

- `emitDistributorRequestProcessing()` on the management contract,
- `emitDTASettlementOpened()` on the settlement contract,

and then expects deterministic enrichment from:

- `getDistributorRequest(bytes32)`,
- `getFundToken(address,bytes32)`.

The mock contracts implement exactly that.

### What is **not** required for the current `certify` flow

The following are part of a fuller DTA environment, but the current certification harness does **not** require them:

- fee manager,
- fund token,
- NAV feed,
- DTADistributor / DTADistributorReceiver,
- allowlist scripts,
- role-grant scripts,
- real DTA request management/settlement deployments.

Those are needed for full DTA business flows, not for this repo’s current certification scenario.

---

## Can I use somebody else’s DTA instance without the deployer private key?

### Short answer

**Sometimes yes, but usually not for this exact certification flow.**

### The precise answer

You do **not** need the original deployer private key just to:

- watch a contract,
- read contract state,
- call a public permissionless method.

So if someone else already deployed the **same mock contracts** and:

- they are wired correctly to each other,
- the emit methods are still public,
- your signer can call them on the target chain,

then yes, you can certify against those addresses without owning the deployer key.

### When the answer is effectively “no”

For the current `certify` package, pointing at an arbitrary real DTA deployment will usually fail because:

1. real DTA contracts do not expose `emitDistributorRequestProcessing()`,
2. real DTA settlement contracts do not expose `emitDTASettlementOpened()`,
3. the test expects deterministic mock enrichment and payment-request behavior.

So for a normal live DTA deployment, you usually **cannot** just reuse it for this certification harness as-is.

### What still requires owner/admin keys

Even with the mock contracts, you need the owner key if you want to:

- call `updateManagementAddr(...)` on the settlement mock,
- call `updateSettlementAddr(...)` on the management mock,
- rewire an incorrectly deployed pair.

With real DTA contracts, you additionally need the right owner/admin/workflow-account permissions for:

- upgrades,
- allowlisting,
- token registration,
- role grants,
- admin/distributor registration flows,
- ownership transfers.

---

## Deploy the minimal certification contracts with Docker + Foundry

These commands assume:

- you are at the root of **this repo**,
- `RPC_URL` points at the target chain,
- `DEPLOYER_KEY` is funded on that chain,
- Docker is installed,
- you want to use Foundry from Docker, not a local `forge` install.

Set a reusable image variable first:

```bash
export FOUNDRY_IMAGE=ghcr.io/foundry-rs/foundry:latest
```

### Deploy the v1 certification pair

This deploys the shared settlement mock first with a placeholder management address, then deploys the v1 management mock with the real settlement address, then patches the settlement mock to point back to the management mock.

```bash
PLACEHOLDER=0x0000000000000000000000000000000000000001; \
SETTLEMENT_ADDR=$(docker run --rm --entrypoint sh -e RPC_URL -e DEPLOYER_KEY -v "$PWD":/src -w /src "$FOUNDRY_IMAGE" -lc "forge create test/MockDTARequestSettlement.sol:MockDTARequestSettlement --rpc-url \$RPC_URL --private-key \$DEPLOYER_KEY --broadcast --constructor-args $PLACEHOLDER | sed -n 's/Deployed to: //p'"); \
MGMT_ADDR=$(docker run --rm --entrypoint sh -e RPC_URL -e DEPLOYER_KEY -v "$PWD":/src -w /src "$FOUNDRY_IMAGE" -lc "forge create test/MockDTARequestManagement.sol:MockDTARequestManagement --rpc-url \$RPC_URL --private-key \$DEPLOYER_KEY --broadcast --constructor-args $SETTLEMENT_ADDR | sed -n 's/Deployed to: //p'"); \
docker run --rm --entrypoint sh -e RPC_URL -e DEPLOYER_KEY -v "$PWD":/src -w /src "$FOUNDRY_IMAGE" -lc "cast send $SETTLEMENT_ADDR 'updateManagementAddr(address)' $MGMT_ADDR --rpc-url \$RPC_URL --private-key \$DEPLOYER_KEY" >/dev/null; \
echo "DTA_REQUEST_MANAGEMENT_ADDRESS=$MGMT_ADDR"; \
echo "DTA_REQUEST_SETTLEMENT_ADDRESS=$SETTLEMENT_ADDR"
```

### Deploy the v2 certification pair

Same flow, but with the v2-compatible management mock:

```bash
PLACEHOLDER=0x0000000000000000000000000000000000000001; \
SETTLEMENT_ADDR=$(docker run --rm --entrypoint sh -e RPC_URL -e DEPLOYER_KEY -v "$PWD":/src -w /src "$FOUNDRY_IMAGE" -lc "forge create test/MockDTARequestSettlement.sol:MockDTARequestSettlement --rpc-url \$RPC_URL --private-key \$DEPLOYER_KEY --broadcast --constructor-args $PLACEHOLDER | sed -n 's/Deployed to: //p'"); \
MGMT_ADDR=$(docker run --rm --entrypoint sh -e RPC_URL -e DEPLOYER_KEY -v "$PWD":/src -w /src "$FOUNDRY_IMAGE" -lc "forge create test/MockDTARequestManagementV2.sol:MockDTARequestManagementV2 --rpc-url \$RPC_URL --private-key \$DEPLOYER_KEY --broadcast --constructor-args $SETTLEMENT_ADDR | sed -n 's/Deployed to: //p'"); \
docker run --rm --entrypoint sh -e RPC_URL -e DEPLOYER_KEY -v "$PWD":/src -w /src "$FOUNDRY_IMAGE" -lc "cast send $SETTLEMENT_ADDR 'updateManagementAddr(address)' $MGMT_ADDR --rpc-url \$RPC_URL --private-key \$DEPLOYER_KEY" >/dev/null; \
echo "DTA_REQUEST_MANAGEMENT_ADDRESS=$MGMT_ADDR"; \
echo "DTA_REQUEST_SETTLEMENT_ADDRESS=$SETTLEMENT_ADDR"
```

### Notes on the deployment flow

- `MockDTARequestSettlement` needs a management address in its constructor.
- `MockDTARequestManagement*` needs a settlement address in its constructor.
- That circular dependency is why one side is deployed with a placeholder and then patched.
- The settlement mock is shared by both v1 and v2.
- The management mock must match the service version:
  - v1 service -> `MockDTARequestManagement.sol`
  - v2 service -> `MockDTARequestManagementV2.sol`

---

## Run certification for V1 workflows

### Required setup

1. Ensure the CRE watcher service `dta.v1` is deployed.
2. Deploy the v1 management mock and shared settlement mock.
3. Fill `certify/.env` with:
   - `DTA_SERVICE=dta.v1`
   - `CHAIN_SELECTOR=<target-selector>`
   - `DTA_REQUEST_MANAGEMENT_ADDRESS=<v1-mock>`
   - `DTA_REQUEST_SETTLEMENT_ADDRESS=<shared-settlement-mock>`
   - CRE credentials
   - funded signer key

### Run it

#### Task shortcut

```bash
task certify:v1
```

#### Direct Go command

```bash
RUN_CREC_CERTIFICATION=true DTA_SERVICE=dta.v1 go test ./certify -run TestCertify -count=1 -v
```

### What happens in the V1 run

The V1 certification will:

1. create a channel,
2. create a CRE wallet,
3. create two watchers bound to service `dta.v1`,
4. verify the V1 Go operations SDK clones deadlines correctly,
5. send an operation to the v1 management mock to emit `DistributorRequestProcessing`,
6. wait for confirmation and watcher event,
7. decode with `github.com/smartcontractkit/crec-sdk-ext-dta/v1`,
8. assert V1 enrichment fields are present,
9. send an operation to the shared settlement mock to emit `DTASettlementOpened`,
10. wait for the watcher event,
11. decode again with the V1 decoder,
12. assert exactly one payment request was produced for settlement.

### What V1 specifically proves

It proves the deployed `dta.v1` service and the local V1 SDK agree on:

- event names,
- event payload decoding,
- fund token enrichment,
- distributor request enrichment,
- settlement payment request generation,
- deadline propagation in the V1 operation builder.

---

## Run certification for V2 workflows

### Required setup

1. Ensure the CRE watcher service `dta.v2` is deployed.
2. Deploy the v2 management mock and shared settlement mock.
3. Fill `certify/.env` with:
   - `DTA_SERVICE=dta.v2`
   - `CHAIN_SELECTOR=<target-selector>`
   - `DTA_REQUEST_MANAGEMENT_ADDRESS=<v2-mock>`
   - `DTA_REQUEST_SETTLEMENT_ADDRESS=<shared-settlement-mock>`
   - CRE credentials
   - funded signer key

### Run it

#### Task shortcut

```bash
task certify:v2
```

#### Direct Go command

```bash
RUN_CREC_CERTIFICATION=true DTA_SERVICE=dta.v2 go test ./certify -run TestCertify -count=1 -v
```

### What happens in the V2 run

The V2 certification follows the same control flow as V1, but uses:

- watcher service `dta.v2`,
- V2 operations builder,
- V2 decode path,
- the V2 management mock with the V2 distributor request shape.

It will:

1. create channel/wallet/watchers,
2. run the V2 deadline probe,
3. trigger `DistributorRequestProcessing`,
4. decode with `github.com/smartcontractkit/crec-sdk-ext-dta/v2`,
5. trigger `DTASettlementOpened`,
6. decode and verify payment request generation.

### What V2 specifically proves

It proves the deployed `dta.v2` service and the local V2 SDK agree on:

- V2 event decoding,
- V2 enrichment payload shape,
- settlement payment request generation,
- deadline propagation in the V2 operations builder.

---

## Expected results

A successful run should end with:

- `go test` exit code `0`,
- `PASS`,
- no watcher activation timeouts,
- no operation confirmation timeouts,
- no decode errors,
- no enrichment assertion failures.

### Management event expectations

For `DistributorRequestProcessing`, the test expects:

- event name = `DistributorRequestProcessing`,
- `FundTokenData` present,
- `DistributorRequest` present,
- `PaymentRequestCount == 0`,
- `shares > 0`,
- `amount > 0`,
- enriched settlement address matches `DTA_REQUEST_SETTLEMENT_ADDRESS`.

### Settlement event expectations

For `DTASettlementOpened`, the test expects:

- event name = `DTASettlementOpened`,
- `FundTokenData` present,
- `DistributorRequest` present,
- `PaymentRequestCount == 1`,
- `shares > 0`,
- `amount > 0`,
- enriched settlement address matches `DTA_REQUEST_SETTLEMENT_ADDRESS`.

---

## Difference between V1 and V2

| Area | V1 | V2 |
|---|---|---|
| CRE service ID | `dta.v1` | `dta.v2` |
| Management mock | `test/MockDTARequestManagement.sol` | `test/MockDTARequestManagementV2.sol` |
| Settlement mock | `test/MockDTARequestSettlement.sol` | `test/MockDTARequestSettlement.sol` |
| Go decode package | `v1` | `v2` |
| Go operations package | `v1/operations` | `v2/operations` |
| Distributor request shape | no `referenceID` | includes `referenceID` |
| Certification deadline probe | V1 builder | V2 builder |

### The most important data-model difference

The V2 management mock returns a `DistributorRequest` that includes `referenceID`.

That is the key reason you must not mix:

- `dta.v1` with `MockDTARequestManagementV2.sol`, or
- `dta.v2` with `MockDTARequestManagement.sol`.

If you do, decoding/enrichment assumptions will diverge.

---

## Common failure modes

### 1. Wrong `DTA_SERVICE`

**Symptom:** watcher activates but decode assertions fail, or the wrong service behavior is observed.

**Fix:** use:
- `dta.v1` with `MockDTARequestManagement.sol`
- `dta.v2` with `MockDTARequestManagementV2.sol`

### 2. Wrong `CHAIN_SELECTOR`

**Symptom:** operation never confirms, or watcher never sees the event.

**Fix:** set the CCIP selector, not the EVM chain ID.

### 3. Management and settlement mocks are not wired together

**Symptom:** enrichment points to the wrong settlement address, or assertions about settlement address fail.

**Fix:** ensure:
- management mock was constructed with the real settlement address,
- settlement mock was updated to point at the real management address.

### 4. Trying to use a real DTA deployment

**Symptom:** operation fails because `emitDistributorRequestProcessing()` or `emitDTASettlementOpened()` does not exist.

**Fix:** use the mock contracts, or extend `certify` to drive real contract methods.

### 5. Signer verification failures

**Symptom:** `watcher.event` signature verification fails.

**Fix:** either:
- configure `CREC_VALID_SIGNERS` correctly, or
- leave it blank to skip verification.

### 6. Wallet or watcher activation timeout

**Symptom:** test fails before any event is triggered.

**Fix:** verify:
- CRE service deployment exists,
- API key is valid,
- courier URL is correct,
- target chain is supported by your CRE environment,
- timeouts are large enough.

---

## Full DTA stack vs current certify flow

The contracts tree you provided from the DTA contracts repo is useful for a **real DTA environment**, but the current `certify` package does not drive that full stack.

### For a full DTA environment you typically need more than the mocks

Examples from the DTA contracts tree include:

- request management deployment,
- request settlement deployment,
- fee manager setup,
- fund token creation,
- NAV feed deployment,
- fund admin/distributor registration,
- allowlisting,
- mint/burn role grants,
- payment-token approvals,
- DTA allow/disallow setup.

### But for this repo’s `certify` package today

You only need:

- version-matched management mock,
- shared settlement mock,
- deployed watcher service,
- correct env.

If later the certification harness is extended to drive real DTA methods instead of mock emitters, that requirement set will expand.

---

## Quick start

### V1

```bash
export FOUNDRY_IMAGE=ghcr.io/foundry-rs/foundry:latest
export RPC_URL=<target-rpc>
export DEPLOYER_KEY=<funded-key>
# deploy mocks with the v1 one-liner above
# update certify/.env with returned addresses
task certify:v1
```

### V2

```bash
export FOUNDRY_IMAGE=ghcr.io/foundry-rs/foundry:latest
export RPC_URL=<target-rpc>
export DEPLOYER_KEY=<funded-key>
# deploy mocks with the v2 one-liner above
# update certify/.env with returned addresses
task certify:v2
```

---

## TL;DR

- `certify` is a **deployed CRE watcher-service certification**, not a full DTA business-flow test.
- Use `dta.v1` with `MockDTARequestManagement.sol`.
- Use `dta.v2` with `MockDTARequestManagementV2.sol`.
- Use `MockDTARequestSettlement.sol` for both.
- `CHAIN_SELECTOR` is a **chain selector**, not chain ID.
- You can reuse someone else’s deployment **only if** it is the same certification mock surface or equivalent and your signer can call it.
- Expected pass criteria:
  - management event decoded with enrichment and no payment requests,
  - settlement event decoded with enrichment and exactly one payment request,
  - overall `go test` passes.
