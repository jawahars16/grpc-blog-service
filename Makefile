proto:
	protoc --go_out=. --go-grpc_out=. internal/**/*.proto

test:
	go test -v ./...

server:
	go run cmd/server/main.go

client:
	go run cmd/client/main.go