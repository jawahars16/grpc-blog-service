proto:
	protoc --go_out=. --go-grpc_out=. internal/**/*.proto

server:
	go run cmd/server/main.go