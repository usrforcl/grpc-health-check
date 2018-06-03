package rpc

import (
	"context"


	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"healthcheck/proto"
)

// generate protobuffs
//   protoc --go_out=plugins=grpc,import_path=proto:. *.proto

type healthClient struct {
	client proto.HealthClient
	conn   *grpc.ClientConn
}

// NewGrpcHealthClient returns a new grpc Client.
func NewGrpcHealthClient(conn *grpc.ClientConn) Health {
	client := new(healthClient)
	client.client = proto.NewHealthClient(conn)
	client.conn = conn
	return client
}

func (c *healthClient) Close() error {
	return c.conn.Close()
}

func (c *healthClient) Check(ctx context.Context,service string) (bool, error) {
	var res *proto.HealthCheckResponse
	var err error
	req := &proto.HealthCheckRequest{Service: service}


	res, err = c.client.Check(ctx, req)
	if err == nil {
		if res.GetStatus() == proto.HealthCheckResponse_SERVING {
			return true, nil
		}
		return false, nil
	}
	switch grpc.Code(err) {
	case
		codes.Aborted,
		codes.DataLoss,
		codes.DeadlineExceeded,
		codes.Internal,
		codes.Unavailable:
		// non-fatal errors
	default:
		return false, err
	}

	return false, err
}
