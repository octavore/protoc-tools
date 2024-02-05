# protoc-tools

This package provides `protoc-gen-go-setter` and `protoc-gen-go-svc-interface` to support generating setters and service interfaces respectively.


## Installation

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/octavore/protoc-tools/cmd/protoc-gen-go-setter@latest
go install github.com/octavore/protoc-tools/cmd/protoc-gen-go-svc-interface@latest
```

## Usage

### CLI

```shell
protoc example.proto --go_out=. --go-setter_out=.
protoc example.proto --go_out=. --go-svc-interface_out=.

protoc example.proto \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-setter_out=. \
  --go-setter_opt=paths=source_relative
```

### Proto

NB: You probably want to just copy the `setter.proto` file somewhere in your project and reference it there 🤷‍♂️.

```proto
syntax = "proto3";

package example.proto3;

option go_package = "github.com/octavore/protoc-gen-setter/example/proto3";

import "google/protobuf/descriptor.proto";
import "setterpb/setter.proto";

message Timestamp {
    option (setter.all_fields) = true;
    int64 created_at = 1;
    int64 updated_at = 2;
}

message Page {
    Timestamp timestamps = 1 [(setter.field).include = true];

    oneof type {
      option (setter.oneof_field).include = true;
      string text = 11;
      int64 number = 12;
    }
}

service PageService {
  rpc GetPage(Timestamp) returns (Page);
}
```

## Development

```
make prepare
make test  # build example protos
```

## Credits

Inspired by https://github.com/confluentinc/proto-go-setter but uses `google.golang.org/protobuf` exclusively instead of `gogo/protobuf`.

## License

MIT