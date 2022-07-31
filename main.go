package main

import (
	"fmt"
	"github.com/ycj-ymatrix/stream-service/server"
	"net"
	"os"
	"os/signal"
	"syscall"

	api "github.com/ycj-ymatrix/stream-service/gen/api/v1"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	api.RegisterStreamDemoServer(srv, server.NewStreamDemo())

	fmt.Printf("stream rpc server starts to service on a random port: %s\n", l.Addr())

	errCh := make(chan error, 1)
	go func() {
		defer close(errCh)
		errCh <- srv.Serve(l)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case e := <-errCh:
		if e != nil {
			panic(e)
		}
	case s := <-sigs:
		panic(s)
	}
}
