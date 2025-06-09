generate:
	protoc --proto_path=/usr/include --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.
