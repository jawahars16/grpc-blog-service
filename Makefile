proto:
	protoc --go_out=. --go-grpc_out=. internal/**/*.proto

test:
	go test -race -v ./...

PORT=9000
server:
	go run cmd/server/main.go --port $(PORT)

client:
	go run cmd/client/main.go --port $(PORT)