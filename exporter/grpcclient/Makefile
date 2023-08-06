
PROTO := protoc
PROTO_DIR := ./testData
PROTO_FILE := data.proto
PROTO_OPT := --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

.PHONY: proto
proto:
	cd $(PROTO_DIR) && $(PROTO) $(PROTO_OPT) $(PROTO_FILE)
