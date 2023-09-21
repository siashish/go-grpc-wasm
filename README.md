// install for mac
`brew install protobuf`
`brew install protoc-gen-go`
or 
`brew install protoc-gen-go-grpc`

// check version
`protoc --version`
`protoc-gen-go --version`
`protoc-gen-go-grpc --version`

to generate the proto stub
`protoc <proto-file-path> --go_out=<output-file-path> --go-grpc_out=<output-file-path>`
`protoc *.proto --go_out=./ --go-grpc_out=./`

to run the server 
`go run server.go`
to run the client
`go run client.go`