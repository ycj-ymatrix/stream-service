package main

import (
	"fmt"
	"time"

	api "github.com/ycj-ymatrix/stream-service/gen/api/v1"
)

type demo struct {
	api.UnimplementedStreamDemoServiceServer
}

func NewStreamDemoService() api.StreamDemoServiceServer {
	return new(demo)
}

func (s *demo) PingPong(req *api.PingPongRequest, stream api.StreamDemoService_PingPongServer) error {
	count := int(req.Count)

	for i := 0; i < count; i++ {
		greet := fmt.Sprintf("stream response: %d", i)

		err := stream.Send(&api.PingPongResponse{Greet: greet})
		if err != nil {
			return err
		}

		time.Sleep(time.Second)
	}

	return nil
}
