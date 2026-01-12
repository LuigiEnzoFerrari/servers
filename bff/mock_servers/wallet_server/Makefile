.PHONY: gen tidy

# Generate Go code from proto files
gen:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/wallet.proto

# Tidy Go modules
tidy:
	go mod tidy
