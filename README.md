# go_algo

## package

go get -u google.golang.org/grpc
go get -u google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

## check version

protoc --version
protoc-gen-go --version
protoc-gen-go-grpc --version

## generate go by protobuf

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/demo.proto
