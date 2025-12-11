module github.com/smartcontractkit/crec-sdk-ext-dta

go 1.25.3

require (
	github.com/ethereum/go-ethereum v1.16.4
	github.com/smartcontractkit/chainlink-protos/cre/go v0.0.0-20251021010742-3f8d3dba17d8
	github.com/smartcontractkit/cre-sdk-go v0.9.0
	github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm v0.9.0
	github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http v0.9.0
	github.com/smartcontractkit/crec-api-go v0.3.7-0.20251209174022-0ae3d2b8fe58
	github.com/smartcontractkit/crec-courier-service v0.0.0
	github.com/smartcontractkit/crec-sdk v0.1.5
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/bits-and-blooms/bitset v1.22.0 // indirect
	github.com/consensys/gnark-crypto v0.18.0 // indirect
	github.com/crate-crypto/go-eth-kzg v1.4.0 // indirect
	github.com/crate-crypto/go-ipa v0.0.0-20240724233137-53bbb0ceb27a // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/deckarep/golang-set/v2 v2.8.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.0 // indirect
	github.com/ethereum/c-kzg-4844/v2 v2.1.3 // indirect
	github.com/ethereum/go-verkle v0.2.2 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/oapi-codegen/runtime v1.1.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rs/zerolog v1.34.0 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/supranational/blst v0.3.16-0.20250831170142-f48500c1fdbe // indirect
	github.com/tklauser/go-sysconf v0.3.15 // indirect
	github.com/tklauser/numcpus v0.10.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	golang.org/x/crypto v0.42.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/smartcontractkit/crec-sdk => ../crec-sdk

replace github.com/smartcontractkit/crec-courier-service => ../crec-courier-service

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
