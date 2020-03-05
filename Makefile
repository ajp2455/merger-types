all:
	go get -u github.com/golang/protobuf github.com/srikrsna/protoc-gen-gotag
	protoc --go_out=. merger.proto
	protoc --gotag_out=. merger.proto
