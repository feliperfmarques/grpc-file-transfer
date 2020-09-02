# gRPC File Transfer Example

Simple example for transfer file(.mp4) using gRPC Stream.

* Server: Go
* Client: Dart


## Dependencies:

1. Go >= 1.5 - For install instructions [Click Here](https://golang.org/doc/install)

2. Dart >= 2.2.2 <= 3.0.0 - For install instructions [Click Here](https://dart.dev/get-dart)

3. protoc (Protobuf Compiler) - For install instructions [Click Here](https://grpc.io/docs/protoc-installation/)

4. protoc-gen-go plugin
```
export GO111MODULE=on
go get github.com/golang/protobuf/protoc-gen-go
```

5. protoc-gen-go plugin
```
export GO111MODULE=on
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

6. protoc-gen-dart plugin
```
pub global activate protoc_plugin
```


## Generating Protos

The last protos for this example already have been genereted and commited, but if you need to regenerate the protos use the instructions below:


### Using `vscode-proto3`:

By default **⇧ + CTRL + P(Linux/Windows)** ou **⇧ + ⌘ + P(Mac)** opens the command prompt. Type one of these commands:

| Command | Description |
|---------|-------------|
| `proto3: Compile All Protos` | Compiles all workspace protos using [configurations](#extension-settings) defined with `protoc.options`. |
| `proto3: Compile This Proto` | Compiles the active proto using [configurations](#extension-settings) defined with `protoc.options`. |

This project has include .vscode/settings with necessary configs for vscode-proto3, open this file and uncomment "--plugin" parameters with absolute path for protoc-gen-dart, protoc-gen-go, protoc-gen-go-grpc case default configs doesn't works properly and get error with unknown path for this plugins.


### Using protoc CLI:

Open terminal and cd for project root:

```
cd PROJECT_ROOT
protoc  --plugin="protoc-gen-go=[absolute_path_to]/protoc-gen-go" \
        --plugin="protoc-gen-go-grpc=[absolute_path_to]/protoc-gen-go-grpc" \
        --plugin="protoc-gen-dart=[absolute_path_to]/protoc-gen-dart" \
        --proto_path=${PWD} \
        --dart_out="grpc:./client/lib/src" \
        --go_out="./server" \
        --go-grpc_out="./server" \
        proto/grpc_file_transfer.proto
```


## Try

1. Run Server:
```
cd server/
go run main.go
```

2. Run Client:
```
cd client/
dart --packages=.dart_tool/package_config.json package:grpcfiletransfer/src/client.dart
```

## Credits

Video Example by [File Examples](https://file-examples-com.github.io/uploads/2017/04/file_example_MP4_1920_18MG.mp4)