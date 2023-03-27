# json-2-protobuf-sample

## Get Started

```
# install protocol compiler
$ brew install protobuf

# install protobuf runtime
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# generate codes
$ protoc ./proto/*.proto --go_out=./pbgo

$ go build

$ ./j2p

# execute with json string
$ ./j2p -d '{ "id": 1, "name": "xxx", "email": "mail@example.com" }'
```

## References

- https://protobuf.dev/getting-started/gotutorial/
- https://github.com/protocolbuffers/protobuf
