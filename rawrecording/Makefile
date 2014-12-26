
GOGOPROTO_ROOT=$(shell go list -f '{{.Root}}' "github.com/gogo/protobuf/gogoproto")
GOGOPROTO_GEN=$(shell go list -f '{{.Target}}' github.com/gogo/protobuf/protoc-gen-gogo)

override PROTOPATH=$(GOGOPROTO_ROOT)/src:$(GOGOPROTO_ROOT)/src/github.com/gogo/protobuf/protobuf:.

all:
	protoc \
		--proto_path=$(PROTOPATH) \
		--plugin=protoc-gen-gogo=$(GOGOPROTO_GEN) \
		--gogo_out=. *.proto

clean:
	rm -f *.pb.go

.PHONY: all clean
