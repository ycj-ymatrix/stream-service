package server

import (
	"fmt"
	"math/rand"
	"time"

	api "github.com/ycj-ymatrix/stream-service/gen/api/v1"
)

type demo struct {
	api.UnimplementedStreamDemoServer
}

func NewStreamDemo() api.StreamDemoServer {
	return new(demo)
}

func (s *demo) PingPong(req *api.PingPong_Request, stream api.StreamDemo_PingPongServer) error {
	count := int(req.Count)
	for i := 1; i <= count; i++ {
		sleep := rand.Intn(3)

		greet := fmt.Sprintf("greet: No.%d, and sleep %d seconds", i, sleep)
		if i == 1 {
			greet = fmt.Sprintf("I will greet %d times\ngreet: No.%d, and sleep %d seconds", count, i, sleep)
		}
		if i == count {
			greet = fmt.Sprintf("greet: No.%d, Bye...", i)
			sleep = 0
		}

		err := stream.Send(&api.PingPong_Response{Greet: greet})
		if err != nil {
			return err
		}

		time.Sleep(time.Duration(sleep) * time.Second)
	}

	return nil
}
