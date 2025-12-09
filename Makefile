.PHONY: test
test:
	go test ./...

.PHONY: generate
generate:
	go run ./v1/gen/main.go

.PHONY: vendor
vendor:
	go mod tidy
	go mod download
	go mod vendor