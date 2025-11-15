dev:
	clear
	go run main.go

migrate.up:
	clear
	cd migration/ && go run main.go UP

migrate.down:
	clear
	cd migration/ && go run main.go DOWN

genproto:
	protoc -I . --go_out=proto/ --go-grpc_out=require_unimplemented_servers=false:proto/ proto/*.proto proto/*.proto

test:
	go test ./...
