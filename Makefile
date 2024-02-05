default:
	go build -o $$PWD/local/bin/protoc-gen-setter ./cmd/protoc-gen-setter

proto:
	PATH=$$PWD/local/bin:$$PATH protoc setterpb/setter.proto --go_out=. --go_opt=paths=source_relative

prepare:
	GOBIN=$$PWD/local/bin go install google.golang.org/protobuf/cmd/protoc-gen-go

test: default example
	go build ./example/...

.PHONY: example
example: default prepare
	PATH=$$PWD/local/bin:$$PATH GOBIN=$$PWD/local/bin \
	protoc \
	-I $$PWD \
	example/proto2/example.proto \
	example/proto3/example.proto \
	--go_out=. --go_opt=paths=source_relative \
	--setter_out=. --setter_opt=paths=source_relative
