gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/processor_message.proto

clean:
	rm -f pb/*.go

run:
	go run main.go