# Weather App with gRPC

```bash
protoc -I. --go_out=api --go_opt=paths=source_relative \
    --go-grpc_out=api --go-grpc_opt=paths=source_relative \
    weather.proto
```
