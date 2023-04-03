.PHONY: install-tools proto clean test tidy run send_auth

install-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

proto:
	@protoc -I ./api/proto \
		--go_out=./pkg/gk \
		--go_opt=paths=source_relative \
		--go-grpc_out=./pkg/gk \
		--go-grpc_opt=paths=source_relative \
		api.proto

clean:
	@rm ./pkg/gk/*

test:
	@go test -race -v ./...

tidy:
	@go mod tidy
	@go mod verify
	@go mod vendor

run:
	@go run ./cmd/main.go

send_auth:
	grpcurl \
		-proto api/proto/api.proto \
		-plaintext -d '{"client_id":"client_id","client_secret":"client_secret"}' \
		localhost:8081 gatekeeper.Api.Authenticate
