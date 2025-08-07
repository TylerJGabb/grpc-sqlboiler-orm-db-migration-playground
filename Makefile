generate:
	@echo "Generating RPC Go code from proto files..."
	protoc --proto_path=proto proto/*.proto --go_out=. --go_grpc_out=.
	@echo "Generating SQL Models..."
	go generate