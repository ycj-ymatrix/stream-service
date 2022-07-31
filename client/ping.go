package client

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	api "github.com/ycj-ymatrix/stream-service/gen/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "127.0.0.1:6789", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	cli := api.NewStreamDemoClient(conn)

	stream, err := cli.PingPong(ctx, &api.PingPong_Request{Count: uint32(16)})
	if err != nil {
		return err
	}

	for {
		pong, err := stream.Recv()
		if err != nil {
			switch err {
			case io.EOF:
				os.Exit(0)
			default:
				return err
			}
		}
		fmt.Printf("%v\n", pong.Greet)
	}
}
