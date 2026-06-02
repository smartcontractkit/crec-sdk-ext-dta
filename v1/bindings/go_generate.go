package bindings

//go:generate go tool abigen --abi ../watcher/bundle/DTARequestManagementU.abi.json --pkg dtarm --out ./dtarm/dtarm.gen.go
//go:generate go tool abigen --abi ../watcher/bundle/DTARequestSettlementU.abi.json --pkg dtars --out ./dtars/dtars.gen.go
