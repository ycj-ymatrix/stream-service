# stream-service
    
This is a demo of generating client code for grpc server streaming service.

## About this demo

### grpc

- grpc definition is in `api/v1`
- It defines a service to respond with a stream of `count` number of response as requested

### server implementation

- server implementation is in `server`
- It implements a simple server of grpc service defined above

## client

- client code generated is in `client`
- It is generated with `buf` and `tool/protoc-gen-go-cli` (which is built from https://github.com/yxd-ym/protoc-gen-go-cli/tree/add-streaming-server-support)

## How to run this demo 
 
### Build server

```
cd server
go build -o ../bin/
```

### Build client

```
cd client
go build -o ../bin/
```

### Start server

```
$ bin/server
stream rpc server starts to service on a random port: [::]:58627
```

The port the server is listening to might be different when you are running.
Here is `58627`.

### Send request with client

A sample request is in `request.json`.

In another terminal whil server is running, run the client
with address of the server running above to `--address`

```
$ bin/client --address localhost:58627 --insecure stream-demo ping-pong -f request.json
{
  "greet": "response: 0"
}
{
  "greet": "response: 1"
}
{
  "greet": "response: 2"
}
{
  "greet": "response: 3"
}
{
  "greet": "response: 4"
}
```
