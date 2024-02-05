default:
	go build -o $$PWD/local/bin/protoc-gen-go-svc-interface ./cmd/protoc-gen-go-svc-interface
	go build -o $$PWD/local/bin/protoc-gen-go-setter ./cmd/protoc-gen-go-setter

setter_proto:
	PATH=$$PWD/local/bin:$$PATH protoc setterpb/setter.proto --go_out=. --go_opt=paths=source_relative

prepare:
	GOBIN=$$PWD/local/bin go install google.golang.org/protobuf/cmd/protoc-gen-go

test: default example
	go build ./example/...

.PHONY: example
example: prepare setter_proto default
	PATH=$$PWD/local/bin:$$PATH GOBIN=$$PWD/local/bin \
	protoc \
	-I $$PWD \
	example/proto2/example.proto \
	example/proto3/example.proto \
	--go_out=. --go_opt=paths=source_relative \
	--go-setter_out=. --go-setter_opt=paths=source_relative \
	--go-svc-interface_out=. --go-svc-interface_opt=paths=source_relative

clean:
	rm example/*/*.pb.go setterpb/*.go local/bin/*