# json-2-protobuf-sample

## Get Started

```
$ brew install protobuf

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

$ protoc ./proto/*.proto --go_out=./pbgo

$ go build

$ ./j2p

$ ./j2p -d '{ "id": 1, "name": "xxx", "email": "mail@example.com" }'
```

## References

- https://protobuf.dev/getting-started/gotutorial/
- https://github.com/protocolbuffers/protobuf
