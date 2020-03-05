all:
	protoc --go_out=. merger.proto
	protoc --gotag_out=. merger.proto
