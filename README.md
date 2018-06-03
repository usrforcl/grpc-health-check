# grpc-health-check

How to write grpc health check interface. See [GRPC Health Checking Protocol](https://github.com/grpc/grpc/blob/master/doc/health-checking.md)

## Start grpc server

```
$ go run server/main.go
```

## Run Health Check

```
$ go run client/grpc_health_check.go  localhost:9000 service-name

```

### Based on
https://github.com/go-training/grpc-health-check 


