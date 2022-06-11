# gRPC quick start with Go

```bash
protoc -I. --go_out=api --go_opt=paths=source_relative \
    --go-grpc_out=api --go-grpc_opt=paths=source_relative \
    weather.proto
```
