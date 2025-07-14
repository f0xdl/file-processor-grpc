combined-env:
	go generate ./internal/client/app.go
	go generate ./internal/fileservice/app.go

proto-gen:
	protoc \
      --proto_path=api/proto \
      --go_out=gen/go --go_opt=paths=source_relative \
      --go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
      api/proto/fileprocessor/*.proto