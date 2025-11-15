dev:
	clear
	go run main.go

genproto:
	protoc -I . --go_out=proto/ --go-grpc_out=require_unimplemented_servers=false:proto/ proto/*.proto proto/*.proto

test:
	go test ./...
	