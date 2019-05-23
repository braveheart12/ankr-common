.PHONY: all common micro grpc build test run deploy local clean echo

GOPATH:=$(shell go env GOPATH)

all: micro grpc


micro:
	@echo "Generating montior micro proto"
	@for type in account notify; do \
		protoc -I. \
			--micro_out=Mcommon/common.proto=github.com/Ankr-network/dccn-common/protos/common:. \
			--go_out=Mcommon/common.proto=github.com/Ankr-network/dccn-common/protos/common:. \
			$$type/v1/*.proto; \
	done

grpc: 
	@echo "Generating montior gPRC proto"
	@for type in account notify; do \
		protoc -I. \
			--go_out=plugins=grpc,Mcommon/common.proto=github.com/Ankr-network/dccn-common/protos/common:. \
			$$type/v1/*.proto; \
	done
